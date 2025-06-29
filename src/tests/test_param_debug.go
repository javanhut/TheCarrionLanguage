package tests

import (
	"fmt"

	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
)

func debug_param() {
	input := `
spell test(param):
    x = param
    return param

print(test(5))
`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Printf("Parser errors: %v\n", p.Errors())
		return
	}

	env := object.NewEnvironment()
	result := evaluator.Eval(program, env, nil)

	if result != nil && result.Type() == object.ERROR_OBJ {
		fmt.Printf("Error: %s\n", result.Inspect())
	}
}

