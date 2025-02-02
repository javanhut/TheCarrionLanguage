package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/javanhut/TheCarrionLanguage/src/ast"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ      = "INTEGER"
	FLOAT_OBJ        = "FLOAT"
	BOOLEAN_OBJ      = "BOOLEAN"
	NONE_OBJ         = "NONE"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	ARRAY_OBJ        = "ARRAY"
	BUILTIN_OBJ      = "BUILTIN"
	HASH_OBJ         = "HASH"
	TUPLE_OBJ        = "TUPLE"
	SPELLBOOK_OBJ    = "SPELLBOOK"
	INSTANCE_OBJ     = "INSTANCE"
	NAMESPACE_OBJ    = "NAMESPACE"
)

var NONE = &None{}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Float struct {
	Value float64
}

func (f *Float) Type() ObjectType { return FLOAT_OBJ }
func (f *Float) Inspect() string  { return fmt.Sprintf("%f", f.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type None struct {
	Value string
}

func (n *None) Type() ObjectType { return NONE_OBJ }
func (n *None) Inspect() string  { return fmt.Sprintf("%s", n.Value) }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

type Function struct {
	Parameters  []*ast.Parameter
	Body        *ast.BlockStatement
	Env         *Environment
	IsAbstract  bool
	IsPrivate   bool
	IsProtected bool
}

func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("spell(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elems := make([]string, 0, len(ao.Elements))
	for _, e := range ao.Elements {
		elems = append(elems, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("]")
	return out.String()
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

type HashPair struct {
	Key   Object
	Value Object
}
type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

type Hashable interface {
	HashKey() HashKey
}

type Tuple struct {
	Elements []Object
}

func (t *Tuple) Type() ObjectType { return TUPLE_OBJ }
func (t *Tuple) Inspect() string {
	var out bytes.Buffer
	elems := []string{}
	for _, e := range t.Elements {
		elems = append(elems, e.Inspect())
	}
	out.WriteString("(")
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString(")")
	return out.String()
}

// Update Object (object.go)
type Spellbook struct {
	Name       string
	Methods    map[string]*Function
	InitMethod *Function
	Inherits   *Spellbook
	Env        *Environment // Add environment to store the spellbook's scope
	IsArcane   bool
}

func (s *Spellbook) Type() ObjectType { return SPELLBOOK_OBJ }
func (s *Spellbook) Inspect() string {
	return fmt.Sprintf("<spellbook %s>", s.Name)
}

// Ensure Instance type implements Object
type Instance struct {
	Spellbook *Spellbook
	Env       *Environment
}

func (i *Instance) Type() ObjectType { return INSTANCE_OBJ }
func (i *Instance) Inspect() string  { return fmt.Sprintf("<instance of %s>", i.Spellbook.Name) }

// object/object.go

type Namespace struct {
	Env *Environment // Holds all exported members of the imported module
}

func (n *Namespace) Type() ObjectType { return "NAMESPACE" }
func (n *Namespace) Inspect() string  { return "<namespace>" }

type Stop struct{}

func (s *Stop) Type() ObjectType { return "STOP" }
func (s *Stop) Inspect() string  { return "stop" }

type Skip struct{}

func (s *Skip) Type() ObjectType { return "SKIP" }
func (s *Skip) Inspect() string  { return "skip" }

var (
	STOP = &Stop{}
	SKIP = &Skip{}
)
