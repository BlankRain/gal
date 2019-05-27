package compiler

import (
	"fmt"
	"runtime"

	"github.com/BlankRain/gal/ast"
	"github.com/BlankRain/gal/llvm/value"
	"github.com/BlankRain/gal/llvm/types"
	irr "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	llvmTypes "github.com/llir/llvm/ir/types"
)
type Compiler struct {
	module *ir.Module

	// functions provided by the OS, such as printf and malloc
	externalFuncs ExternalFuncs

	// functions provided by the language, such as println
	globalFuncs map[string]*types.Function

	packages           map[string]*types.PackageInstance
	currentPackage     *types.PackageInstance
	currentPackageName string

	contextFunc *types.Function

	// Stack of return values pointers, is used both used if a function returns more
	// than one value (arg pointers), and single stack based returns
	contextFuncRetVals [][]value.Value

	contextBlock *ir.Block

	// Stack of variables that are in scope
	contextBlockVariables []map[string]value.Value

	// What a break or continue should resolve to
	contextLoopBreak    []*ir.Block
	contextLoopContinue []*ir.Block

	// Where a condition should jump when done
	contextCondAfter []*ir.Block

	// What type the current assign operation is assigning to.
	// Is used when evaluating what type an integer constant should be.
	contextAssignDest []value.Value

	// Stack of Alloc instructions
	// Is used to decide if values should be stack or heap allocated
	contextAlloc []*parser.AllocNode

	stringConstants map[string]*ir.Global
}

var (
	i8  = types.I8
	i32 = types.I32
	i64 = types.I64
)

func NewCompiler() *Compiler {
	c := &Compiler{
		module:      ir.NewModule(),
		globalFuncs: make(map[string]*types.Function),

		packages: make(map[string]*types.PackageInstance),

		contextFuncRetVals: make([][]value.Value, 0),

		contextBlockVariables: make([]map[string]value.Value, 0),

		contextLoopBreak:    make([]*ir.Block, 0),
		contextLoopContinue: make([]*ir.Block, 0),
		contextCondAfter:    make([]*ir.Block, 0),

		contextAssignDest: make([]value.Value, 0),

		stringConstants: make(map[string]*ir.Global),
	}

	c.createExternalPackage()
	c.addGlobal()

	// Triple examples:
	// x86_64-apple-macosx10.13.0
	// x86_64-pc-linux-gnu
	var targetTriple [2]string

	switch runtime.GOARCH {
	case "amd64":
		targetTriple[0] = "x86_64"
	default:
		panic("unsupported GOARCH: " + runtime.GOARCH)
	}

	switch runtime.GOOS {
	case "darwin":
		targetTriple[1] = "apple-macosx10.13.0"
	case "linux":
		targetTriple[1] = "pc-linux-gnu"
	case "windows":
		targetTriple[1] = "pc-windows"
	default:
		panic("unsupported GOOS: " + runtime.GOOS)
	}

	c.module.TargetTriple = fmt.Sprintf("%s-%s", targetTriple[0], targetTriple[1])

	return c
}


func (c *Compiler)  GetIR() string {
	return module.String()
}
func (c *Compiler)  Compile(node ast.Node) {
	switch node := node.(type) {
	case *ast.Program:
		c.compileProgram(node)
	case *ast.ExpressionStatement:
		c.Compile(node.Expression)
	case *ast.IntegerLiteral:
		c.compileInteger(node)
	case *ast.Boolean:
		c.compileNativeBoolen(node)
	case *ast.PrefixExpression:
		c.compilePrefixExpression(node)
	case *ast.InfixExpression:
		c.compileInfixExpression(node)
	case *ast.BlockStatement:
		c.compileBlockStatement(node)
	case *ast.IfExpression:
		c.compileIfExpression(node)
	case *ast.ReturnStatement:
		c.compileReturnExpression(node)
	case *ast.FunctionLiteral:
		c.compileFunction(node)
	case *ast.LetStatement:
		c.compileLetStatement(node)
	case *ast.CallExpression:
		c.compileApplyFunction(node)
	case *ast.Identifier:
		c.compileIdentifier(node)
	case *ast.StringLiteral:
		//  &object.String{Value: node.Value}
		c.compileString(node)
	case *ast.ArrayLiteral:
		c.compileArray(node)
	case *ast.IndexExpression:
		c.compileIndexExpression(node)
	}
}

func  (c *Compiler) compileProgram(prog *ast.Program) {
	for _, statement := range prog.Statements {
		Compile(statement)
	}
}

func (c *Compiler)  compileInteger(node *ast.IntegerLiteral) value.Value {
	return value.Value{
		Value:      constant.NewInt(llvmTypes.64, node.Value),
		Type:       types.,
		IsVariable: false,
	}
}

func (c *Compiler)  compileNativeBoolen(node *ast.Boolean) {
	return value.Value{
		Value:      constant.NewInt(llvmTypes.I1, node.Value),
		Type:       types.Bool,
		IsVariable: false,
	}
}
func  (c *Compiler) compileInfixExpression(node *ast.InfixExpression) {
}
func  (c *Compiler) compileBlockStatement(node *ast.BlockStatement) {

}
func  (c *Compiler) compilePrefixExpression(Operator string) {

}

func  (c *Compiler) compileIfExpression(node *ast.IfExpression) {

}

func (c *Compiler)  compileReturnExpression(node *ast.ReturnStatement) {

}

func (c *Compiler)  compileFunction(node *ast.FunctionLiteral) {

}
func  (c *Compiler) compileLetStatement(node *ast.LetStatement) {

}
func  (c *Compiler) compileExpressions(node []ast.Expression) {

}
func (c *Compiler)  compileApplyFunction(node *ast.CallExpression) {

}

func  (c *Compiler) compileIdentifier(node *ast.Identifier) {

}
func  (c *Compiler) compileString(node *ast.StringLiteral) {
	var constString *ir.Global

	// Reuse the *ir.Global if it has already been created
	if reusedConst, ok := stringConstants[v.ValueStr]; ok {
		constString = reusedConst
	} else {
		constString = module.NewGlobalDef(strings.NextStringName(), strings.Constant(v.ValueStr))
		constString.Immutable = true
		stringConstants[v.ValueStr] = constString
	}

	alloc := contextBlock.NewAlloca(typeConvertMap["string"].LLVM())

	// Save length of the string
	lenItem := contextBlock.NewGetElementPtr(alloc, constant.NewInt(llvmTypes.I32, 0), constant.NewInt(llvmTypes.I32, 0))
	contextBlock.NewStore(constant.NewInt(llvmTypes.I64, int64(len(v.ValueStr))), lenItem)

	// Save i8* version of string
	strItem := contextBlock.NewGetElementPtr(alloc, constant.NewInt(llvmTypes.I32, 0), constant.NewInt(llvmTypes.I32, 1))
	contextBlock.NewStore(strings.Toi8Ptr(contextBlock, constString), strItem)

	return value.Value{
		Value:      contextBlock.NewLoad(alloc),
		Type:       types.String,
		IsVariable: false,
	}

}

func  (c *Compiler) compileArray(node *ast.ArrayLiteral) {

}
func  (c *Compiler) compileIndexExpression(node *ast.IndexExpression) {
}
