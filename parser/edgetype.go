package parser

import (
	"github.com/BlankRain/gal/ast"
	"github.com/BlankRain/gal/token"
)

/**

EdgeType  referTo(s:Person,t:Person) @filter(has(s.Name)) @reverse {
    id :int,
    lable: string as s.Name,
}

**/
func (p *Parser) parseEdgeTypeLiteral() ast.Expression {
	// EdgeType
	ex := &ast.EdgeTypeLiteral{
		Token: p.curToken,
	}
	// referTo
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	ex.EdgeName = p.curToken.Literal
	if !p.expectPeek(token.LPAREN) { // (
		return nil
	}
	if !p.expectPeek(token.IDENT) { // s
		return nil
	}
	ex.SourceRefer = p.curToken.Literal
	if !p.expectPeek(token.SHOW) { //:
		return nil
	}
	if !p.expectPeek(token.IDENT) { // Person
		return nil
	}
	ex.SourceType = p.curToken.Literal
	if !p.expectPeek(token.COMMA) { //,
		return nil
	}

	if !p.expectPeek(token.IDENT) { // t
		return nil
	}
	ex.TargetRefer = p.curToken.Literal
	if !p.expectPeek(token.SHOW) { //:
		return nil
	}
	if !p.expectPeek(token.IDENT) { // Person
		return nil
	}
	ex.TargetType = p.curToken.Literal
	if !p.expectPeek(token.RPAREN) { // )
		return nil
	}
	for !p.peekTokenIs(token.LBRACE) {
		if p.peekTokenIs(token.AT) { //@
			p.nextToken()
			if !p.expectPeek(token.IDENT) {
				return nil
			}
			//reverse or filter
			if p.curToken.Literal == "reverse" { //@reverse
				ex.IsReverse = true
			} else if p.curToken.Literal == "filter" { //@filter
				ex.Filter = p.parseEdgeFilter()
			} else {
				p.addError("unknow token" + p.curToken.Literal)
			}
		}
	}
	ex.Properties = p.parseEdgeProperties()
	return ex
}
func (p *Parser) parseEdgeProperties() []*ast.EdgeProperty {
	eps := []*ast.EdgeProperty{}
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	edgeProperty := &ast.EdgeProperty{
		Name: p.curToken.Literal,
	}
	//: type
	if !p.expectPeek(token.SHOW) {
		return eps
	}
	if !p.expectPeek(token.IDENT) {
		return eps
	}
	edgeProperty.Type = p.curToken.Literal
	if p.peekTokenIs(token.AS) {
		edgeProperty.As = p.parseEdgeAsExpression()
	}
	eps = append(eps, edgeProperty)
	// append others
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		if p.peekTokenIs(token.RBRACE) {
			break
		}
		p.nextToken()
		edgeProperty := &ast.EdgeProperty{
			Name: p.curToken.Literal,
		}
		if !p.expectPeek(token.SHOW) {
			return eps
		}
		if !p.expectPeek(token.IDENT) {
			return eps
		}
		edgeProperty.Type = p.curToken.Literal
		if p.peekTokenIs(token.AS) {
			edgeProperty.As = p.parseEdgeAsExpression()
		}
		eps = append(eps, edgeProperty)
	}
	//
	if !p.expectPeek(token.RBRACE) {
		return nil
	}
	return eps
}

func (p *Parser) parseEdgeAsExpression() *ast.AsExpression {
	ae := &ast.AsExpression{}
	p.nextToken()
	literal := ""
	tks := []token.Token{}
	for !p.peekTokenIs(token.COMMA) {
		p.nextToken()
		literal += p.curToken.Literal
		tks = append(tks, p.curToken)
	}
	ae.Literal = literal
	ae.Tokens = tks
	return ae
}

func (p *Parser) parseEdgeFilter() *ast.EdgeFilter {
	ef := &ast.EdgeFilter{}
	if !p.expectPeek(token.LPAREN) { //(
		return nil
	}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	ef.FuncName = p.curToken.Literal
	ef.FuncParams = p.parseFilterParameters()
	if !p.expectPeek(token.RPAREN) { //)
		return nil
	}
	return ef
}
func (p *Parser) parseFilterParameters() []*ast.Identifier {
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	ids := []*ast.Identifier{}
	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return ids
	}
	p.nextToken()
	ident := &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	ids = append(ids, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
		ids = append(ids, ident)
	}
	if !p.expectPeek(token.RPAREN) {
		return nil
	}
	return ids
}
