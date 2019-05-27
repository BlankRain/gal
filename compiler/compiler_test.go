package compiler

import (
	"testing"

	"github.com/BlankRain/gal/lexer"
	ir "github.com/BlankRain/gal/llvm"
	"github.com/BlankRain/gal/parser"
)

func TestStringIR(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
		ir       string
	}{
		{"10", 10, "i32 10"},
		// {"if (true) { 10 }", 10},
		// {"if (false) { 10 }", nil},
		// {"if (1) { 10 }", 10},
		// {"if (1 < 2) { 10 }", 10},
		// {"if (1 > 2) { 10 }", nil}, {"if (1 > 2) { 10 } else { 20 }", 20},
		// {"if (1 < 2) { 10 } else { 20 }", 10},
	}
	for _, tt := range tests {
		evaluated := testCompile(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer), tt.ir)
		} else {
			testNullObject(t, evaluated)
		}
	}
}
func testCompile(input string) ir.IRObject {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	return Compile(program)
}

func testIntegerObject(t *testing.T, obj ir.IRObject, expected int64, expectedIr string) bool {
	result, ok := obj.(*ir.IntegerObject)
	if obj.IR() != expectedIr {
		t.Errorf("IRcode is %v, want %s", obj, expectedIr)
		return false
	}
	if !ok {
		t.Errorf("object is %T (%+v), want Integer", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("Value is %d, want %d", result.Value, expected)
		return false
	}

	return true
}

func testNullObject(t *testing.T, obj ir.IRObject) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}
