package parser

import (
	"go/token"

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

func (p *Parser) ParseProgram() {
	return nil
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
