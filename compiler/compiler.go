package compiler

import (
	"fmt"

	"github.com/BlankRain/gal/ast"
	ir "github.com/BlankRain/gal/llvm"
)

var (
	NULL = &ir.NULL{}
)

func Compile(node ast.Node) ir.IRObject {
	switch node := node.(type) {
	case *ast.Program:
		return compileProgram(node)
	case *ast.ExpressionStatement:
		return Compile(node.Expression)
	case *ast.IntegerLiteral:
		return compileInteger(node)
	case *ast.Boolean:
		return compileNativeBoolen(node)
	case *ast.PrefixExpression:
		right := Compile(node.Right)
		if isError(right) {
			return right
		}
		return compilePrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Compile(node.Left)
		if isError(left) {
			return left
		}
		right := Compile(node.Right)
		if isError(right) {
			return right
		}
		return compileInfixExpression(node.Operator, left, right)
	case *ast.BlockStatement:
		return compileBlockStatement(node)
	case *ast.IfExpression:
		return compileIfExpression(node)
	case *ast.ReturnStatement:
		return compileReturnExpression(node)
	case *ast.FunctionLiteral:
		return compileFunction(node)
	case *ast.LetStatement:
		val := Compile(node.Value)
		if isError(val) {
			return val
		}
		// env.Set(node.Name.Value, val)
		compileLetStatement(node, val)
	case *ast.CallExpression:
		function := Compile(node.Function)
		if isError(function) {
			return function
		}
		args := compileExpressions(node.Arguments)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return compileApplyFunction(function, args)
	case *ast.Identifier:
		return compileIdentifier(node)
	case *ast.StringLiteral:
		// return &object.String{Value: node.Value}
		return compileString(node)
	case *ast.ArrayLiteral:
		elements := compileExpressions(node.Elements)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return compileArray(elements)
	case *ast.IndexExpression:
		left := Compile(node.Left)
		if isError(left) {
			return left
		}
		index := Compile(node.Index)
		if isError(index) {
			return index
		}
		return compileIndexExpression(left, index)
	}
	return nil
}

func isError(obj ir.IRObject) bool {
	if obj != nil {
		return obj.Type() == ir.ERROR_OBJ
	}
	return false
}

func compileProgram(prog *ast.Program) ir.IRObject {
	var ret ir.IRObject
	for _, statement := range prog.Statements {
		ret = Compile(statement)
	}
	return ret
}

func compileInteger(node *ast.IntegerLiteral) ir.IRObject {
	fmt.Println(node)
	return &ir.IntegerObject{Value: node.Value}
}

func compileNativeBoolen(node *ast.Boolean) ir.IRObject {
	return nil
}
func compileInfixExpression(
	Operator string,
	left, right ir.IRObject) ir.IRObject {
	return nil
}
func compileBlockStatement(node *ast.BlockStatement) ir.IRObject {
	return nil
}
func compilePrefixExpression(Operator string, right ir.IRObject) ir.IRObject {
	return nil
}

func compileIfExpression(node *ast.IfExpression) ir.IRObject {
	return nil
}

func compileReturnExpression(node *ast.ReturnStatement) ir.IRObject {
	return nil
}

func compileFunction(node *ast.FunctionLiteral) ir.IRObject {
	return nil
}
func compileLetStatement(node *ast.LetStatement, val ir.IRObject) ir.IRObject {
	return nil
}
func compileExpressions(node []ast.Expression) []ir.IRObject {
	return nil
}
func compileApplyFunction(function ir.IRObject, args []ir.IRObject) ir.IRObject {
	return nil
}

func compileIdentifier(node *ast.Identifier) ir.IRObject {
	return nil
}
func compileString(node *ast.StringLiteral) ir.IRObject {
	return nil
}

func compileArray(array []ir.IRObject) ir.IRObject {
	return nil
}
func compileIndexExpression(left, right ir.IRObject) ir.IRObject {
	return nil
}
