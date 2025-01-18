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

	p.registerPrefix(token.LPAREN, p.parseParenExpression)
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

	return p
}

func (p *Parser) parseDotExpression(left ast.Expression) ast.Expression {
	exp := &ast.DotExpression{ // Create a new DotExpression node
		Token: p.currToken,
		Left:  left, // The object being accessed (e.g., Test)
	}

	if !p.expectPeek(token.IDENT) {
		return nil // Expecting a method or attribute name
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
	case token.SPELL:
		return p.parseFunctionDefinition()
	default:
		if fn, ok := p.statementParseFns[p.currToken.Type]; ok {
			return fn()
		}
	}
	// Handle assignment statements
	if p.currToken.Type == token.IDENT && p.peekTokenIs(token.ASSIGN) {
		return p.parseAssignmentStatement()
	}

	return p.parseExpressionStatement()
}

func (p *Parser) parseAssignmentStatement() *ast.AssignStatement {
	stmt := &ast.AssignStatement{Token: p.currToken}

	stmt.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) &&
		!p.expectPeek(token.INCREMENT) &&
		!p.expectPeek(token.DECREMENT) &&
		!p.expectPeek(token.MULTASSGN) &&
		!p.expectPeek(token.DIVASSGN) {
		return nil
	}

	stmt.Operator = p.currToken.Literal

	p.nextToken()
	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.NEWLINE) {
		p.nextToken()
	}

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
	} else {
		p.peekError(t)
		fmt.Printf("expectPeek: expected %s, and got %s\n", t, p.peekToken.Type)
		return false
	}
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

	// 1) Parse condition (with optional parentheses)
	//    If next token is LPAREN, parse condition inside ( ), else parse bare expression
	if p.peekTokenIs(token.LPAREN) {
		// consume the 'if' token
		p.nextToken() // move to '('
		p.nextToken() // now at first token of condition
		stmt.Condition = p.parseExpression(LOWEST)
		if !p.expectPeek(token.RPAREN) {
			return nil
		}
	} else {
		// no parentheses
		p.nextToken()
		stmt.Condition = p.parseExpression(LOWEST)
	}

	// 2) Expect a colon
	if !p.expectPeek(token.COLON) {
		return nil
	}

	// 3) Single-line or multi-line?
	if p.peekTokenIs(token.NEWLINE) {
		// MULTI-LINE mode
		p.nextToken() // consume NEWLINE

		if p.peekTokenIs(token.INDENT) {
			// parse an indented block
			p.nextToken() // consume INDENT
			stmt.Consequence = p.parseBlockStatement()
		} else {
			// next statement on a new line, but not indented => single statement
			singleStmt := p.parseStatement()
			stmt.Consequence = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{singleStmt},
			}
		}
	} else {
		// SINGLE-LINE (inline) form => e.g. `if ...: return 10`
		p.nextToken() // move to first token of that single statement
		singleStmt := p.parseStatement()
		stmt.Consequence = &ast.BlockStatement{
			Token:      p.currToken,
			Statements: []ast.Statement{singleStmt},
		}
	}

	// 4) Handle ANY number of `otherwise(...)` branches
	//    each 'otherwise' is like an "elseif"
	for p.peekTokenIs(token.OTHERWISE) {
		p.nextToken() // consume 'otherwise'
		branch := ast.OtherwiseBranch{Token: p.currToken}

		// optional parentheses again
		if p.peekTokenIs(token.LPAREN) {
			p.nextToken() // consume '('
			p.nextToken() // now at condition expression
			branch.Condition = p.parseExpression(LOWEST)
			if !p.expectPeek(token.RPAREN) {
				return nil
			}
		} else {
			// bare expression
			p.nextToken()
			branch.Condition = p.parseExpression(LOWEST)
		}

		// expect a colon
		if !p.expectPeek(token.COLON) {
			return nil
		}

		// single-line or multi-line again
		if p.peekTokenIs(token.NEWLINE) {
			// MULTI-LINE
			p.nextToken() // consume NEWLINE
			if p.peekTokenIs(token.INDENT) {
				p.nextToken() // consume INDENT
				branch.Consequence = p.parseBlockStatement()
			} else {
				singleStmt := p.parseStatement()
				branch.Consequence = &ast.BlockStatement{
					Token:      p.currToken,
					Statements: []ast.Statement{singleStmt},
				}
			}
		} else {
			// SINGLE-LINE
			p.nextToken()
			singleStmt := p.parseStatement()
			branch.Consequence = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{singleStmt},
			}
		}

		// append this 'otherwise' branch
		stmt.OtherwiseBranches = append(stmt.OtherwiseBranches, branch)

	}

	// 5) Finally, handle `else:` if present
	if p.peekTokenIs(token.ELSE) {
		p.nextToken() // consume else
		if !p.expectPeek(token.COLON) {
			return nil
		}

		if p.peekTokenIs(token.NEWLINE) {
			// MULTI-LINE
			p.nextToken() // consume NEWLINE
			if p.peekTokenIs(token.INDENT) {
				p.nextToken() // consume INDENT
				stmt.Alternative = p.parseBlockStatement()
			} else {
				singleStmt := p.parseStatement()
				stmt.Alternative = &ast.BlockStatement{
					Token:      p.currToken,
					Statements: []ast.Statement{singleStmt},
				}
			}
		} else {
			// SINGLE-LINE
			p.nextToken()
			singleStmt := p.parseStatement()
			stmt.Alternative = &ast.BlockStatement{
				Token:      p.currToken,
				Statements: []ast.Statement{singleStmt},
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

	// Expect the function name
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	// Expect the parameter list
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	stmt.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.COLON) {
		return nil
	}

	// Check the next token after COLON
	p.nextToken()

	if p.currTokenIs(token.NEWLINE) {
		// We expect an INDENT if it's a multiline function
		if !p.expectPeek(token.INDENT) {
			return nil
		}
		stmt.Body = p.parseBlockStatement()
	} else {
		// Inline statement (function with a single statement)
		stmt.Body = &ast.BlockStatement{
			Statements: []ast.Statement{p.parseStatement()},
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

	// Expect the spellbook name
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	// Expect a colon
	if !p.expectPeek(token.COLON) {
		return nil
	}

	// Parse methods and init
	stmt.Methods = []*ast.FunctionDefinition{}
	for p.peekTokenIs(token.SPELL) || p.peekTokenIs(token.INIT) {
		p.nextToken()

		if p.currToken.Type == token.INIT {
			if stmt.InitMethod != nil {
				p.errors = append(p.errors, "Duplicate init method")
				return nil
			}

			// Assert the type of parseFunctionDefinition to *ast.FunctionDefinition
			fn, ok := p.parseFunctionDefinition().(*ast.FunctionDefinition)
			if !ok {
				p.errors = append(p.errors, "Invalid function definition in init method")
				return nil
			}
			stmt.InitMethod = fn
		} else {
			// Assert the type of parseFunctionDefinition to *ast.FunctionDefinition
			fn, ok := p.parseFunctionDefinition().(*ast.FunctionDefinition)
			if !ok {
				p.errors = append(p.errors, "Invalid function definition in spellbook method")
				return nil
			}
			stmt.Methods = append(stmt.Methods, fn)
		}
	}

	return stmt
}
