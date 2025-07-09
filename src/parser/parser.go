package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/javanhut/TheCarrionLanguage/src/ast"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/token"
)

const (
	_      int = iota
	LOWEST int = iota
	ASSIGN
	LOGICAL_OR
	LOGICAL_AND
	EQUALS
	COMPARSION
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
	POSTFIX
	INDEX
)

var precedences = map[token.TokenType]int{
	token.ASSIGN:          ASSIGN,
	token.INCREMENT:       ASSIGN,
	token.DECREMENT:       ASSIGN,
	token.MULTASSGN:       ASSIGN,
	token.DIVASSGN:        ASSIGN,
	token.EQ:              EQUALS,
	token.NOT_EQ:          EQUALS,
	token.LT:              LESSGREATER,
	token.GT:              LESSGREATER,
	token.LE:              LESSGREATER,
	token.GE:              LESSGREATER,
	token.PLUS:            SUM,
	token.MINUS:           SUM,
	token.SLASH:           PRODUCT,
	token.INTDIV:          PRODUCT,
	token.ASTERISK:        PRODUCT,
	token.MOD:             PRODUCT,
	token.EXPONENT:        PRODUCT,
	token.PLUS_INCREMENT:  POSTFIX,
	token.MINUS_DECREMENT: POSTFIX,
	token.LPAREN:          CALL,
	token.DOT:             CALL,
	token.LBRACK:          INDEX,
	token.OR:              LOGICAL_OR,
	token.AND:             LOGICAL_AND,
	token.IN:              COMPARSION,
	token.NOT_IN:          COMPARSION,

	token.LSHIFT:    6,
	token.RSHIFT:    6,
	token.AMPERSAND: 5,
	token.XOR:       4,
	token.PIPE:      3,
}

type (
	prefixParseFn  func() ast.Expression
	infixParseFn   func(ast.Expression) ast.Expression
	postfixParseFn func(ast.Expression) ast.Expression
)

type Parser struct {
	l                 *lexer.Lexer
	currToken         token.Token
	peekToken         token.Token
	errors            []string
	contextStack      []string
	parsingParameters bool
	inTypeHintContext bool
	prefixParseFns    map[token.TokenType]prefixParseFn
	infixParseFns     map[token.TokenType]infixParseFn
	postfixParseFns   map[token.TokenType]postfixParseFn
	statementParseFns map[token.TokenType]func() ast.Statement
	controlStack      []struct {
		Type        string
		IndentLevel int
		HasElse     bool
		Token       token.Token
	}
}

func (p *Parser) isInsideGrimoire() bool {
	if len(p.contextStack) == 0 {
		return false
	}
	return p.contextStack[len(p.contextStack)-1] == "grim"
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l, errors: []string{},
		controlStack: []struct {
			Type        string
			IndentLevel int
			HasElse     bool
			Token       token.Token
		}{},
	}
	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.postfixParseFns = make(map[token.TokenType]postfixParseFn)
	p.statementParseFns = make(map[token.TokenType]func() ast.Statement)

	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.NOT, p.parsePrefixExpression)
	p.registerPrefix(token.PLUS_INCREMENT, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS_DECREMENT, p.parsePrefixExpression)
	p.registerPrefix(token.CASE, func() ast.Expression { return nil })
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.COLON, func() ast.Expression { return nil })
	p.registerPrefix(token.TILDE, p.parsePrefixExpression)
	p.registerPrefix(token.UNDERSCORE, func() ast.Expression { return nil })
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.LBRACK, p.parseArrayLiteral)
	p.registerPrefix(token.LBRACE, p.parseHashLiteral)
	p.registerPrefix(token.NONE, p.parseNoneLiteral)
	p.registerPrefix(token.NEWLINE, func() ast.Expression { return nil })
	p.registerPrefix(token.INDENT, func() ast.Expression { return nil })
	p.registerPrefix(token.DEDENT, func() ast.Expression { return nil })
	p.registerPrefix(token.EOF, func() ast.Expression { return nil })
	// OTHERWISE is handled as part of if-statement parsing, not as a prefix expression
	p.registerPrefix(token.ENSNARE, func() ast.Expression { return nil })
	p.registerPrefix(token.AS, func() ast.Expression { return nil })
	p.registerPrefix(token.AT, func() ast.Expression { return nil })
	p.registerPrefix(token.ARCANESPELL, func() ast.Expression { return nil })
	p.registerPrefix(token.LPAREN, p.parseParenExpression)
	p.registerPrefix(token.SELF, p.parseSelf)
	p.registerPrefix(token.SUPER, p.parseSuperExpression)
	p.registerPrefix(token.FSTRING, p.parseFStringLiteral)
	p.registerPrefix(token.INTERP, p.parseStringInterpolationLiteral)
	p.registerPrefix(token.DOCSTRING, p.parseDocStringLiteral)
	p.registerPrefix(token.INIT, func() ast.Expression {
		return &ast.Identifier{
			Token: token.Token{Type: token.INIT, Literal: "init"},
			Value: "init",
		}
	})

	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.INTDIV, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.MOD, p.parseInfixExpression)
	p.registerInfix(token.EXPONENT, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.LE, p.parseInfixExpression)
	p.registerInfix(token.GE, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.OR, p.parseInfixExpression)
	p.registerInfix(token.AND, p.parseInfixExpression)
	p.registerInfix(token.INCREMENT, p.parseInfixExpression)
	p.registerInfix(token.DECREMENT, p.parseInfixExpression)
	p.registerInfix(token.MULTASSGN, p.parseInfixExpression)
	p.registerInfix(token.DIVASSGN, p.parseInfixExpression)
	p.registerInfix(token.LBRACK, p.parseIndexExpression)
	p.registerInfix(token.DOT, p.parseDotExpression)
	p.registerInfix(token.COMMA, p.parseCommaExpression)
	p.registerInfix(token.LSHIFT, p.parseInfixExpression)
	p.registerInfix(token.RSHIFT, p.parseInfixExpression)
	p.registerInfix(token.AMPERSAND, p.parseInfixExpression)
	p.registerInfix(token.XOR, p.parseInfixExpression)
	p.registerInfix(token.PIPE, p.parseInfixExpression)
	p.registerInfix(token.IN, p.parseInfixExpression)
	p.registerInfix(token.NOT_IN, p.parseInfixExpression)

	p.registerPostfix(token.PLUS_INCREMENT, p.parsePostfixExpression)
	p.registerPostfix(token.MINUS_DECREMENT, p.parsePostfixExpression)

	p.registerStatement(token.RETURN, p.parseReturnStatement)
	p.registerStatement(token.IF, p.parseIfStatement)
	p.registerStatement(token.FOR, p.parseForStatement)
	p.registerStatement(token.SPELL, p.parseFunctionDefinition)
	p.registerStatement(token.IMPORT, p.parseImportStatement)
	p.registerStatement(token.MATCH, p.parseMatchStatement)
	p.registerStatement(token.ATTEMPT, p.parseAttemptStatement)
	p.registerStatement(token.RESOLVE, p.parseResolveStatement)
	p.registerStatement(token.ENSNARE, p.parseEnsnareStatement)
	p.registerStatement(token.RAISE, p.parseRaiseStatement)
	p.registerStatement(token.ARCANE, p.parseArcaneGrimoire)
	p.registerStatement(token.IGNORE, p.parseIgnoreStatement)
	p.registerStatement(token.STOP, p.parseStopStatement)
	p.registerStatement(token.SKIP, p.parseSkipStatement)
	p.registerStatement(token.CHECK, p.parseCheckStatement)
	p.registerStatement(token.GLOBAL, p.parseGlobalStatement)
	p.registerStatement(token.AUTOCLOSE, p.parseWithStatement)

	return p
}

func tokenFromExpression(expr ast.Expression) token.Token {
	switch node := expr.(type) {
	case *ast.Identifier:
		return node.Token
	case *ast.IntegerLiteral:
		return node.Token
	case *ast.StringLiteral:
		return node.Token
	case *ast.FloatLiteral:
		return node.Token
	case *ast.DotExpression:
		return node.Token
	default:
		return token.Token{Type: token.ILLEGAL, Literal: ""}
	}
}

func (p *Parser) parseExpressionTuple(firstExpr ast.Expression, isLHS bool) ast.Expression {
	if firstExpr == nil {
		if isLHS {
			p.errors = append(p.errors, "expected assignable expression")
		} else {
			p.errors = append(p.errors, "expected expression")
		}
		return nil
	}

	if !p.peekTokenIs(token.COMMA) {
		return firstExpr
	}

	var expressions []ast.Expression
	expressions = append(expressions, firstExpr)

	errorPrefix := "expected expression after comma in assignment"
	if isLHS {
		errorPrefix = "expected assignable expression after comma in assignment"
	}

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		for p.currToken.Type == token.COMMA {
			p.nextToken()
		}

		if p.isAtEnd() || !p.canStartExpression(p.currToken.Type) {
			p.errors = append(p.errors, "unexpected trailing comma in assignment")
			return nil
		}

		nextExpr := p.parseExpression(LOWEST)
		if nextExpr == nil {
			p.errors = append(p.errors, errorPrefix)
			return nil
		}
		expressions = append(expressions, nextExpr)
	}

	if len(expressions) == 1 && !isLHS {
		return expressions[0]
	}

	return &ast.TupleLiteral{
		Token:    tokenFromExpression(expressions[0]),
		Elements: expressions,
	}
}

func (p *Parser) parseAssignmentLHS() ast.Expression {
	expr := p.parseExpression(LOWEST)
	return p.parseExpressionTuple(expr, true)
}

func (p *Parser) parseAssignmentRHS() ast.Expression {
	exp := p.parseExpression(LOWEST)
	return p.parseExpressionTuple(exp, false)
}

func (p *Parser) isAtEnd() bool {
	return p.currToken.Type == token.EOF
}

func (p *Parser) canStartExpression(tokenType token.TokenType) bool {
	switch tokenType {
	case token.IDENT, token.INT, token.FLOAT, token.STRING, token.TRUE, token.FALSE, token.NONE,
		token.LPAREN, token.LBRACK, token.LBRACE, token.SPELL, token.IF, token.SELF:
		return true
	default:
		return false
	}
}

func (p *Parser) parseCommaExpression(left ast.Expression) ast.Expression {
	var elements []ast.Expression
	if tuple, ok := left.(*ast.TupleLiteral); ok {
		elements = tuple.Elements
	} else {
		elements = append(elements, left)
	}

	p.nextToken()

	right := p.parseExpression(LOWEST)
	if right != nil {
		elements = append(elements, right)
	}

	return &ast.TupleLiteral{
		Token:    p.currToken,
		Elements: elements,
	}
}

func (p *Parser) parseStopStatement() ast.Statement {
	return &ast.StopStatement{Token: p.currToken}
}

func (p *Parser) parseSkipStatement() ast.Statement {
	return &ast.SkipStatement{Token: p.currToken}
}

func (p *Parser) parseCheckStatement() ast.Statement {
	stmt := &ast.CheckStatement{Token: p.currToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)
	if stmt.Condition == nil {
		p.errors = append(p.errors, "expected condition expression in check statement")
		return nil
	}

	if p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		stmt.Message = p.parseExpression(LOWEST)

		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	} else {
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	}

	return stmt
}

func (p *Parser) parseDocStringLiteral() ast.Expression {
	return &ast.StringLiteral{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}
}

func (p *Parser) parseFStringLiteral() ast.Expression {
	raw := p.currToken.Literal

	fslit := &ast.FStringLiteral{
		Token: p.currToken,
		Parts: []ast.FStringPart{},
	}

	var builder strings.Builder
	i := 0
	for i < len(raw) {
		ch := raw[i]

		if ch == '{' {
			if builder.Len() > 0 {
				fslit.Parts = append(fslit.Parts, &ast.FStringText{Value: builder.String()})
				builder.Reset()
			}

			end := findMatchingBrace(raw, i+1)
			if end < 0 {
				p.errors = append(p.errors, "Unclosed brace in f-string")
				return fslit
			}
			exprStr := raw[i+1 : end]

			expr := p.parseFStringExpression(exprStr)
			fslit.Parts = append(fslit.Parts, &ast.FStringExpr{Expr: expr})

			i = end + 1
		} else {
			builder.WriteByte(ch)
			i++
		}
	}

	if builder.Len() > 0 {
		fslit.Parts = append(fslit.Parts, &ast.FStringText{Value: builder.String()})
	}

	return fslit
}

func findMatchingBrace(s string, start int) int {
	for i := start; i < len(s); i++ {
		if s[i] == '}' {
			return i
		}
	}
	return -1
}

func (p *Parser) parseFStringExpression(exprStr string) ast.Expression {
	l := lexer.New(exprStr)
	subParser := New(l)

	program := subParser.ParseProgram()

	if len(subParser.Errors()) > 0 {
		p.errors = append(p.errors, "error parsing f-string expression: "+subParser.Errors()[0])
		return nil
	}

	if len(program.Statements) == 1 {
		if es, ok := program.Statements[0].(*ast.ExpressionStatement); ok {
			return es.Expression
		}
	}

	p.errors = append(p.errors, "invalid f-string expression")
	return nil
}

func (p *Parser) parseSuperExpression() ast.Expression {
	return &ast.Identifier{
		Token: p.currToken,
		Value: "super",
	}
}

func (p *Parser) parseIgnoreStatement() ast.Statement {
	return &ast.IgnoreStatement{Token: p.currToken}
}

func (p *Parser) parseArcaneGrimoire() ast.Statement {
	stmt := &ast.ArcaneGrimoire{Token: p.currToken}

	if !p.expectPeek(token.GRIMOIRE) {
		return nil
	}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeek(token.COLON) {
		return nil
	}

	p.skipNewlines()
	if p.peekTokenIs(token.INDENT) {
		p.nextToken()
	}

	stmt.Methods = []*ast.ArcaneSpell{}

	for p.peekTokenIs(token.AT) {
		method := p.parseArcaneMethod()
		if method != nil {
			stmt.Methods = append(stmt.Methods, method)
		}

		p.skipNewlines()
	}

	return stmt
}

func (p *Parser) parseArcaneMethod() *ast.ArcaneSpell {
	p.nextToken()
	if !p.expectPeek(token.ARCANESPELL) {
		p.errors = append(p.errors, "expected 'arcanespell' after '@'")
		return nil
	}
	if p.currToken.Literal != "arcanespell" {
		p.errors = append(p.errors, "expected 'arcanespell' after '@', got "+p.currToken.Literal)
		return nil
	}

	for p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
	}

	if !p.peekTokenIs(token.SPELL) && !p.peekTokenIs(token.INIT) {
		p.errors = append(p.errors, "expected 'spell' or 'init' after '@arcanespell'")
		return nil
	}
	p.nextToken()

	arcMethod := &ast.ArcaneSpell{Token: p.currToken}
	if p.currToken.Type == token.SPELL {
		if !p.expectPeek(token.IDENT) {
			p.errors = append(p.errors, "expected method name after 'spell'")
			return nil
		}
		arcMethod.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	} else if p.currToken.Type == token.INIT {
		arcMethod.Name = &ast.Identifier{Token: p.currToken, Value: "init"}
	}

	if p.peekTokenIs(token.LPAREN) {
		p.nextToken()
		arcMethod.Parameters = p.parseFunctionParameters()
	}

	if p.peekTokenIs(token.COLON) {
		p.nextToken()
	}

	return arcMethod
}

func (p *Parser) parseRaiseStatement() ast.Statement {
	stmt := &ast.RaiseStatement{Token: p.currToken}

	p.nextToken()
	stmt.Error = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.NEWLINE) || p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseAttemptStatement parses an attempt block with ensnare and resolve clauses
func (p *Parser) parseAttemptStatement() ast.Statement {
	stmt := &ast.AttemptStatement{Token: p.currToken}

	currentIndent := p.getCurrentIndent()
	p.controlStack = append(p.controlStack, struct {
		Type        string
		IndentLevel int
		HasElse     bool
		Token       token.Token
	}{
		Type:        "attempt",
		IndentLevel: currentIndent,
		HasElse:     false,
		Token:       p.currToken,
	})

	// Expect ':' after 'attempt'
	if !p.expectPeek(token.COLON) {
		return nil
	}

	// Parse the try block body
	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
		if p.peekTokenIs(token.INDENT) {
			p.nextToken() // Move to INDENT token
			p.nextToken() // Move past INDENT to first statement token
			stmt.TryBlock = p.parseBlockStatement()
		} else {
			stmt.TryBlock = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}
	} else {
		p.nextToken()
		stmt.TryBlock = &ast.BlockStatement{
			Token:      p.currToken,
			Statements: []ast.Statement{p.parseStatement()},
		}
	}

	// Parse ensnare clauses
	stmt.EnsnareClauses = []*ast.EnsnareClause{}
	for p.peekTokenIs(token.ENSNARE) {
		p.nextToken() // ENSNARE
		clause := &ast.EnsnareClause{Token: p.currToken}

		// Optional condition or alias in parentheses
		if p.peekTokenIs(token.LPAREN) {
			p.nextToken() // consume '('
			p.nextToken() // move to the content

			// Check if it's a simple identifier (alias) or an expression (condition)
			if p.currTokenIs(token.IDENT) && p.peekTokenIs(token.RPAREN) {
				// It's an alias like ensnare (e):
				clause.Alias = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
			} else {
				// It's a condition expression
				clause.Condition = p.parseExpression(LOWEST)
			}

			if !p.expectPeek(token.RPAREN) {
				return nil
			}
		}

		// Expect ':'
		if !p.expectPeek(token.COLON) {
			return nil
		}

		// Parse ensnare body
		if p.peekTokenIs(token.NEWLINE) {
			p.nextToken()
			if p.peekTokenIs(token.INDENT) {
				p.nextToken() // Move to INDENT token
				p.nextToken() // Move past INDENT to first statement token
				clause.Consequence = p.parseBlockStatement()
			} else {
				clause.Consequence = &ast.BlockStatement{
					Token:      p.currToken,
					Statements: []ast.Statement{p.parseStatement()},
				}
			}
		} else {
			p.nextToken()
			clause.Consequence = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}

		stmt.EnsnareClauses = append(stmt.EnsnareClauses, clause)
	}

	// Parse resolve clause
	if p.peekTokenIs(token.RESOLVE) {
		p.nextToken() // RESOLVE
		if !p.expectPeek(token.COLON) {
			return nil
		}

		// Parse resolve body
		if p.peekTokenIs(token.NEWLINE) {
			p.nextToken()
			if p.peekTokenIs(token.INDENT) {
				p.nextToken() // Move to INDENT token
				p.nextToken() // Move past INDENT to first statement token
				stmt.ResolveBlock = p.parseBlockStatement()
			} else {
				stmt.ResolveBlock = &ast.BlockStatement{
					Token:      p.currToken,
					Statements: []ast.Statement{p.parseStatement()},
				}
			}
		} else {
			p.nextToken()
			stmt.ResolveBlock = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}
	}

	// Clean up control stack
	if len(p.controlStack) > 0 {
		p.controlStack = p.controlStack[:len(p.controlStack)-1]
	}

	return stmt
}

func (p *Parser) parseEnsnareStatement() ast.Statement {
	p.errors = append(p.errors, "Unexpected 'ensnare' outside of 'attempt' block")
	return nil
}

func (p *Parser) parseResolveStatement() ast.Statement {
	p.errors = append(p.errors, "Unexpected 'resolve' outside of 'attempt' block")
	return nil
}

func (p *Parser) parseNoneLiteral() ast.Expression {
	return &ast.NoneLiteral{Token: p.currToken}
}

func (p *Parser) parseMatchStatement() ast.Statement {
	currentIndent := p.getCurrentIndent()
	p.controlStack = append(p.controlStack, struct {
		Type        string
		IndentLevel int
		HasElse     bool
		Token       token.Token
	}{
		Type:        "match",
		IndentLevel: currentIndent,
		HasElse:     false,
		Token:       p.currToken,
	})
	stmt := &ast.MatchStatement{Token: p.currToken}

	p.nextToken()
	stmt.MatchValue = p.parseExpression(LOWEST)

	if !p.expectPeek(token.COLON) {
		return nil
	}

	p.skipNewlines()
	if p.peekTokenIs(token.INDENT) {
		p.nextToken()
	}

	stmt.Cases = []*ast.CaseClause{}

	for p.peekTokenIs(token.CASE) {
		p.nextToken()
		caseClause := &ast.CaseClause{Token: p.currToken}

		p.nextToken()

		// Check if this is a wildcard case (case _:)
		if p.currTokenIs(token.UNDERSCORE) {
			// This is a wildcard case, treat it as default
			caseClause.Condition = &ast.WildcardExpression{Token: p.currToken}

			if !p.expectPeek(token.COLON) {
				return nil
			}

			if p.peekTokenIs(token.NEWLINE) {
				p.nextToken()
				if p.peekTokenIs(token.INDENT) {
					p.nextToken()
					caseClause.Body = p.parseBlockStatement()
				} else {
					caseClause.Body = &ast.BlockStatement{
						Token:      p.currToken,
						Statements: []ast.Statement{p.parseStatement()},
					}
				}
			} else {
				p.nextToken()
				caseClause.Body = &ast.BlockStatement{
					Token:      p.currToken,
					Statements: []ast.Statement{p.parseStatement()},
				}
			}

			// Set this as the default case instead of adding to regular cases
			stmt.Default = caseClause
			continue
		}

		caseClause.Condition = p.parseExpression(LOWEST)

		if !p.expectPeek(token.COLON) {
			return nil
		}

		if p.peekTokenIs(token.NEWLINE) {
			p.nextToken()
			if p.peekTokenIs(token.INDENT) {
				p.nextToken()
				caseClause.Body = p.parseBlockStatement()
			} else {
				caseClause.Body = &ast.BlockStatement{
					Token:      p.currToken,
					Statements: []ast.Statement{p.parseStatement()},
				}
			}
		} else {
			p.nextToken()
			caseClause.Body = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}

		stmt.Cases = append(stmt.Cases, caseClause)
	}

	if p.peekTokenIs(token.UNDERSCORE) {
		p.nextToken()
		defaultClause := &ast.CaseClause{Token: p.currToken}

		if !p.expectPeek(token.COLON) {
			return nil
		}

		if p.peekTokenIs(token.NEWLINE) {
			p.nextToken()
			if p.peekTokenIs(token.INDENT) {
				p.nextToken()
				defaultClause.Body = p.parseBlockStatement()
			} else {
				defaultClause.Body = &ast.BlockStatement{
					Token:      p.currToken,
					Statements: []ast.Statement{p.parseStatement()},
				}
			}
		} else {
			p.nextToken()
			defaultClause.Body = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}

		stmt.Default = defaultClause
	}
	if len(p.controlStack) > 0 {
		p.controlStack = p.controlStack[:len(p.controlStack)-1]
	}
	return stmt
}

func (p *Parser) skipNewlines() {
	for p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
	}
}

func (p *Parser) parseOtherwise() ast.Expression {
	p.errors = append(p.errors, "'otherwise' can only be used as part of an if statement")
	return nil
}

func (p *Parser) parseSelf() ast.Expression {
	return &ast.Identifier{
		Token: p.currToken,
		Value: "self",
	}
}

func (p *Parser) parseDotExpression(left ast.Expression) ast.Expression {
	exp := &ast.DotExpression{
		Token: p.currToken,
		Left:  left,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	exp.Right = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	return exp
}

func (p *Parser) parseParenExpression() ast.Expression {
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return &ast.TupleLiteral{
			Token:    p.currToken,
			Elements: []ast.Expression{},
		}
	}

	p.nextToken()
	firstExpr := p.parseExpression(LOWEST)
	if firstExpr == nil {
		return nil
	}

	if p.peekTokenIs(token.COMMA) {
		elements := []ast.Expression{firstExpr}

		for p.peekTokenIs(token.COMMA) {
			p.nextToken()
			p.nextToken()
			nextExpr := p.parseExpression(LOWEST)
			if nextExpr != nil {
				elements = append(elements, nextExpr)
			}
		}

		if !p.expectPeek(token.RPAREN) {
			return nil
		}

		return &ast.TupleLiteral{
			Token:    p.currToken,
			Elements: elements,
		}
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return firstExpr
}

func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.currToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)

	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()

		commaFn := p.infixParseFns[token.COMMA]
		delete(p.infixParseFns, token.COMMA)
		key := p.parseExpression(LOWEST)
		if commaFn != nil {
			p.infixParseFns[token.COMMA] = commaFn
		}

		if !p.expectPeek(token.COLON) {
			return nil
		}
		p.nextToken()

		commaFn = p.infixParseFns[token.COMMA]
		delete(p.infixParseFns, token.COMMA)
		value := p.parseExpression(LOWEST)
		if commaFn != nil {
			p.infixParseFns[token.COMMA] = commaFn
		}

		hash.Pairs[key] = value

		if !p.peekTokenIs(token.RBRACE) && !p.expectPeek(token.COMMA) {
			return nil
		}
	}

	if !p.expectPeek(token.RBRACE) {
		return nil
	}
	return hash
}

func (p *Parser) parseTupleLiteral() ast.Expression {
	tuple := &ast.TupleLiteral{Token: p.currToken}
	tuple.Elements = p.parseExpressionList(token.RPAREN)
	return tuple
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{Token: p.currToken}
	value, err := strconv.ParseFloat(p.currToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as float", p.currToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: p.currToken, Left: left}
	p.nextToken()
	exp.Index = p.parseExpression(LOWEST)
	if !p.expectPeek(token.RBRACK) {
		return nil
	}
	return exp
}

func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}
	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}
	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}
	if !p.expectPeek(end) {
		return nil
	}
	return list
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.currToken}
	array.Elements = p.parseExpressionList(token.RBRACK)
	return array
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) registerPostfix(tokenType token.TokenType, fn postfixParseFn) {
	p.postfixParseFns[tokenType] = fn
}

func (p *Parser) registerStatement(tokenType token.TokenType, fn func() ast.Statement) {
	p.statementParseFns[tokenType] = fn
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.currToken, Value: p.currToken.Literal}
}

func (p *Parser) parseBoolean() ast.Expression {
	value := (p.currToken.Type == token.TRUE)
	return &ast.Boolean{Token: p.currToken, Value: value}
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		// Skip blank lines, indentation dedents, and semicolons
		for p.currToken.Type == token.NEWLINE ||
			p.currToken.Type == token.INDENT ||
			p.currToken.Type == token.DEDENT ||
			p.currToken.Type == token.SEMICOLON {
			p.nextToken()
			if p.currToken.Type == token.EOF {
				break
			}
		}

		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	if p.currToken.Type == token.NEWLINE || p.currToken.Type == token.EOF ||
		p.currToken.Type == token.INDENT || p.currToken.Type == token.DEDENT {
		return nil
	}
	switch p.currToken.Type {
	case token.IF:
		return p.parseIfStatement()
	case token.ELSE:
		if !validateElseStatement(p) {
			p.errors = append(p.errors, "Unexpected 'else' without matching 'if'")
			return nil
		}
		return p.parseElseStatement()
	case token.WHILE:
		return p.parseWhileStatement()
	case token.GRIMOIRE:
		return p.parseGrimoireDefinition()
	case token.SPELL, token.INIT:
		return p.parseFunctionDefinition()
	case token.FOR:
		return p.parseForStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.IMPORT:
		return p.parseImportStatement()
	case token.MATCH:
		return p.parseMatchStatement()
	case token.ATTEMPT:
		return p.parseAttemptStatement()
	case token.MAIN:
		return p.parseMainStatement()
	case token.RESOLVE:
		return p.parseResolveStatement()
	case token.RAISE:
		return p.parseRaiseStatement()
	case token.ARCANE:
		return p.parseArcaneGrimoire()
	case token.IGNORE:
		return p.parseIgnoreStatement()
	case token.SKIP:
		return p.parseSkipStatement()
	case token.STOP:
		return p.parseStopStatement()
	case token.CHECK:
		return p.parseCheckStatement()
	case token.GLOBAL:
		return p.parseGlobalStatement()
	case token.AUTOCLOSE:
		return p.parseWithStatement()
	}
	leftExpr := p.parseAssignmentLHS()
	if leftExpr == nil {
		return nil
	}
	if p.peekTokenIs(token.COLON) || p.peekTokenIs(token.ASSIGN) ||
		p.peekTokenIs(token.INCREMENT) ||
		p.peekTokenIs(token.DECREMENT) ||
		p.peekTokenIs(token.MULTASSGN) ||
		p.peekTokenIs(token.DIVASSGN) {
		return p.finishAssignmentStatement(leftExpr)
	}

	stmt := &ast.ExpressionStatement{
		Token:      p.currToken,
		Expression: leftExpr,
	}

	if p.peekTokenIs(token.NEWLINE) || p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseElseStatement() ast.Statement {
	elseStmt := &ast.ElseStatement{Token: p.currToken}

	if !p.expectPeek(token.COLON) {
		return nil
	}

	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
		if p.peekTokenIs(token.INDENT) {
			p.nextToken()
			elseStmt.Body = p.parseBlockStatement()
		} else {
			elseStmt.Body = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}
	} else {
		p.nextToken()
		elseStmt.Body = &ast.BlockStatement{
			Token:      p.currToken,
			Statements: []ast.Statement{p.parseStatement()},
		}
	}

	return elseStmt
}

func validateElseStatement(p *Parser) bool {
	currentIndent := p.getCurrentIndent()

	validElse := false
	elseIndex := -1

	for i := len(p.controlStack) - 1; i >= 0; i-- {
		ctrl := p.controlStack[i]
		if ctrl.Type == "if" && ctrl.IndentLevel == currentIndent && !ctrl.HasElse {
			validElse = true
			elseIndex = i
			break
		}

		if ctrl.IndentLevel < currentIndent {
			break
		}
	}

	if validElse && elseIndex >= 0 {
		ctrl := p.controlStack[elseIndex]
		ctrl.HasElse = true
		p.controlStack[elseIndex] = ctrl
	}

	return validElse
}

func (p *Parser) finishAssignmentStatement(leftExpr ast.Expression) ast.Statement {
	var typeHint ast.Expression = nil

	if _, ok := leftExpr.(*ast.Identifier); ok {
		if p.peekTokenIs(token.COLON) {
			p.nextToken()
			if !p.expectPeek(token.IDENT) {
				return nil
			}
			typeHint = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
		}
	}

	if !p.peekTokenIs(token.ASSIGN) {
		return nil
	}
	p.nextToken()

	stmt := &ast.AssignStatement{
		Token:    p.currToken,
		Name:     leftExpr,
		Operator: p.currToken.Literal,
		TypeHint: typeHint,
	}

	p.nextToken()
	stmt.Value = p.parseAssignmentRHS()
	return stmt
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	// Modified parameter context handling
	if p.parsingParameters {
		// Handle type hints without IN expectation
		if t == token.COLON && p.peekTokenIs(token.COLON) {
			p.inTypeHintContext = true
			return true
		}

		// Allow commas in parameter lists
		if t == token.COMMA && p.peekTokenIs(token.COMMA) {
			return true
		}
	}

	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
	return false
}

func (p *Parser) parseAssignmentStatement() *ast.AssignStatement {
	stmt := &ast.AssignStatement{Token: p.currToken}

	if p.currToken.Type == token.IDENT || p.currToken.Type == token.SELF {
		stmt.Name = p.parseExpression(LOWEST)
	} else {
		p.errors = append(p.errors, "Invalid assignment target")
		return nil
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	stmt.Operator = p.currToken.Literal

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	return stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.currToken}
	p.nextToken()
	stmt.ReturnValue = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currToken.Type)
		return nil
	}

	leftExp := prefix()
	for !p.peekTokenIs(token.NEWLINE) &&
		!p.peekTokenIs(token.SEMICOLON) &&
		!p.peekTokenIs(token.EOF) &&
		!p.peekTokenIs(token.COMMA) &&
		precedence < p.peekPrecedence() {

		if postfixFn, ok := p.postfixParseFns[p.peekToken.Type]; ok {
			p.nextToken()
			leftExp = postfixFn(leftExp)
			continue
		}

		infixFn := p.infixParseFns[p.peekToken.Type]
		if infixFn == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infixFn(leftExp)
	}
	return leftExp
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) currPrecedence() int {
	if p, ok := precedences[p.currToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.currToken}

	value, err := strconv.ParseInt(p.currToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
	}

	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
		Left:     left,
	}

	precedence := p.currPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	if expression.Right == nil {
		msg := fmt.Sprintf("no right-hand expression for infix operator %q", expression.Operator)
		p.errors = append(p.errors, msg)
		return nil
	}
	return expression
}

func (p *Parser) parsePostfixExpression(left ast.Expression) ast.Expression {
	return &ast.PostfixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Literal,
		Left:     left,
	}
}

func (p *Parser) parseIfStatement() ast.Statement {
	stmt := &ast.IfStatement{Token: p.currToken}

	currentIndent := p.getCurrentIndent()
	p.controlStack = append(p.controlStack, struct {
		Type        string
		IndentLevel int
		HasElse     bool
		Token       token.Token
	}{
		Type:        "if",
		IndentLevel: currentIndent,
		HasElse:     false,
		Token:       p.currToken,
	})

	if p.peekTokenIs(token.LPAREN) {
		p.nextToken()
		p.nextToken()
		stmt.Condition = p.parseExpression(LOWEST)
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	} else {
		p.nextToken()
		stmt.Condition = p.parseExpression(LOWEST)
	}

	if !p.expectPeek(token.COLON) {
		return nil
	}

	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
		if p.peekTokenIs(token.INDENT) {
			p.nextToken() // Move to INDENT token
			p.nextToken() // Move past INDENT to first statement token
			stmt.Consequence = p.parseBlockStatement()
		} else {
			stmt.Consequence = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}
	} else {
		p.nextToken()
		stmt.Consequence = &ast.BlockStatement{
			Token:      p.currToken,
			Statements: []ast.Statement{p.parseStatement()},
		}
	}

	for p.peekTokenIs(token.OTHERWISE) {
		p.nextToken()
		branch := ast.OtherwiseBranch{Token: p.currToken}

		if p.peekTokenIs(token.LPAREN) {
			p.nextToken()
			p.nextToken()
			branch.Condition = p.parseExpression(LOWEST)
			if !p.expectPeek(token.RPAREN) {
				return nil
			}
		} else {
			p.nextToken()
			branch.Condition = p.parseExpression(LOWEST)
		}

		if !p.expectPeek(token.COLON) {
			return nil
		}

		if p.peekTokenIs(token.NEWLINE) {
			p.nextToken()
			if p.peekTokenIs(token.INDENT) {
				p.nextToken() // Move to INDENT token
				p.nextToken() // Move past INDENT to first statement token
				branch.Consequence = p.parseBlockStatement()
			} else {
				branch.Consequence = &ast.BlockStatement{
					Token:      p.currToken,
					Statements: []ast.Statement{p.parseStatement()},
				}
			}
		} else {
			p.nextToken()
			branch.Consequence = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}

		stmt.OtherwiseBranches = append(stmt.OtherwiseBranches, branch)
	}

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		if !p.expectPeek(token.COLON) {
			return nil
		}

		if p.peekTokenIs(token.NEWLINE) {
			p.nextToken()
			if p.peekTokenIs(token.INDENT) {
				p.nextToken() // Move to INDENT token
				p.nextToken() // Move past INDENT to first statement token
				stmt.Alternative = p.parseBlockStatement()
			} else {
				stmt.Alternative = &ast.BlockStatement{
					Token:      p.currToken,
					Statements: []ast.Statement{p.parseStatement()},
				}
			}
		} else {
			p.nextToken()
			stmt.Alternative = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}
	}

	if len(p.controlStack) > 0 {
		p.controlStack = p.controlStack[:len(p.controlStack)-1]
	}

	return stmt
}

func (p *Parser) getCurrentIndent() int {
	return len(p.contextStack)
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.currToken}
	block.Statements = []ast.Statement{}
	prevIndent := p.getCurrentIndent()

	for !p.currTokenIs(token.DEDENT) && !p.currTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	currentIndent := p.getCurrentIndent()
	if currentIndent < prevIndent {
		for i := len(p.controlStack) - 1; i >= 0; i-- {
			if p.controlStack[i].IndentLevel >= prevIndent {
				p.controlStack = append(p.controlStack[:i], p.controlStack[i+1:]...)
			} else {
				break
			}
		}
	}
	return block
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.currToken, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return args
}

func (p *Parser) parseForStatement() ast.Statement {
	currentIndent := p.getCurrentIndent()
	p.controlStack = append(p.controlStack, struct {
		Type        string
		IndentLevel int
		HasElse     bool
		Token       token.Token
	}{
		Type:        "for",
		IndentLevel: currentIndent,
		HasElse:     false,
		Token:       p.currToken,
	})

	// Create a new ForStatement.
	fs := &ast.ForStatement{Token: p.currToken}

	// --- Parse the loop variable(s) - support both single variables and tuple unpacking ---
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	
	// Start with the first identifier
	firstVar := &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	
	// Check if there are more variables (tuple unpacking)
	if p.peekTokenIs(token.COMMA) {
		// Multiple variables - create a tuple literal
		variables := []ast.Expression{firstVar}
		
		// Parse remaining variables
		for p.peekTokenIs(token.COMMA) {
			p.nextToken() // consume comma
			if !p.expectPeek(token.IDENT) {
				return nil
			}
			nextVar := &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
			variables = append(variables, nextVar)
		}
		
		// Create tuple literal for multiple variables
		tupleLiteral := &ast.TupleLiteral{
			Token:    firstVar.Token,
			Elements: variables,
		}
		fs.Variable = tupleLiteral
	} else {
		// Single variable
		fs.Variable = firstVar
	}

	// --- Expect the 'in' keyword ---
	if !p.expectPeek(token.IN) {
		return nil
	}

	// Advance to the first token of the iterable expression.
	p.nextToken()
	fs.Iterable = p.parseExpression(LOWEST)

	// --- Expect a colon ---
	if !p.expectPeek(token.COLON) {
		return nil
	}

	// --- Parse the loop body ---
	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
		if p.peekTokenIs(token.INDENT) {
			p.nextToken() // Move to INDENT token
			p.nextToken() // Move past INDENT to first statement token
			fs.Body = p.parseBlockStatement()
		}
	} else {
		p.nextToken()
		fs.Body = &ast.BlockStatement{
			Token:      p.currToken,
			Statements: []ast.Statement{p.parseStatement()},
		}
	}

	if len(p.controlStack) > 0 {
		p.controlStack = p.controlStack[:len(p.controlStack)-1]
	}

	// Return fs as an ast.Statement (note: *ast.ForStatement implements ast.Statement).
	return fs
}

func (p *Parser) parseFunctionDefinition() ast.Statement {
	currentIndent := p.getCurrentIndent()
	p.controlStack = append(p.controlStack, struct {
		Type        string
		IndentLevel int
		HasElse     bool
		Token       token.Token
	}{
		Type:        "spell",
		IndentLevel: currentIndent,
		HasElse:     false,
		Token:       p.currToken,
	})
	stmt := &ast.FunctionDefinition{Token: p.currToken}

	if p.currTokenIs(token.SPELL) {
		p.nextToken()
	}

	if p.currToken.Type == token.INIT {
		stmt.Name = &ast.Identifier{
			Token: p.currToken,
			Value: "init",
		}
	} else if p.currToken.Type == token.IDENT {
		stmt.Name = &ast.Identifier{
			Token: p.currToken,
			Value: p.currToken.Literal,
		}
	} else {
		p.errors = append(p.errors, "Expected function name or 'init' after 'spell'")
		return nil
	}

	if !p.expectPeek(token.LPAREN) {
		p.errors = append(p.errors, "Expected '(' after function name")
		return nil
	}

	stmt.Parameters = p.parseFunctionParameters()

	// CRITICAL FIX: Explicitly exit parameter parsing mode
	p.parsingParameters = false

	if !p.expectPeek(token.COLON) {
		p.errors = append(p.errors, "Expected ':' after parameter list")
		return nil
	}

	// Keep track of the line where the colon is (before advancing)
	colonLine := p.currToken.Line

	p.nextToken()

	// Parse function body - handle both single-line and multi-line functions
	if p.currTokenIs(token.NEWLINE) {
		if p.peekTokenIs(token.INDENT) {
			p.nextToken() // Move to INDENT token
			p.nextToken() // Move past INDENT to first statement token
			stmt.Body = p.parseBlockStatement()
		} else {
			singleStmt := p.parseStatement()
			stmt.Body = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{singleStmt},
			}
		}
	} else {
		// Either single-line function or multi-line without NEWLINE token
		stmt.Body = &ast.BlockStatement{Token: p.currToken, Statements: []ast.Statement{}}

		// Check if we're on the same line as the colon (single-line function)
		if p.currToken.Line == colonLine {
			// Single-line function
			singleStmt := p.parseStatement()
			if singleStmt != nil {
				stmt.Body.Statements = append(stmt.Body.Statements, singleStmt)
			}
		} else {
			// Multi-line function - parse until we see clear end markers
			for !p.currTokenIs(token.EOF) {
				// Stop at column 1 on a different line (top-level statement)
				if p.currToken.Column <= 1 && p.currToken.Line > colonLine {
					break
				}

				// Parse statement
				s := p.parseStatement()
				if s != nil {
					stmt.Body.Statements = append(stmt.Body.Statements, s)
				}

				// Check if next token is at top level
				if p.peekToken.Column <= 1 || p.peekTokenIs(token.EOF) {
					break
				}

				p.nextToken()
			}
		}
	}

	if len(stmt.Body.Statements) > 0 {
		if exprStmt, ok := stmt.Body.Statements[0].(*ast.ExpressionStatement); ok {
			if strLit, ok := exprStmt.Expression.(*ast.StringLiteral); ok {
				if strLit.Token.Type == token.DOCSTRING {
					stmt.DocString = strLit
					stmt.Body.Statements = stmt.Body.Statements[1:]
				}
			}
		}
	}

	if len(p.controlStack) > 0 {
		p.controlStack = p.controlStack[:len(p.controlStack)-1]
	}
	return stmt
}

func (p *Parser) parseFunctionParameters() []ast.Expression {
	p.parsingParameters = true
	defer func() {
		p.parsingParameters = false
		p.inTypeHintContext = false
	}()

	parameters := []ast.Expression{}
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return parameters
	}

	p.nextToken()

	param := &ast.Parameter{
		Name: &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal},
	}

	if p.peekTokenIs(token.COLON) {
		p.nextToken()
		if !p.expectPeek(token.IDENT) {
			// Error already added by expectPeek, return empty slice
			return []ast.Expression{}
		}
		param.TypeHint = &ast.Identifier{
			Token: p.currToken,
			Value: p.currToken.Literal,
		}
	}

	if p.peekTokenIs(token.ASSIGN) {
		p.nextToken()
		p.nextToken()
		param.DefaultValue = p.parseExpression(LOWEST)
	}

	// Append identifier for simple params, otherwise full Parameter node
	if param.TypeHint == nil && param.DefaultValue == nil {
		parameters = append(parameters, param.Name)
	} else {
		parameters = append(parameters, param)
	}

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()

		param := &ast.Parameter{
			Name: &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal},
		}

		if p.peekTokenIs(token.COLON) {
			p.nextToken()
			if !p.expectPeek(token.IDENT) {
				// Error already added by expectPeek, return empty slice
				return []ast.Expression{}
			}
			param.TypeHint = &ast.Identifier{
				Token: p.currToken,
				Value: p.currToken.Literal,
			}
		}

		if p.peekTokenIs(token.ASSIGN) {
			p.nextToken()
			p.nextToken()
			param.DefaultValue = p.parseExpression(LOWEST)
		}

		// Append identifier for simple params, otherwise full Parameter node
		if param.TypeHint == nil && param.DefaultValue == nil {
			parameters = append(parameters, param.Name)
		} else {
			parameters = append(parameters, param)
		}
	}

	if !p.expectPeek(token.RPAREN) {
		// Error already added by expectPeek, return empty slice
		return []ast.Expression{}
	}

	return parameters
}

func (p *Parser) parseWhileStatement() ast.Statement {
	currentIndent := p.getCurrentIndent()
	p.controlStack = append(p.controlStack, struct {
		Type        string
		IndentLevel int
		HasElse     bool
		Token       token.Token
	}{
		Type:        "while",
		IndentLevel: currentIndent,
		HasElse:     false,
		Token:       p.currToken,
	})
	stmt := &ast.WhileStatement{Token: p.currToken}

	if p.peekTokenIs(token.LPAREN) {
		p.nextToken()
		p.nextToken()
		stmt.Condition = p.parseExpression(LOWEST)

		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	} else {
		p.nextToken()
		stmt.Condition = p.parseExpression(LOWEST)
	}

	if !p.expectPeek(token.COLON) {
		return nil
	}

	for p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
	}

	if p.peekTokenIs(token.INDENT) {
		p.nextToken() // Move to INDENT token
		p.nextToken() // Move past INDENT to first statement token
		stmt.Body = p.parseBlockStatement()
	} else {
		stmt.Body = p.parseBlockStatement()
	}
	if len(p.controlStack) > 0 {
		p.controlStack = p.controlStack[:len(p.controlStack)-1]
	}
	return stmt
}

func (p *Parser) parseGrimoireDefinition() ast.Statement {
	currentIndent := p.getCurrentIndent()
	p.controlStack = append(p.controlStack, struct {
		Type        string
		IndentLevel int
		HasElse     bool
		Token       token.Token
	}{
		Type:        "grim",
		IndentLevel: currentIndent,
		HasElse:     false,
		Token:       p.currToken,
	})
	stmt := &ast.GrimoireDefinition{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if p.peekTokenIs(token.LPAREN) {
		p.nextToken()
		p.nextToken()
		stmt.Inherits = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	}

	if !p.expectPeek(token.COLON) {
		return nil
	}

	stmt.Methods = []*ast.FunctionDefinition{}

	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
		if p.peekTokenIs(token.INDENT) {
			p.nextToken() // Move to INDENT token

			p.contextStack = append(p.contextStack, "grim")
			defer func() {
				p.contextStack = p.contextStack[:len(p.contextStack)-1]
			}()

			block := p.parseBlockStatement()

			if len(block.Statements) > 0 {
				if exprStmt, ok := block.Statements[0].(*ast.ExpressionStatement); ok {
					if strLit, ok := exprStmt.Expression.(*ast.StringLiteral); ok {
						if strLit.Token.Type == token.DOCSTRING {
							stmt.DocString = strLit
							block.Statements = block.Statements[1:]
						}
					}
				}
			}

			for _, s := range block.Statements {
				if exprStmt, ok := s.(*ast.ExpressionStatement); ok {
					if strLit, ok := exprStmt.Expression.(*ast.StringLiteral); ok {
						if strLit.Token.Type == token.DOCSTRING {
							continue
						}
					}
				}

				if fnDef, ok := s.(*ast.FunctionDefinition); ok {
					if fnDef.Name.Value == "init" {
						stmt.InitMethod = fnDef
					} else {
						stmt.Methods = append(stmt.Methods, fnDef)
					}
				} else {
					continue
				}
			}
		}
	} else {
		for p.peekTokenIs(token.SPELL) || p.peekTokenIs(token.INIT) {
			p.nextToken()
			fnStmt := p.parseFunctionDefinition()
			if fnStmt == nil {
				p.errors = append(p.errors, "Invalid function definition in single-line grim")
				return stmt
			}
			fnDef := fnStmt.(*ast.FunctionDefinition)
			if fnDef.Name.Value == "init" {
				stmt.InitMethod = fnDef
			} else {
				stmt.Methods = append(stmt.Methods, fnDef)
			}
		}
	}
	if len(p.controlStack) > 0 {
		p.controlStack = p.controlStack[:len(p.controlStack)-1]
	}
	return stmt
}

func (p *Parser) parseImportStatement() ast.Statement {
	stmt := &ast.ImportStatement{Token: p.currToken}

	if !p.expectPeek(token.STRING) {
		p.errors = append(p.errors, "expected file path string after 'import'")
		return nil
	}
	
	importPath := p.currToken.Literal
	
	// Check if this is a grimoire-only import (single name starting with uppercase)
	if !strings.Contains(importPath, "/") && !strings.Contains(importPath, ".") && 
		len(importPath) > 0 && importPath[0] >= 'A' && importPath[0] <= 'Z' {
		// This is a grimoire name import like import "HelloWorld"
		// We'll treat it as a selective import where the evaluator will search for the grimoire
		stmt.ClassName = &ast.Identifier{
			Token: p.currToken,
			Value: importPath,
		}
		// FilePath will be empty, signaling to search for this grimoire
		stmt.FilePath = &ast.StringLiteral{
			Token: p.currToken,
			Value: "", // Empty to indicate grimoire search
		}
	} else {
		// Check if the import path ends with a dot notation for selective imports
		// We need to check if the LAST dot represents a grimoire selection
		lastDotIndex := strings.LastIndex(importPath, ".")
		if lastDotIndex != -1 && lastDotIndex < len(importPath)-1 {
			// Check if the part after the last dot looks like a grimoire name (starts with uppercase)
			potentialGrimoire := importPath[lastDotIndex+1:]
			if len(potentialGrimoire) > 0 && potentialGrimoire[0] >= 'A' && potentialGrimoire[0] <= 'Z' {
				// This looks like a selective import: "module/path.GrimoireName"
				modulePath := importPath[:lastDotIndex]
				stmt.FilePath = &ast.StringLiteral{
					Token: p.currToken,
					Value: modulePath, // The module path without the grimoire name
				}
				stmt.ClassName = &ast.Identifier{
					Token: p.currToken,
					Value: potentialGrimoire, // The specific grimoire name
				}
			} else {
				// Regular import path that happens to have dots
				stmt.FilePath = &ast.StringLiteral{
					Token: p.currToken,
					Value: importPath,
				}
			}
		} else {
			stmt.FilePath = &ast.StringLiteral{
				Token: p.currToken,
				Value: importPath,
			}
		}
	}

	if p.peekTokenIs(token.AS) {
		p.nextToken()
		if !p.expectPeek(token.IDENT) {
			p.errors = append(p.errors, "expected alias name after 'as'")
			return nil
		}
		stmt.Alias = &ast.Identifier{
			Token: p.currToken,
			Value: p.currToken.Literal,
		}
	}

	return stmt
}

func (p *Parser) parseStringInterpolationLiteral() ast.Expression {
	raw := p.currToken.Literal

	si := &ast.StringInterpolation{
		Token: p.currToken,
		Parts: []ast.StringPart{},
	}

	var currentText strings.Builder
	i := 0
	for i < len(raw) {
		ch := raw[i]

		if ch == '$' && i+1 < len(raw) && raw[i+1] == '{' {
			if currentText.Len() > 0 {
				si.Parts = append(si.Parts, &ast.StringText{Value: currentText.String()})
				currentText.Reset()
			}

			exprStart := i + 2
			braceDepth := 1
			exprEnd := exprStart
			formatSpecStart := -1

			for exprEnd < len(raw) && braceDepth > 0 {
				if raw[exprEnd] == '{' {
					braceDepth++
				} else if raw[exprEnd] == '}' {
					braceDepth--
					if braceDepth == 0 {
						break
					}
				} else if raw[exprEnd] == ':' && braceDepth == 1 {
					formatSpecStart = exprEnd
				}
				exprEnd++
			}

			if braceDepth > 0 {
				p.errors = append(p.errors, "Unclosed brace in string interpolation")
				return si
			}

			var exprStr string
			var formatSpec string

			if formatSpecStart != -1 {
				exprStr = raw[exprStart:formatSpecStart]
				formatSpec = raw[formatSpecStart+1 : exprEnd]
			} else {
				exprStr = raw[exprStart:exprEnd]
			}

			expr := p.parseStringInterpolationExpression(exprStr)

			stringExpr := &ast.StringExpr{Expr: expr}

			if formatSpec != "" {
				parseFormatSpec(stringExpr, formatSpec)
			}

			si.Parts = append(si.Parts, stringExpr)
			i = exprEnd + 1
		} else {
			currentText.WriteByte(ch)
			i++
		}
	}

	if currentText.Len() > 0 {
		si.Parts = append(si.Parts, &ast.StringText{Value: currentText.String()})
	}

	return si
}

func parseFormatSpec(se *ast.StringExpr, spec string) {
	se.FormatSpec = spec

	i := 0
	if i+1 < len(spec) && (spec[i+1] == '<' || spec[i+1] == '>' || spec[i+1] == '^') {
		se.FillChar = spec[i]
		se.Alignment = spec[i+1]
		i += 2
	} else if i < len(spec) && (spec[i] == '<' || spec[i] == '>' || spec[i] == '^') {
		se.FillChar = ' '
		se.Alignment = spec[i]
		i++
	}

	widthStart := i
	for i < len(spec) && isDigit(spec[i]) {
		i++
	}
	if i > widthStart {
		width, _ := strconv.Atoi(spec[widthStart:i])
		se.Width = width
	}

	if i < len(spec) && spec[i] == '.' {
		i++
		precStart := i
		for i < len(spec) && isDigit(spec[i]) {
			i++
		}
		if i > precStart {
			prec, _ := strconv.Atoi(spec[precStart:i])
			se.Precision = prec
		}
	}
}

func (p *Parser) parseMainStatement() ast.Statement {
	stmt := &ast.MainStatement{Token: p.currToken}

	if !p.expectPeek(token.COLON) {
		p.errors = append(p.errors, "expected ':' after 'main'")
		return nil
	}

	// Handle newline after colon
	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken() // consume the newline
	}

	// Expect indented block
	if !p.peekTokenIs(token.INDENT) {
		p.errors = append(p.errors, "expected indented block after 'main:'")
		return nil
	}

	p.nextToken() // Move to INDENT token
	p.nextToken() // Move past INDENT to first statement token
	stmt.Body = p.parseBlockStatement()
	return stmt
}

func (p *Parser) parseStringInterpolationExpression(exprStr string) ast.Expression {
	l := lexer.New(exprStr)
	subParser := New(l)
	program := subParser.ParseProgram()

	if len(program.Statements) == 1 {
		if es, ok := program.Statements[0].(*ast.ExpressionStatement); ok {
			return es.Expression
		}
	}
	return nil
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (p *Parser) parseGlobalStatement() ast.Statement {
	stmt := &ast.GlobalStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	name := &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	stmt.Names = []*ast.Identifier{name}

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		if !p.expectPeek(token.IDENT) {
			return nil
		}
		name := &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
		stmt.Names = append(stmt.Names, name)
	}

	return stmt
}

func (p *Parser) parseWithStatement() ast.Statement {
	stmt := &ast.WithStatement{Token: p.currToken}
	
	// autoclose <expression> as <variable>:
	if !p.expectPeek(token.IDENT) && !p.currTokenIs(token.LPAREN) {
		return nil
	}
	
	// Parse the expression (e.g., open("file.txt", "r"))
	stmt.Expression = p.parseExpression(LOWEST)
	
	// Expect 'as'
	if !p.expectPeek(token.AS) {
		return nil
	}
	
	// Parse the variable name
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Variable = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	
	// Expect colon
	if !p.expectPeek(token.COLON) {
		return nil
	}
	
	// Expect newline and indent
	if !p.expectPeek(token.NEWLINE) {
		return nil
	}
	if !p.expectPeek(token.INDENT) {
		return nil
	}
	
	// Parse the body
	stmt.Body = p.parseBlockStatement()
	
	return stmt
}
