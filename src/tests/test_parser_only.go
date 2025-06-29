package main

import (
	"fmt"
	"os"

	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run test_parser_only.go <file.crl>")
		os.Exit(1)
	}

	filename := os.Args[1]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	l := lexer.NewWithFilename(string(content), filename)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Printf("Parse errors in %s:\n", filename)
		for i, err := range p.Errors() {
			fmt.Printf("  Error %d: %s\n", i+1, err)
		}
		os.Exit(1)
	}

	fmt.Printf("Successfully parsed %s with %d statements\n", filename, len(program.Statements))
}