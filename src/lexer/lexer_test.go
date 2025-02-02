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
  [1, 2]
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// Line 1: five = 5
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},

		// Line 2: two spaces + ten = 10
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		// Line 3: two spaces + spell add(x , y):
		{token.SPELL, "spell"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.COLON, ":"},
		// Line 4: four spaces + return x + y
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		// Line 6: two spaces + result = add(five, ten)
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},

		// result greater than or equal to 16
		{token.IDENT, "result"},
		{token.GE, ">="},
		{token.INT, "16"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACK, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACK, "]"},
		// End of input: dedent to base and EOF
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
