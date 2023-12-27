package parser

import (
	"fmt"

	"github.com/ricci2511/riccilang/ast"
	"github.com/ricci2511/riccilang/lexer"
	"github.com/ricci2511/riccilang/token"
)

// The parser of the riccilang language. It's responsible for producing the AST from the lexer's tokens.
type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken() // First call only initializes peekToken
	p.nextToken() // Call a second time to initialize currToken
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Parse statements from the lexer's tokens until EOF
	for !p.currTokenIs(token.EOF) {
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
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	// TODO: Add more statements parsing cases
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}

	// Second token of a let statement must be an identifier
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	// Third token of a let statement must be an assign token
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: Handle expressions parsing (e.g. let x = 5 + 5 * 10;)
	// For now, just move forward until we encounter a semicolon
	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currToken}

	// TODO: Handle expressions parsing
	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// Helper functions

func (p *Parser) currTokenIs(tt token.TokenType) bool {
	return p.currToken.Type == tt
}

func (p *Parser) peekTokenIs(tt token.TokenType) bool {
	return p.peekToken.Type == tt
}

func (p *Parser) expectPeek(tt token.TokenType) bool {
	if p.peekTokenIs(tt) {
		p.nextToken() // Move forward if the next token is the expected one
		return true
	}
	p.peekError(tt) // Add parse error if the given token doesn't match the expected one
	return false
}

func (p *Parser) peekError(tt token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", tt, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
