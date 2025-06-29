package main

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/ast"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/object"
)

func testCode(name, code string) {
	fmt.Printf("\n=== %s ===\n", name)
	fmt.Printf("Code:\n%s\n", code)
	
	l := lexer.New(code)
	p := parser.New(l)
	program := p.ParseProgram()
	
	if len(p.Errors()) > 0 {
		fmt.Printf("Parser errors: %v\n", p.Errors())
		return
	}
	
	// Check parsing
	for _, stmt := range program.Statements {
		if fn, ok := stmt.(*ast.FunctionDefinition); ok {
			fmt.Printf("Function %s has %d statements in body\n", 
				fn.Name.Value, len(fn.Body.Statements))
		}
	}
	
	// Evaluate
	env := object.NewEnvironment()
	result := evaluator.Eval(program, env, nil)
	if result != nil {
		fmt.Printf("Result: %s\n", result.Inspect())
	}
}

func main() {
	// Test 1: Multi-statement function with parameter access
	testCode("Multi-statement function", `
spell factorial(n):
    if n == 0:
        return 1
    result = n * factorial(n - 1)
    return result

factorial(5)`)

	// Test 2: Simple multi-statement function
	testCode("Simple multi-statement", `
spell add_and_double(x):
    temp = x + 1
    result = temp * 2
    return result

add_and_double(5)`)

	// Test 3: Parameter access in different positions
	testCode("Parameter access", `
spell test(a, b):
    x = a + b
    return x

test(3, 2)`)
}