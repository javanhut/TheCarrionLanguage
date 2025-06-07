package main

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/token"
)

func main() {
	input := `spell test(a, b):
    x = a + b
    return x`
	
	l := lexer.New(input)
	
	fmt.Println("=== TOKENS ===")
	for i := 0; i < 30; i++ {
		tok := l.NextToken()
		fmt.Printf("%d: %s '%s' line=%d col=%d\n", 
			i, tok.Type, tok.Literal, tok.Line, tok.Column)
		if tok.Type == token.EOF {
			break
		}
	}
}