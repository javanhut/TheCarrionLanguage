package parser

import (
	"fmt"
	"testing"
	"thecarrionlang/ast"
	"thecarrionlang/lexer"
)

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf(
			"program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0],
		)
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5", literal.TokenLiteral())
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input    string
		leftVal  int64
		operator string
		rightVal int64
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
		{"5 > 5", 5, ">", 5},
		{"5 < 5", 5, "<", 5},
		{"5 == 5", 5, "==", 5},
		{"5 != 5", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf(
				"program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0],
			)
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, tt.leftVal) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.rightVal) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
		return false
	}

	return true
}

func TestParserErrors(t *testing.T) {
	input := "5 +"

	l := lexer.New(input)
	p := New(l)
	p.ParseProgram()

	errors := p.Errors()
	if len(errors) == 0 {
		t.Fatalf("expected parser errors, but got none")
	}

	for _, err := range errors {
		t.Logf("parser error: %s", err)
	}
}

func TestParsingIfStatement(t *testing.T) {
	input := `
if (x < y):
    return x
else:
    return y
`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program.Statements does not contain %d statements. got=%d\n",
			1,
			len(program.Statements),
		)
	}

	_, ok := program.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.IfStatement. got=%T", program.Statements[0])
	}

	// Further checks on stmt.Condition, stmt.Consequence, stmt.Alternative...
}

func TestPrintAST(t *testing.T) {
	input := "5 + 5 * (10 - 2)"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	fmt.Println(program.String())
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"1 + 2 + 3",
			"((1 + 2) + 3)",
		},
		{
			"1 + 2 * 3",
			"(1 + (2 * 3))",
		},
		{
			"1 - 2 * 3",
			"(1 - (2 * 3))",
		},
		{
			"(1 + 2) * 3",
			"((1 + 2) * 3)",
		},
		{
			"1 + 2 * 3 + 4 / 2",
			"((1 + (2 * 3)) + (4 / 2))",
		},
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},

		{
			"!(true == true)",
			"(!(true == true))",
		},
		{
			"a++",
			"(a++)",
		},
		{
			"++a",
			"(++a)",
		},
		{
			"a++ + b",
			"((a++) + b)",
		},
		{
			"--b * a",
			"((--b) * a)",
		},
		{
			"a + b * c-- - d",
			"((a + (b * (c--))) - d)",
		},
		{
			"a * (b + c) * d",
			"((a * (b + c)) * d)",
		},
		{
			"a == b or c != d",
			"((a == b) or (c != d))",
		},
		{
			"(a == b) != (c == d)",
			"((a == b) != (c == d))",
		},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("test[%d] - expected=%q, got=%q", i, tt.expected, actual)
		}
	}
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	return true
}

func TestFunctionDefinitionParsing(t *testing.T) {
	input := `
spell add(x, y):
    return x + y
`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program.Statements does not contain %d statements. got=%d\n",
			1,
			len(program.Statements),
		)
	}

	stmt, ok := program.Statements[0].(*ast.FunctionDefinition)
	if !ok {
		t.Fatalf(
			"program.Statements[0] is not ast.FunctionDefinition. got=%T",
			program.Statements[0],
		)
	}

	if stmt.Name.Value != "add" {
		t.Errorf("Function name wrong. Expected 'add', got '%s'", stmt.Name.Value)
	}

	if len(stmt.Parameters) != 2 {
		t.Fatalf("Function 'add' parameters wrong. Expected 2, got=%d", len(stmt.Parameters))
	}

	testLiteralExpression(t, stmt.Parameters[0], "x")
	testLiteralExpression(t, stmt.Parameters[1], "y")

	if len(stmt.Body.Statements) != 1 {
		t.Fatalf("Function body does not contain 1 statement. got=%d", len(stmt.Body.Statements))
	}

	bodyStmt, ok := stmt.Body.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf(
			"Function body statement is not ast.ReturnStatement. got=%T",
			stmt.Body.Statements[0],
		)
	}

	// Test the expression x + y
	exp, ok := bodyStmt.ReturnValue.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Return value is not ast.InfixExpression. got=%T", bodyStmt.ReturnValue)
	}

	if !testIdentifier(t, exp.Left, "x") {
		return
	}

	if exp.Operator != "+" {
		t.Errorf("Operator is not '+'. got=%s", exp.Operator)
	}

	if !testIdentifier(t, exp.Right, "y") {
		return
	}
}

func testLiteralExpression(t *testing.T, expr ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, expr, int64(v))
	case int64:
		return testIntegerLiteral(t, expr, v)
	case string:
		return testIdentifier(t, expr, v)
	default:
		t.Errorf("type of expr not handled. got=%T", expr)
		return false
	}
}

func TestFunctionDefinitionInlineParsing(t *testing.T) {
	input := `
spell add(x, y): return x + y
`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program.Statements does not contain %d statements. got=%d\n",
			1,
			len(program.Statements),
		)
	}

	stmt, ok := program.Statements[0].(*ast.FunctionDefinition)
	if !ok {
		t.Fatalf(
			"program.Statements[0] is not ast.FunctionDefinition. got=%T",
			program.Statements[0],
		)
	}

	if stmt.Name.Value != "add" {
		t.Errorf("Function name wrong. Expected 'add', got '%s'", stmt.Name.Value)
	}

	if len(stmt.Parameters) != 2 {
		t.Fatalf("Function 'add' parameters wrong. Expected 2, got=%d", len(stmt.Parameters))
	}

	testLiteralExpression(t, stmt.Parameters[0], "x")
	testLiteralExpression(t, stmt.Parameters[1], "y")

	if len(stmt.Body.Statements) != 1 {
		t.Fatalf("Function body does not contain 1 statement. got=%d", len(stmt.Body.Statements))
	}

	bodyStmt, ok := stmt.Body.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf(
			"Function body statement is not ast.ReturnStatement. got=%T",
			stmt.Body.Statements[0],
		)
	}

	// Test the expression x + y
	exp, ok := bodyStmt.ReturnValue.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Return value is not ast.InfixExpression. got=%T", bodyStmt.ReturnValue)
	}

	if !testIdentifier(t, exp.Left, "x") {
		return
	}

	if exp.Operator != "+" {
		t.Errorf("Operator is not '+'. got=%s", exp.Operator)
	}

	if !testIdentifier(t, exp.Right, "y") {
		return
	}
}
