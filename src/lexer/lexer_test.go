// lexer/lexer_test.go
package lexer

import (
	"testing"

	"github.com/javanhut/TheCarrionLanguage/src/token"
)

func TestNextToken(t *testing.T) {
	input := `five = 5
  ten = 10
  spell add(x , y):
    return x + y

  result = add(five, ten)
  result >= 16
  "foobar"
  "foo bar"
  [1, 2]`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NEWLINE, ""},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.NEWLINE, ""},
		{token.INDENT, ""},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.NEWLINE, ""},
		{token.NEWLINE, ""},
		{token.SPELL, "spell"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.COLON, ":"},
		{token.NEWLINE, ""},
		{token.INDENT, ""},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.NEWLINE, ""},
		{token.DEDENT, ""},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.NEWLINE, ""},
		{token.NEWLINE, ""},
		{token.IDENT, "result"},
		{token.GE, ">="},
		{token.INT, "16"},
		{token.NEWLINE, ""},
		{token.NEWLINE, ""},
		{token.STRING, "foobar"},
		{token.NEWLINE, ""},
		{token.NEWLINE, ""},
		{token.STRING, "foo bar"},
		{token.NEWLINE, ""},
		{token.NEWLINE, ""},
		{token.LBRACK, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACK, "]"},
		{token.NEWLINE, ""},
		{token.DEDENT, ""},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		// Check Token Type
		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - Token Type Wrong.\nExpected Type: %q\nActual Type:   %q\n",
				i, tt.expectedType, tok.Type)
		}

		// Check Token Literal
		if tok.Literal != tt.expectedLiteral {
			t.Errorf(
				"tests[%d] - Token Literal Wrong.\nExpected Literal: %q\nActual Literal:   %q\n",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}
