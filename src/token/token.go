// token/token.go
package token

type TokenType string

type Token struct {
	Type     TokenType
	Literal  string
	Filename string
	Line     int
	Column   int
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
	INTDIV          TokenType = "//"
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
	GRIMOIRE    TokenType = "GRIM"
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
	INTERP      TokenType = "INTERP"
	CHECK       TokenType = "CHECK"
	NONE        TokenType = "NONE"
	AND         TokenType = "AND"
	OR          TokenType = "OR"
	NOT         TokenType = "NOT"
	NOT_IN      TokenType = "NOT_IN"
	MAIN        TokenType = "MAIN"
	GLOBAL      TokenType = "GLOBAL"
)

var keywords = map[string]TokenType{
	"import":      IMPORT,
	"match":       MATCH,
	"case":        CASE,
	"var":         VAR,
	"spell":       SPELL,
	"self":        SELF,
	"init":        INIT,
	"grim":        GRIMOIRE,
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
	"not in":      NOT_IN,
	"main":        MAIN,
	"global":      GLOBAL,

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

func NewToken(tokenType TokenType, literal string, filename string, line int, column int) Token {
	return Token{
		Type:     tokenType,
		Literal:  literal,
		Filename: filename,
		Line:     line,
		Column:   column,
	}
}

// For compatibility with existing code that creates tokens simply
func SimpleToken(tokenType TokenType, ch byte) Token {
	return Token{
		Type:     tokenType,
		Literal:  string(ch),
		Filename: "",
		Line:     0,
		Column:   0,
	}
}
