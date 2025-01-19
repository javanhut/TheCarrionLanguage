package ast

import (
	"bytes"
	"fmt"
	"strings"

	"thecarrionlanguage/token"
)

type AssignStatement struct {
	Token    token.Token // The token for the assignment operator (e.g. '=' or '+=')
	Name     Expression  // The LHS of the assignment (Identifier, DotExpression, etc.)
	Operator string      // e.g. "=" or "+="
	Value    Expression  // The expression on the RHS
}

// Ensure AssignStatement satisfies Statement interface:
func (as *AssignStatement) statementNode()       {}
func (as *AssignStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssignStatement) String() string {
	return fmt.Sprintf("%s %s %s", as.Name.String(), as.Operator, as.Value.String())
}

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

type IfStatement struct {
	Token             token.Token // The 'if' token
	Condition         Expression
	Consequence       *BlockStatement
	OtherwiseBranches []OtherwiseBranch
	Alternative       *BlockStatement
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IfStatement) String() string {
	var out strings.Builder

	out.WriteString("if ")
	out.WriteString(is.Condition.String())
	out.WriteString(":\n")
	out.WriteString(is.Consequence.String())

	if is.Alternative != nil {
		out.WriteString("else:\n")
		out.WriteString(is.Alternative.String())
	}

	return out.String()
}

type OtherwiseBranch struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
}

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

type FunctionDefinition struct {
	Token      token.Token // The 'SPELL' token
	Name       *Identifier
	Parameters []*Identifier
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

type WhileStatement struct {
	Token     token.Token // The 'while' token
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

type SpellbookDefinition struct {
	Token      token.Token           // The 'spellbook' token
	Name       *Identifier           // Spellbook name
	Methods    []*FunctionDefinition // List of methods (spells)
	InitMethod *FunctionDefinition   // Optional `init` method
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
