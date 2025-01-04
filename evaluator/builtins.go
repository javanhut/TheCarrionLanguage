package evaluator

import (
	"fmt"
	"strconv"

	"thecarrionlanguage/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"print": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Print(arg.Inspect(), " ")
			}
			return &object.None{}
		},
	},

	"type": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			obj := args[0]
			return &object.String{Value: string(obj.Type())}
		},
	},
	"int": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				value, err := strconv.Atoi(arg.Value)
				if err != nil {
					return newError("cannot convert string to int: %s", err)
				}
				return &object.Integer{Value: int64(value)}
			case *object.Float:
				return &object.Integer{Value: int64(arg.Value)}
			case *object.Integer:
				return arg
			default:
				return newError("cannot convert %s to int", arg.Type())
			}
		},
	},
	"float": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				value, err := strconv.ParseFloat(arg.Value, 64)
				if err != nil {
					return newError("cannot convert string to float: %s", err)
				}
				return &object.Float{Value: value}
			case *object.Integer:
				return &object.Float{Value: float64(arg.Value)}
			case *object.Float:
				return arg
			default:
				return newError("cannot convert %s to float", arg.Type())
			}
		},
	},
	"string": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			return &object.String{Value: args[0].Inspect()}
		},
	},
	"list": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				elements := make([]object.Object, len(arg.Value))
				for i, char := range arg.Value {
					elements[i] = &object.String{Value: string(char)}
				}
				return &object.Array{Elements: elements}
			case *object.Tuple:
				return &object.Array{Elements: arg.Elements}
			default:
				return newError("cannot convert %s to list", arg.Type())
			}
		},
	},
	"tuple": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Tuple{Elements: arg.Elements}
			case *object.Tuple:
				return arg
			default:
				return newError("cannot convert %s to tuple", arg.Type())
			}
		},
	},
}
