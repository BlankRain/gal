package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (i *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

func (i *Boolean) Inspect() string {
	return fmt.Sprintf("%t", i.Value)
}

type NULL struct {
}

func (i *NULL) Type() ObjectType { return NULL_OBJ }

func (i *NULL) Inspect() string {
	return "null"
}

type ReturnValue struct {
	Value Object
}

func (i *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

func (i *ReturnValue) Inspect() string {
	return i.Value.Inspect()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }
