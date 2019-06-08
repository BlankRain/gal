package parser

import (
	"fmt"

	"github.com/BlankRain/gal/ast"
	"github.com/dgraph-io/dgraph/gql"
)

func (p *Parser) parseGraphQLLiteral() ast.Expression {
	ex := &ast.GraphQLLiteral{
		Token: p.curToken,
		Body:  p.curToken.Literal,
	}
	r, e := gql.Parse(gql.Request{Str: ex.Body})
	if e != nil {
		msg := fmt.Sprintf("could not parse %q as graphql", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	ex.Result = r
	return ex
}
