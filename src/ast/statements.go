package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/javanhut/TheCarrionLanguage/src/token"
)

type AssignStatement struct {
	Token    token.Token
	Name     Expression
	Operator string
	TypeHint Expression
	Value    Expression
}

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

type OtherwiseBranch struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
}

type IfStatement struct {
	Token             token.Token
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

	for _, branch := range is.OtherwiseBranches {
		out.WriteString("otherwise ")
		out.WriteString(branch.Condition.String())
		out.WriteString(":\n")
		out.WriteString(branch.Consequence.String())
	}

	if is.Alternative != nil {
		out.WriteString("else:\n")
		out.WriteString(is.Alternative.String())
	}

	return out.String()
}

type ForStatement struct {
	Token       token.Token
	Variable    Expression // Now supports identifiers, tuple literals, etc.
	Iterable    Expression
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
	TypeHint     Expression
	DefaultValue Expression
}

func (p *Parameter) expressionNode()      {}
func (p *Parameter) TokenLiteral() string { return "Parameter" }

func (p *Parameter) String() string {
	if p.DefaultValue != nil {
		return fmt.Sprintf("%s=%s", p.Name.String(), p.DefaultValue.String())
	}
	return p.Name.String()
}

type FunctionDefinition struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Parameter
	Body       *BlockStatement
	DocString  *StringLiteral
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

type SpellbookDefinition struct {
	Token      token.Token
	Name       *Identifier
	Inherits   *Identifier
	Methods    []*FunctionDefinition
	InitMethod *FunctionDefinition
	DocString  *StringLiteral
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

type ImportStatement struct {
	Token     token.Token
	FilePath  *StringLiteral
	ClassName *Identifier
	Alias     *Identifier
}

func (is *ImportStatement) statementNode()       {}
func (is *ImportStatement) TokenLiteral() string { return is.Token.Literal }
func (is *ImportStatement) String() string {
	if is.ClassName != nil {
		if is.Alias != nil {
			return fmt.Sprintf(
				"import %s.%s as %s",
				is.FilePath.Value,
				is.ClassName.Value,
				is.Alias.Value,
			)
		}
		return fmt.Sprintf("import %s.%s", is.FilePath.Value, is.ClassName.Value)
	}
	return fmt.Sprintf("import %s", is.FilePath.Value)
}

type MatchStatement struct {
	Token      token.Token
	MatchValue Expression
	Cases      []*CaseClause
	Default    *CaseClause
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
	Token     token.Token
	Condition Expression
	Body      *BlockStatement
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

type AttemptStatement struct {
	Token          token.Token
	TryBlock       *BlockStatement
	EnsnareClauses []*EnsnareClause
	ResolveBlock   *BlockStatement
}

type EnsnareClause struct {
	Token       token.Token
	Condition   Expression
	Alias       *Identifier
	Consequence *BlockStatement
}

func (ec *EnsnareClause) statementNode()       {}
func (ec *EnsnareClause) TokenLiteral() string { return ec.Token.Literal }
func (ec *EnsnareClause) String() string {
	var out strings.Builder
	out.WriteString("ensnare")
	if ec.Condition != nil {
		out.WriteString(" (")
		out.WriteString(ec.Condition.String())
		out.WriteString(")")
	}
	if ec.Alias != nil {
		out.WriteString(" as ")
		out.WriteString(ec.Alias.Value)
	}
	out.WriteString(":\n")
	if ec.Consequence != nil {
		out.WriteString(ec.Consequence.String())
	}
	return out.String()
}

func (as *AttemptStatement) statementNode()       {}
func (as *AttemptStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AttemptStatement) String() string {
	var out strings.Builder

	out.WriteString("attempt:\n")
	if as.TryBlock != nil {
		out.WriteString(as.TryBlock.String())
	}

	for _, e := range as.EnsnareClauses {
		out.WriteString("ensnare")
		if e.Condition != nil {
			out.WriteString(" (")
			out.WriteString(e.Condition.String())
			out.WriteString(")")
		}
		out.WriteString(":\n")
		if e.Consequence != nil {
			out.WriteString(e.Consequence.String())
		}
	}

	if as.ResolveBlock != nil {
		out.WriteString("resolve:\n")
		out.WriteString(as.ResolveBlock.String())
	}

	return out.String()
}

type RaiseStatement struct {
	Token token.Token
	Error Expression
}

func (rs *RaiseStatement) statementNode()       {}
func (rs *RaiseStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *RaiseStatement) String() string {
	return fmt.Sprintf("raise %s", rs.Error.String())
}

type ArcaneSpell struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Parameter
	Body       *BlockStatement
}

func (as *ArcaneSpell) expressionNode()      {}
func (as *ArcaneSpell) TokenLiteral() string { return as.Token.Literal }
func (as *ArcaneSpell) String() string {
	var out bytes.Buffer
	out.WriteString("@arcanespell ")
	out.WriteString(as.Name.String())
	out.WriteString("(")
	params := []string{}
	for _, p := range as.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	return out.String()
}

type ArcaneSpellbook struct {
	Token      token.Token
	Name       *Identifier
	Methods    []*ArcaneSpell
	InitMethod *FunctionDefinition
}

func (asb *ArcaneSpellbook) statementNode()       {}
func (asb *ArcaneSpellbook) TokenLiteral() string { return asb.Token.Literal }
func (asb *ArcaneSpellbook) String() string {
	var out bytes.Buffer
	out.WriteString("arcane spellbook ")
	out.WriteString(asb.Name.String())
	out.WriteString(":\n")
	for _, method := range asb.Methods {
		out.WriteString("    ")
		out.WriteString(method.String())
		out.WriteString("\n")
	}
	return out.String()
}

type IgnoreStatement struct {
	Token token.Token
}

func (is *IgnoreStatement) statementNode()       {}
func (is *IgnoreStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IgnoreStatement) String() string       { return "ignore" }

type StopStatement struct {
	Token token.Token
}

func (ss *StopStatement) statementNode()       {}
func (ss *StopStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *StopStatement) String() string       { return "stop" }

type SkipStatement struct {
	Token token.Token
}

func (s *SkipStatement) statementNode()       {}
func (s *SkipStatement) TokenLiteral() string { return s.Token.Literal }
func (s *SkipStatement) String() string       { return "skip" }

type CheckStatement struct {
	Token     token.Token
	Condition Expression
	Message   Expression
}

func (cs *CheckStatement) statementNode()       {}
func (cs *CheckStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *CheckStatement) String() string {
	var out bytes.Buffer
	out.WriteString("check (")
	out.WriteString(cs.Condition.String())
	out.WriteString(")")
	if cs.Message != nil {
		out.WriteString(" : ")
		out.WriteString(cs.Message.String())
	}
	return out.String()
}
