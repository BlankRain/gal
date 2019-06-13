package evaluator

import (
	"github.com/BlankRain/gal/ast"
	"github.com/BlankRain/gal/object"
)

func evalMakeLiteral(node *ast.MakeLiteral, env *object.Environment) object.Object {
	params := []*ast.Identifier{}
	for _, p := range node.Params {
		params = append(params, &ast.Identifier{
			Value: p.Value,
		})
	}
	val := &object.Function{Parameters: params, Env: env, Body: node.Body}
	env.Set(node.Name.Value, val)
	return val
}
