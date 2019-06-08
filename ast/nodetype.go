package ast

import (
	"bytes"
	"fmt"

	"github.com/BlankRain/gal/token"
	"github.com/dgraph-io/dgraph/gql"
)

/**
NodeType Page{
    URL: string @index(exact, fulltext) @count  @required @updated   @filter(StartWith('hello')),
    Name:   string @required(false)    ,
}@Query{
    has(website){
        uid
        url
        name
    }
}

**/
type NodeTypeLiteral struct {
	Token      token.Token
	NodeName   string
	Properties []*Property
	Query      *Query
}
type Property struct {
	Name string
	Type string
}
type Query struct {
	Body   string
	Result gql.Result
}

func (node *NodeTypeLiteral) expressionNode() {}
func (node *NodeTypeLiteral) TokenLiteral() string {
	return node.Token.Literal
}
func (node *NodeTypeLiteral) String() string {
	var out bytes.Buffer
	out.WriteString("NodeType  ")
	out.WriteString(node.NodeName)
	out.WriteString("{\n")
	// append property
	for _, property := range node.Properties {
		out.WriteString(property.Name + " ")
		out.WriteString(":" + property.Type + " ,\n")
	}
	out.WriteString("}")
	// append @query
	out.WriteString("@Query{\n")
	out.WriteString(fmt.Sprintf("%v", node.Query.Result))
	out.WriteString("\n}\n")
	return out.String()
}
func (node *NodeTypeLiteral) statementNode() {}
