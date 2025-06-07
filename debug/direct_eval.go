package main

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/object"
)

func main() {
	input := `spell test(param):
    x = param + 1
    return x

result = test(5)`
	
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	
	if len(p.Errors()) > 0 {
		fmt.Printf("Parser errors: %v\n", p.Errors())
		return
	}
	
	env := object.NewEnvironment()
	result := evaluator.Eval(program, env, nil)
	fmt.Printf("Result: %s\n", result.Inspect())
	
	// Check if the function worked
	if val, ok := env.Get("result"); ok {
		fmt.Printf("result variable: %s\n", val.Inspect())
	} else {
		fmt.Println("result variable not found")
	}
}