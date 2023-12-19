package parser

import (
	"testing"

	"github.com/ricci2511/riccilang/ast"
	"github.com/ricci2511/riccilang/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	let 88924;
	`

	l := lexer.New(input) // Lexer
	p := New(l)           // Parser

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Expected 3 statements, got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return // Stop on failure
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return // All good
	}

	t.Errorf("Parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Errorf("Expected token literal 'let', got %q", stmt.TokenLiteral())
		return false
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("Expected *ast.LetStatement, got %T", stmt)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("Expected letStmt.Name.Value to be %q, got %q", name, letStmt.Name.Value)
		return false
	}

	// TODO: Check letStmt.Value field for integer literals

	return true
}
