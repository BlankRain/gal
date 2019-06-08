package parser

import (
	"fmt"

	"github.com/BlankRain/gal/ast"
	"github.com/BlankRain/gal/token"
	"github.com/dgraph-io/dgraph/gql"
)

func (p *Parser) parseNodeTypeLiteral() ast.Expression {
	// package a.b.c.d ;
	ex := &ast.NodeTypeLiteral{
		Token: p.curToken,
	}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	ex.NodeName = p.curToken.Literal
	// {
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	ex.Properties = p.parseNodeTypeProperties()
	// }
	//@Query{}
	ex.Query = p.parseNodeTypeQuery()
	return ex
}

func (p *Parser) parseNodeTypeProperties() []*ast.Property {
	ret := []*ast.Property{}
	if !p.expectPeek(token.IDENT) {
		return ret
	}
	property := &ast.Property{
		Name: p.curToken.Literal,
	}
	if !p.expectPeek(token.SHOW) {
		return ret
	}
	if !p.expectPeek(token.IDENT) {
		return ret
	}
	property.Type = p.curToken.Literal
	ret = append(ret, property)
	// append others
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		if p.peekTokenIs(token.RBRACE) {
			break
		}
		p.nextToken()
		property := &ast.Property{
			Name: p.curToken.Literal,
		}
		if !p.expectPeek(token.SHOW) {
			return ret
		}
		if !p.expectPeek(token.IDENT) {
			return ret
		}
		property.Type = p.curToken.Literal
		ret = append(ret, property)
	}
	//
	if !p.expectPeek(token.RBRACE) {
		return nil
	}
	return ret
}
func (p *Parser) parseNodeTypeQuery() *ast.GraphQLLiteral {
	//@
	if !p.expectPeek(token.AT) {
		return nil
	}
	//Query
	if !p.expectPeek(token.IDENT) && p.curToken.Literal == "Query" {
		return nil
	}
	//
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	if !p.expectPeek(token.GQL) {
		return nil
	}
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
	if !p.expectPeek(token.RBRACE) {
		return nil
	}
	return ex
}
