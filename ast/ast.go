package ast

import (
	"github.com/shengng325/monkey/token"
)

// The base Node interface
type Node interface {
	TokenLiteral() string
	// String() string
}

// All statement nodes implement this
type Statement interface {
	Node
	statementNode()
}

// All expression nodes implement this
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// func (p *Program) String() string {
// 	var out bytes.Buffer

// 	for _, s := range p.Statements {
// 		out.WriteString(s.String())
// 	}

// 	return out.String()
// }

// Implements node statement
type LetStatement struct { //let x = 5;
	Token token.Token // the token.LET token
	Name  *Identifier // x
	Value Expression  // 5
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Implements Expression interfacde
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
