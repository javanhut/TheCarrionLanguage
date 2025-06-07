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
	
	// Step through parsing manually to see what tokens we get
	fmt.Println("Starting parse...")
	
	// Parse the program
	program := p.ParseProgram()
	
	if len(p.Errors()) > 0 {
		fmt.Printf("Parser errors: %v\n", p.Errors())
	} else {
		fmt.Printf("Parsed successfully\n")
		fmt.Printf("Number of statements: %d\n", len(program.Statements))
		
		fmt.Printf("Total program statements: %d\n", len(program.Statements))
		for i, stmt := range program.Statements {
			fmt.Printf("Program statement %d: %T\n", i, stmt)
			if fn, ok := stmt.(*ast.FunctionDefinition); ok {
				fmt.Printf("  Function name: %s\n", fn.Name.Value)
				fmt.Printf("  Function body statements: %d\n", len(fn.Body.Statements))
				for j, s := range fn.Body.Statements {
					fmt.Printf("    Body statement %d: %T\n", j, s)
				}
			}
		}
	}
}