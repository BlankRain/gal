package parser

import (
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
func (p *Parser) parseNodeTypeQuery() *ast.Query {
	q := &ast.Query{}
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
	if !p.expectPeek(token.GQLSTRING) {
		return nil
	}
	q.Body = p.curToken.Literal
	r, _ := gql.Parse(gql.Request{Str: q.Body})
	// if e != nil {
	// 	return nil
	// }
	q.Result = r
	if !p.expectPeek(token.RBRACE) {
		return nil
	}
	return q
}
