//go:build js && wasm
// +build js,wasm

package main

import (
	"strings"
	"syscall/js"

	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
)

var globalEnv *object.Environment
var stdlibError string

func init() {
	// Register WASM-specific builtins that override incompatible ones
	registerWASMBuiltins()

	// Initialize the global environment
	globalEnv = object.NewEnvironment()

	// Load the standard library
	if err := evaluator.LoadMuninStdlib(globalEnv); err != nil {
		stdlibError = err.Error()
	}
}

// registerWASMBuiltins registers WASM-compatible versions of builtins
func registerWASMBuiltins() {
	// Override OS builtins with WASM-safe versions
	evaluator.OverrideBuiltin("osRunCommand", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "osRunCommand is not available in the browser playground"}
		},
	})

	evaluator.OverrideBuiltin("osGetEnv", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.String{Value: ""}
		},
	})

	evaluator.OverrideBuiltin("osSetEnv", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.None{}
		},
	})

	evaluator.OverrideBuiltin("osGetCwd", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.String{Value: "/playground"}
		},
	})

	evaluator.OverrideBuiltin("osChdir", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.None{}
		},
	})

	evaluator.OverrideBuiltin("osSleep", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// In WASM, we can't do blocking sleep
			return &object.None{}
		},
	})

	evaluator.OverrideBuiltin("osListDir", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.Array{Elements: []object.Object{}}
		},
	})

	evaluator.OverrideBuiltin("osRemove", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "osRemove is not available in the browser playground"}
		},
	})

	evaluator.OverrideBuiltin("osMkdir", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "osMkdir is not available in the browser playground"}
		},
	})

	evaluator.OverrideBuiltin("osExpandEnv", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "osExpandEnv requires 1 argument"}
			}
			if str, ok := args[0].(*object.String); ok {
				return str
			}
			return &object.Error{Message: "osExpandEnv argument must be a string"}
		},
	})

	// Override file builtins
	evaluator.OverrideBuiltin("fileRead", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "fileRead is not available in the browser playground"}
		},
	})

	evaluator.OverrideBuiltin("fileWrite", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "fileWrite is not available in the browser playground"}
		},
	})

	evaluator.OverrideBuiltin("fileAppend", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "fileAppend is not available in the browser playground"}
		},
	})

	evaluator.OverrideBuiltin("fileExists", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.Boolean{Value: false}
		},
	})

	// Override input - not supported in browser without async
	evaluator.OverrideBuiltin("input", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			return &object.Error{Message: "input() is not available in the browser playground"}
		},
	})
}

// evaluate runs Carrion code and returns the result
func evaluate(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return map[string]interface{}{
			"success": false,
			"error":   "No code provided",
			"output":  "",
		}
	}

	code := args[0].String()

	// Create output capture
	var outputBuilder strings.Builder

	// Set up output capture for print statements
	evaluator.SetOutputWriter(&outputBuilder)

	// Lexer
	l := lexer.New(code)

	// Parser
	p := parser.New(l)
	program := p.ParseProgram()

	// Check for parse errors
	if len(p.Errors()) > 0 {
		errorMsg := strings.Join(p.Errors(), "\n")
		return map[string]interface{}{
			"success": false,
			"error":   errorMsg,
			"output":  "",
		}
	}

	// Evaluate (pass nil for CallContext at top level)
	result := evaluator.Eval(program, globalEnv, nil)

	// Check for runtime errors
	if errObj, ok := result.(*object.Error); ok {
		return map[string]interface{}{
			"success": false,
			"error":   errObj.Message,
			"output":  outputBuilder.String(),
		}
	}

	// Get the final result
	resultStr := ""
	if result != nil && result.Type() != object.NONE_OBJ {
		resultStr = result.Inspect()
	}

	return map[string]interface{}{
		"success": true,
		"error":   "",
		"output":  outputBuilder.String(),
		"result":  resultStr,
	}
}

// resetEnvironment creates a fresh environment
func resetEnvironment(this js.Value, args []js.Value) interface{} {
	globalEnv = object.NewEnvironment()
	if err := evaluator.LoadMuninStdlib(globalEnv); err != nil {
		stdlibError = err.Error()
	} else {
		stdlibError = ""
	}
	return nil
}

// getVersion returns the Carrion version
func getVersion(this js.Value, args []js.Value) interface{} {
	return "0.1.9"
}

// getStdlibStatus returns the stdlib loading status
func getStdlibStatus(this js.Value, args []js.Value) interface{} {
	if stdlibError != "" {
		return map[string]interface{}{
			"loaded": false,
			"error":  stdlibError,
		}
	}
	return map[string]interface{}{
		"loaded": true,
		"error":  "",
	}
}

func main() {
	c := make(chan struct{}, 0)

	// Register JavaScript functions
	js.Global().Set("carrionEval", js.FuncOf(evaluate))
	js.Global().Set("carrionReset", js.FuncOf(resetEnvironment))
	js.Global().Set("carrionVersion", js.FuncOf(getVersion))
	js.Global().Set("carrionStdlibStatus", js.FuncOf(getStdlibStatus))

	// Signal that WASM is ready
	js.Global().Call("dispatchEvent", js.Global().Get("CustomEvent").New("carrionReady"))

	// Keep the Go runtime alive
	<-c
}
