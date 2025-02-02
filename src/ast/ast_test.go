package ast

import (
	"testing"

	"github.com/javanhut/TheCarrionLanguage/src/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&AssignStatement{
				Token: token.Token{Type: token.ASSIGN, Literal: "="},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	expected := "myVar = anotherVar"
	if program.String() != expected {
		t.Errorf("program.String() wrong. got=%q, want=%q", program.String(), expected)
	}
}
