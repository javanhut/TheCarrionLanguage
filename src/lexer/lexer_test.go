// lexer/lexer_test.go
package lexer

import (
	"strings"
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

// TestIndentationSpacesOnly tests that a file using only spaces for indentation works correctly
func TestIndentationSpacesOnly(t *testing.T) {
	input := `spell test():
    x = 1
    return x`

	l := New(input)

	// Collect all tokens and check no INDENT_ERROR
	for {
		tok := l.NextToken()
		if tok.Type == token.INDENT_ERROR {
			t.Errorf("Unexpected INDENT_ERROR: %s", tok.Literal)
		}
		if tok.Type == token.EOF {
			break
		}
	}
}

// TestIndentationTabsOnly tests that a file using only tabs for indentation works correctly
func TestIndentationTabsOnly(t *testing.T) {
	input := "spell test():\n\tx = 1\n\treturn x"

	l := New(input)

	// Collect all tokens and check no INDENT_ERROR
	for {
		tok := l.NextToken()
		if tok.Type == token.INDENT_ERROR {
			t.Errorf("Unexpected INDENT_ERROR: %s", tok.Literal)
		}
		if tok.Type == token.EOF {
			break
		}
	}
}

// TestIndentationMixedInSameLine tests that mixing tabs and spaces in the same line produces an error
func TestIndentationMixedInSameLine(t *testing.T) {
	// Mix of tab and spaces on the same indentation line
	input := "spell test():\n\t x = 1"  // tab followed by space

	l := New(input)

	foundError := false
	for {
		tok := l.NextToken()
		if tok.Type == token.INDENT_ERROR {
			foundError = true
			if !strings.Contains(tok.Literal, "mixed tabs and spaces") {
				t.Errorf("Expected 'mixed tabs and spaces' error, got: %s", tok.Literal)
			}
		}
		if tok.Type == token.EOF {
			break
		}
	}

	if !foundError {
		t.Errorf("Expected INDENT_ERROR for mixed tabs and spaces in same line")
	}
}

// TestIndentationSpacesThenTabs tests that using spaces first then tabs produces an error
func TestIndentationSpacesThenTabs(t *testing.T) {
	// First use spaces, then use tabs
	input := "spell test():\n    x = 1\n\ty = 2"  // 4 spaces, then tab

	l := New(input)

	foundError := false
	for {
		tok := l.NextToken()
		if tok.Type == token.INDENT_ERROR {
			foundError = true
			if !strings.Contains(tok.Literal, "inconsistent indentation") {
				t.Errorf("Expected 'inconsistent indentation' error, got: %s", tok.Literal)
			}
			if !strings.Contains(tok.Literal, "expected spaces") {
				t.Errorf("Expected error to mention 'expected spaces', got: %s", tok.Literal)
			}
			if !strings.Contains(tok.Literal, "got tabs") {
				t.Errorf("Expected error to mention 'got tabs', got: %s", tok.Literal)
			}
		}
		if tok.Type == token.EOF {
			break
		}
	}

	if !foundError {
		t.Errorf("Expected INDENT_ERROR for spaces then tabs")
	}
}

// TestIndentationTabsThenSpaces tests that using tabs first then spaces produces an error
func TestIndentationTabsThenSpaces(t *testing.T) {
	// First use tabs, then use spaces
	input := "spell test():\n\tx = 1\n    y = 2"  // tab, then 4 spaces

	l := New(input)

	foundError := false
	for {
		tok := l.NextToken()
		if tok.Type == token.INDENT_ERROR {
			foundError = true
			if !strings.Contains(tok.Literal, "inconsistent indentation") {
				t.Errorf("Expected 'inconsistent indentation' error, got: %s", tok.Literal)
			}
			if !strings.Contains(tok.Literal, "expected tabs") {
				t.Errorf("Expected error to mention 'expected tabs', got: %s", tok.Literal)
			}
			if !strings.Contains(tok.Literal, "got spaces") {
				t.Errorf("Expected error to mention 'got spaces', got: %s", tok.Literal)
			}
		}
		if tok.Type == token.EOF {
			break
		}
	}

	if !foundError {
		t.Errorf("Expected INDENT_ERROR for tabs then spaces")
	}
}

// TestIndentationNoIndent tests that files with no indentation work correctly
func TestIndentationNoIndent(t *testing.T) {
	input := `x = 1
y = 2
z = 3`

	l := New(input)

	// Collect all tokens and check no INDENT_ERROR
	for {
		tok := l.NextToken()
		if tok.Type == token.INDENT_ERROR {
			t.Errorf("Unexpected INDENT_ERROR: %s", tok.Literal)
		}
		if tok.Type == token.EOF {
			break
		}
	}
}
