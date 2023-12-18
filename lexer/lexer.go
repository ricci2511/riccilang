package lexer

import "github.com/ricci2511/riccilang/token"

type Lexer struct {
	input        string // The input string to be tokenized
	position     int    // Current position in input (points to current char)
	readPosition int    // Current reading position in input (after current char)
	ch           byte   // Current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Initializes the lexer's fields (position, readPosition, ch)
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" (null)
	} else {
		l.ch = l.input[l.readPosition] // Get next char
	}
	l.position = l.readPosition // Update current position
	l.readPosition += 1         // Update reading position
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Create the correct token based on the current char
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = "" // End of file
		tok.Type = token.EOF
	}

	l.readChar() // Read next char to keep the lexer moving
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
