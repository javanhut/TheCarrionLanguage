package parser

import (
	"fmt"
	"testing"

	"github.com/javanhut/TheCarrionLanguage/src/ast"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
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
		leftVal  interface{}
		operator string
		rightVal interface{}
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
		{"5 > 5", 5, ">", 5},
		{"5 < 5", 5, "<", 5},
		{"5 == 5", 5, "==", 5},
		{"5 != 5", 5, "!=", 5},
		{"True == True", true, "==", true},
		{"True != False", true, "!=", false},
		{"False == False", false, "==", false},
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

		if !testLiteralExpression(t, exp.Left, tt.leftVal) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.rightVal) {
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
	input := `5 + `

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
otherwise (x > y):
    return y
else:
    return z
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
		{
			"true",
			"true",
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
			"a * [1, 2, 3, 4][b * c] * d",
			"((a * ([1, 2, 3, 4][(b * c)])) * d)",
		},
		{
			"add(a * b[2], b[1], 2 * [1, 2][1])",
			"add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))",
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
	case bool:
		return testBooleanLiteral(t, expr, v)
	default:
		t.Errorf("type of expr not handled. got=%T", expr)
		return false
	}
}

func testBooleanLiteral(t *testing.T, expr ast.Expression, value bool) bool {
	bo, ok := expr.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast.Boolean. got=%T", expr)
		return false
	}
	if bo.Value != value {
		t.Errorf("bo.Value not %t. got=%t", value, bo.Value)
		return false
	}
	if bo.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("bo.TokenLiteral not %t. got=%s", value, bo.TokenLiteral())
		return false
	}
	return true
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

func testInfixExpression(
	t *testing.T,
	exp ast.Expression,
	left interface{},
	operator string,
	right interface{},
) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.OperatorExpression. got=%T(%s)", exp, exp)
		return false
	}
	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}
	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s' . got=%q", operator, opExp.Operator)
		return false
	}
	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	return true
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
		{"!True", "!", true},
		{"!False", "!", false},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
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
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("exp is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("exp is not a  ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}
		if !testLiteralExpression(t, exp.Right, tt.value) {
			return
		}
	}
}

func TestStringLiteralExpression(t *testing.T) {
	input := `"hello world"`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	literal, ok := stmt.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("exp not *ast.StringLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != "hello world" {
		t.Errorf("literal.Value not %q. got=%q", "hello world", literal.Value)
	}
}

func TestParsingArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	if !ok {
		t.Fatalf("exp not ast.ArrayLiteral. got=%T", stmt.Expression)
	}
	if len(array.Elements) != 3 {
		t.Fatalf("len(array.Elements) not 3. got=%d", len(array.Elements))
	}
	testIntegerLiteral(t, array.Elements[0], 1)
	testInfixExpression(t, array.Elements[1], 2, "*", 2)
	testInfixExpression(t, array.Elements[2], 3, "+", 3)
}

func TestParsingIndexExpressions(t *testing.T) {
	input := "myArray[1 + 1]"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	indexExp, ok := stmt.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("exp not *ast.IndexExpression. got=%T", stmt.Expression)
	}
	if !testIdentifier(t, indexExp.Left, "myArray") {
		return
	}
	if !testInfixExpression(t, indexExp.Index, 1, "+", 1) {
		return
	}
}

func TestParsingHashLiteralsStringKeys(t *testing.T) {
	input := `{"one": 1, "two": 2, "three": 3}`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	stmt := program.Statements[0].(*ast.ExpressionStatement)
	hash, ok := stmt.Expression.(*ast.HashLiteral)
	if !ok {
		t.Fatalf("exp is not ast.HashLiteral. got=%T", stmt.Expression)
	}
	if len(hash.Pairs) != 3 {
		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.Pairs))
	}
	expected := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for key, value := range hash.Pairs {
		literal, ok := key.(*ast.StringLiteral)
		if !ok {
			t.Errorf("key is not ast.StringLiteral. got=%T", key)
		}
		expectedValue := expected[literal.String()]
		testIntegerLiteral(t, value, expectedValue)
	}
}

func TestParsingEmptyHashLiteral(t *testing.T) {
	input := "{}"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	stmt := program.Statements[0].(*ast.ExpressionStatement)
	hash, ok := stmt.Expression.(*ast.HashLiteral)
	if !ok {
		t.Fatalf("exp is not ast.HashLiteral. got=%T", stmt.Expression)
	}
	if len(hash.Pairs) != 0 {
		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.Pairs))
	}
}

func TestParsingHashLiteralsWithExpressions(t *testing.T) {
	input := `{"one": 0 + 1, "two": 10 - 8, "three": 15 / 5}`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	stmt := program.Statements[0].(*ast.ExpressionStatement)
	hash, ok := stmt.Expression.(*ast.HashLiteral)
	if !ok {
		t.Fatalf("exp is not ast.HashLiteral. got=%T", stmt.Expression)
	}
	if len(hash.Pairs) != 3 {
		t.Errorf("hash.Pairs has wrong length. got=%d", len(hash.Pairs))
	}
	tests := map[string]func(ast.Expression){
		"one": func(e ast.Expression) {
			testInfixExpression(t, e, 0, "+", 1)
		},
		"two": func(e ast.Expression) {
			testInfixExpression(t, e, 10, "-", 8)
		},
		"three": func(e ast.Expression) {
			testInfixExpression(t, e, 15, "/", 5)
		},
	}
	for key, value := range hash.Pairs {
		literal, ok := key.(*ast.StringLiteral)
		if !ok {
			t.Errorf("key is not ast.StringLiteral. got=%T", key)
			continue
		}
		testFunc, ok := tests[literal.String()]
		if !ok {
			t.Errorf("No test function for key %q found", literal.String())
			continue
		}
		testFunc(value)
	}
}

func TestAttemptStatement(t *testing.T) {
	input := `
attempt:
    x = 5
    do_something()
ensnare(TypeError):
    print("Got a type error")
ensnare:
    print("Got something else")
resolve:
    print("Done with attempt")
`

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p) // You can define a helper function to check parser errors.

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d",
			len(program.Statements))
	}

	// Assert the single statement is an *ast.AttemptStatement
	attemptStmt, ok := program.Statements[0].(*ast.AttemptStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.AttemptStatement. got=%T",
			program.Statements[0])
	}

	// Check the TryBlock
	if attemptStmt.TryBlock == nil {
		t.Fatalf("attemptStmt.TryBlock is nil; expected a block")
	}
	// Optionally, assert how many statements are in the TryBlock
	if len(attemptStmt.TryBlock.Statements) != 2 {
		t.Fatalf("expected 2 statements in the try block, got=%d",
			len(attemptStmt.TryBlock.Statements))
	}

	// Check the EnsnareClauses
	if len(attemptStmt.EnsnareClauses) != 2 {
		t.Fatalf("expected 2 ensnare clauses, got=%d",
			len(attemptStmt.EnsnareClauses))
	}

	// 1st ensnare: ensnare(TypeError)
	firstEnsnare := attemptStmt.EnsnareClauses[0]
	if firstEnsnare.Condition == nil {
		t.Errorf("firstEnsnare.Condition is nil; expected an expression for `TypeError`")
	}
	if firstEnsnare.Consequence == nil {
		t.Errorf("firstEnsnare.Consequence is nil; expected a block statement")
	}

	// 2nd ensnare: ensnare:
	secondEnsnare := attemptStmt.EnsnareClauses[1]
	if secondEnsnare.Condition == nil {
		t.Logf("secondEnsnare.Condition is nil as expected for a bare ensnare")
	}
	if secondEnsnare.Consequence == nil {
		t.Errorf("secondEnsnare.Consequence is nil; expected a block statement")
	}

	// Check the resolve block
	if attemptStmt.ResolveBlock == nil {
		t.Fatalf("attemptStmt.ResolveBlock is nil; expected a resolve block")
	}
	if len(attemptStmt.ResolveBlock.Statements) != 1 {
		t.Fatalf("expected 1 statement in resolve block, got=%d",
			len(attemptStmt.ResolveBlock.Statements))
	}
}
