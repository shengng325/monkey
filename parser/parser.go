package parser

import (
	"github.com/shengng325/monkey/ast"
	"github.com/shengng325/monkey/lexer"
	"github.com/shengng325/monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// To intialise curToken and peekToken
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	// var statement ast.Statement
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// declare a Let statement
	// assign curToken into LetStatement.Token
	stmt := &ast.LetStatement{Token: p.curToken}

	// check if next token is Identifier
	// if true, run nextToken()
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// curToken is Identifier,
	// assign curToken into LetStatement.name
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// check if next token is "="
	// if true, run nextToken()
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// curTokenIs checks if current token is of specific type
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs peeks if the next token is of specific type
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek return true if the next token is of specific type,
// or vice versa
// it also call the nextToken function
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
