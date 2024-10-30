package parser

import (
	"testing"

	"thecarrionlang/ast"
	"thecarrionlang/lexer"
)

func TestAssignmentStatements(t *testing.T) {
	input := `
  x = 5
  y = 11
  foobar = 838383
  `
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)
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

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error : %q", msg)
	}
	t.FailNow()
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

func TestReturnStatements(t *testing.T) {
	input := `
  return 5
  return 10
  return 993322
  `

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}
