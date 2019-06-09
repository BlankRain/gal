package ast

import (
	"bytes"

	"github.com/BlankRain/gal/token"
)

type FromGraphLiteral struct {
	Token     token.Token
	NodeTypes []*Identifier
	EdgeTypes []*Identifier
	As        *AsExpression
}

func (node *FromGraphLiteral) expressionNode() {}
func (node *FromGraphLiteral) TokenLiteral() string {
	return node.Token.Literal
}
func (node *FromGraphLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(node.Token.Literal)
	out.WriteString(" Graph(")
	if node.NodeTypes != nil {
		out.WriteString("[")
		for i, nt := range node.NodeTypes {
			out.WriteString(nt.Value)
			if i != len(node.NodeTypes)-1 {
				out.WriteString(", ")
			}
		}
		out.WriteString("]")
		if node.EdgeTypes != nil {
			out.WriteString(", [")
			for i, et := range node.EdgeTypes {
				out.WriteString(et.Value)
				if i != len(node.EdgeTypes)-1 {
					out.WriteString(", ")
				}
			}
			out.WriteString("]")
		}
	}

	out.WriteString(") ")
	out.WriteString(node.As.Literal)
	return out.String()
}
