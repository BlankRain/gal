package ast

import (
	"bytes"

	"github.com/BlankRain/gal/token"
	"github.com/dgraph-io/dgraph/gql"
)

type GraphQLLiteral struct {
	Token  token.Token
	Body   string
	Result gql.Result
}

func (node *GraphQLLiteral) expressionNode() {}
func (node *GraphQLLiteral) TokenLiteral() string {
	return node.Token.Literal
}
func (node *GraphQLLiteral) String() string {
	var out bytes.Buffer
	out.WriteString(node.Body)
	return out.String()
}
func (node *GraphQLLiteral) statementNode() {}
