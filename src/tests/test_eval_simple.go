package main

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/object"
)

func main() {
	// Test 1: Single statement function (should work)
	input1 := `
spell add_one(x): return x + 1

result = add_one(5)
print("Result is:")
print(result)
result
`
	
	fmt.Println("=== TEST 1: Single statement function ===")
	l1 := lexer.New(input1)
	p1 := parser.New(l1)
	program1 := p1.ParseProgram()
	
	if len(p1.Errors()) > 0 {
		fmt.Printf("Parser errors: %v\n", p1.Errors())
	} else {
		env1 := object.NewEnvironment()
		result1 := evaluator.Eval(program1, env1, nil)
		fmt.Printf("Result: %s\n", result1.Inspect())
		fmt.Printf("Result type: %T\n", result1)
	}
	
	// Test 2: Multi-statement function (parameter scoping issue)
	input2 := `
spell test(param):
    x = param + 1
    return x

test(5)
`
	
	fmt.Println("\n=== TEST 2: Multi-statement function ===")
	l2 := lexer.New(input2)
	p2 := parser.New(l2)
	program2 := p2.ParseProgram()
	
	if len(p2.Errors()) > 0 {
		fmt.Printf("Parser errors: %v\n", p2.Errors())
	} else {
		env2 := object.NewEnvironment()
		result2 := evaluator.Eval(program2, env2, nil)
		fmt.Printf("Result: %s\n", result2.Inspect())
	}
}