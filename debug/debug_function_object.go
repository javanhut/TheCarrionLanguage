package main

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/object"
)

func main() {
	// Test single statement function
	fmt.Println("=== SINGLE STATEMENT FUNCTION ===")
	input1 := `spell test(param): return param`
	l1 := lexer.New(input1)
	p1 := parser.New(l1)
	program1 := p1.ParseProgram()
	env1 := object.NewEnvironment()
	result1 := evaluator.Eval(program1, env1, nil)
	
	// Get the function from environment
	if fn, ok := env1.Get("test"); ok {
		if fnObj, ok := fn.(*object.Function); ok {
			fmt.Printf("Function Body Type: %T\n", fnObj.Body)
			fmt.Printf("Function Body: %s\n", fnObj.Body.String())
			fmt.Printf("Function Body Statements: %d\n", len(fnObj.Body.Statements))
		}
	}
	
	fmt.Printf("Result: %s\n", result1.Inspect())
	
	// Test multi statement function
	fmt.Println("\n=== MULTI STATEMENT FUNCTION ===")
	input2 := `spell test(param):
    x = param + 1
    return param`
	l2 := lexer.New(input2)
	p2 := parser.New(l2)
	program2 := p2.ParseProgram()
	env2 := object.NewEnvironment()
	result2 := evaluator.Eval(program2, env2, nil)
	
	// Get the function from environment
	if fn, ok := env2.Get("test"); ok {
		if fnObj, ok := fn.(*object.Function); ok {
			fmt.Printf("Function Body Type: %T\n", fnObj.Body)
			fmt.Printf("Function Body: %s\n", fnObj.Body.String())
			fmt.Printf("Function Body Statements: %d\n", len(fnObj.Body.Statements))
		}
	}
	
	fmt.Printf("Result: %s\n", result2.Inspect())
}