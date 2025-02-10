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

	indentResolved bool
}

func New(input string) *Lexer {
	rawLines := strings.Split(input, "\n")

	l := &Lexer{
		lines:       rawLines,
		indentStack: []int{0},
	}
	if len(l.lines) == 0 {
		l.finished = true
	} else {
		l.currLine = l.lines[0]
	}
	return l
}

func (l *Lexer) NextToken() token.Token {
	if l.finished {
		return token.Token{Type: token.EOF, Literal: ""}
	}

	if l.charIndex == 0 && !l.indentResolved {
		l.indentResolved = true
		newIndent := measureIndent(l.currLine)
		return l.handleIndentChange(newIndent)
	}

	if l.charIndex >= len(l.currLine) {
		tok := token.Token{Type: token.NEWLINE, Literal: "\\n"}
		l.advanceLine()
		return tok
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

		if l.peekCharIsLetterOrDigitOrUnderscore() {
			return l.readIdentifier()
		} else {

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
		if l.peekChar() == '<' { // check for left-shift
			l.charIndex += 2
			return token.Token{Type: token.LSHIFT, Literal: "<<"}
		} else if l.peekChar() == '=' { // less than or equal
			l.charIndex += 2
			return token.Token{Type: token.LE, Literal: "<="}
		}
		l.charIndex++
		return newToken(token.LT, '<')

	case '>':
		if l.peekChar() == '>' { // check for right-shift
			l.charIndex += 2
			return token.Token{Type: token.RSHIFT, Literal: ">>"}
		} else if l.peekChar() == '=' { // greater than or equal
			l.charIndex += 2
			return token.Token{Type: token.GE, Literal: ">="}
		}
		l.charIndex++
		return newToken(token.GT, '>')

	case '^':
		l.charIndex++
		return newToken(token.XOR, '^')

	case '~':
		l.charIndex++
		return newToken(token.TILDE, '~')

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

		return l.readString()
	case '\'':
		return l.readString()

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

func (l *Lexer) handleIndentChange(newIndent int) token.Token {
	currentIndent := l.indentStack[len(l.indentStack)-1]

	if newIndent == currentIndent {

		l.charIndex = newIndent

		return token.Token{Type: token.NEWLINE, Literal: ""}

	}

	if newIndent > currentIndent {

		l.indentStack = append(l.indentStack, newIndent)
		l.charIndex = newIndent
		return token.Token{Type: token.INDENT, Literal: ""}
	}

	l.indentStack = l.indentStack[:len(l.indentStack)-1]

	return token.Token{Type: token.DEDENT, Literal: ""}
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
	tokType := token.LookupIdent(literal)
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
