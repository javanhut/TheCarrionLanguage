package main

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
)

func main() {
	input := `spell test(param):
    x = param + 1
    return x`
	
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	
	fmt.Printf("Parser errors: %v\n", p.Errors())
	fmt.Printf("Statements: %d\n", len(program.Statements))
}