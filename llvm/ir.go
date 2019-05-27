package llvm

import (
	"fmt"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

const (
	ERROR_OBJ   = "COMPILE ERROR"
	INTEGER_OBJ = "INTEGER"
	NULL_OBJ    = "NULL"
)

type IRObjectType string
type IRObject interface {
	Type() IRObjectType
	IR() string
}
type IntegerObject struct {
	Value int64
}

func (io *IntegerObject) Type() IRObjectType {
	return INTEGER_OBJ
}
func (io *IntegerObject) IR() string {
	i := constant.NewInt(types.I32, io.Value)
	return fmt.Sprintf("%v", i)
}

type NULL struct {
}

func (io *NULL) Type() IRObjectType {
	return NULL_OBJ
}
func (io *NULL) IR() string {
	return "null"
}
