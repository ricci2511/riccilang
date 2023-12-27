package ast

import "github.com/ricci2511/riccilang/token"

// The base interface that all AST nodes implement
type Node interface {
	TokenLiteral() string
}

// An AST node that doesn't produce a value
type Statement interface {
	Node
	statementNode()
}

// An AST node that produces a value
type Expression interface {
	Node
	expressionNode()
}

// The root node of the AST, a riccilang program consists of a series of statements
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// The expression node that represents the identifier of a binding
type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// The statement node that represents a let binding (e.g. let x = 5;)
type LetStatement struct {
	Token token.Token // The token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// The statement node that represents a return statement (e.g. return 5;)
type ReturnStatement struct {
	Token       token.Token // The token.RETURN token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
