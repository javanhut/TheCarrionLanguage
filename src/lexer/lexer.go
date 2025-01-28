package lexer

import (
	"strings"
	"thecarrionlanguage/src/token"
	"unicode"
)

type Lexer struct {
	lines       []string
	lineIndex   int
	charIndex   int
	indentStack []int
	currLine    string
	finished    bool

	// We keep track if we already returned an INDENT or DEDENT at the *start* of this line
	// so we don't keep re-checking indentation in a loop.
	// We'll see below how we use this to avoid recursion.
	indentResolved bool
}

func New(input string) *Lexer {
	rawLines := strings.Split(input, "\n")

	l := &Lexer{
		lines:       rawLines,
		indentStack: []int{0}, // base indent = 0
	}
	if len(l.lines) == 0 {
		l.finished = true
	} else {
		l.currLine = l.lines[0]
	}
	return l
}

func (l *Lexer) NextToken() token.Token {
	// 1) If we've finished all lines, emit EOF
	if l.finished {
		return token.Token{Type: token.EOF, Literal: ""}
	}

	// 2) If we’re at the *start* of a line (charIndex=0) and indentation not yet resolved
	if l.charIndex == 0 && !l.indentResolved {
		l.indentResolved = true // Mark that we’ve handled indentation this line
		newIndent := measureIndent(l.currLine)
		return l.handleIndentChange(newIndent)
	}

	// 3) If we've reached the end of the current line -> emit NEWLINE, go to the next line
	if l.charIndex >= len(l.currLine) {
		tok := token.Token{Type: token.NEWLINE, Literal: "\\n"}
		l.advanceLine()
		return tok
	}

	// 4) Otherwise, we are mid-line -> scan a token
	ch := l.currLine[l.charIndex]

	// Skip horizontal whitespace
	if isHorizontalWhitespace(ch) {
		l.charIndex++
		return l.NextToken() // skip
	}

	switch ch {
	case '=':
		if l.peekChar() == '=' {
			l.charIndex += 2
			return token.Token{Type: token.EQ, Literal: "=="}
		}
		l.charIndex++
		return newToken(token.ASSIGN, '=')

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
		return newToken(token.PLUS, '+')

	case '-':
		nxt := l.peekChar()
		if nxt == '-' {
			l.charIndex += 2
			return token.Token{Type: token.MINUS_DECREMENT, Literal: "--"}
		} else if nxt == '=' {
			l.charIndex += 2
			return token.Token{Type: token.DECREMENT, Literal: "-="}
		}
		l.charIndex++
		return newToken(token.MINUS, '-')

	case '*':
		if l.peekChar() == '=' {
			l.charIndex += 2
			return token.Token{Type: token.MULTASSGN, Literal: "*="}
		} else if l.peekChar() == '*' {
			l.charIndex += 2
			return token.Token{Type: token.EXPONENT, Literal: "**"}
		}
		l.charIndex++
		return newToken(token.ASTERISK, '*')
	case '_':
		// Check if next character continues an identifier
		if l.peekCharIsLetterOrDigitOrUnderscore() {
			// If so, read the entire identifier from here (including the initial '_')
			return l.readIdentifier()
		} else {
			// It's a single underscore token
			l.charIndex++
			return token.Token{Type: token.UNDERSCORE, Literal: "_"}
		}

	case '/':
		next := l.peekChar()
		if next == '=' {
			l.charIndex += 2
			return token.Token{Type: token.DIVASSGN, Literal: "/="}
		} else if next == '/' {
			l.skipLineComment()
			return l.NextToken()
		} else if next == '*' {
			l.skipBlockComment()
			return l.NextToken()
		}
		l.charIndex++
		return newToken(token.SLASH, '/')

	case '%':
		l.charIndex++
		return newToken(token.MOD, '%')

	case '<':
		if l.peekChar() == '=' {
			l.charIndex += 2
			return token.Token{Type: token.LE, Literal: "<="}
		}
		l.charIndex++
		return newToken(token.LT, '<')

	case '>':
		if l.peekChar() == '=' {
			l.charIndex += 2
			return token.Token{Type: token.GE, Literal: ">="}
		}
		l.charIndex++
		return newToken(token.GT, '>')

	case '!':
		if l.peekChar() == '=' {
			l.charIndex += 2
			return token.Token{Type: token.NOT_EQ, Literal: "!="}
		}
		l.charIndex++
		return newToken(token.BANG, '!')

	case ',':
		l.charIndex++
		return newToken(token.COMMA, ',')

	case ':':
		l.charIndex++
		return newToken(token.COLON, ':')

	case ';':
		l.charIndex++
		return newToken(token.SEMICOLON, ';')
	case '(':
		l.charIndex++
		return newToken(token.LPAREN, '(')

	case ')':
		l.charIndex++
		return newToken(token.RPAREN, ')')

	case '[':
		l.charIndex++
		return newToken(token.LBRACK, '[')

	case ']':
		l.charIndex++
		return newToken(token.RBRACK, ']')

	case '{':
		l.charIndex++
		return newToken(token.LBRACE, '{')

	case '}':
		l.charIndex++
		return newToken(token.RBRACE, '}')

	case '.':
		l.charIndex++
		return newToken(token.DOT, '.')

	case '#':
		l.charIndex++
		return newToken(token.HASH, '#')

	case '&':
		l.charIndex++
		return newToken(token.AMPERSAND, '&')

	case '|':
		l.charIndex++
		return newToken(token.PIPE, '|')

	case '@':
		l.charIndex++
		return newToken(token.AT, '@')

	case '"':
		// read string
		return l.readString()

	default:
		if isLetter(ch) {
			return l.readIdentifier()
		} else if isDigit(ch) {
			return l.readNumber()
		} else {
			// unknown char
			l.charIndex++
			return token.Token{Type: token.ILLEGAL, Literal: string(ch)}
		}
	}
}

func (l *Lexer) peekCharIsLetterOrDigitOrUnderscore() bool {
	nxt := l.peekChar()
	// If nxt == 0, means end of line => not a letter/digit
	if nxt == 0 {
		return false
	}
	return isLetterOrDigit(nxt) || nxt == '_'
}

// skipLineComment moves charIndex to the end of the current line
func (l *Lexer) skipLineComment() {
	l.charIndex = len(l.currLine)
}

// skipBlockComment consumes everything until '*/' or EOF (end-of-file)
func (l *Lexer) skipBlockComment() {
	// We've already seen the '/*', so move past those two chars
	l.charIndex += 2

	for {
		// If we reach end of this line, go to the next line
		if l.charIndex >= len(l.currLine) {
			l.advanceLine() // move to next line, reset charIndex=0, etc.
			if l.finished {
				// We've hit EOF in the middle of a block comment
				return
			}
			continue
		}

		// If we see '*/', consume it and return
		if l.currLine[l.charIndex] == '*' && l.peekChar() == '/' {
			l.charIndex += 2
			return
		}

		// Otherwise, just move forward
		l.charIndex++
	}
}

func (l *Lexer) handleIndentChange(newIndent int) token.Token {
	// Get the current top of indentation stack
	currentIndent := l.indentStack[len(l.indentStack)-1]

	if newIndent == currentIndent {
		// same indentation: skip those spaces, proceed scanning
		l.charIndex = newIndent
		// Return the next token from the line (NOT calling NextToken again recursively).
		// Instead, we just continue in NextToken after returning a "NOOP" or some placeholder.
		// But we need to return a real token. Let's do a small trick: we forcibly skip
		// returning an indentation token altogether and jump straight to scanning the line.

		// The simplest trick: we forcibly move past indentation, then call NextToken again
		// *but* from the calling function’s perspective (not recursive). We can do that by
		// returning a special “SKIP_INDENT” token or something, and let NextToken re-check.
		//
		// Or we can do a small one-shot token that means "no indent change."
		return token.Token{Type: token.NEWLINE, Literal: ""}
		// Alternatively, you might do an invisible no-op token or something.
		// The point is we do NOT call l.NextToken() here, to avoid recursion.
	}

	if newIndent > currentIndent {
		// We have one or more new indent levels. Typically, we only do ONE indent at a time.
		// But in case the user jumped e.g. from 0 spaces to 8 spaces, we either
		// handle that in increments or just push a single new indent.
		// Let’s do a single level approach:
		l.indentStack = append(l.indentStack, newIndent)
		l.charIndex = newIndent
		return token.Token{Type: token.INDENT, Literal: ""}
	}

	// else if newIndent < currentIndent => DEDENT
	// pop once
	l.indentStack = l.indentStack[:len(l.indentStack)-1]
	// If we still have more to dedent, we handle it on subsequent calls to NextToken.
	// We'll return one DEDENT at a time each time NextToken gets called at line start.
	return token.Token{Type: token.DEDENT, Literal: ""}
}

func (l *Lexer) advanceLine() {
	l.lineIndex++
	l.indentResolved = false // So we handle indentation on the new line
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

// measureIndent counts how many leading spaces/tabs in the line.
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
	// Move past the opening double-quote
	l.charIndex++

	var sb strings.Builder

loop:
	for {
		// If we run out of line, treat it as an unclosed string
		if l.charIndex >= len(l.currLine) {
			break loop
		}

		ch := l.currLine[l.charIndex]

		// Closing quote?
		if ch == '"' {
			l.charIndex++
			break loop
		}

		// Check for escape sequences
		if ch == '\\' {
			// Skip the backslash
			l.charIndex++
			if l.charIndex >= len(l.currLine) {
				// Nothing after the backslash => incomplete escape
				break loop
			}
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
			case '"':
				sb.WriteByte('"')
			default:
				// For unrecognized escapes (e.g. \q), just add the raw character
				sb.WriteByte(esc)
			}
		} else {
			// Normal character, add it to the string
			sb.WriteByte(ch)
		}

		l.charIndex++
	}

	return token.Token{
		Type:    token.STRING,
		Literal: sb.String(),
	}
}

func (l *Lexer) readIdentifier() token.Token {
	start := l.charIndex
	for l.charIndex < len(l.currLine) && isLetterOrDigit(l.currLine[l.charIndex]) {
		l.charIndex++
	}
	literal := l.currLine[start:l.charIndex]
	tokType := token.LookupIdent(literal) // check if it’s a keyword or else IDENT
	return token.Token{Type: tokType, Literal: literal}
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
				// second dot => break
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

func newToken(tt token.TokenType, ch byte) token.Token {
	return token.Token{Type: tt, Literal: string(ch)}
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
