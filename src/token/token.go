// token/token.go
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"
	NEWLINE TokenType = "NEWLINE"
	INDENT  TokenType = "INDENT"
	DEDENT  TokenType = "DEDENT"

	// Identifiers and Literals
	IDENT     TokenType = "IDENT"
	INT       TokenType = "INT"
	FLOAT     TokenType = "FLOAT"
	STRING    TokenType = "STRING"
	DOCSTRING TokenType = "DOCSTRING"

	// Operators
	ASSIGN          TokenType = "="
	PLUS            TokenType = "+"
	MINUS           TokenType = "-"
	ASTERISK        TokenType = "*"
	SLASH           TokenType = "/"
	MOD             TokenType = "%"
	EXPONENT        TokenType = "**"
	INCREMENT       TokenType = "+="
	DECREMENT       TokenType = "-="
	MULTASSGN       TokenType = "*="
	DIVASSGN        TokenType = "/="
	PLUS_INCREMENT  TokenType = "++"
	MINUS_DECREMENT TokenType = "--"
	EQ              TokenType = "=="
	NOT_EQ          TokenType = "!="
	LT              TokenType = "<"
	GT              TokenType = ">"
	LE              TokenType = "<="
	GE              TokenType = ">="
	BANG            TokenType = "!"
	AMPERSAND       TokenType = "&"
	HASH            TokenType = "#"
	AT              TokenType = "@"
	// DUNDER          TokenType = "__"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	COLON     TokenType = ":"
	PIPE      TokenType = "|"
	DOT       TokenType = "."
	LSHIFT    TokenType = "<<"
	RSHIFT    TokenType = ">>"
	XOR       TokenType = "^"
	TILDE     TokenType = "~"

	LPAREN     TokenType = "("
	RPAREN     TokenType = ")"
	LBRACE     TokenType = "{"
	RBRACE     TokenType = "}"
	LBRACK     TokenType = "["
	RBRACK     TokenType = "]"
	UNDERSCORE TokenType = "_"

	// Keywords
	VAR         TokenType = "VAR"
	INIT        TokenType = "INIT"
	SELF        TokenType = "SELF"
	SPELL       TokenType = "SPELL"
	SPELLBOOK   TokenType = "SPELLBOOK"
	TRUE        TokenType = "TRUE"
	FALSE       TokenType = "FALSE"
	IF          TokenType = "IF"
	OTHERWISE   TokenType = "OTHERWISE"
	ELSE        TokenType = "ELSE"
	FOR         TokenType = "FOR"
	IN          TokenType = "IN"
	WHILE       TokenType = "WHILE"
	STOP        TokenType = "STOP"
	SKIP        TokenType = "SKIP"
	IGNORE      TokenType = "IGNORE"
	RETURN      TokenType = "RETURN"
	IMPORT      TokenType = "IMPORT"
	MATCH       TokenType = "MATCH"
	CASE        TokenType = "CASE"
	ATTEMPT     TokenType = "ATTEMPT"
	RESOLVE     TokenType = "RESOLVE"
	ENSNARE     TokenType = "ENSNARE"
	RAISE       TokenType = "RAISE"
	AS          TokenType = "AS"
	ARCANE      TokenType = "ARCANE"
	ARCANESPELL TokenType = "ARCANESPELL"
	SUPER       TokenType = "SUPER"
	FSTRING     TokenType = "FSTRING"
	CHECK       TokenType = "CHECK"
	NONE        TokenType = "NONE"
	AND         TokenType = "AND"
	OR          TokenType = "OR"
	NOT         TokenType = "NOT"
)

var keywords = map[string]TokenType{
	"import":      IMPORT,
	"match":       MATCH,
	"case":        CASE,
	"var":         VAR,
	"spell":       SPELL,
	"self":        SELF,
	"init":        INIT,
	"spellbook":   SPELLBOOK,
	"True":        TRUE,
	"False":       FALSE,
	"if":          IF,
	"otherwise":   OTHERWISE,
	"else":        ELSE,
	"for":         FOR,
	"in":          IN,
	"while":       WHILE,
	"stop":        STOP,
	"skip":        SKIP,
	"ignore":      IGNORE,
	"and":         AND,
	"or":          OR,
	"not":         NOT,
	"return":      RETURN,
	"attempt":     ATTEMPT,
	"resolve":     RESOLVE,
	"ensnare":     ENSNARE,
	"raise":       RAISE,
	"as":          AS,
	"arcane":      ARCANE,
	"arcanespell": ARCANESPELL,
	"super":       SUPER,
	"check":       CHECK,

	//"range":     RANGE,
	"None": NONE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// LookupIndent determines the TokenType based on the indentation string.
func LookupIndent(indent string) TokenType {
	indentLevels := map[int]TokenType{
		0: DEDENT,
		4: INDENT, // 4 spaces
		8: INDENT, // 8 spaces, etc.
		// Add more levels as needed
	}

	length := len(indent)
	if tok, ok := indentLevels[length]; ok {
		return tok
	}
	return ILLEGAL
}
