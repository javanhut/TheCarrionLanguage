package parser

import (
	"testing"

	"thecarrionlang/ast"
	"thecarrionlang/lexer"
)

func TestAssignmentStatements(t *testing.T) {
	input := `
  x = 5
  y= 10
  foobar = 838383
  `
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 Statements! got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifer string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testAssignmentStatment(t, stmt, tt.expectedIdentifer) {
			return
		}
	}
}

func testAssignmentStatment(t *testing.T, s ast.Statement, name string) bool {
	assignStmt, ok := s.(*ast.AssignStatement)
	if !ok {
		t.Errorf("s is not *ast.AssignStatement. git %T", s)
	}

	if assignStmt.Name.Value != name {
		t.Errorf("assignStmt.Name.Value not '%s'. got=%s", name, assignStmt.Name.Value)
		return false
	}
	if assignStmt.Name.TokenLiteral() != name {
		t.Errorf("assignStmt.Name.TokenLiteral not '%s'. got=%s", name, assignStmt.Name.TokenLiteral())
		return false
	}
	return true
}
