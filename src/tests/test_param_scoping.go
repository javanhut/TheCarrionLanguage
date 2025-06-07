package main

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/ast"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/object"
)

func main() {
	input := `
spell test(param):
    x = param + 1
    return x

test(5)
`
	
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	
	if len(p.Errors()) > 0 {
		fmt.Printf("Parser errors: %v\n", p.Errors())
		return
	}
	
	fmt.Printf("Program statements: %d\n", len(program.Statements))
	if len(program.Statements) > 0 {
		if fn, ok := program.Statements[0].(*ast.FunctionDefinition); ok {
			fmt.Printf("Function %s has %d statements\n", fn.Name.Value, len(fn.Body.Statements))
		}
	}
	
	env := object.NewEnvironment()
	result := evaluator.Eval(program, env, nil)
	fmt.Printf("Result: %s\n", result.Inspect())
}