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

// Only supports ASCII for now
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

	l.eatWhitespace() // Skip whitespaces (spaces, tabs, newlines, etc.)

	// Generate the correct token based on the current char
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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readLiteral(isLetter)
			tok.Type = token.LookupIdent(tok.Literal) // Letters can be either keywords or user-defined identifiers
			return tok                                // Return early to avoid reading the next char since readLiteral() already did that
		} else if isDigit(l.ch) {
			tok.Literal = l.readLiteral(isDigit)
			tok.Type = token.INT
			return tok // Return early to avoid reading the next char since readLiteral() already did that
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar() // Read next char to keep the lexer moving
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar() // Keep reading until we hit a non-whitespace char
	}
}

type LiteralReader func(ch byte) bool

func (l *Lexer) readLiteral(reader LiteralReader) string {
	startingPos := l.position
	for reader(l.ch) {
		l.readChar() // Keep reading until we hit a char that doesn't match the reader's requirements
	}
	return l.input[startingPos:l.position]
}

// Checks if a given char is a letter (a-z, A-Z, _), satisfies the LiteralReader func type
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Checks if a given char is a digit (0-9), satisfies the LiteralReader func type
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
