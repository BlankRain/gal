package parser

import (
	"github.com/BlankRain/gal/ast"
	"github.com/BlankRain/gal/token"
)

func (p *Parser) parseGraphQLLiteral() ast.Expression {
	ex := &ast.GraphQLLiteral{
		Token: p.curToken,
	}
	p.nextToken()
	if !p.peekTokenIs(token.GQLSTRING) {
		return nil
	}
	ex.Value = p.curToken.Literal
	return ex
}
