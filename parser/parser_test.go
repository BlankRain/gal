package parser

import (
	"testing"

	"github.com/BlankRain/gal/ast"
	"github.com/BlankRain/gal/lexer"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5; 
	let y = 10;
	let hello = 12345;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contains 3 statements .Got %v", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"hello"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetstatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetstatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral() is not 'let',Got %v", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not a *ast.LetStatement, Got %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("name not equal. want %s  , got %s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not equal . want %s, got %s ", name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parse error : %q", msg)
	}
	t.FailNow()

}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return abc;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statement. got =%d", len(program.Statements))
	}
	for _, stmt := range program.Statements {
		retStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not a *ast.ReturnStatemnt .But %T", stmt)
			continue
		}
		if retStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral() is not return,but %q", retStmt.TokenLiteral())
		}

	}
}
