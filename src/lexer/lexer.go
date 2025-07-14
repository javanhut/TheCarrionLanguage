package lexer

import (
	"strings"
	"unicode"

	"github.com/javanhut/TheCarrionLanguage/src/token"
)

type Lexer struct {
	lines       []string
	lineIndex   int
	charIndex   int
	indentStack []int
	currLine    string
	finished    bool
	sourceFile  string // Source file name for error reporting

	indentResolved bool
	tokenQueue     []token.Token // Queue for pending DEDENT tokens
}

func New(input string) *Lexer {
	return NewWithFilename(input, "<input>")
}

func NewWithFilename(input string, sourceFile string) *Lexer {
	rawLines := strings.Split(input, "\n")

	l := &Lexer{
		lines:       rawLines,
		indentStack: []int{0},
		sourceFile:  sourceFile, // Use sourceFile instead of filename to avoid confusion
	}
	if len(l.lines) == 0 {
		l.finished = true
	} else {
		l.currLine = l.lines[0]
	}
	return l
}

func (l *Lexer) NextToken() token.Token {
   // First check if there are queued tokens to return
   if len(l.tokenQueue) > 0 {
       token := l.tokenQueue[0]
       l.tokenQueue = l.tokenQueue[1:]
       return token
   }

   if l.finished {
       // At EOF, unwind any remaining indents
       if len(l.indentStack) > 1 {
           l.indentStack = l.indentStack[:len(l.indentStack)-1]
           return l.newToken(token.DEDENT, "")
       }
       return l.newToken(token.EOF, "")
   }

   // Handle indentation changes at the start of a new line, skipping simple newlines/indents
   // Skip blank or whitespace-only lines entirely
   if l.charIndex == 0 && strings.TrimSpace(l.currLine) == "" {
       l.advanceLine()
       return l.NextToken()
   }
   // Handle indentation changes at the start of a new line, but be selective about tokens
   if l.charIndex == 0 && !l.indentResolved {
       l.indentResolved = true
       newIndent := measureIndent(l.currLine)
       tok := l.handleIndentChange(newIndent)
       // Return all indentation-related tokens
       if tok.Type == token.NEWLINE || tok.Type == token.INDENT || tok.Type == token.DEDENT {
           return tok
       }
       // Continue to next token for other cases
       return l.NextToken()
   }

   // Generate NEWLINE token at end of line
   if l.charIndex >= len(l.currLine) {
       l.advanceLine()
       // Return NEWLINE token for explicit line break
       return l.newToken(token.NEWLINE, "")
   }

	ch := l.currLine[l.charIndex]

	if isHorizontalWhitespace(ch) {
		l.charIndex++
		return l.NextToken()
	}

	if ch == 'f' {
		next := l.peekChar()
		if next == '"' || next == '\'' {
			l.charIndex++
			return l.readFString()
		}
		return l.readIdentifier()
	}
	if ch == 'i' {
		next := l.peekChar()
		if next == '"' || next == '\'' {
			l.charIndex++
			return l.readStringInterpolation()
		}
		return l.readIdentifier()
	}

	switch ch {
	case '=':
		if l.peekChar() == '=' {
			l.charIndex += 2
			return token.Token{Type: token.EQ, Literal: "=="}
		}
		l.charIndex++
		return l.newToken(token.ASSIGN, "=")

	case '+':
		nxt := l.peekChar()
		if nxt == '+' {
			l.charIndex += 2
			return token.Token{Type: token.PLUS_INCREMENT, Literal: "++"}
		} else if nxt == '=' {
			l.charIndex += 2
			return token.Token{Type: token.INCREMENT, Literal: "+="}
		}
		l.charIndex++
		return l.newToken(token.PLUS, "+")

	case '-':
		nxt := l.peekChar()
		if nxt == '-' {
			l.charIndex += 2
			return token.Token{Type: token.MINUS_DECREMENT, Literal: "--"}
		} else if nxt == '=' {
			l.charIndex += 2
			return token.Token{Type: token.DECREMENT, Literal: "-="}
		} else if nxt == '>' {
			l.charIndex += 2
			return token.Token{Type: token.ARROW, Literal: "->"}
		}
		l.charIndex++
		return l.newToken(token.MINUS, "-")

	case '*':
		if l.peekChar() == '=' {
			l.charIndex += 2
			return token.Token{Type: token.MULTASSGN, Literal: "*="}
		} else if l.peekChar() == '*' {
			l.charIndex += 2
			return token.Token{Type: token.EXPONENT, Literal: "**"}
		}
		l.charIndex++
		return l.newToken(token.ASTERISK, "*")
	case '_':
		if l.peekCharIsLetterOrDigitOrUnderscore() {
			return l.readIdentifier()
		} else {
			l.charIndex++
			return token.Token{Type: token.UNDERSCORE, Literal: "_"}
		}
	case '#':
		l.skipLineComment()
		return l.NextToken()
	case '/':
		next := l.peekChar()
		if next == '=' {
			l.charIndex += 2
			return token.Token{Type: token.DIVASSGN, Literal: "/="}
		} else if next == '/' {
			// Always treat // as integer division operator
			l.charIndex += 2
			return token.Token{Type: token.INTDIV, Literal: "//"}
		} else if next == '*' {
			l.skipBlockComment()
			return l.NextToken()
		}
		l.charIndex++
		return l.newToken(token.SLASH, "/")

	case '%':
		l.charIndex++
		return l.newToken(token.MOD, "%")

	case '<':
		if l.peekChar() == '<' { // check for left-shift
			l.charIndex += 2
			return token.Token{Type: token.LSHIFT, Literal: "<<"}
		} else if l.peekChar() == '=' { // less than or equal
			l.charIndex += 2
			return token.Token{Type: token.LE, Literal: "<="}
		} else if l.peekChar() == '-' { // unpack operator
			l.charIndex += 2
			return token.Token{Type: token.UNPACK, Literal: "<-"}
		}
		l.charIndex++
		return l.newToken(token.LT, "<")

	case '>':
		if l.peekChar() == '>' { // check for right-shift
			l.charIndex += 2
			return token.Token{Type: token.RSHIFT, Literal: ">>"}
		} else if l.peekChar() == '=' { // greater than or equal
			l.charIndex += 2
			return token.Token{Type: token.GE, Literal: ">="}
		}
		l.charIndex++
		return l.newToken(token.GT, ">")

	case '^':
		l.charIndex++
		return l.newToken(token.XOR, "^")

	case '~':
		l.charIndex++
		return l.newToken(token.TILDE, "~")

	case '!':
		if l.peekChar() == '=' {
			l.charIndex += 2
			return token.Token{Type: token.NOT_EQ, Literal: "!="}
		}
		l.charIndex++
		return l.newToken(token.BANG, "!")

	case ',':
		l.charIndex++
		return l.newToken(token.COMMA, ",")

	case ':':
		l.charIndex++
		return l.newToken(token.COLON, ":")

	case ';':
		l.charIndex++
		return l.newToken(token.SEMICOLON, ";")
	case '(':
		l.charIndex++
		return l.newToken(token.LPAREN, "(")

	case ')':
		l.charIndex++
		return l.newToken(token.RPAREN, ")")

	case '[':
		l.charIndex++
		return l.newToken(token.LBRACK, "[")

	case ']':
		l.charIndex++
		return l.newToken(token.RBRACK, "]")

	case '{':
		l.charIndex++
		return l.newToken(token.LBRACE, "{")

	case '}':
		l.charIndex++
		return l.newToken(token.RBRACE, "}")

	case '.':
		l.charIndex++
		return l.newToken(token.DOT, ".")


	case '&':
		l.charIndex++
		return l.newToken(token.AMPERSAND, "&")

	case '|':
		l.charIndex++
		return l.newToken(token.PIPE, "|")

	case '@':
		l.charIndex++
		return l.newToken(token.AT, "@")

	case '"':
		return l.readString()
	case '\'':
		return l.readString()

	case '`':
		// Check for triple backtick multiline comments
		if l.peekChar() == '`' && l.peekCharAt(2) == '`' {
			l.skipTripleBacktickComment()
			return l.NextToken()
		}
		// Single backtick is treated as illegal for now
		l.charIndex++
		return token.Token{Type: token.ILLEGAL, Literal: "`"}

	default:
		if isLetter(ch) {
			return l.readIdentifier()
		} else if isDigit(ch) {
			return l.readNumber()
		} else {
			l.charIndex++
			return token.Token{Type: token.ILLEGAL, Literal: string(ch)}
		}
	}
}

// Add this helper function to check if a specific word follows the current position
func (l *Lexer) wordFollows(word string) bool {
	// Check if there's enough characters left
	if l.charIndex+len(word) > len(l.currLine) {
		return false
	}

	// Check if the substring matches the word
	wordStart := l.charIndex
	wordEnd := l.charIndex + len(word)
	possibleWord := l.currLine[wordStart:wordEnd]

	// Ensure it's a complete word by checking if it's followed by a non-identifier character
	isComplete := wordEnd >= len(l.currLine) || !isLetterOrDigit(l.currLine[wordEnd])

	return possibleWord == word && isComplete
}

func (l *Lexer) readFString() token.Token {
	if l.charIndex >= len(l.currLine) {
		return token.Token{Type: token.ILLEGAL, Literal: "unexpected end of line after f"}
	}
	openingQuote := l.currLine[l.charIndex]
	l.charIndex++

	isTriple := false
	if l.charIndex+1 < len(l.currLine) &&
		l.currLine[l.charIndex] == openingQuote &&
		l.currLine[l.charIndex+1] == openingQuote {
		isTriple = true
		l.charIndex += 2
	}

	var sb strings.Builder

	if isTriple {
		for {
			if l.charIndex >= len(l.currLine) {
				sb.WriteByte('\n')
				l.advanceLine()
				if l.finished {
					break
				}
				continue
			}

			if l.charIndex+2 < len(l.currLine) &&
				l.currLine[l.charIndex] == openingQuote &&
				l.currLine[l.charIndex+1] == openingQuote &&
				l.currLine[l.charIndex+2] == openingQuote {
				l.charIndex += 3
				break
			}
			ch := l.currLine[l.charIndex]

			if ch == '\\' {
				l.charIndex++
				if l.charIndex < len(l.currLine) {
					esc := l.currLine[l.charIndex]
					switch esc {
					case 'n':
						sb.WriteByte('\n')
					case 't':
						sb.WriteByte('\t')
					case 'r':
						sb.WriteByte('\r')
					case '\\':
						sb.WriteByte('\\')
					case openingQuote:
						sb.WriteByte(openingQuote)
					default:
						sb.WriteByte(esc)
					}
				}
			} else {
				sb.WriteByte(ch)
			}
			l.charIndex++
		}
	} else {
		for {
			if l.charIndex >= len(l.currLine) {
				break
			}
			ch := l.currLine[l.charIndex]

			if ch == openingQuote {
				l.charIndex++
				break
			}
			if ch == '\\' {
				l.charIndex++
				if l.charIndex < len(l.currLine) {
					esc := l.currLine[l.charIndex]
					switch esc {
					case 'n':
						sb.WriteByte('\n')
					case 't':
						sb.WriteByte('\t')
					case 'r':
						sb.WriteByte('\r')
					case '\\':
						sb.WriteByte('\\')
					case openingQuote:
						sb.WriteByte(openingQuote)
					default:
						sb.WriteByte(esc)
					}
				}
			} else {
				sb.WriteByte(ch)
			}
			l.charIndex++
		}
	}

	return token.Token{
		Type:    token.FSTRING,
		Literal: sb.String(),
	}
}

func (l *Lexer) peekCharIsLetterOrDigitOrUnderscore() bool {
	nxt := l.peekChar()

	if nxt == 0 {
		return false
	}
	return isLetterOrDigit(nxt) || nxt == '_'
}

func (l *Lexer) skipLineComment() {
	l.charIndex = len(l.currLine)
}

func (l *Lexer) skipBlockComment() {
	l.charIndex += 2

	for {
		if l.charIndex >= len(l.currLine) {
			l.advanceLine()
			if l.finished {
				return
			}
			continue
		}

		if l.currLine[l.charIndex] == '*' && l.peekChar() == '/' {
			l.charIndex += 2
			return
		}

		l.charIndex++
	}
}

func (l *Lexer) skipTripleBacktickComment() {
	// Skip the opening ```
	l.charIndex += 3

	for {
		if l.charIndex >= len(l.currLine) {
			l.advanceLine()
			if l.finished {
				return
			}
			continue
		}

		// Check for closing ``` (need exactly 3 characters from current position)
		if l.charIndex+2 < len(l.currLine) &&
			l.currLine[l.charIndex] == '`' &&
			l.currLine[l.charIndex+1] == '`' &&
			l.currLine[l.charIndex+2] == '`' {
			l.charIndex += 3
			return
		}

		l.charIndex++
	}
}

func (l *Lexer) handleIndentChange(newIndent int) token.Token {
	currentIndent := l.indentStack[len(l.indentStack)-1]

	if newIndent == currentIndent {
		l.charIndex = newIndent
		return l.newToken(token.NEWLINE, "")
	}

	if newIndent > currentIndent {
		l.indentStack = append(l.indentStack, newIndent)
		l.charIndex = newIndent
		return l.newToken(token.INDENT, "")
	}

	// Handle multiple DEDENT levels - generate multiple DEDENT tokens
	dedentCount := 0
	for len(l.indentStack) > 1 && l.indentStack[len(l.indentStack)-1] > newIndent {
		l.indentStack = l.indentStack[:len(l.indentStack)-1]
		dedentCount++
	}
	
	// Set charIndex to the new indentation level
	l.charIndex = newIndent
	
	// Queue additional DEDENT tokens if needed
	for i := 1; i < dedentCount; i++ {
		l.tokenQueue = append(l.tokenQueue, l.newToken(token.DEDENT, ""))
	}
	
	// Return the first DEDENT token
	return l.newToken(token.DEDENT, "")
}

func (l *Lexer) advanceLine() {
	l.lineIndex++
	l.indentResolved = false
	l.charIndex = 0
	if l.lineIndex >= len(l.lines) {
		l.finished = true
		l.currLine = ""
		return
	}
	l.currLine = l.lines[l.lineIndex]
}

func (l *Lexer) peekChar() byte {
	if l.charIndex+1 >= len(l.currLine) {
		return 0
	}
	return l.currLine[l.charIndex+1]
}

func (l *Lexer) peekCharAt(offset int) byte {
	if l.charIndex+offset >= len(l.currLine) {
		return 0
	}
	return l.currLine[l.charIndex+offset]
}

func measureIndent(line string) int {
	count := 0
	for _, ch := range line {
		if ch == ' ' {
			count++
		} else if ch == '\t' {
			count += 4
		} else {
			break
		}
	}
	return count
}

func (l *Lexer) readString() token.Token {
	quoteChar := l.currLine[l.charIndex]

	l.charIndex++

	isTriple := false
	if l.charIndex+1 < len(l.currLine) &&
		l.currLine[l.charIndex] == quoteChar &&
		l.currLine[l.charIndex+1] == quoteChar {
		isTriple = true

		l.charIndex += 2
	}

	var sb strings.Builder

	if isTriple {
		for {
			if l.charIndex >= len(l.currLine) {
				sb.WriteByte('\n')
				l.advanceLine()
				if l.finished {
					break
				}
				continue
			}

			if l.charIndex+2 < len(l.currLine) &&
				l.currLine[l.charIndex] == quoteChar &&
				l.currLine[l.charIndex+1] == quoteChar &&
				l.currLine[l.charIndex+2] == quoteChar {
				l.charIndex += 3
				break
			}
			ch := l.currLine[l.charIndex]
			if ch == '\\' {
				l.charIndex++
				if l.charIndex < len(l.currLine) {
					esc := l.currLine[l.charIndex]
					switch esc {
					case 'n':
						sb.WriteByte('\n')
					case 't':
						sb.WriteByte('\t')
					case 'r':
						sb.WriteByte('\r')
					case '\\':
						sb.WriteByte('\\')
					case byte(quoteChar):
						sb.WriteByte(quoteChar)
					default:
						sb.WriteByte(esc)
					}
				}
			} else {
				sb.WriteByte(ch)
			}
			l.charIndex++
		}
		return token.Token{
			Type:    token.DOCSTRING,
			Literal: sb.String(),
		}
	} else {
		for {
			if l.charIndex >= len(l.currLine) {
				break
			}
			ch := l.currLine[l.charIndex]
			if ch == quoteChar {
				l.charIndex++
				break
			}
			if ch == '\\' {
				l.charIndex++
				if l.charIndex < len(l.currLine) {
					esc := l.currLine[l.charIndex]
					switch esc {
					case 'n':
						sb.WriteByte('\n')
					case 't':
						sb.WriteByte('\t')
					case 'r':
						sb.WriteByte('\r')
					case '\\':
						sb.WriteByte('\\')
					case byte(quoteChar):
						sb.WriteByte(quoteChar)
					default:
						sb.WriteByte(esc)
					}
				}
			} else {
				sb.WriteByte(ch)
			}
			l.charIndex++
		}
		return token.Token{
			Type:    token.STRING,
			Literal: sb.String(),
		}
	}
}

func (l *Lexer) readIdentifier() token.Token {
	start := l.charIndex
	for l.charIndex < len(l.currLine) && isLetterOrDigit(l.currLine[l.charIndex]) {
		l.charIndex++
	}
	literal := l.currLine[start:l.charIndex]

	// Handle "not in" as a special case
	if literal == "not" {
		// Save current position
		savedCharIndex := l.charIndex

		// Skip any whitespace
		for l.charIndex < len(l.currLine) && isHorizontalWhitespace(l.currLine[l.charIndex]) {
			l.charIndex++
		}

		// Check if "in" follows
		if l.wordFollows("in") {
			// Consume "in"
			// oldCharIndex := l.charIndex
			l.charIndex += 2 // Length of "in"

			// Return "not in" token
			return token.Token{
				Type:     token.NOT_IN,
				Literal:  "not in",
				Filename: l.sourceFile,
				Line:     l.lineIndex + 1,
				Column:   start + 1,
			}
		}

		// If "in" doesn't follow, restore position and continue normally
		l.charIndex = savedCharIndex
	}

	tokType := token.LookupIdent(literal)
	return token.Token{
		Type:     tokType,
		Literal:  literal,
		Filename: l.sourceFile,
		Line:     l.lineIndex + 1,
		Column:   start + 1,
	}
}

func isLetterOrDigit(ch byte) bool {
	return isLetter(ch) || unicode.IsDigit(rune(ch))
}

func (l *Lexer) readNumber() token.Token {
	start := l.charIndex
	isFloat := false
	for l.charIndex < len(l.currLine) {
		ch := l.currLine[l.charIndex]
		if ch == '.' {
			if isFloat {
				break
			}
			isFloat = true
		} else if !isDigit(ch) {
			break
		}
		l.charIndex++
	}
	literal := l.currLine[start:l.charIndex]
	if isFloat {
		return token.Token{Type: token.FLOAT, Literal: literal}
	}
	return token.Token{Type: token.INT, Literal: literal}
}

func (l *Lexer) newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{
		Type:     tokenType,
		Literal:  literal,
		Filename: l.sourceFile,
		Line:     l.lineIndex + 1, // Make line numbers 1-based for user-friendliness
		Column:   l.charIndex + 1, // Make column numbers 1-based
	}
}

func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

func isDigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}

func isHorizontalWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t'
}

func (l *Lexer) readStringInterpolation() token.Token {
	if l.charIndex >= len(l.currLine) {
		return token.Token{Type: token.ILLEGAL, Literal: "unexpected end of line after i"}
	}

	openingQuote := l.currLine[l.charIndex]
	l.charIndex++

	isTriple := false
	if l.charIndex+1 < len(l.currLine) &&
		l.currLine[l.charIndex] == openingQuote &&
		l.currLine[l.charIndex+1] == openingQuote {
		isTriple = true
		l.charIndex += 2
	}

	var sb strings.Builder
	isBraceOpen := false
	braceDepth := 0
	exprStart := 0

	processChar := func(ch byte) bool {
		if ch == '$' && l.peekChar() == '{' && !isBraceOpen {
			// Start of expression
			l.charIndex++ // Skip the '{'
			isBraceOpen = true
			braceDepth = 1
			exprStart = l.charIndex + 1
			return true
		} else if isBraceOpen {
			if ch == '{' {
				braceDepth++
			} else if ch == '}' {
				braceDepth--
				if braceDepth == 0 {
					// End of expression
					exprStr := l.currLine[exprStart:l.charIndex]
					sb.WriteString("${" + exprStr + "}")
					isBraceOpen = false
					return true
				}
			}
		}

		if !isBraceOpen {
			if ch == '\\' {
				l.charIndex++
				if l.charIndex < len(l.currLine) {
					esc := l.currLine[l.charIndex]
					switch esc {
					case 'n':
						sb.WriteByte('\n')
					case 't':
						sb.WriteByte('\t')
					case 'r':
						sb.WriteByte('\r')
					case '\\':
						sb.WriteByte('\\')
					case openingQuote:
						sb.WriteByte(openingQuote)
					case '$':
						sb.WriteByte('$')
					default:
						sb.WriteByte(esc)
					}
				}
				return true
			}

			if ch == openingQuote && !isTriple {
				// End of string
				return false
			} else if isTriple &&
				l.charIndex+2 < len(l.currLine) &&
				l.currLine[l.charIndex] == openingQuote &&
				l.currLine[l.charIndex+1] == openingQuote &&
				l.currLine[l.charIndex+2] == openingQuote {
				// End of triple-quoted string
				l.charIndex += 2
				return false
			}
		}

		sb.WriteByte(ch)
		return true
	}

	for l.charIndex < len(l.currLine) {
		if !processChar(l.currLine[l.charIndex]) {
			break
		}
		l.charIndex++
	}

	// Skip the closing quote
	if l.charIndex < len(l.currLine) && l.currLine[l.charIndex] == openingQuote {
		l.charIndex++
	}

	return token.Token{
		Type:    token.INTERP,
		Literal: sb.String(),
	}
}
