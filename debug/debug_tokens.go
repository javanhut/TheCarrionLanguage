package main

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/token"
)

func main() {
	input := `spell test(param):
    return param + 1

result = test(5)`
	
	l := lexer.New(input)
	
	for {
		tok := l.NextToken()
		fmt.Printf("Token: %s, Literal: %q\n", tok.Type, tok.Literal)
		if tok.Type == token.EOF {
			break
		}
	}
}