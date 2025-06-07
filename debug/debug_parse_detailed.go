package main

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/ast"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
)

func main() {
	input := `spell test(param):
    x = param + 1
    return x

result = test(5)`
	
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	
	fmt.Printf("Parser errors: %v\n", p.Errors())
	fmt.Printf("Statements: %d\n", len(program.Statements))
	
	if len(program.Statements) > 0 {
		if fn, ok := program.Statements[0].(*ast.FunctionDefinition); ok {
			fmt.Printf("Function: %s\n", fn.Name.Value)
			fmt.Printf("Body statements: %d\n", len(fn.Body.Statements))
			for i, stmt := range fn.Body.Statements {
				fmt.Printf("  Statement %d: %s\n", i, stmt.String())
			}
		}
	}
}