package parser

import (
	"fmt"
	"strconv"

	"thecarrionlanguage/ast"
	"thecarrionlanguage/lexer"
	"thecarrionlanguage/token"
)

const (
	_           int = iota
	LOWEST      int = iota
	ASSIGN          // =
	LOGICAL_OR      // or
	LOGICAL_AND     // and
	EQUALS          // ==, !=
	LESSGREATER     // >, <
	SUM             // +, -
	PRODUCT         // *, /, %
	PREFIX          // -X, !X, ++X, --X
	CALL            // myFunction(X)
	POSTFIX         // X++, X--
	INDEX
)

var precedences = map[token.TokenType]int{
	token.ASSIGN:          ASSIGN,
	token.INCREMENT:       ASSIGN, // +=
	token.DECREMENT:       ASSIGN, // -=
	token.MULTASSGN:       ASSIGN, // *=
	token.DIVASSGN:        ASSIGN, // /=
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
	token.PLUS_INCREMENT:  POSTFIX,
	token.MINUS_DECREMENT: POSTFIX,
	token.LPAREN:          CALL,
	token.DOT:             CALL,
	token.LBRACK:          INDEX,
	token.OR:              LOGICAL_OR,
	token.AND:             LOGICAL_AND,
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

func (p *Parser) isInsideSpellbook() bool {
	if len(p.contextStack) == 0 {
		return false
	}
	return p.contextStack[len(p.contextStack)-1] == "spellbook"
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.postfixParseFns = make(map[token.TokenType]postfixParseFn)
	p.statementParseFns = make(map[token.TokenType]func() ast.Statement)

	// Register prefix parsers
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.PLUS_INCREMENT, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS_DECREMENT, p.parsePrefixExpression)
	// p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.COLON, func() ast.Expression {
		return nil
	})

	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.LBRACK, p.parseArrayLiteral)
	p.registerPrefix(token.LBRACE, p.parseHashLiteral)
	p.registerPrefix(token.NEWLINE, func() ast.Expression { return nil })
	p.registerPrefix(token.INDENT, func() ast.Expression { return nil })
	p.registerPrefix(token.DEDENT, func() ast.Expression { return nil })
	p.registerPrefix(token.EOF, func() ast.Expression { return nil })
	p.registerPrefix(token.OTHERWISE, p.parseOtherwise)
	p.registerPrefix(token.LPAREN, p.parseParenExpression)
	p.registerPrefix(token.SELF, p.parseSelf)
	p.registerPrefix(token.INIT, func() ast.Expression {
		return &ast.Identifier{
			Token: token.Token{Type: token.INIT, Literal: "init"},
			Value: "init",
		}
	})
	// Register infix parsers
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.MOD, p.parseInfixExpression)
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
	// Register postfix parsers
	p.registerPostfix(token.PLUS_INCREMENT, p.parsePostfixExpression)
	p.registerPostfix(token.MINUS_DECREMENT, p.parsePostfixExpression)
	// Register statement parsers
	p.registerStatement(token.RETURN, p.parseReturnStatement)
	p.registerStatement(token.IF, p.parseIfStatement)
	p.registerStatement(token.FOR, p.parseForStatement)
	p.registerStatement(token.SPELL, p.parseFunctionDefinition)
	p.registerStatement(token.IMPORT, p.parseImportStatement)

	return p
}

func (p *Parser) parseOtherwise() ast.Expression {
	// Not directly parsed as an expression, handled in `parseIfStatement`
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

	// Expect the method or property name
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	exp.Right = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
	return exp
}

func (p *Parser) parseParenExpression() ast.Expression {
	// If next token is RPAREN => empty tuple
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken() // consume ')'
		return &ast.TupleLiteral{
			Token:    p.currToken, // this is the ')'
			Elements: []ast.Expression{},
		}
	}

	// 1) Parse the first expression after '('
	p.nextToken() // move past '(' to the first token inside
	firstExpr := p.parseExpression(LOWEST)
	if firstExpr == nil {
		return nil
	}

	// 2) If we see a comma after parsing the first expression, parse the rest => it’s a tuple
	if p.peekTokenIs(token.COMMA) {
		// We already have one expression in `firstExpr`
		elements := []ast.Expression{firstExpr}

		// Parse remaining expressions separated by commas
		for p.peekTokenIs(token.COMMA) {
			p.nextToken() // consume the comma
			p.nextToken() // move to the next expression
			nextExpr := p.parseExpression(LOWEST)
			if nextExpr != nil {
				elements = append(elements, nextExpr)
			}
		}

		// Expect closing parenthesis
		if !p.expectPeek(token.RPAREN) {
			return nil
		}

		// Return a tuple literal
		return &ast.TupleLiteral{
			Token:    p.currToken, // this is the ')'
			Elements: elements,
		}
	}

	// 3) Otherwise, we expect exactly one expression => grouped expression
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	// Return that single expression
	return firstExpr
}

func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.currToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)
	for !p.peekTokenIs(token.RBRACE) {
		p.nextToken()
		key := p.parseExpression(LOWEST)
		if !p.expectPeek(token.COLON) {
			return nil
		}
		p.nextToken()
		value := p.parseExpression(LOWEST)
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
		// Skip any leading newlines, dedents, etc.
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
	// Handle special statements first:
	switch p.currToken.Type {
	case token.IF:
		return p.parseIfStatement()
	case token.ELSE:
		p.errors = append(p.errors, "Unexpected 'else' without matching 'if'")
		return nil
	case token.WHILE:
		return p.parseWhileStatement()
	case token.SPELLBOOK:
		return p.parseSpellbookDefinition()
	case token.SPELL, token.INIT:
		// function definition at top level
		return p.parseFunctionDefinition()
	case token.FOR:
		return p.parseForStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	case token.IMPORT:
		return p.parseImportStatement()
		// Add any other top-level statements here...
	}

	leftExpr := p.parseExpression(LOWEST)

	// If next token is an assignment operator, finish as an assignment:
	if p.peekTokenIs(token.ASSIGN) ||
		p.peekTokenIs(token.INCREMENT) ||
		p.peekTokenIs(token.DECREMENT) ||
		p.peekTokenIs(token.MULTASSGN) ||
		p.peekTokenIs(token.DIVASSGN) {
		return p.finishAssignmentStatement(leftExpr)
	}

	// Else, it’s just an expression statement:
	stmt := &ast.ExpressionStatement{
		Token:      p.currToken, // or whatever token you want
		Expression: leftExpr,
	}

	if p.peekTokenIs(token.NEWLINE) || p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) finishAssignmentStatement(leftExpr ast.Expression) ast.Statement {
	// 1) Advance to the operator token
	p.nextToken()
	assignOp := p.currToken.Literal // e.g. "=" or "+="

	// 2) Build the AssignStatement node, using 'Name' for the leftExpr
	stmt := &ast.AssignStatement{
		Token:    p.currToken, // The token for '=' (or '+=')
		Name:     leftExpr,    // LHS expression (Identifier, DotExpression, etc.)
		Operator: assignOp,
	}

	// 3) Parse the right-hand side
	p.nextToken() // move past '='
	stmt.Value = p.parseExpression(LOWEST)

	// 4) Optionally consume trailing newline or semicolon
	if p.peekTokenIs(token.NEWLINE) || p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseAssignmentStatement() *ast.AssignStatement {
	stmt := &ast.AssignStatement{Token: p.currToken}

	// Parse left-hand side (allow Identifier or DotExpression)
	if p.currToken.Type == token.IDENT || p.currToken.Type == token.SELF {
		stmt.Name = p.parseExpression(LOWEST)
	} else {
		p.errors = append(p.errors, "Invalid assignment target")
		return nil
	}

	// Parse the assignment operator
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	stmt.Operator = p.currToken.Literal

	// Parse the right-hand side
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
		Token:    p.currToken, // The `++` token
		Operator: p.currToken.Literal,
		Left:     left,
	}
}

func (p *Parser) parseIfStatement() ast.Statement {
	stmt := &ast.IfStatement{Token: p.currToken} // token.IF

	// Parse condition (with optional parentheses)
	if p.peekTokenIs(token.LPAREN) {
		p.nextToken() // consume '('
		p.nextToken() // move to first token of condition
		stmt.Condition = p.parseExpression(LOWEST)
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	} else {
		// Parse bare condition without parentheses
		p.nextToken()
		stmt.Condition = p.parseExpression(LOWEST)
	}

	// Expect colon
	if !p.expectPeek(token.COLON) {
		return nil
	}

	// Parse the consequence block (single-line or multi-line)
	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken() // consume NEWLINE
		if p.peekTokenIs(token.INDENT) {
			p.nextToken() // consume INDENT
			stmt.Consequence = p.parseBlockStatement()
		} else {
			// Single statement on a new line but not indented
			stmt.Consequence = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}
	} else {
		// Inline consequence, e.g., `if ...: return x`
		p.nextToken()
		stmt.Consequence = &ast.BlockStatement{
			Token:      p.currToken,
			Statements: []ast.Statement{p.parseStatement()},
		}
	}

	// Parse any number of `otherwise` branches
	for p.peekTokenIs(token.OTHERWISE) {
		p.nextToken() // consume 'otherwise'
		branch := ast.OtherwiseBranch{Token: p.currToken}

		// Parse condition for otherwise (optional parentheses)
		if p.peekTokenIs(token.LPAREN) {
			p.nextToken() // consume '('
			p.nextToken() // move to condition
			branch.Condition = p.parseExpression(LOWEST)
			if !p.expectPeek(token.RPAREN) {
				return nil
			}
		} else {
			// No parentheses; parse bare expression
			p.nextToken()
			branch.Condition = p.parseExpression(LOWEST)
		}

		// Expect colon
		if !p.expectPeek(token.COLON) {
			return nil
		}

		// Parse block or inline statement
		if p.peekTokenIs(token.NEWLINE) {
			p.nextToken() // consume NEWLINE
			if p.peekTokenIs(token.INDENT) {
				p.nextToken() // consume INDENT
				branch.Consequence = p.parseBlockStatement()
			} else {
				branch.Consequence = &ast.BlockStatement{
					Token:      p.currToken,
					Statements: []ast.Statement{p.parseStatement()},
				}
			}
		} else {
			// Inline single statement
			p.nextToken()
			branch.Consequence = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{p.parseStatement()},
			}
		}

		stmt.OtherwiseBranches = append(stmt.OtherwiseBranches, branch)
	}

	// Parse optional `else` block
	if p.peekTokenIs(token.ELSE) {
		p.nextToken() // consume 'else'
		if !p.expectPeek(token.COLON) {
			return nil
		}

		if p.peekTokenIs(token.NEWLINE) {
			p.nextToken() // consume NEWLINE
			if p.peekTokenIs(token.INDENT) {
				p.nextToken() // consume INDENT
				stmt.Alternative = p.parseBlockStatement()
			} else {
				stmt.Alternative = &ast.BlockStatement{
					Token:      p.currToken,
					Statements: []ast.Statement{p.parseStatement()},
				}
			}
		} else {
			// Inline else block
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

	// Consume the INDENT token
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

	// Parse loop variable
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Variable = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	// Expect the 'in' keyword
	if !p.expectPeek(token.IN) {
		return nil
	}

	// Parse the iterable expression
	p.nextToken()
	stmt.Iterable = p.parseExpression(LOWEST)

	// Expect colon after the iterable
	if !p.expectPeek(token.COLON) {
		return nil
	}

	// Parse loop body
	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken() // consume NEWLINE
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

	// Parse optional else block
	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		if !p.expectPeek(token.COLON) {
			return nil
		}
		if p.peekTokenIs(token.NEWLINE) {
			p.nextToken() // consume NEWLINE
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

	// If current token is SPELL or INIT, we move one token ahead so we can see
	// either 'init' or the actual function name. (In a bare 'init' case, p.currToken is already INIT.)
	if p.currTokenIs(token.SPELL) {
		// Move forward one token => could be INIT or IDENT
		p.nextToken()
	}

	// Now p.currToken might be `INIT` or `IDENT`.
	if p.currToken.Type == token.INIT {
		// It's the special "init" function
		stmt.Name = &ast.Identifier{
			Token: p.currToken,
			Value: "init",
		}
	} else if p.currToken.Type == token.IDENT {
		// Normal named function
		stmt.Name = &ast.Identifier{
			Token: p.currToken,
			Value: p.currToken.Literal,
		}
	} else {
		// We expected either INIT or IDENT
		p.errors = append(p.errors, "Expected function name or 'init' after 'spell'")
		return nil
	}

	// Expect '(' next
	if !p.expectPeek(token.LPAREN) {
		p.errors = append(p.errors, "Expected '(' after function name")
		return nil
	}
	// Parse the parameter list
	stmt.Parameters = p.parseFunctionParameters()

	// Expect ':'
	if !p.expectPeek(token.COLON) {
		p.errors = append(p.errors, "Expected ':' after parameter list")
		return nil
	}

	// Now parse the function body (single-line or multi-line)
	p.nextToken() // consume ':' => move on

	if p.currTokenIs(token.NEWLINE) {
		// We likely have an indented block
		if p.peekTokenIs(token.INDENT) {
			p.nextToken() // consume INDENT
			stmt.Body = p.parseBlockStatement()
		} else {
			// We have a newline but no INDENT => maybe a single statement
			singleStmt := p.parseStatement()
			stmt.Body = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{singleStmt},
			}
		}
	} else {
		// No newline => single-line function body
		singleStmt := p.parseStatement()
		stmt.Body = &ast.BlockStatement{
			Token:      p.currToken,
			Statements: []ast.Statement{singleStmt},
		}
	}

	return stmt
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{
			Token: p.currToken,
			Value: p.currToken.Literal,
		}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseWhileStatement() ast.Statement {
	// Create an empty WhileStatement node.
	stmt := &ast.WhileStatement{Token: p.currToken}

	// 1) Parse the condition
	if p.peekTokenIs(token.LPAREN) {
		p.nextToken() // skip '('
		p.nextToken() // move to first token in condition
		stmt.Condition = p.parseExpression(LOWEST)
		// Expect closing parenthesis
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	} else {
		p.nextToken() // move to the expression start
		stmt.Condition = p.parseExpression(LOWEST)
	}

	// 2) Expect a colon
	if !p.expectPeek(token.COLON) {
		return nil
	}
	// 3) Consume optional newline(s)
	// -------------------------
	for p.peekTokenIs(token.NEWLINE) {
		p.nextToken() // skip the newline
	}

	if p.peekTokenIs(token.INDENT) {
		p.nextToken()                       // consume INDENT
		stmt.Body = p.parseBlockStatement() // parse statements until DEDENT
	} else {
		stmt.Body = p.parseBlockStatement()
	}

	return stmt
}

func (p *Parser) parseSpellbookDefinition() ast.Statement {
	stmt := &ast.SpellbookDefinition{Token: p.currToken}

	// Expect "spellbook Name"
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	// Expect colon => "spellbook Name:"
	if !p.expectPeek(token.COLON) {
		return nil
	}

	// Prepare to collect methods
	stmt.Methods = []*ast.FunctionDefinition{}

	// Check if next token is NEWLINE => possible multi-line block
	if p.peekTokenIs(token.NEWLINE) {
		// Consume the newline
		p.nextToken()

		// If we now see an INDENT, treat it as a multi-line block
		if p.peekTokenIs(token.INDENT) {
			// Consume INDENT
			p.nextToken()

			// Push "spellbook" context
			p.contextStack = append(p.contextStack, "spellbook")
			defer func() {
				// Pop the context once we exit
				p.contextStack = p.contextStack[:len(p.contextStack)-1]
			}()

			// Parse each statement until DEDENT/EOF.
			// We expect only `spell` or `init` inside the block, but
			// the parser will keep going until we see DEDENT.
			for !p.currTokenIs(token.DEDENT) && !p.currTokenIs(token.EOF) {
				if p.currTokenIs(token.SPELL) || p.currTokenIs(token.INIT) {
					fnStmt := p.parseFunctionDefinition()
					if fnStmt == nil {
						p.errors = append(p.errors, "Invalid function definition in spellbook")
					} else {
						fnDef := fnStmt.(*ast.FunctionDefinition)
						if fnDef.Name.Value == "init" {
							// Only one init allowed
							if stmt.InitMethod != nil {
								p.errors = append(p.errors, "Duplicate init method in spellbook")
							} else {
								stmt.InitMethod = fnDef
							}
						} else {
							stmt.Methods = append(stmt.Methods, fnDef)
						}
					}
				} else {
					// Something else in the spellbook block, which we don't allow
					if p.currToken.Type != token.NEWLINE &&
						p.currToken.Type != token.INDENT &&
						p.currToken.Type != token.DEDENT {
						msg := fmt.Sprintf(
							"Unexpected token %q inside spellbook; expected 'spell' or 'init'.",
							p.currToken.Literal)
						p.errors = append(p.errors, msg)
					}
				}
				p.nextToken()
			}
			// Done with multi‐line block
		} else {
			// We got a newline but no indent => treat as an empty block
		}
	} else {
		// Single-line style: "spellbook Person: spell init(...): ... etc."
		// Let’s parse as many `spell` or `init` definitions as appear on the same line
		for p.peekTokenIs(token.SPELL) || p.peekTokenIs(token.INIT) {
			p.nextToken()
			fnStmt := p.parseFunctionDefinition()
			if fnStmt == nil {
				p.errors = append(p.errors, "Invalid function definition in single-line spellbook")
				return stmt
			}
			fnDef := fnStmt.(*ast.FunctionDefinition)
			if fnDef.Name.Value == "init" {
				if stmt.InitMethod != nil {
					p.errors = append(p.errors, "Duplicate init method in spellbook")
				} else {
					stmt.InitMethod = fnDef
				}
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

	// Parse the file path
	filePath := p.currToken.Literal
	stmt.FilePath = &ast.StringLiteral{
		Token: p.currToken,
		Value: filePath,
	}

	// Check for the optional class name
	if p.peekTokenIs(token.DOT) {
		p.nextToken() // Consume '.'
		if !p.expectPeek(token.IDENT) {
			p.errors = append(p.errors, "expected class name after '.' in import statement")
			return nil
		}
		stmt.ClassName = &ast.Identifier{
			Token: p.currToken,
			Value: p.currToken.Literal,
		}
	}

	return stmt
}
