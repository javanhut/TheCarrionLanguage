package main

import (
	"fmt"

	"thecarrionlang/lexer"
	"thecarrionlang/parser"
	"thecarrionlang/token"
)

func main() {
	input := "(5 + (5 * (10 - 2)))"
	l := lexer.New(input)

	// Print out tokens for debugging
	fmt.Println("----- Debugging Lexer Tokens -----")
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}

	// Reinitialize the lexer since the above loop consumed all tokens
	l = lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	fmt.Println("----- Debugging Parser Output -----")
	// Print out the AST
	fmt.Println(program.String())
}
