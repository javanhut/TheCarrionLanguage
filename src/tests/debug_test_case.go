package main

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/object"
)

func main() {
	input := `
spell add(x, y):
    return x + y

result = add(2, 3)
result
`
	
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	
	fmt.Printf("Parser errors: %v\n", p.Errors())
	fmt.Printf("Program statements: %d\n", len(program.Statements))
	
	env := object.NewEnvironment()
	result := evaluator.Eval(program, env, nil)
	fmt.Printf("Result: %s\n", result.Inspect())
	
	// Check if result variable was set
	if val, ok := env.Get("result"); ok {
		fmt.Printf("result variable: %s\n", val.Inspect())
	}
}