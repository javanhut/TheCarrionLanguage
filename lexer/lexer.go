// lexer/lexer.go
package lexer

import (
	"bytes" //"fmt"
	"unicode"

	"thecarrionlanguage/token"
)

// Lexer represents a lexical scanner.
type Lexer struct {
	input        string
	position     int  // Current position in input (points to current char)
	readPosition int  // Current reading position in input (after current char)
	ch           rune // Current char under examination
	tokens       []token.Token
	indentStack  []int // Stack to track indentation levels
}

// New initializes a new Lexer with the provided input string.
func New(input string) *Lexer {
	l := &Lexer{
		input:       input,
		indentStack: []int{0}, // Initialize stack with a base indentation level of 0
	}
	l.readChar()
	return l
}

// readChar reads the next character and advances the positions.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for NUL, signifies EOF
	} else {
		l.ch = rune(l.input[l.readPosition])
	}
	l.position = l.readPosition
	l.readPosition++
}

// peekChar returns the next character without advancing the position.
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return rune(l.input[l.readPosition])
}

// NextToken lexes and returns the next token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '+':
		if l.peekChar() == '+' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.PLUS_INCREMENT, Literal: literal}
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.INCREMENT, Literal: literal}
		} else {
			tok = newToken(token.PLUS, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.MULTASSGN, Literal: literal}
		} else {
			tok = newToken(token.ASTERISK, l.ch)
		}
	case '-':
		if l.peekChar() == '-' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.MINUS_DECREMENT, Literal: literal}
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.DECREMENT, Literal: literal}
		} else {
			tok = newToken(token.MINUS, l.ch)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '.':
		tok = newToken(token.DOT, l.ch)
	case '/':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.DIVASSGN, Literal: literal}
		} else {
			tok = newToken(token.SLASH, l.ch)
		}
	case '%':
		tok = newToken(token.MOD, l.ch)
	case '[':
		tok = newToken(token.LBRACK, l.ch)
	case ']':
		tok = newToken(token.RBRACK, l.ch)
	case '|':
		tok = newToken(token.PIPE, l.ch)
	case '&':
		tok = newToken(token.AMPERSAND, l.ch)
	case '#':
		tok = newToken(token.HASH, l.ch)
	case '@':
		tok = newToken(token.AT, l.ch)
	case '\n':
		tok = token.Token{Type: token.NEWLINE, Literal: "\\n"}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LE, Literal: literal}
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.GE, Literal: literal}
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case '"':
		tok = token.Token{Type: token.STRING, Literal: l.readString()}
		return tok

	case 0:
		// Emit remaining DEDENT tokens before EOF
		if len(l.indentStack) > 1 {
			l.indentStack = l.indentStack[:len(l.indentStack)-1]
			tok = token.Token{Type: token.DEDENT, Literal: ""}
		} else {
			tok = token.Token{Type: token.EOF, Literal: ""}
		}
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tok.Type = token.LookupIdent(literal)
			tok.Literal = literal
			return tok
		} else if isDigit(l.ch) {
			literal, tokType := l.readNumber()
			tok.Type = tokType
			tok.Literal = literal
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	// fmt.Printf("[DEBUG] NextToken -> Type: %s, Literal: %q\n", tok.Type, tok.Literal)

	l.readChar()
	return tok
}

// skipWhiteSpace skips over spaces, tabs, and handles newlines for indentation.
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		if l.ch == '\n' {
			l.emitNewline()
			l.readChar()
			l.handleIndentation()
		} else {
			l.readChar()
		}
	}
}

func (l *Lexer) readString() string {
	var out bytes.Buffer

	// Skip the opening quote mark
	l.readChar()

	for {
		if l.ch == '"' {
			// Found the closing quote
			l.readChar() // Move past the closing quote
			return out.String()
		}

		if l.ch == 0 {
			// EOF before string termination
			return out.String()
		}

		if l.ch == '\\' {
			l.readChar()
			switch l.ch {
			case 'n':
				out.WriteByte('\n')
			case 't':
				out.WriteByte('\t')
			case 'r':
				out.WriteByte('\r')
			case '\\':
				out.WriteByte('\\')
			case '"':
				out.WriteByte('"')
			default:
				out.WriteByte('\\')
				out.WriteByte(byte(l.ch))
			}
		} else {
			out.WriteByte(byte(l.ch))
		}
		l.readChar()
	}
}

// emitNewline emits a NEWLINE token.
func (l *Lexer) emitNewline() {
	tok := token.Token{Type: token.NEWLINE, Literal: "\n"}
	l.tokens = append(l.tokens, tok)
}

// handleIndentation manages INDENT and DEDENT tokens based on the current indentation.
func (l *Lexer) handleIndentation() {
	startPos := l.position
	// Count the number of consecutive spaces or tabs for indentation
	indent := 0
	for l.ch == ' ' || l.ch == '\t' {
		if l.ch == ' ' {
			indent += 1
		} else if l.ch == '\t' {
			indent += 4 // Assuming a tab equals 4 spaces; adjust as needed
		}
		l.readChar()
	}

	indentString := l.input[startPos:l.position]
	currentIndent := len(indentString)

	lastIndent := l.indentStack[len(l.indentStack)-1]

	if currentIndent > lastIndent {
		// Increased indentation level
		l.indentStack = append(l.indentStack, currentIndent)
		l.emitIndentToken(token.INDENT)
	} else if currentIndent < lastIndent {
		// Decreased indentation level
		for len(l.indentStack) > 0 && currentIndent < lastIndent {
			l.indentStack = l.indentStack[:len(l.indentStack)-1]
			l.emitIndentToken(token.DEDENT)
			if len(l.indentStack) > 0 {
				lastIndent = l.indentStack[len(l.indentStack)-1]
			}
		}
		if currentIndent != lastIndent {
			// Indentation does not match any previous level
			l.emitIndentToken(token.ILLEGAL)
		}
	}
	// If currentIndent == lastIndent, do nothing
}

// emitIndentToken appends an INDENT, DEDENT, or ILLEGAL token to the tokens slice.
func (l *Lexer) emitIndentToken(tokenType token.TokenType) {
	tok := token.Token{Type: tokenType, Literal: ""}
	l.tokens = append(l.tokens, tok)
}

// readIdentifier reads an identifier starting with a letter or underscore.
func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

// readNumber reads a numeric literal (integer).
func (l *Lexer) readNumber() (string, token.TokenType) {
	start := l.position
	isFloat := false

	for isDigit(l.ch) || l.ch == '.' {
		if l.ch == '.' {
			if isFloat {
				break
			}
			isFloat = true
		}
		l.readChar()
	}
	literal := l.input[start:l.position]
	if isFloat {
		return literal, token.FLOAT
	}
	return literal, token.INT
}

// newToken creates a new Token instance.
func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter checks if the character is a letter or underscore.
func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

// isDigit checks if the character is a digit.
func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}
