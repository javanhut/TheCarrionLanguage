package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	//Idenfiers and Literals
	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	//Operators
	ASSIGN   TokenType = "="
	PLUS     TokenType = "+"
	MINUS    TokenType = "-"
	ASTERICK TokenType = "*"
	SLASH    TokenType = "/"
	MOD      TokenType = "%"
	EQ       TokenType = "=="
	NOT_EQ   TokenType = "!="
	LT       TokenType = "<"
	GT       TokenType = ">"

	//Delimters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	COLON     TokenType = ":"
	PIPE      TokenType = "|"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"
	LBRACK TokenType = "["
	RBRACK TokenType = "]"

	//Keywords
	VAR    TokenType = "VAR"
	FLOAT  TokenType = "FLOAT"
	SPELLS TokenType = "SPELLS"
)
