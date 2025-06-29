package tests

import (
	"fmt"

	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/token"
)

func debug_tokens() {
	input := `spell test(param):
    x = param + 1
    return x

result = test(5)`

	l := lexer.New(input)

	fmt.Println("=== ALL TOKENS ===")
	tokenNum := 0
	for {
		tok := l.NextToken()
		fmt.Printf("%d: Token: %s, Literal: %q, Line: %d, Col: %d\n",
			tokenNum, tok.Type, tok.Literal, tok.Line, tok.Column)
		tokenNum++
		if tok.Type == token.EOF {
			break
		}
	}
}

