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

func TestNodeType(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&NodeTypeLiteral{
				Token:      token.Token{Type: token.NODETYPE, Literal: "NodeType"},
				NodeName:   "Page",
				Properties: []property{},
				Query:      query{},
			},
		},
	}
	if program.String() != "NodeType  Page{}" {
		t.Fatalf("expect 'NodeType Page{};'but,got  '%v'", program.String())
	}
}
