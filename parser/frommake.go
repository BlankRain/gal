package parser

import (
	"github.com/BlankRain/gal/ast"
	"github.com/BlankRain/gal/token"
)

/**
from Graph([Page],[referTo])  g
make function  PageRank()  {
    print(g)
}
------------
from Graph([Page],[referTo]  g
make function  ShortestPath(sourePage Page ,endPage Page)
Path([Page]){
    print(g)
}
**/
func (p *Parser) parseFromGraphLiteral() ast.Expression {
	fm := &ast.FromGraphLiteral{
		Token: p.curToken,
	}
	if !(p.peekTokenIs(token.IDENT) && p.peekToken.Literal == "Graph") {
		p.addError("need from Graph, but got ...")
		return nil
	}
	p.nextToken()
	if !p.expectPeek(token.LPAREN) { //(
		return nil
	}
	fm.NodeTypes = p.parseFromGraphNodeTypes()
	fm.EdgeTypes = p.parseFromGraphEdgeTypes()
	if !p.expectPeek(token.RPAREN) { //)
		return nil
	}
	if !p.expectPeek(token.IDENT) { //as
		return nil
	}
	fm.As = &ast.AsExpression{
		Literal: p.curToken.Literal,
		Tokens:  []token.Token{p.curToken},
	}
	return fm
}

func (p *Parser) parseIdentifierArrays() []*ast.Identifier {
	nts := []*ast.Identifier{}
	if p.peekTokenIs(token.RPAREN) { // return )
		p.nextToken()
		return nts
	}
	if !p.expectPeek(token.LBRACKET) {
		return nil
	}
	p.nextToken()
	if p.curTokenIs(token.RBRACKET) {
		return nts
	}
	nt := &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	nts = append(nts, nt)
	// check , []
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		if p.peekTokenIs(token.RBRACKET) {
			break
		}
		if p.peekTokenIs(token.COMMA) {
			continue
		}
		p.nextToken()

		nt := &ast.Identifier{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
		nts = append(nts, nt)
	}
	if !p.expectPeek(token.RBRACKET) {
		return nil
	}
	return nts
}
func (p *Parser) parseFromGraphEdgeTypes() []*ast.Identifier {
	if p.peekTokenIs(token.RPAREN) {
		return nil
	}
	if p.peekTokenIs(token.COMMA) {
		p.nextToken()
	}
	return p.parseIdentifierArrays()
}
func (p *Parser) parseFromGraphNodeTypes() []*ast.Identifier {
	return p.parseIdentifierArrays()
}

func (p *Parser) parseMakeLiteral() ast.Expression {
	ml := &ast.MakeLiteral{
		Token: p.curToken,
	}
	return ml
}
