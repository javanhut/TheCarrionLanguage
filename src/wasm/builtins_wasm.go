//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/javanhut/TheCarrionLanguage/src/object"
)

// WASM-incompatible builtins that return errors
var wasmUnsupportedBuiltins = map[string]*object.Builtin{
	"osRunCommand": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "osRunCommand is not available in the browser playground"}
		},
	},
	"osGetEnv": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "osGetEnv is not available in the browser playground"}
		},
	},
	"osSetEnv": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "osSetEnv is not available in the browser playground"}
		},
	},
	"osGetCwd": {
		Fn: func(args ...object.Object) object.Object {
			return &object.String{Value: "/playground"}
		},
	},
	"osChdir": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "osChdir is not available in the browser playground"}
		},
	},
	"osSleep": {
		Fn: func(args ...object.Object) object.Object {
			// In WASM, we can't do blocking sleep, just return immediately
			return &object.None{}
		},
	},
	"osListDir": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "osListDir is not available in the browser playground"}
		},
	},
	"osRemove": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "osRemove is not available in the browser playground"}
		},
	},
	"osMkdir": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "osMkdir is not available in the browser playground"}
		},
	},
	"osExpandEnv": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "osExpandEnv requires 1 argument"}
			}
			// Just return the string unchanged in WASM
			if str, ok := args[0].(*object.String); ok {
				return str
			}
			return &object.Error{Message: "osExpandEnv argument must be a string"}
		},
	},
	"fileRead": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "fileRead is not available in the browser playground"}
		},
	},
	"fileWrite": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "fileWrite is not available in the browser playground"}
		},
	},
	"fileAppend": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "fileAppend is not available in the browser playground"}
		},
	},
	"fileExists": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Boolean{Value: false}
		},
	},
	"input": {
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "input() is not available in the browser playground. Use the code editor instead."}
		},
	},
}

// GetWASMUnsupportedBuiltins returns the map of WASM-incompatible builtins
func GetWASMUnsupportedBuiltins() map[string]*object.Builtin {
	return wasmUnsupportedBuiltins
}
