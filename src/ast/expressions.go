package ast

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/javanhut/TheCarrionLanguage/src/token"
)

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return strconv.FormatInt(il.Value, 10) }

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) String() string       { return strconv.FormatFloat(fl.Value, 'f', -1, 64) }

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", pe.Operator, pe.Right.String())
}

type InfixExpression struct {
	Token    token.Token
	Operator string
	Left     Expression
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", ie.Left.String(), ie.Operator, ie.Right.String())
}

type PostfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
}

func (pe *PostfixExpression) expressionNode()      {}
func (pe *PostfixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PostfixExpression) String() string {
	return fmt.Sprintf("(%s%s)", pe.Left.String(), pe.Operator)
}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string {
	if b.Value {
		return "true"
	}
	return "false"
}
func (b *Boolean) String() string { return b.TokenLiteral() }

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elems := make([]string, 0, len(al.Elements))
	for _, e := range al.Elements {
		elems = append(elems, e.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("]")
	return out.String()
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}

type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

type TupleLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (tl *TupleLiteral) expressionNode()      {}
func (tl *TupleLiteral) TokenLiteral() string { return tl.Token.Literal }
func (tl *TupleLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range tl.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("(")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString(")")
	return out.String()
}

type DotExpression struct {
	Token token.Token
	Left  Expression
	Right *Identifier
}

func (de *DotExpression) expressionNode()      {}
func (de *DotExpression) TokenLiteral() string { return de.Token.Literal }
func (de *DotExpression) String() string {
	return fmt.Sprintf("(%s.%s)", de.Left.String(), de.Right.String())
}

type NoneLiteral struct {
	Token token.Token
}

func (nl *NoneLiteral) expressionNode()      {}
func (nl *NoneLiteral) TokenLiteral() string { return nl.Token.Literal }
func (nl *NoneLiteral) String() string       { return "None" }

type FStringLiteral struct {
	Token token.Token
	Parts []FStringPart
}

type FStringPart interface {
	partNode()
	String() string
}

type FStringText struct {
	Value string
}

func (ft *FStringText) partNode()      {}
func (ft *FStringText) String() string { return ft.Value }

type FStringExpr struct {
	Expr Expression
}

func (fe *FStringExpr) partNode() {}
func (fe *FStringExpr) String() string {
	if fe.Expr != nil {
		return fe.Expr.String()
	}
	return ""
}

func (fsl *FStringLiteral) expressionNode() {}

func (fsl *FStringLiteral) TokenLiteral() string {
	return fsl.Token.Literal
}

func (fsl *FStringLiteral) String() string {
	var out bytes.Buffer
	for _, part := range fsl.Parts {
		out.WriteString(part.String())
	}
	return out.String()
}

type StringInterpolation struct {
	Token token.Token
	Parts []StringPart
}
type StringPart interface {
	partNode()
	String() string
}

type StringText struct {
	Value string
}

func (st *StringText) partNode()      {}
func (st *StringText) String() string { return st.Value }

type StringExpr struct {
	Expr       Expression
	FormatSpec string
	Width      int
	Precision  int
	Alignment  byte
	FillChar   byte
}

func (se *StringExpr) partNode() {}
func (se *StringExpr) String() string {
	if se.Expr != nil {
		return se.Expr.String()
	}
	return ""
}

func (si *StringInterpolation) expressionNode() {}

func (si *StringInterpolation) TokenLiteral() string {
	return si.Token.Literal
}

func (si *StringInterpolation) String() string {
	var out bytes.Buffer
	out.WriteString("i\"")
	for _, part := range si.Parts {
		out.WriteString(part.String())
	}
	out.WriteString("\"")
	return out.String()
}
