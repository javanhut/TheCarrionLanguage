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

	token.LSHIFT: 6,

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
	prefixParseFns    map[token.TokenType]prefixParseFn
	infixParseFns     map[token.TokenType]infixParseFn
	postfixParseFns   map[token.TokenType]postfixParseFn
	statementParseFns map[token.TokenType]func() ast.Statement
}

func (p *Parser) isInsideGrimoire() bool {
	if len(p.contextStack) == 0 {
		return false
	}
	return p.contextStack[len(p.contextStack)-1] == "grim"
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
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
	p.registerPrefix(token.COLON, func() ast.Expression {
		return nil
	})
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
	p.registerPrefix(token.OTHERWISE, p.parseOtherwise)
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
	// Add more cases as needed for other expression types.
	default:
		// As a fallback, return the current token from the parser,
		// or you could return an ILLEGAL token.
		return token.Token{Type: token.ILLEGAL, Literal: ""}
	}
}

// parseExpressionTuple parses a comma-separated list of expressions into a tuple.
// If only a single expression is present, it returns that expression directly.
// The errorPrefix is used to customize error messages based on context (LHS or RHS).
func (p *Parser) parseExpressionTuple(firstExpr ast.Expression, isLHS bool) ast.Expression {
	if firstExpr == nil {
		return nil
	}

	// Check if there is a comma following; if not, return the expression as-is.
	if !p.peekTokenIs(token.COMMA) {
		return firstExpr
	}

	var expressions []ast.Expression
	expressions = append(expressions, firstExpr)

	// Customize error message based on context
	errorPrefix := "expected expression after comma in assignment"
	if isLHS {
		errorPrefix = "expected assignable expression after comma in assignment"
	}

	// While there is a comma, consume it and parse the next expression.
	for p.peekTokenIs(token.COMMA) {
		p.nextToken() // Consume the comma.
		// Skip any consecutive commas.
		for p.currToken.Type == token.COMMA {
			p.nextToken()
		}

		// If we've reached the end of input or a token that cannot start an expression,
		// then we have a trailing comma which is an error.
		if p.isAtEnd() || !p.canStartExpression(p.currToken.Type) {
			p.errors = append(p.errors, "unexpected trailing comma in assignment")
			return nil
		}

		// At this point, p.currToken should be the first token of the next expression.
		nextExpr := p.parseExpression(LOWEST)
		if nextExpr == nil {
			p.errors = append(p.errors, errorPrefix)
			return nil
		}
		expressions = append(expressions, nextExpr)
	}

	// For RHS, return the single expression directly instead of wrapping in a tuple
	if len(expressions) == 1 && !isLHS {
		return expressions[0]
	}

	return &ast.TupleLiteral{
		Token:    tokenFromExpression(expressions[0]),
		Elements: expressions,
	}
}

// parseAssignmentLHS parses an assignment left-hand side.
// It first parses a generic expression and, if a comma is encountered,
// aggregates additional expressions into a tuple literal.
func (p *Parser) parseAssignmentLHS() ast.Expression {
	// Parse a generic expression; this handles identifiers, dot expressions, etc.
	expr := p.parseExpression(LOWEST)
	return p.parseExpressionTuple(expr, true)
}

// parseAssignmentRHS parses an assignment right-hand side,
// gathering one or more expressions into a tuple literal if a comma is present.
func (p *Parser) parseAssignmentRHS() ast.Expression {
	exp := p.parseExpression(LOWEST)
	return p.parseExpressionTuple(exp, false)
}

// isAtEnd checks if we've reached the end of the token stream
func (p *Parser) isAtEnd() bool {
	return p.currToken.Type == token.EOF
}

// canStartExpression checks if the given token type can start an expression
func (p *Parser) canStartExpression(tokenType token.TokenType) bool {
	// This is a simplification - you should include all token types that can start an expression
	switch tokenType {
	case token.IDENT, token.INT, token.FLOAT, token.STRING, token.TRUE, token.FALSE, token.NONE,
		token.LPAREN, token.LBRACK, token.LBRACE, token.SPELL, token.IF:
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
	l := lexer.New(
		exprStr,
	)
	subParser := New(l)

	program := subParser.ParseProgram()

	if len(program.Statements) == 1 {
		if es, ok := program.Statements[0].(*ast.ExpressionStatement); ok {
			return es.Expression
		}
	}

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

func (p *Parser) parseAttemptStatement() ast.Statement {
	stmt := &ast.AttemptStatement{Token: p.currToken}

	if !p.expectPeek(token.COLON) {
		return nil
	}
	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
		if p.peekTokenIs(token.INDENT) {
			p.nextToken()
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

	for p.peekTokenIs(token.ENSNARE) {
		p.nextToken()
		ensnareClause := &ast.EnsnareClause{Token: p.currToken}

		if p.peekTokenIs(token.LPAREN) {
			p.nextToken()
			p.nextToken()
			ensnareClause.Condition = p.parseExpression(LOWEST)
			if !p.expectPeek(token.RPAREN) {
				return nil
			}
		} else {
			p.nextToken()
			ensnareClause.Condition = p.parseExpression(LOWEST)
		}

		if p.peekTokenIs(token.AS) {
			p.nextToken()
			if !p.expectPeek(token.IDENT) {
				return nil
			}
			ensnareClause.Alias = &ast.Identifier{
				Token: p.currToken,
				Value: p.currToken.Literal,
			}
		}

		if !p.expectPeek(token.COLON) {
			return nil
		}
		if p.peekTokenIs(token.NEWLINE) {
			p.nextToken()
			if p.peekTokenIs(token.INDENT) {
				p.nextToken()
				ensnareClause.Consequence = p.parseBlockStatement()
			} else {
				ensnareClause.Consequence = &ast.BlockStatement{
					Token:      p.currToken,
					Statements: []ast.Statement{p.parseStatement()},
				}
			}
		} else {
			p.nextToken()
			ensnareClause.Consequence = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}

		stmt.EnsnareClauses = append(stmt.EnsnareClauses, ensnareClause)
	}

	if p.peekTokenIs(token.RESOLVE) {
		p.nextToken()
		if !p.expectPeek(token.COLON) {
			return nil
		}
		if p.peekTokenIs(token.NEWLINE) {
			p.nextToken()
			if p.peekTokenIs(token.INDENT) {
				p.nextToken()
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

	return stmt
}

func (p *Parser) skipNewlines() {
	for p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
	}
}

func (p *Parser) parseOtherwise() ast.Expression {
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

		for p.currToken.Type == token.NEWLINE ||
			p.currToken.Type == token.INDENT ||
			p.currToken.Type == token.DEDENT {
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
	if p.currToken.Type == token.NEWLINE || p.currToken.Type == token.EOF {
		return nil
	}
	switch p.currToken.Type {
	case token.IF:
		return p.parseIfStatement()
	case token.ELSE:
		p.errors = append(p.errors, "Unexpected 'else' without matching 'if'")
		return nil
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

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
	return false
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
		!p.peekTokenIs(token.COMMA) && // stop if a comma is encountered
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
			p.nextToken()
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
				p.nextToken()
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
				p.nextToken()
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

	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.currToken}
	block.Statements = []ast.Statement{}

	/*if p.currTokenIs(token.INDENT) {
		p.nextToken()
	}*/

	for !p.currTokenIs(token.DEDENT) && !p.currTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
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
	stmt := &ast.ForStatement{Token: p.currToken}

	p.nextToken()

	var loopVars []ast.Expression

	expr := p.parseExpression(LOWEST)
	if expr == nil {
		return nil
	}
	loopVars = append(loopVars, expr)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		expr = p.parseExpression(LOWEST)
		if expr == nil {
			return nil
		}
		loopVars = append(loopVars, expr)
	}

	if len(loopVars) == 1 {
		stmt.Variable = loopVars[0]
	} else {
		stmt.Variable = &ast.TupleLiteral{
			Token:    p.currToken,
			Elements: loopVars,
		}
	}

	if !p.expectPeek(token.IN) {
		return nil
	}

	p.nextToken()
	stmt.Iterable = p.parseExpression(LOWEST)

	if !p.expectPeek(token.COLON) {
		return nil
	}

	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
		if !p.expectPeek(token.INDENT) {
			return nil
		}
		stmt.Body = p.parseBlockStatement()
	} else {
		p.nextToken()
		stmt.Body = &ast.BlockStatement{
			Token:      p.currToken,
			Statements: []ast.Statement{p.parseStatement()},
		}
	}

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		if !p.expectPeek(token.COLON) {
			return nil
		}
		if p.peekTokenIs(token.NEWLINE) {
			p.nextToken()
			if !p.expectPeek(token.INDENT) {
				return nil
			}
			stmt.Alternative = p.parseBlockStatement()
		} else {
			p.nextToken()
			stmt.Alternative = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}
	}

	return stmt
}

func (p *Parser) parseFunctionDefinition() ast.Statement {
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

	if !p.expectPeek(token.COLON) {
		p.errors = append(p.errors, "Expected ':' after parameter list")
		return nil
	}

	p.nextToken()

	if p.currTokenIs(token.NEWLINE) {
		if p.peekTokenIs(token.INDENT) {
			p.nextToken()
			stmt.Body = p.parseBlockStatement()
		} else {

			singleStmt := p.parseStatement()
			stmt.Body = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{singleStmt},
			}
		}
	} else {

		singleStmt := p.parseStatement()
		stmt.Body = &ast.BlockStatement{
			Token:      p.currToken,
			Statements: []ast.Statement{singleStmt},
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

	return stmt
}

func (p *Parser) parseFunctionParameters() []*ast.Parameter {
	commaInfix := p.infixParseFns[token.COMMA]
	delete(p.infixParseFns, token.COMMA)

	defer func() {
		if commaInfix != nil {
			p.infixParseFns[token.COMMA] = commaInfix
		}
	}()

	parameters := []*ast.Parameter{}
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
			return nil
		}
		param.TypeHint = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	}
	if p.peekTokenIs(token.ASSIGN) {
		p.nextToken()
		p.nextToken()
		param.DefaultValue = p.parseExpression(LOWEST)
	}
	parameters = append(parameters, param)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		param := &ast.Parameter{
			Name: &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal},
		}
		if p.peekTokenIs(token.COLON) {
			p.nextToken()
			if !p.expectPeek(token.IDENT) {
				return nil
			}
			param.TypeHint = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
		}
		if p.peekTokenIs(token.ASSIGN) {
			p.nextToken()
			p.nextToken()
			param.DefaultValue = p.parseExpression(LOWEST)
		}
		parameters = append(parameters, param)
	}
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return parameters
}

func (p *Parser) parseWhileStatement() ast.Statement {
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
		p.nextToken()
		stmt.Body = p.parseBlockStatement()
	} else {
		stmt.Body = p.parseBlockStatement()
	}

	return stmt
}

func (p *Parser) parseGrimoireDefinition() ast.Statement {
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
			p.nextToken()

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

	return stmt
}

func (p *Parser) parseImportStatement() ast.Statement {
	stmt := &ast.ImportStatement{Token: p.currToken}

	if !p.expectPeek(token.STRING) {
		p.errors = append(p.errors, "expected file path string after 'import'")
		return nil
	}
	stmt.FilePath = &ast.StringLiteral{
		Token: p.currToken,
		Value: p.currToken.Literal,
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
			// If we have accumulated text, add it as a StringText part
			if currentText.Len() > 0 {
				si.Parts = append(si.Parts, &ast.StringText{Value: currentText.String()})
				currentText.Reset()
			}

			// Find the matching closing brace
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
					// Format specifier found
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

			// Parse the expression
			expr := p.parseStringInterpolationExpression(exprStr)

			// Create a StringExpr with formatting information
			stringExpr := &ast.StringExpr{Expr: expr}

			// Parse format spec if present
			if formatSpec != "" {
				parseFormatSpec(stringExpr, formatSpec)
			}

			si.Parts = append(si.Parts, stringExpr)
			i = exprEnd + 1 // Skip past the closing brace
		} else {
			currentText.WriteByte(ch)
			i++
		}
	}

	// Add any remaining text
	if currentText.Len() > 0 {
		si.Parts = append(si.Parts, &ast.StringText{Value: currentText.String()})
	}

	return si
}

func parseFormatSpec(se *ast.StringExpr, spec string) {
	se.FormatSpec = spec

	i := 0
	// Check for fill and align
	if i+1 < len(spec) && (spec[i+1] == '<' || spec[i+1] == '>' || spec[i+1] == '^') {
		se.FillChar = spec[i]
		se.Alignment = spec[i+1]
		i += 2
	} else if i < len(spec) && (spec[i] == '<' || spec[i] == '>' || spec[i] == '^') {
		se.FillChar = ' ' // Default fill is space
		se.Alignment = spec[i]
		i++
	}

	// Parse width
	widthStart := i
	for i < len(spec) && isDigit(spec[i]) {
		i++
	}
	if i > widthStart {
		width, _ := strconv.Atoi(spec[widthStart:i])
		se.Width = width
	}

	// Parse precision
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
