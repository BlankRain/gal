package ast

import (
	"testing"

	"github.com/BlankRain/gal/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "someVar"},
					Value: "someVar",
				},
			},
		},
	}
	if program.String() != "let myVar = someVar;" {
		t.Fatalf("expect 'let myVar = someVar;'but,got  '%v'", program.String())
	}
}
