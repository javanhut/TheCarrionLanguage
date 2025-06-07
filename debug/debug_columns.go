package main

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/token"
)

func main() {
	input := `spell test(param):
    print("hello")
    return param`
	
	l := lexer.New(input)
	
	for i := 0; i < 20; i++ {
		tok := l.NextToken()
		fmt.Printf("%d: %s '%s' at line %d, col %d\n", 
			i, tok.Type, tok.Literal, tok.Line, tok.Column)
		if tok.Type == token.EOF {
			break
		}
	}
}