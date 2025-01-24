package ast

import (
	"bytes"
	"fmt"
	"strings"
	"thecarrionlanguage/src/token"
)

// AssignStatement represents assignments like `x = 5` or `obj.field += 1`.
type AssignStatement struct {
	Token    token.Token // The token for the assignment operator (e.g. '=', '+=', '-=', etc.)
	Name     Expression  // The LHS of the assignment (Identifier, DotExpression, etc.)
	Operator string      // e.g. "=" or "+="
	Value    Expression  // The expression on the RHS
}

func (as *AssignStatement) statementNode()       {}
func (as *AssignStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssignStatement) String() string {
	return fmt.Sprintf("%s %s %s", as.Name.String(), as.Operator, as.Value.String())
}

// ReturnStatement represents a `return` statement.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out strings.Builder
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	return out.String()
}

// BlockStatement represents a block of statements enclosed by { } or newlines/indent in a block.
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out strings.Builder
	for _, s := range bs.Statements {
		out.WriteString(s.String())
		out.WriteString("\n")
	}
	return out.String()
}

// ExpressionStatement wraps an expression in a statement context.
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// OtherwiseBranch holds the condition and consequence for an 'otherwise' branch.
type OtherwiseBranch struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
}

// IfStatement represents an if/otherwise/.../else structure.
type IfStatement struct {
	Token             token.Token // The 'if' token
	Condition         Expression
	Consequence       *BlockStatement
	OtherwiseBranches []OtherwiseBranch
	Alternative       *BlockStatement
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }

// String() now prints "otherwise" instead of "elif".
func (is *IfStatement) String() string {
	var out strings.Builder

	// "if condition:"
	out.WriteString("if ")
	out.WriteString(is.Condition.String())
	out.WriteString(":\n")
	out.WriteString(is.Consequence.String())

	// Handle any 'otherwise' branch
	for _, branch := range is.OtherwiseBranches {
		out.WriteString("otherwise ")
		out.WriteString(branch.Condition.String())
		out.WriteString(":\n")
		out.WriteString(branch.Consequence.String())
	}

	// Handle an "else:" branch if present
	if is.Alternative != nil {
		out.WriteString("else:\n")
		out.WriteString(is.Alternative.String())
	}

	return out.String()
}

// ForStatement represents a for-loop: for x in iterable: ...
type ForStatement struct {
	Token       token.Token
	Variable    *Identifier // Loop variable
	Iterable    Expression  // Iterable expression (e.g., range())
	Body        *BlockStatement
	Alternative *BlockStatement
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	var out strings.Builder

	out.WriteString("for ")
	out.WriteString(fs.Variable.String())
	out.WriteString(" in ")
	out.WriteString(fs.Iterable.String())
	out.WriteString(":\n")
	out.WriteString(fs.Body.String())

	if fs.Alternative != nil {
		out.WriteString("else:\n")
		out.WriteString(fs.Alternative.String())
	}

	return out.String()
}

type Parameter struct {
	Name         *Identifier
	DefaultValue Expression
}

func (p *Parameter) String() string {
	if p.DefaultValue != nil {
		return fmt.Sprintf("%s=%s", p.Name.String(), p.DefaultValue.String())
	}
	return p.Name.String()
}

// FunctionDefinition represents a named function definition.
type FunctionDefinition struct {
	Token      token.Token // The 'SPELL' token
	Name       *Identifier
	Parameters []*Parameter
	Body       *BlockStatement
}

func (fd *FunctionDefinition) statementNode()       {}
func (fd *FunctionDefinition) TokenLiteral() string { return fd.Token.Literal }
func (fd *FunctionDefinition) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fd.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fd.TokenLiteral() + " ")
	out.WriteString(fd.Name.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString("):\n")
	out.WriteString(fd.Body.String())

	return out.String()
}

// WhileStatement represents a while-loop.
type WhileStatement struct {
	Token     token.Token
	Condition Expression
	Body      *BlockStatement
}

func (ws *WhileStatement) statementNode()       {}
func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }
func (ws *WhileStatement) String() string {
	var out strings.Builder
	out.WriteString("while ")
	out.WriteString(ws.Condition.String())
	out.WriteString(":\n")
	out.WriteString(ws.Body.String())
	return out.String()
}

// SpellbookDefinition represents a grouping of methods in a named 'spellbook'.
type SpellbookDefinition struct {
	Token      token.Token
	Name       *Identifier
	Methods    []*FunctionDefinition
	InitMethod *FunctionDefinition
}

func (sb *SpellbookDefinition) statementNode()       {}
func (sb *SpellbookDefinition) TokenLiteral() string { return sb.Token.Literal }
func (sb *SpellbookDefinition) String() string {
	var out bytes.Buffer
	out.WriteString("spellbook ")
	out.WriteString(sb.Name.String())
	out.WriteString(":\n")

	if sb.InitMethod != nil {
		out.WriteString("    ")
		out.WriteString(sb.InitMethod.String())
		out.WriteString("\n")
	}
	for _, method := range sb.Methods {
		out.WriteString("    ")
		out.WriteString(method.String())
		out.WriteString("\n")
	}
	return out.String()
}

// ImportStatement represents an import statement.
type ImportStatement struct {
	Token     token.Token
	FilePath  *StringLiteral
	ClassName *Identifier
}

func (is *ImportStatement) statementNode()       {}
func (is *ImportStatement) TokenLiteral() string { return is.Token.Literal }
func (is *ImportStatement) String() string {
	if is.ClassName != nil {
		return fmt.Sprintf("import %s.%s", is.FilePath.Value, is.ClassName.String())
	}
	return fmt.Sprintf("import %s", is.FilePath.Value)
}

type MatchStatement struct {
	Token      token.Token   // The "match" token
	MatchValue Expression    // The value being matched
	Cases      []*CaseClause // List of case clauses
	Default    *CaseClause   // Default case (optional)
}

func (ms *MatchStatement) statementNode()       {}
func (ms *MatchStatement) TokenLiteral() string { return ms.Token.Literal }
func (ms *MatchStatement) String() string {
	var out bytes.Buffer
	out.WriteString("match ")
	out.WriteString(ms.MatchValue.String())
	out.WriteString(":\n")
	for _, c := range ms.Cases {
		out.WriteString(c.String())
	}
	if ms.Default != nil {
		out.WriteString(ms.Default.String())
	}
	return out.String()
}

type CaseClause struct {
	Token     token.Token     // The "case" or "_" token
	Condition Expression      // The case condition
	Body      *BlockStatement // The body of the case
}

func (cc *CaseClause) statementNode()       {}
func (cc *CaseClause) TokenLiteral() string { return cc.Token.Literal }
func (cc *CaseClause) String() string {
	var out bytes.Buffer
	out.WriteString("case ")
	out.WriteString(cc.Condition.String())
	out.WriteString(":\n")
	out.WriteString(cc.Body.String())
	return out.String()
}
