package evaluator

import (
	"testing"

	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
)

func benchEval(b *testing.B, input string) {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		b.Fatalf("parse errors: %v", p.Errors())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		env := object.NewEnvironment()
		ctx := &CallContext{
			FunctionName:      "<bench>",
			Node:              program,
			Parent:            nil,
			IsDirectExecution: true,
			env:               env,
		}
		Eval(program, env, ctx)
	}
}

func BenchmarkIntegerLoop(b *testing.B) {
	benchEval(b, `
i = 0
while i < 10000:
    i = i + 1
`)
}

func BenchmarkArithmetic(b *testing.B) {
	benchEval(b, `
x = 0
i = 0
while i < 5000:
    x = x + i * 2 - 1
    i = i + 1
`)
}

func BenchmarkFibonacci(b *testing.B) {
	benchEval(b, `
spell fib(n):
    if n < 2:
        return n
    return fib(n - 1) + fib(n - 2)
fib(20)
`)
}

func BenchmarkArrayBuild(b *testing.B) {
	benchEval(b, `
arr = []
i = 0
while i < 1000:
    arr = arr + [i]
    i = i + 1
`)
}

func BenchmarkStringConcat(b *testing.B) {
	benchEval(b, `
s = ""
i = 0
while i < 1000:
    s = s + "a"
    i = i + 1
`)
}
