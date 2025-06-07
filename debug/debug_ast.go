package main

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
)

func main() {
	// Test single statement function
	fmt.Println("=== SINGLE STATEMENT FUNCTION ===")
	input1 := `spell test(param): return param`
	l1 := lexer.New(input1)
	p1 := parser.New(l1)
	program1 := p1.ParseProgram()
	fmt.Printf("AST: %s\n", program1.String())
	
	// Test multi statement function
	fmt.Println("\n=== MULTI STATEMENT FUNCTION ===")
	input2 := `spell test(param):
    print("hello")
    return param`
	l2 := lexer.New(input2)
	p2 := parser.New(l2)
	program2 := p2.ParseProgram()
	fmt.Printf("AST: %s\n", program2.String())
}