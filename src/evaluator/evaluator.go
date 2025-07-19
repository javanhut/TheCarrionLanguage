package evaluator

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/javanhut/TheCarrionLanguage/src/ast"
	"github.com/javanhut/TheCarrionLanguage/src/debug"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
	"github.com/javanhut/TheCarrionLanguage/src/token"
)

// Debug flag for primitive wrapping debug output
var debugPrimitiveWrapping = os.Getenv("CARRION_DEBUG_WRAPPING") == "1"

var (
	NONE                        = &object.None{Value: "None"}
	TRUE                        = &object.Boolean{Value: true}
	FALSE                       = &object.Boolean{Value: false}
	importedFiles               = map[string]bool{}
	MAX_CALL_DEPTH              = 1000
	CurrentContext *CallContext = nil
)

// CallContext tracks function call state for better error reporting
type CallContext struct {
	FunctionName string
	Node         ast.Node
	Parent       *CallContext
	env          *object.Environment
	depth        int
	IsDirectExecution bool  // True when file is run directly, false when imported
	MethodGrimoire *object.Grimoire // The grimoire that owns the current method
}

// A map to track call stack depth for recursive functions
var callStack = make(map[*object.Function]*CallContext)

// Global map to track recursion depth per function (defined later in file)
var recursionDepths = make(map[*ast.BlockStatement]int)

// CleanupGlobalState clears all global state maps to prevent memory leaks
func CleanupGlobalState() {
	// Clear imported files
	importedFiles = make(map[string]bool)
	
	// Clear call stack
	for k := range callStack {
		delete(callStack, k)
	}
	
	// Clear recursion depths
	for k := range recursionDepths {
		delete(recursionDepths, k)
	}
	
	// Reset current context
	CurrentContext = nil
	
	// Cleanup goroutine manager
	CleanupGoroutineManager()
}

// CleanupGoroutineManager waits for goroutines to finish and resets the manager
func CleanupGoroutineManager() {
	
	// Create a channel to signal completion
	done := make(chan bool, 1)
	
	go func() {
		// Wait for all named goroutines to finish
		namedGoroutines := globalGoroutineManager.GetAllNamedGoroutines()
		for _, goroutine := range namedGoroutines {
			if goroutine.IsRunning {
				select {
				case <-goroutine.Done:
					// Goroutine finished normally
				case <-time.After(100 * time.Millisecond):
					// Continue to next goroutine after short timeout
				}
			}
		}
		
		// Wait for all anonymous goroutines to finish
		anonymousGoroutines := globalGoroutineManager.GetAllAnonymousGoroutines()
		for _, goroutine := range anonymousGoroutines {
			if goroutine.IsRunning {
				select {
				case <-goroutine.Done:
					// Goroutine finished normally
				case <-time.After(100 * time.Millisecond):
					// Continue to next goroutine after short timeout
				}
			}
		}
		
		done <- true
	}()
	
	// Wait for completion or timeout after 5 seconds
	select {
	case <-done:
		// All goroutines finished or timed out individually
	case <-time.After(5 * time.Second):
		// Global timeout reached
	}
	
	// Reset the global goroutine manager to a fresh state
	globalGoroutineManager.Reset()
}

// CleanupCallStack removes entries from call stack for specific function
func CleanupCallStack(fn *object.Function) {
	if fn != nil {
		delete(callStack, fn)
	}
}

// CleanupRecursionDepth removes recursion tracking for specific AST node
func CleanupRecursionDepth(node *ast.BlockStatement) {
	if node != nil {
		delete(recursionDepths, node)
	}
}

func getSourcePosition(node ast.Node) object.SourcePosition {
	pos := object.SourcePosition{
		Filename: "unknown",
		Line:     0,
		Column:   0,
	}

	token := getNodeToken(node)
	if token != nil {
		pos.Line = token.Line
		pos.Column = token.Column
		if token.Filename != "" {
			pos.Filename = token.Filename
		}
	}

	return pos
}

func getNodeToken(node ast.Node) *token.Token {
	switch n := node.(type) {
	case *ast.Program:
		// For programs, try to get filename from the first statement
		var filename string
		if len(n.Statements) > 0 {
			if firstToken := getNodeToken(n.Statements[0]); firstToken != nil {
				filename = firstToken.Filename
			}
		}
		return &token.Token{
			Type:     token.NEWLINE,
			Literal:  "",
			Line:     1,
			Column:   1,
			Filename: filename,
		}
	case *ast.ExpressionStatement:
		return getNodeToken(n.Expression)
	case *ast.IntegerLiteral:
		return &n.Token
	case *ast.FloatLiteral:
		return &n.Token
	case *ast.StringLiteral:
		return &n.Token
	case *ast.Boolean:
		return &n.Token
	case *ast.Identifier:
		return &n.Token
	case *ast.PrefixExpression:
		return &n.Token
	case *ast.InfixExpression:
		return &n.Token
	case *ast.PostfixExpression:
		return &n.Token
	case *ast.IfStatement:
		return &n.Token
	case *ast.BlockStatement:
		return &n.Token
	case *ast.FunctionDefinition:
		return &n.Token
	case *ast.CallExpression:
		return &n.Token
	case *ast.ReturnStatement:
		return &n.Token
	case *ast.AssignStatement:
		return &n.Token
	case *ast.DotExpression:
		return &n.Token
	case *ast.IndexExpression:
		return &n.Token
	case *ast.ForStatement:
		return &n.Token
	case *ast.WhileStatement:
		return &n.Token
	case *ast.ArrayLiteral:
		return &n.Token
	case *ast.HashLiteral:
		return &n.Token
	case *ast.TupleLiteral:
		return &n.Token
	case *ast.MatchStatement:
		return &n.Token
	case *ast.GrimoireDefinition:
		return &n.Token
	case *ast.AttemptStatement:
		return &n.Token
	case *ast.WithStatement:
		return &n.Token
	case *ast.MainStatement:
		return &n.Token
	case *ast.RaiseStatement:
		return &n.Token
	case *ast.IgnoreStatement:
		return &n.Token
	case *ast.StopStatement:
		return &n.Token
	case *ast.SkipStatement:
		return &n.Token
	case *ast.DivergeStatement:
		return &n.Token
	case *ast.ConvergeStatement:
		return &n.Token
	case *ast.CheckStatement:
		return &n.Token
	case *ast.ElseStatement:
		return &n.Token
	case *ast.GlobalStatement:
		return &n.Token
	case *ast.SliceExpression:
		return &n.Token
	case *ast.WildcardExpression:
		return &n.Token
	default:
		// Return a synthetic token for unknown nodes to prevent nil pointer issues
		return &token.Token{
			Type:     token.ILLEGAL,
			Literal:  "unknown",
			Line:     1,
			Column:   1,
			Filename: "",
		}
	}
}

func newErrorWithTrace(
	format string,
	node ast.Node,
	ctx *CallContext,
	args ...interface{},
) object.Object {
	pos := getSourcePosition(node)
	message := fmt.Sprintf(format, args...)
	err := &object.ErrorWithTrace{
		ErrorType:     object.ERROR_OBJ,
		Message:       message,
		Position:      pos,
		Stack:         []object.StackTraceEntry{},
		CustomDetails: make(map[string]object.Object),
	}

	// Build stack trace from context
	currentCtx := ctx
	for currentCtx != nil {
		if currentCtx.Node != nil {
			nodePos := getSourcePosition(currentCtx.Node)
			
			// If current node has unknown filename, try to get it from parent context
			if (nodePos.Filename == "unknown" || nodePos.Filename == "") && currentCtx.Parent != nil && currentCtx.Parent.Node != nil {
				parentPos := getSourcePosition(currentCtx.Parent.Node)
				if parentPos.Filename != "unknown" && parentPos.Filename != "" {
					nodePos.Filename = parentPos.Filename
				}
			}
			
			entry := object.StackTraceEntry{
				FunctionName: currentCtx.FunctionName,
				Position:     nodePos,
			}
			
			// Skip duplicate consecutive entries with same function name and unknown location
			shouldSkip := false
			if len(err.Stack) > 0 {
				lastEntry := err.Stack[len(err.Stack)-1]
				if lastEntry.FunctionName == entry.FunctionName && 
				   (entry.Position.Filename == "unknown" || entry.Position.Filename == "") &&
				   (lastEntry.Position.Filename != "unknown" && lastEntry.Position.Filename != "") {
					shouldSkip = true
				}
			}
			
			if !shouldSkip {
				err.Stack = append(err.Stack, entry)
			}
		}
		currentCtx = currentCtx.Parent
	}

	return err
}

// newCustomErrorWithTrace creates a custom error with trace information
func newCustomErrorWithTrace(
	name string,
	message string,
	node ast.Node,
	ctx *CallContext,
	details map[string]object.Object,
) object.Object {
	pos := getSourcePosition(node)
	err := &object.ErrorWithTrace{
		ErrorType:     object.CUSTOM_ERROR_OBJ, // Use ErrorType instead of Type
		Message:       fmt.Sprintf("%s: %s", name, message),
		Position:      pos,
		Stack:         []object.StackTraceEntry{},
		CustomDetails: details,
	}
	
	// Debug: ensure details is not nil
	if err.CustomDetails == nil {
		err.CustomDetails = make(map[string]object.Object)
	}

	// Build stack trace from context
	currentCtx := ctx
	for currentCtx != nil {
		if currentCtx.Node != nil {
			entry := object.StackTraceEntry{
				FunctionName: currentCtx.FunctionName,
				Position:     getSourcePosition(currentCtx.Node),
			}
			err.Stack = append(err.Stack, entry)
		}
		currentCtx = currentCtx.Parent
	}

	return err
}

func newError(format string, args ...interface{}) object.Object {
	msg := fmt.Sprintf(format, args...)
	err := &object.Error{Message: msg}

	// For better compatibility with stack traces, convert to ErrorWithTrace
	// when we're inside the evaluator
	if CurrentContext != nil {
		return newErrorWithTrace(msg, CurrentContext.Node, CurrentContext)
	}

	return err
}

func isPrimitiveLiteral(obj object.Object) bool {
	switch obj.Type() {
	case object.INTEGER_OBJ, object.FLOAT_OBJ, object.STRING_OBJ, object.BOOLEAN_OBJ:
		return true
	default:
		return false
	}
}

func areValuesEqual(a, b object.Object) bool {
	if a.Type() != b.Type() {
		return false
	}

	switch objA := a.(type) {
	case *object.Integer:
		objB := b.(*object.Integer)
		return objA.Value == objB.Value
	case *object.String:
		objB := b.(*object.String)
		return objA.Value == objB.Value
	case *object.Boolean:
		objB := b.(*object.Boolean)
		return objA.Value == objB.Value
	case *object.Float:
		objB := b.(*object.Float)
		const epsilon = 1e-9
		return math.Abs(objA.Value-objB.Value) < epsilon
	default:
		return a.Inspect() == b.Inspect()
	}
}

func wrapPrimitive(obj object.Object, env *object.Environment, ctx *CallContext) object.Object {
	if debugPrimitiveWrapping {
		fmt.Fprintf(os.Stderr, "WRAP: Evaluating %T, ctx=%s, hasSelf=%t\n", obj, getContextName(ctx), hasSelfInEnv(env))
	}
	
	// Don't wrap if we're inside an instance method or initializer
	// Check current environment and traverse up parent environments
	if hasSelfInEnv(env) {
		if debugPrimitiveWrapping {
			fmt.Fprintf(os.Stderr, "WRAP: Found self in environment, not wrapping\n")
		}
		return obj
	}
	
	// Don't wrap if we're in a method context (check the call context chain)
	if isInMethodContext(ctx) {
		if debugPrimitiveWrapping {
			fmt.Fprintf(os.Stderr, "WRAP: In method context %s, not wrapping\n", getContextName(ctx))
		}
		return obj
	}
	
	// Don't wrap if this is in a function call context that expects primitives
	if ctx != nil && ctx.FunctionName != "" {
		// Don't wrap arguments to builtin functions (except for input functions that should return String instances and pairs() that should return Array instances)
		if isBuiltinFunction(ctx.FunctionName) && !shouldWrapStringResult(ctx.FunctionName) && ctx.FunctionName != "pairs" {
			if debugPrimitiveWrapping {
				fmt.Fprintf(os.Stderr, "WRAP: Builtin function %s, not wrapping\n", ctx.FunctionName)
			}
			return obj
		}
		// Don't wrap arguments to grimoire constructors
		if isGrimoireConstructor(ctx.FunctionName, env) {
			if debugPrimitiveWrapping {
				fmt.Fprintf(os.Stderr, "WRAP: Grimoire constructor %s, not wrapping\n", ctx.FunctionName)
			}
			return obj
		}
	}
	
	if debugPrimitiveWrapping {
		fmt.Fprintf(os.Stderr, "WRAP: Wrapping %T in context %s\n", obj, getContextName(ctx))
	}
	
	var grimName string

	switch obj.Type() {
	case object.INTEGER_OBJ:
		grimName = "Integer"
	case object.FLOAT_OBJ:
		grimName = "Float"
	case object.STRING_OBJ:
		grimName = "String"
	case object.BOOLEAN_OBJ:
		grimName = "Boolean"
	case object.ARRAY_OBJ:
		grimName = "Array"
	default:
		return obj // Not a primitive, return as is
	}

	// Try to find the grimoire
	if grimObj, ok := env.Get(grimName); ok {
		if grimoire, isGrim := grimObj.(*object.Grimoire); isGrim {
			// Create instance exactly like the normal grimoire constructor
			instance := &object.Instance{
				Grimoire: grimoire,
				Env:      object.NewEnclosedEnvironment(grimoire.Env),
			}

			// Set self reference
			instance.Env.Set("self", instance)
			
			// Handle different types appropriately
			if grimName == "Array" {
				// For arrays, set the elements directly
				if arrayObj, isArray := obj.(*object.Array); isArray {
					instance.Env.Set("elements", arrayObj)
				}
			} else {
				// For other primitives, set the value
				instance.Env.Set("value", obj)
			}

			return instance
		}
	}

	// If grimoire not found, return the original object
	return obj
}

// unwrapPrimitive extracts the primitive value from a wrapped instance if applicable
func unwrapPrimitive(obj object.Object) object.Object {
	if instance, ok := obj.(*object.Instance); ok {
		// Check for primitive values wrapped in "value" field (String, Integer, Float, Boolean)
		if value, exists := instance.Env.Get("value"); exists {
			return value
		}
		// Check for Array instances which store data in "elements" field
		if instance.Grimoire.Name == "Array" {
			if elements, exists := instance.Env.Get("elements"); exists {
				// If elements is a direct Array, return it
				if arr, isArray := elements.(*object.Array); isArray {
					return arr
				}
				// If elements is another instance, try to unwrap it recursively
				if elemInstance, isInstance := elements.(*object.Instance); isInstance {
					if innerValue, innerExists := elemInstance.Env.Get("value"); innerExists {
						if arr, isArray := innerValue.(*object.Array); isArray {
							return arr
						}
					}
				}
			}
		}
	}
	return obj
}

// isBuiltinFunction checks if a function name is a builtin
func isBuiltinFunction(name string) bool {
	_, isBuiltin := builtins[name]
	return isBuiltin
}

// isGrimoireConstructor checks if a function name is a grimoire constructor
func isGrimoireConstructor(name string, env *object.Environment) bool {
	if obj, exists := env.Get(name); exists {
		_, isGrimoire := obj.(*object.Grimoire)
		return isGrimoire
	}
	return false
}

// shouldWrapStringResult determines if string results from a function should be wrapped in String grimoire instances
func shouldWrapStringResult(functionName string) bool {
	switch functionName {
	case "input", "fileRead", "osGetEnv", "osExpandEnv", "chr":
		return true
	default:
		return false
	}
}

// hasSelfInEnv checks if 'self' exists in the environment hierarchy
func hasSelfInEnv(env *object.Environment) bool {
	_, hasSelf := env.Get("self")
	return hasSelf
}

// isInMethodContext checks if the current context is inside a grimoire method
func isInMethodContext(ctx *CallContext) bool {
	// Traverse the context chain to find any method context
	current := ctx
	for current != nil {
		if current.FunctionName != "" && strings.Contains(current.FunctionName, ".") {
			return true
		}
		current = current.Parent
	}
	return false
}

// getContextName returns a debug-friendly context name
func getContextName(ctx *CallContext) string {
	if ctx == nil {
		return "nil"
	}
	if ctx.FunctionName == "" {
		return "unnamed"
	}
	return ctx.FunctionName
}

func isErrorWithTrace(obj object.Object) bool {
	_, ok := obj.(*object.ErrorWithTrace)
	return ok
}

// Main evaluation function
func Eval(node ast.Node, env *object.Environment, ctx *CallContext) object.Object {
	// Add nil parameter validation
	if node == nil {
		return &object.Error{Message: "cannot evaluate nil node"}
	}
	if env == nil {
		return newErrorWithTrace("environment cannot be nil", node, ctx)
	}
	// Add ctx validation
	if ctx == nil {
		ctx = &CallContext{
			FunctionName: "<unknown>",
			Node:         node,
			Parent:       nil,
			env:          env,
		}
	}
	
	// Debug AST node types for // operations
	if debugPrimitiveWrapping {
		if infixNode, ok := node.(*ast.InfixExpression); ok && infixNode.Operator == "//" {
			fmt.Fprintf(os.Stderr, "EVAL: InfixExpression // in %s\n", getContextName(ctx))
		}
	}
	
	oldContext := CurrentContext
	CurrentContext = ctx
	defer func() { CurrentContext = oldContext }()
	// Create a new call context if node is a function call
	if callExp, ok := node.(*ast.CallExpression); ok {
		funcName := ""
		if ident, ok := callExp.Function.(*ast.Identifier); ok {
			funcName = ident.Value
		} else {
			funcName = "<anonymous function>"
		}

		newCtx := &CallContext{
			FunctionName: funcName,
			Node:         node,
			Parent:       ctx,
			env:          env,
			MethodGrimoire: ctx.MethodGrimoire, // Inherit from parent
		}
		ctx = newCtx
	}

	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env, ctx)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env, ctx)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env, ctx)
	case *ast.IfStatement:
		return evalIfExpression(node, env, ctx)
	case *ast.MainStatement:
		return evalBlockStatement(node.Body, env, ctx)

	case *ast.StopStatement:
		return object.STOP
	case *ast.SkipStatement:
		return object.SKIP
	case *ast.DivergeStatement:
		return evalDivergeStatement(node, env, ctx)
	case *ast.ConvergeStatement:
		return evalConvergeStatement(node, env, ctx)
	case *ast.CheckStatement:
		cond := Eval(node.Condition, env, ctx)
		if isError(cond) {
			return cond
		}
		
		// Special handling for check(value, expected) pattern
		if node.Message != nil {
			expected := Eval(node.Message, env, ctx)
			if isError(expected) {
				return expected
			}
			
			// Check if this looks like check(actual, expected)
			if cond.Type() == expected.Type() {
				if !areValuesEqual(cond, expected) {
					msg := fmt.Sprintf("Value %s didn't Match Value %s, Expected %s to Equal %s got %s instead", cond.Inspect(), expected.Inspect(), cond.Inspect(), expected.Inspect(), cond.Inspect())
					details := make(map[string]object.Object)
					details["actual"] = cond
					details["expected"] = expected
					return newCustomErrorWithTrace("Assertion Check Failed", msg, node, ctx, details)
				}
				return object.NONE
			} else {
				// Types don't match - this is also an assertion failure
				msg := fmt.Sprintf("Value %s didn't Match Value %s, Expected %s to Equal %s got %s instead", cond.Inspect(), expected.Inspect(), cond.Inspect(), expected.Inspect(), cond.Inspect())
				details := make(map[string]object.Object)
				details["actual"] = cond
				details["expected"] = expected
				return newCustomErrorWithTrace("Assertion Check Failed", msg, node, ctx, details)
			}
		}
		
		// Standard boolean check
		if !isTruthy(cond) {
			msg := "Assertion failed: " + node.Condition.String()
			if node.Message != nil {
				m := Eval(node.Message, env, ctx)
				if !isError(m) {
					msg = m.Inspect()
				}
			}

			details := make(map[string]object.Object)
			details["expression"] = &object.String{Value: node.Condition.String()}
			return newCustomErrorWithTrace("Assertion Check Failed", msg, node, ctx, details)
		}
		return object.NONE

	case *ast.PrefixExpression:
		if node.Operator == "++" || node.Operator == "--" {
			return evalPrefixIncrementDecrement(node.Operator, node, env, ctx)
		}
		right := Eval(node.Right, env, ctx)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, node, env, ctx)

	case *ast.InfixExpression:
		if node.Operator == "+=" || node.Operator == "-=" ||
			node.Operator == "*=" || node.Operator == "/=" {
			return evalCompoundAssignment(node, env, ctx)
		}

		if node.Operator == "and" {
			left := Eval(node.Left, env, ctx)
			if isError(left) {
				return left
			}
			if !isTruthy(left) {
				return left
			}
			return Eval(node.Right, env, ctx)
		}

		if node.Operator == "or" {
			left := Eval(node.Left, env, ctx)
			if isError(left) {
				return left
			}
			if isTruthy(left) {
				return left
			}
			return Eval(node.Right, env, ctx)
		}

		right := Eval(node.Right, env, ctx)
		if isError(right) {
			if isErrorWithTrace(right) {
				return right
			}
			if errorObj, ok := right.(*object.Error); ok {
				return newErrorWithTrace("Error in right operand: %s", node, ctx, errorObj.Message)
			}
			return newErrorWithTrace("Error in right operand", node, ctx)
		}

		left := Eval(node.Left, env, ctx)
		if isError(left) {
			if isErrorWithTrace(left) {
				return left
			}
			if errorObj, ok := left.(*object.Error); ok {
				return newErrorWithTrace("Error in left operand: %s", node, ctx, errorObj.Message)
			}
			return newErrorWithTrace("Error in left operand", node, ctx)
		}

		if debugPrimitiveWrapping && node.Operator == "//" {
			fmt.Fprintf(os.Stderr, "CALLING evalInfixExpression for %s in %s\n", node.Operator, getContextName(ctx))
		}
		result := evalInfixExpression(node.Operator, left, right, node, ctx)
		return result

	case *ast.PostfixExpression:
		return evalPostfixIncrementDecrement(node.Operator, node, env, ctx)

	case *ast.IntegerLiteral:
		primitive := &object.Integer{Value: node.Value}
		return wrapPrimitive(primitive, env, ctx)
	case *ast.FloatLiteral:
		primitive := &object.Float{Value: node.Value}
		return wrapPrimitive(primitive, env, ctx)
	case *ast.FStringLiteral:
		return evalFStringLiteral(node, env, ctx)
	case *ast.StringInterpolation:
		return evalStringInterpolation(node, env, ctx)
	case *ast.NoneLiteral:
		return object.NONE
	case *ast.WildcardExpression:
		return &object.String{Value: "_"}
	case *ast.ReturnStatement:
		var val object.Object = NONE
		if node.ReturnValue != nil {
			val = Eval(node.ReturnValue, env, ctx)
			if isError(val) {
				return val
			}
		}
		return &object.ReturnValue{Value: val}
	case *ast.Boolean:
		primitive := nativeBoolToBooleanObject(node.Value)
		return wrapPrimitive(primitive, env, ctx)
	case *ast.AssignStatement:
		return evalAssignStatement(node, env, ctx)
	case *ast.GlobalStatement:
		return evalGlobalStatement(node, env, ctx)
	case *ast.WhileStatement:
		return evalWhileStatement(node, env, ctx)
	case *ast.ForStatement:
		return evalForStatement(node, env, ctx)
	case *ast.ImportStatement:
		return evalImportStatement(node, env, ctx)
	case *ast.MatchStatement:
		return evalMatchStatement(node, env, ctx)
	case *ast.RaiseStatement:
		return evalRaiseStatement(node, env, ctx)
	case *ast.ArcaneGrimoire:
		return evalArcaneGrimoire(node, env, ctx)
	case *ast.Identifier:
		return evalIdentifier(node, env, ctx)
	case *ast.ArrayLiteral:
		elements := evalExpressions(node.Elements, env, ctx)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		
		// Create Array instance if Array grimoire exists
		if grimObj, ok := env.Get("Array"); ok {
			if grimoire, isGrim := grimObj.(*object.Grimoire); isGrim {
				instance := &object.Instance{
					Grimoire: grimoire,
					Env:      object.NewEnclosedEnvironment(grimoire.Env),
				}
				// Set the elements directly in the instance environment
				instance.Env.Set("elements", &object.Array{Elements: elements})
				return instance
			}
		}
		
		// Fallback to regular array
		return &object.Array{Elements: elements}

	case *ast.StringLiteral:
		primitive := &object.String{Value: node.Value}
		return wrapPrimitive(primitive, env, ctx)
	case *ast.TupleLiteral:
		return evalTupleLiteral(node, env, ctx)
	case *ast.HashLiteral:
		return evalHashLiteral(node, env, ctx)
	case *ast.FunctionDefinition:
		fnObj := &object.Function{
			Parameters: node.Parameters,
			ReturnType: node.ReturnType,
			Body:       node.Body,
			Env:        env,
		}
		env.Set(node.Name.Value, fnObj)
		return fnObj
	case *ast.DotExpression:
		return evalDotExpression(node, env, ctx)
	case *ast.IndexExpression:
		left := Eval(node.Left, env, ctx)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env, ctx)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index, node, ctx)
	case *ast.SliceExpression:
		left := Eval(node.Left, env, ctx)
		if isError(left) {
			return left
		}
		var start, end object.Object
		if node.Start != nil {
			start = Eval(node.Start, env, ctx)
			if isError(start) {
				return start
			}
		}
		if node.End != nil {
			end = Eval(node.End, env, ctx)
			if isError(end) {
				return end
			}
		}
		return evalSliceExpression(left, start, end, node, ctx)
	case *ast.GrimoireDefinition:
		return evalGrimoireDefinition(node, env, ctx)
	case *ast.AttemptStatement:
		return evalAttemptStatement(node, env, ctx)
	case *ast.WithStatement:
		return evalWithStatement(node, env, ctx)
	case *ast.UnpackStatement:
		return evalUnpackStatement(node, env, ctx)
	case *ast.IgnoreStatement:
		return object.NONE
	case *ast.CallExpression:
		fnObj := Eval(node.Function, env, ctx)
		if isError(fnObj) {
			return fnObj
		} // ← new

		argObjs := evalExpressions(node.Arguments, env, ctx)
		if err := promoteErrors(argObjs); err != nil {
			return err // ← new (optional helper)
		}

		return evalCallExpression(fnObj, argObjs, env, ctx)
	}

	return NONE
}

// EvalWithDebug evaluates an AST node with debug output
func EvalWithDebug(node ast.Node, env *object.Environment, ctx *CallContext, debugConfig *debug.Config) object.Object {
	// If no context provided, create one for direct execution
	if ctx == nil {
		ctx = &CallContext{
			FunctionName: "main",
			Node:         node,
			Parent:       nil,
			env:          env,
			depth:        0,
			IsDirectExecution: true,
		}
	}
	
	if debugConfig == nil || !debugConfig.ShouldDebugEvaluator() {
		return Eval(node, env, ctx)
	}
	
	// Log the node being evaluated
	fmt.Fprintf(os.Stderr, "evaluator: Evaluating %T", node)
	if n, ok := node.(ast.Node); ok && n != nil {
		fmt.Fprintf(os.Stderr, " - %s", n.String())
	}
	fmt.Fprintf(os.Stderr, "\n")
	
	// Evaluate the node
	result := Eval(node, env, ctx)
	
	// Log the result
	if result != nil {
		fmt.Fprintf(os.Stderr, "evaluator: Result: %s (type: %s)\n", result.Inspect(), result.Type())
	} else {
		fmt.Fprintf(os.Stderr, "evaluator: Result: nil\n")
	}
	
	return result
}

func promoteErrors(objs []object.Object) object.Object {
	for _, o := range objs {
		if isError(o) {
			return o
		}
	}
	return nil
}

func evalStringInterpolation(
	si *ast.StringInterpolation,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	var sb strings.Builder

	for _, part := range si.Parts {
		switch p := part.(type) {
		case *ast.StringText:
			sb.WriteString(p.Value)
		case *ast.StringExpr:
			val := Eval(p.Expr, env, ctx)
			if isError(val) {
				return val
			}

			// Format the value according to the format specifier
			formattedVal := formatValue(val, p)
			sb.WriteString(formattedVal)
		}
	}

	return &object.String{Value: sb.String()}
}

// Helper to format values according to specified format
func formatValue(val object.Object, exprPart *ast.StringExpr) string {
	var formatted string

	// Basic conversion to string
	switch obj := val.(type) {
	case *object.Integer:
		formatted = strconv.FormatInt(obj.Value, 10)
	case *object.Float:
		if exprPart.Precision > 0 {
			formatted = strconv.FormatFloat(obj.Value, 'f', exprPart.Precision, 64)
		} else {
			formatted = strconv.FormatFloat(obj.Value, 'f', -1, 64)
		}
	case *object.Boolean:
		formatted = strconv.FormatBool(obj.Value)
	case *object.String:
		formatted = obj.Value
	case *object.None:
		formatted = "None"
	default:
		formatted = val.Inspect()
	}

	// Apply width and alignment
	if exprPart.Width > 0 {
		fillChar := ' '
		if exprPart.FillChar != 0 {
			fillChar = rune(exprPart.FillChar)
		}

		padding := exprPart.Width - len(formatted)
		if padding <= 0 {
			return formatted // No padding needed
		}

		switch exprPart.Alignment {
		case '<': // Left align
			return formatted + strings.Repeat(string(fillChar), padding)
		case '>': // Right align
			return strings.Repeat(string(fillChar), padding) + formatted
		case '^': // Center align
			leftPad := padding / 2
			rightPad := padding - leftPad
			return strings.Repeat(
				string(fillChar),
				leftPad,
			) + formatted + strings.Repeat(
				string(fillChar),
				rightPad,
			)
		default: // Default to left align
			return formatted + strings.Repeat(string(fillChar), padding)
		}
	}

	return formatted
}

func evalFStringLiteral(
	fslit *ast.FStringLiteral,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	var sb strings.Builder

	for _, part := range fslit.Parts {
		switch p := part.(type) {
		case *ast.FStringText:
			sb.WriteString(p.Value)
		case *ast.FStringExpr:
			val := Eval(p.Expr, env, ctx)
			if isError(val) {
				return val
			}
			sb.WriteString(val.Inspect())
		}
	}

	return &object.String{Value: sb.String()}
}

func evalArcaneGrimoire(
	node *ast.ArcaneGrimoire,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	methods := make(map[string]*object.Function)

	for _, method := range node.Methods {
		methods[method.Name.Value] = &object.Function{
			Parameters: method.Parameters,
			Body:       method.Body,
			Env:        env,
		}
	}

	grimoire := &object.Grimoire{
		Name:     node.Name.Value,
		Methods:  methods,
		Env:      env,
		IsArcane: true,
	}

	env.Set(node.Name.Value, grimoire)
	return grimoire
}

func evalRaiseStatement(
	node *ast.RaiseStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	errObj := Eval(node.Error, env, ctx)
	if isError(errObj) {
		return errObj
	}

	if instance, ok := errObj.(*object.Instance); ok {
		message := ""
		if msg, ok := instance.Env.Get("message"); ok {
			if msgStr, ok := msg.(*object.String); ok {
				message = msgStr.Value
			}
		}
		
		// Check for String wrapper instances
		if instance.Grimoire.Name == "String" {
			if value, ok := instance.Env.Get("value"); ok {
				if strVal, ok := value.(*object.String); ok {
					message = strVal.Value
				}
			}
		}

		details := make(map[string]object.Object)
		details["errorType"] = &object.String{Value: instance.Grimoire.Name}
		details["instance"] = instance

		return newCustomErrorWithTrace(instance.Grimoire.Name, message, node, ctx, details)
	}

	if str, ok := errObj.(*object.String); ok {
		details := make(map[string]object.Object)
		details["errorType"] = &object.String{Value: "String"}
		details["instance"] = str
		return newCustomErrorWithTrace("Error", str.Value, node, ctx, details)
	}

	return newErrorWithTrace("cannot raise non-error object: %s", node, ctx, getObjectTypeString(errObj))
}

func evalExpressions(
	exps []ast.Expression,
	env *object.Environment,
	ctx *CallContext,
) []object.Object {
	result := make([]object.Object, 0, len(exps))

	for _, e := range exps {
		if e == nil {
			return []object.Object{&object.Error{Message: "cannot evaluate nil expression"}}
		}
		evaluated := Eval(e, env, ctx)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalAttemptStatement(
	node *ast.AttemptStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	var result object.Object

	// Create a new context for the try block
	tryCtx := &CallContext{
		FunctionName: "attempt",
		Node:         node.TryBlock,
		Parent:       ctx,
		env:          env,
		MethodGrimoire: ctx.MethodGrimoire, // Inherit from parent
	}
	tryResult := Eval(node.TryBlock, env, tryCtx)

	if isError(tryResult) {
		for _, ensnare := range node.EnsnareClauses {
			var shouldCatch bool
			var ensnareEnv *object.Environment
			
			// Check if this ensnare clause should catch the error
			if ensnare.Alias != nil {
				// If there's an alias, catch all errors and bind the error to the alias
				shouldCatch = true
				ensnareEnv = object.NewEnclosedEnvironment(env)
				// Bind the error to the alias, wrapped to prevent propagation
				if isError(tryResult) {
					caughtError := &object.CaughtError{OriginalError: tryResult}
					ensnareEnv.Set(ensnare.Alias.Value, caughtError)
				} else {
					// Handle non-error case - this shouldn't happen but is defensive
					ensnareEnv.Set(ensnare.Alias.Value, tryResult)
				}
			} else if ensnare.Condition != nil {
				// If there's a condition, evaluate it to check if we should catch
				condition := Eval(ensnare.Condition, env, ctx)
				if isError(condition) {
					// Check if this is an "identifier not found" error, which means it should be treated as an alias
					if errWithTrace, ok := condition.(*object.ErrorWithTrace); ok {
						if strings.Contains(errWithTrace.Message, "identifier not found") {
							// Treat this as an alias - catch all errors and bind to the identifier
							if identifier, ok := ensnare.Condition.(*ast.Identifier); ok {
								shouldCatch = true
								ensnareEnv = object.NewEnclosedEnvironment(env)
								if isError(tryResult) {
									caughtError := &object.CaughtError{OriginalError: tryResult}
									ensnareEnv.Set(identifier.Value, caughtError)
								} else {
									ensnareEnv.Set(identifier.Value, tryResult)
								}
							}
						} else {
							result = condition
							break
						}
					} else {
						result = condition
						break
					}
				} else {
					// Check different error types
					if customErr, ok := tryResult.(*object.CustomError); ok {
						if grimoire, ok := condition.(*object.Grimoire); ok {
							shouldCatch = customErr.ErrorType == grimoire
						} else if str, ok := condition.(*object.String); ok {
							shouldCatch = customErr.Name == str.Value
						}
					} else if errWithTrace, ok := tryResult.(*object.ErrorWithTrace); ok {
						if grimoire, ok := condition.(*object.Grimoire); ok {
							// Check if the error type matches the grimoire name
							if errWithTrace.CustomDetails != nil {
								if errorType, exists := errWithTrace.CustomDetails["errorType"]; exists {
									if errorTypeStr, ok := errorType.(*object.String); ok {
										shouldCatch = errorTypeStr.Value == grimoire.Name
									}
								}
							}
						} else if str, ok := condition.(*object.String); ok {
							// Check if the error type matches the string or if message starts with it
							if errWithTrace.CustomDetails != nil {
								if errorType, exists := errWithTrace.CustomDetails["errorType"]; exists {
									if errorTypeStr, ok := errorType.(*object.String); ok {
										shouldCatch = errorTypeStr.Value == str.Value
									}
								}
							}
							if !shouldCatch {
								shouldCatch = strings.HasPrefix(errWithTrace.Message, str.Value)
							}
						}
					}
				}
				if ensnareEnv == nil {
					ensnareEnv = env
				}
			} else {
				// No condition or alias, catch all errors
				shouldCatch = true
				ensnareEnv = env
			}
			
			if shouldCatch {
				ensnareCtx := &CallContext{
					FunctionName: "ensnare",
					Node:         ensnare.Consequence,
					Parent:       ctx,
					env:          ensnareEnv,
				}
				result = Eval(ensnare.Consequence, ensnareEnv, ensnareCtx)
				break
			}
		}

		if result == nil {
			result = tryResult
		}
	} else {
		result = tryResult
	}

	if node.ResolveBlock != nil {
		resolveCtx := &CallContext{
			FunctionName: "resolve",
			Node:         node.ResolveBlock,
			Parent:       ctx,
			env:          env,
		}
		resolveResult := Eval(node.ResolveBlock, env, resolveCtx)
		if isError(resolveResult) {
			return resolveResult
		}
	}

	return result
}

func evalWithStatement(
	ws *ast.WithStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	// Evaluate the expression to get the resource
	resource := Eval(ws.Expression, env, ctx)
	if isError(resource) {
		return resource
	}

	// Create a new environment for the with block
	withEnv := object.NewEnclosedEnvironment(env)
	
	// Bind the resource to the specified variable
	if ws.Variable != nil {
		withEnv.Set(ws.Variable.Value, resource)
	}

	// Create a new context for the with block
	withCtx := &CallContext{
		FunctionName: "with_block",
		Node:         ws.Body,
		Parent:       ctx,
		env:          withEnv,
	}

	// Execute the body
	result := Eval(ws.Body, withEnv, withCtx)

	// Check if the resource has a close method and call it
	if resource != nil {
		if instance, ok := resource.(*object.Instance); ok {
			if closeMethod, exists := instance.Env.Get("close"); exists {
				if fn, isFn := closeMethod.(*object.Function); isFn {
					// Call the close method
					closeCtx := &CallContext{
						FunctionName: "close",
						Node:         ws,
						Parent:       ctx,
						env:          instance.Env,
					}
					evalCallExpression(fn, []object.Object{}, instance.Env, closeCtx)
				}
			}
		}
	}

	return result
}

func evalMatchStatement(
	ms *ast.MatchStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	matchValue := Eval(ms.MatchValue, env, ctx)
	if isError(matchValue) {
		return matchValue
	}

	for _, caseClause := range ms.Cases {
		caseCondition := Eval(caseClause.Condition, env, ctx)
		if isError(caseCondition) {
			return caseCondition
		}

		if isEqual(matchValue, caseCondition) {
			caseCtx := &CallContext{
				FunctionName: "case",
				Node:         caseClause.Body,
				Parent:       ctx,
				env:          env,
			}
			return Eval(caseClause.Body, env, caseCtx)
		}
	}

	if ms.Default != nil {
		defaultCtx := &CallContext{
			FunctionName: "default_case",
			Node:         ms.Default.Body,
			Parent:       ctx,
			env:          env,
		}
		return Eval(ms.Default.Body, env, defaultCtx)
	}

	return NONE
}

func isEqual(obj1, obj2 object.Object) bool {
	// Unwrap primitives if they are wrapped in instances
	obj1 = unwrapPrimitive(obj1)
	obj2 = unwrapPrimitive(obj2)
	
	switch obj1 := obj1.(type) {
	case *object.Integer:
		if obj2, ok := obj2.(*object.Integer); ok {
			return obj1.Value == obj2.Value
		}
	case *object.String:
		if obj2, ok := obj2.(*object.String); ok {
			return obj1.Value == obj2.Value
		}
	case *object.Float:
		if obj2, ok := obj2.(*object.Float); ok {
			return obj1.Value == obj2.Value
		}
	case *object.Boolean:
		if obj2, ok := obj2.(*object.Boolean); ok {
			return obj1.Value == obj2.Value
		}
	default:
		return false
	}
	return false
}

func evalAssignStatement(
	node *ast.AssignStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	switch target := node.Name.(type) {
	case *ast.Identifier:
		val := Eval(node.Value, env, ctx)
		if isError(val) {
			return val
		}

		// Check type hint if present
		if node.TypeHint != nil {
			// Get the type name from the type hint expression
			typeHintIdent, ok := node.TypeHint.(*ast.Identifier)
			if !ok {
				return newErrorWithTrace("invalid type hint: %s", node, ctx, node.TypeHint.String())
			}
			expectedType := typeHintIdent.Value

			// Validate the type
			if !checkType(val, expectedType) {
				return newErrorWithTrace("type mismatch: cannot assign %s to variable '%s' with type hint %s", node, ctx, getObjectTypeString(val), target.Value, expectedType)
			}

			// Store the type hint for future validations
			typeHintKey := "__type_hint__" + target.Value
			env.Set(typeHintKey, &object.String{Value: expectedType})
		} else {
			// Check if variable already has a type hint from previous assignment
			typeHintKey := "__type_hint__" + target.Value
			if existingTypeHint, exists := env.Get(typeHintKey); exists {
				if typeHintStr, ok := existingTypeHint.(*object.String); ok {
					if !checkType(val, typeHintStr.Value) {
						return newErrorWithTrace("type mismatch: cannot assign %s to variable '%s' with type hint %s", node, ctx, getObjectTypeString(val), target.Value, typeHintStr.Value)
					}
				}
			}
		}

		// Don't wrap primitives - this breaks arithmetic operations
		// Wrapping should only happen for explicit method calls on literals

		env.SetWithGlobalCheck(target.Value, val)
		return val

	case *ast.DotExpression:
		left := Eval(target.Left, env, ctx)
		if isError(left) {
			return left
		}
		instance, ok := left.(*object.Instance)
		if !ok {
			return newErrorWithTrace("invalid assignment target: %s", target, ctx, left.Type())
		}
		val := Eval(node.Value, env, ctx)
		if isError(val) {
			return val
		}
		instance.Env.Set(target.Right.Value, val)
		return val

	case *ast.TupleLiteral:
		val := Eval(node.Value, env, ctx)
		if isError(val) {
			return val
		}

		var values []object.Object
		switch v := val.(type) {
		case *object.Tuple:
			values = v.Elements
		case *object.Array:
			values = v.Elements
		default:
			return newErrorWithTrace("cannot unpack non-iterable type: %s", node, ctx, val.Type())
		}

		if len(target.Elements) != len(values) {
			return newErrorWithTrace("unpacking mismatch: expected %d values, got %d",
				node, ctx, len(target.Elements), len(values))
		}

		for i, expr := range target.Elements {
			ident, ok := expr.(*ast.Identifier)
			if !ok {
				return newErrorWithTrace("invalid assignment target in tuple assignment", node, ctx)
			}
			env.Set(ident.Value, values[i])
		}
		return val
	case *ast.IndexExpression:
		left := Eval(target.Left, env, ctx)
		if isError(left) {
			return left
		}

		index := Eval(target.Index, env, ctx)
		if isError(index) {
			return index
		}

		val := Eval(node.Value, env, ctx)
		if isError(val) {
			return val
		}

		return evalIndexAssignment(left, index, val, node, ctx)

	default:
		return newErrorWithTrace("invalid assignment target: %T", node, ctx, node.Name)
	}
}

func evalIndexAssignment(
	array, index, value object.Object,
	node ast.Node,
	ctx *CallContext,
) object.Object {
	switch array := array.(type) {
	case *object.Array:
		intIndex, ok := index.(*object.Integer)
		if !ok {
			return newErrorWithTrace("array index must be INTEGER, got %s", node, ctx, index.Type())
		}

		idx := intIndex.Value
		maxIndex := int64(len(array.Elements) - 1)

		// Handle negative indices Python-style
		if idx < 0 {
			idx = int64(len(array.Elements)) + idx
		}

		if idx < 0 || idx > maxIndex {
			return newErrorWithTrace("index out of bounds: %d (array length: %d)",
				node, ctx, idx, maxIndex+1)
		}

		// Perform the assignment
		array.Elements[idx] = value
		return value

	case *object.Hash:
		// Unwrap primitives that may have been wrapped in Instance objects
		unwrappedIndex := unwrapPrimitive(index)
		key, ok := unwrappedIndex.(object.Hashable)
		if !ok {
			return newErrorWithTrace("unusable as hash key: %s", node, ctx, unwrappedIndex.Type())
		}

		pair := object.HashPair{Key: unwrappedIndex, Value: value}
		array.Pairs[key.HashKey()] = pair
		return value

	default:
		return newErrorWithTrace("index assignment not supported: %s", node, ctx, array.Type())
	}
}

func checkType(val object.Object, expectedType string) bool {
	switch expectedType {
	case "str":
		// Check both primitive STRING and String grimoire instances
		if val.Type() == object.STRING_OBJ {
			return true
		}
		if instance, ok := val.(*object.Instance); ok && instance.Grimoire.Name == "String" {
			return true
		}
		return false
	case "int":
		// Check both primitive INTEGER and Integer grimoire instances
		if val.Type() == object.INTEGER_OBJ {
			return true
		}
		if instance, ok := val.(*object.Instance); ok && instance.Grimoire.Name == "Integer" {
			return true
		}
		return false
	case "float":
		// Check both primitive FLOAT and Float grimoire instances
		if val.Type() == object.FLOAT_OBJ {
			return true
		}
		if instance, ok := val.(*object.Instance); ok && instance.Grimoire.Name == "Float" {
			return true
		}
		return false
	case "bool":
		// Check both primitive BOOLEAN and Boolean grimoire instances
		if val.Type() == object.BOOLEAN_OBJ {
			return true
		}
		if instance, ok := val.(*object.Instance); ok && instance.Grimoire.Name == "Boolean" {
			return true
		}
		return false
	case "list":
		return val.Type() == object.ARRAY_OBJ
	case "dict":
		return val.Type() == object.MAP_OBJ
	case "None":
		return val.Type() == object.NONE_OBJ
	case "any":
		return true
	default:
		// For custom grimoire types, check if the value is an instance of that type
		if instance, ok := val.(*object.Instance); ok {
			return instance.Grimoire.Name == expectedType
		}
		return true
	}
}

func getGlobalEnv(env *object.Environment, ctx *CallContext) *object.Environment {
	for env.GetOuter() != nil {
		env = env.GetOuter()
	}
	return env
}


func evalGrimoireDefinition(
	node *ast.GrimoireDefinition,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	methods := map[string]*object.Function{}

	var parentGrimoire *object.Grimoire
	if node.Inherits != nil {
		parentObj, ok := env.Get(node.Inherits.Value)
		if !ok {
			return newErrorWithTrace(
				"parent grimoire '%s' not found",
				node,
				ctx,
				node.Inherits.Value,
			)
		}
		parentGrimoire, ok = parentObj.(*object.Grimoire)
		if !ok {
			return newErrorWithTrace("'%s' is not a grimoire", node, ctx, node.Inherits.Value)
		}

		for name, method := range parentGrimoire.Methods {
			methods[name] = method
		}
	}

	for _, method := range node.Methods {
		fn := &object.Function{
			Parameters: method.Parameters,
			Body:       method.Body,
			Env:        env.Clone(),
		}
		if strings.HasPrefix(method.Name.Value, "__") {
			fn.IsPrivate = true
		} else if strings.HasPrefix(method.Name.Value, "_") {
			fn.IsProtected = true
		}

		if method.Token.Type == token.ARCANESPELL {
			fn.IsAbstract = true
		}
		methods[method.Name.Value] = fn
	}

	if parentGrimoire != nil {
		for name, method := range parentGrimoire.Methods {
			if method.IsAbstract {
				if _, ok := methods[name]; !ok {
					return newErrorWithTrace(
						"grimoire '%s' must implement abstract method '%s'",
						node, ctx, node.Name.Value, name)
				}
			}
		}
	}

	grimoire := &object.Grimoire{
		Name:       node.Name.Value,
		Methods:    methods,
		InitMethod: nil,
		Env:        env.Clone(),
		Inherits:   parentGrimoire,
		IsArcane:   false,
	}

	if node.Token.Type == token.ARCANE {
		grimoire.IsArcane = true
	}
	if node.InitMethod != nil {
		initFn := &object.Function{
			Parameters: node.InitMethod.Parameters,
			Body:       node.InitMethod.Body,
			Env:        env.Clone(),
		}
		grimoire.InitMethod = initFn
	}

	env.Set(node.Name.Value, grimoire)
	return grimoire
}

// evalStaticMethodCall executes a static method call on a grimoire
// Sets MethodGrimoire to the grimoire itself since static methods belong to the class
func evalStaticMethodCall(
	grimoire *object.Grimoire,
	methodName string,
	args []object.Object,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	method, ok := grimoire.Methods[methodName]
	if !ok {
		return newErrorWithTrace("static method '%s' not found on %s",
			ctx.Node, ctx, methodName, grimoire.Name)
	}

	if method.IsPrivate && !sameClass(env, grimoire) {
		return newErrorWithTrace("private static method '%s' not accessible outside its defining class",
			ctx.Node, ctx, methodName)
	}

	if method.IsProtected && !sameOrSubclass(env, grimoire) {
		return newErrorWithTrace("protected static method '%s' not accessible here",
			ctx.Node, ctx, methodName)
	}

	// Create isolated method environment (no instance, no self)
	methodEnv := object.NewEnclosedEnvironment(grimoire.Env)
	
	// Add the grimoire constructor to the environment for self-instantiation
	methodEnv.Set(grimoire.Name, grimoire)

	// Create method context
	methodCtx := &CallContext{
		FunctionName: grimoire.Name + "." + methodName,
		Node:         ctx.Node,
		Parent:       ctx,
		env:          methodEnv,
		MethodGrimoire: grimoire,
	}

	// Bind arguments: support simple identifiers or full Parameter nodes
	// For static methods, we don't skip 'self' parameter
	argIndex := 0
	for _, pExpr := range method.Parameters {
		switch param := pExpr.(type) {
		case *ast.Identifier:
			name := param.Value
			if argIndex < len(args) {
				methodEnv.Set(name, args[argIndex])
				argIndex++
			} else {
				methodEnv.Set(name, NONE)
			}
		case *ast.Parameter:
			name := param.Name.Value
			if argIndex < len(args) {
				methodEnv.Set(name, args[argIndex])
				argIndex++
			} else if param.DefaultValue != nil {
				methodEnv.Set(name, Eval(param.DefaultValue, method.Env, methodCtx))
			} else {
				methodEnv.Set(name, NONE)
			}
		default:
			// unsupported parameter type
		}
	}

	// Execute with bounds checking for recursive calls
	return evalWithRecursionLimit(method.Body, methodEnv, method, methodCtx, 0)
}

// bindMethodParameters binds arguments to method parameters, handling default values and self parameter
func bindMethodParameters(
	method *object.Function,
	args []object.Object,
	methodEnv *object.Environment,
	methodCtx *CallContext,
	skipSelf bool,
) {
	argIndex := 0
	for _, pExpr := range method.Parameters {
		switch param := pExpr.(type) {
		case *ast.Identifier:
			name := param.Value
			// Skip 'self' parameter if requested (for instance methods)
			if skipSelf && name == "self" {
				continue
			}
			if argIndex < len(args) {
				methodEnv.Set(name, args[argIndex])
				argIndex++
			} else {
				methodEnv.Set(name, NONE)
			}
		case *ast.Parameter:
			name := param.Name.Value
			// Skip 'self' parameter if requested (for instance methods)
			if skipSelf && name == "self" {
				continue
			}
			if argIndex < len(args) {
				methodEnv.Set(name, args[argIndex])
				argIndex++
			} else if param.DefaultValue != nil {
				methodEnv.Set(name, Eval(param.DefaultValue, method.Env, methodCtx))
			} else {
				methodEnv.Set(name, NONE)
			}
		default:
			// unsupported parameter type
		}
	}
}

// evalGrimoireMethodCall executes a method call on a grimoire instance
// Sets MethodGrimoire to the method's defining class for proper super resolution
func evalGrimoireMethodCall(
	instance *object.Instance,
	methodName string,
	args []object.Object,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	method, ok := instance.Grimoire.Methods[methodName]
	if !ok {
		return newErrorWithTrace("method '%s' not found on %s",
			ctx.Node, ctx, methodName, instance.Grimoire.Name)
	}

	if method.IsPrivate && !sameClass(env, instance.Grimoire) {
		return newErrorWithTrace("private method '%s' not accessible outside its defining class",
			ctx.Node, ctx, methodName)
	}

	if method.IsProtected && !sameOrSubclass(env, instance.Grimoire) {
		return newErrorWithTrace("protected method '%s' not accessible here",
			ctx.Node, ctx, methodName)
	}

	// Create isolated method environment
	methodEnv := object.NewEnclosedEnvironment(instance.Env)
	methodEnv.Set("self", instance)

	// Find which grimoire owns this method for proper super resolution
	// This enables multi-level inheritance by tracking method ownership
	methodOwner := findMethodOwner(instance, methodName, method)
	
	// Create method context
	methodCtx := &CallContext{
		FunctionName: methodOwner.Name + "." + methodName,
		Node:         ctx.Node,
		Parent:       ctx,
		env:          methodEnv,
		MethodGrimoire: methodOwner,
	}

	// Bind arguments using the common helper function
	bindMethodParameters(method, args, methodEnv, methodCtx, true)

	// Execute with bounds checking for recursive calls
	return evalWithRecursionLimit(method.Body, methodEnv, method, methodCtx, 0)
}

// findMethodOwner finds which grimoire in the inheritance chain owns the given method
// This is crucial for proper super resolution in multi-level inheritance, ensuring
// that super calls resolve to the parent of the method's defining class, not the instance's class
func findMethodOwner(instance *object.Instance, methodName string, method *object.Function) *object.Grimoire {
	current := instance.Grimoire
	for current != nil {
		// Check if this grimoire has the method
		if methodName == "init" {
			if current.InitMethod == method {
				return current
			}
		} else {
			if m, exists := current.Methods[methodName]; exists && m == method {
				return current
			}
		}
		current = current.Inherits
	}
	// Fallback: return the instance's grimoire if we can't find the owner
	return instance.Grimoire
}

// evalBoundMethodCall executes a bound method call with proper context tracking
// Sets MethodGrimoire to enable correct super resolution in inheritance hierarchies
func evalBoundMethodCall(
	boundMethod *object.BoundMethod,
	args []object.Object,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	// For bound methods, use the stored method directly with proper environment
	method := boundMethod.Method
	instance := boundMethod.Instance
	methodName := boundMethod.Name

	// Create isolated method environment from the method's original environment, not instance env
	methodEnv := object.NewEnclosedEnvironment(method.Env)
	methodEnv.Set("self", instance)

	// Find which grimoire owns this method for proper super resolution
	// This enables multi-level inheritance by tracking method ownership
	methodOwner := findMethodOwner(instance, methodName, method)
	
	
	// Create method context
	methodCtx := &CallContext{
		FunctionName: methodOwner.Name + "." + methodName,
		Node:         ctx.Node,
		Parent:       ctx,
		env:          methodEnv,
		depth:        ctx.depth + 1,
		MethodGrimoire: methodOwner,
	}

	// Bind arguments using the common helper function
	bindMethodParameters(method, args, methodEnv, methodCtx, true)

	// Execute with bounds checking for recursive calls
	return evalWithRecursionLimit(method.Body, methodEnv, method, methodCtx, 0)
}

// recursionDepths is defined at the top of the file

func evalWithRecursionLimit(
	body *ast.BlockStatement,
	env *object.Environment,
	method *object.Function,
	ctx *CallContext,
	depth int,
) object.Object {
	// Get current depth or start at provided depth
	currentDepth, exists := recursionDepths[body]
	if !exists {
		currentDepth = depth
	}

	// Increment and check
	recursionDepths[body] = currentDepth + 1
	if recursionDepths[body] > MAX_CALL_DEPTH {
		recursionDepths[body]-- // Clean up
		return newErrorWithTrace("maximum recursion depth exceeded (limit: %d)",
			body, ctx, MAX_CALL_DEPTH)
	}

	// Evaluate with depth tracking
	result := Eval(body, env, ctx)

	// Clean up
	recursionDepths[body]--
	if recursionDepths[body] <= 0 {
		delete(recursionDepths, body)
	}

	return unwrapReturnValue(result)
}

func evalCallExpression(
	fn object.Object,
	args []object.Object,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	// ── 1.  Flatten a single‑tuple argument safely ────────────────────────────
	if len(args) == 1 && args[0] != nil {
		if tup, ok := args[0].(*object.Tuple); ok {
			args = tup.Elements
		}
	}

	// ── 2.  Propagate existing errors, don’t try to “call” them ───────────────
	if isError(fn) {
		return fn
	}
	for _, a := range args {
		if isError(a) {
			return a
		}
	}

	// ── 3.  Normal dispatch ───────────────────────────────────────────────────
	switch fnTyped := fn.(type) {

	case *object.Function:
		// Type checking before function call
		if typeErr := checkParameterTypes(fnTyped, args, ctx); typeErr != nil {
			return typeErr
		}
		
		// use the correctly typed value as the map key
		fun := fnTyped // alias for brevity

		callCtx, ok := callStack[fun] // was callStack[fn]
		if !ok {
			callCtx = &CallContext{depth: 0, env: env}
			callStack[fun] = callCtx // was callStack[fn]
		}
		callCtx.depth++
		if callCtx.depth > MAX_CALL_DEPTH {
			callCtx.depth--
			return newErrorWithTrace(
				"maximum recursion depth exceeded (%d)", ctx.Node, ctx, MAX_CALL_DEPTH)
		}

		global := getGlobalEnv(fun.Env, ctx)
		extended := extendFunctionEnv(fun, args, global, ctx)

		// we don’t know the spell’s name here; fall back to the caller’s ctx
		funcName := ctx.FunctionName
		if funcName == "" {
			funcName = "<anonymous>"
		}

		fnCtx := &CallContext{
			FunctionName: funcName,
			Node:         fun.Body,
			Parent:       ctx,
			env:          extended,
		}

		evaluated := Eval(fun.Body, extended, fnCtx)

		callCtx.depth--
		if callCtx.depth == 0 {
			delete(callStack, fun) // was delete(callStack, fn)
		}
		return unwrapReturnValue(evaluated)
	case *object.BoundMethod:
		// Use a special version of evalGrimoireMethodCall that uses the stored method
		return evalBoundMethodCall(fnTyped, args, env, ctx)

	case *object.StaticMethod:
		return evalStaticMethodCall(fnTyped.Grimoire, fnTyped.Name, args, env, ctx)

	case *object.Grimoire:
		if fnTyped.IsArcane {
			return newErrorWithTrace(
				"cannot instantiate arcane grimoire: %s", ctx.Node, ctx, fnTyped.Name)
		}

		instance := &object.Instance{
			Grimoire: fnTyped,
			Env:      object.NewEnclosedEnvironment(fnTyped.Env),
		}

		if fnTyped.InitMethod != nil {
			global := getGlobalEnv(fnTyped.Env, ctx)
			extended := extendFunctionEnv(fnTyped.InitMethod, args, global, ctx)
			extended.Set("self", instance)

			initCtx := &CallContext{
				FunctionName: fnTyped.Name + ".init",
				Node:         fnTyped.InitMethod.Body,
				Parent:       ctx,
				env:          extended,
				MethodGrimoire: fnTyped,
			}
			result := Eval(fnTyped.InitMethod.Body, extended, initCtx)
			if isError(result) {
				return result
			}
		}
		return instance

	case *object.Builtin:
		res := fnTyped.Fn(args...)
		if err, ok := res.(*object.Error); ok {
			return newErrorWithTrace(err.Message, ctx.Node, ctx)
		}
		// Wrap string results from input functions in String grimoire instances
		if shouldWrapStringResult(ctx.FunctionName) {
			if stringObj, isString := res.(*object.String); isString {
				return wrapPrimitive(stringObj, env, ctx)
			}
		}
		// Wrap array results from pairs() function so they have access to methods
		if ctx.FunctionName == "pairs" {
			if arrayObj, isArray := res.(*object.Array); isArray {
				return wrapPrimitive(arrayObj, env, ctx)
			}
		}
		return res

	default:
		return newErrorWithTrace(
			"not a function: %s (in file %s)",
			ctx.Node, ctx, fn.Type(), getSourcePosition(ctx.Node).Filename)
	}
}

func evalDotExpression(
	node *ast.DotExpression,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	// Note: Removed literal blocking code - primitives are now automatically wrapped
	leftObj := Eval(node.Left, env, ctx)
	if isError(leftObj) {
		return leftObj
	}

	// Handle super object access
	if superObj, ok := leftObj.(*object.Super); ok {
		var parentMethod *object.Function
		var methodExists bool

		// Check if it's the init method
		if node.Right.Value == "init" {
			parentMethod = superObj.Parent.InitMethod
			methodExists = (parentMethod != nil)
		} else {
			// Check regular methods
			parentMethod, methodExists = superObj.Parent.Methods[node.Right.Value]
		}

		if !methodExists {
			return newErrorWithTrace(
				"no method '%s' found in parent class",
				node,
				ctx,
				node.Right.Value,
			)
		}


		return &object.BoundMethod{
			Instance: superObj.Instance,
			Method:   parentMethod,
			Name:     node.Right.Value,
		}
	}

	// Handle super string-based access (fallback for original syntax)
	if node.Left.String() == "super" {
		instance, ok := env.Get("self")
		if !ok || instance == nil {
			return newErrorWithTrace("'super' can only be used in an instance method", node, ctx)
		}

		inst, ok := instance.(*object.Instance)
		if !ok {
			return newErrorWithTrace(
				"'super' must be used in an instance of a grimoire",
				node,
				ctx,
			)
		}

		if inst.Grimoire == nil || inst.Grimoire.Inherits == nil {
			return newErrorWithTrace("no parent class found for 'super'", node, ctx)
		}

		var parentMethod *object.Function
		var methodExists bool

		// Check if it's the init method
		if node.Right.Value == "init" {
			parentMethod = inst.Grimoire.Inherits.InitMethod
			methodExists = (parentMethod != nil)
		} else {
			// Check regular methods
			parentMethod, methodExists = inst.Grimoire.Inherits.Methods[node.Right.Value]
		}

		if !methodExists {
			return newErrorWithTrace(
				"no method '%s' found in parent class",
				node,
				ctx,
				node.Right.Value,
			)
		}

		return &object.BoundMethod{
			Instance: inst,
			Method:   parentMethod,
			Name:     node.Right.Value,
		}
	}

	// Handle namespace access (for import aliases)
	if namespace, ok := leftObj.(*object.Namespace); ok {
		fieldOrMethodName := node.Right.Value
		if val, found := namespace.Env.Get(fieldOrMethodName); found {
			return val
		}
		return newErrorWithTrace("undefined member in namespace: %s", node, ctx, fieldOrMethodName)
	}

	// Handle static method calls on grimoire classes
	if grimoire, ok := leftObj.(*object.Grimoire); ok {
		methodName := node.Right.Value
		method, exists := grimoire.Methods[methodName]
		if !exists {
			return newErrorWithTrace("undefined static method: %s", node, ctx, methodName)
		}

		// Return a static method wrapper that doesn't require self
		return &object.StaticMethod{
			Grimoire: grimoire,
			Method:   method,
			Name:     methodName,
		}
	}

	// Handle CaughtError access
	if caughtErr, ok := leftObj.(*object.CaughtError); ok {
		switch node.Right.Value {
		case "message":
			return &object.String{Value: caughtErr.GetMessage()}
		case "type":
			if errWithTrace, ok := caughtErr.OriginalError.(*object.ErrorWithTrace); ok {
				if errWithTrace.CustomDetails != nil {
					if errorType, exists := errWithTrace.CustomDetails["errorType"]; exists {
						return errorType
					}
				}
			}
			if customErr, ok := caughtErr.OriginalError.(*object.CustomError); ok {
				return &object.String{Value: customErr.Name}
			}
			return &object.String{Value: "Error"}
		default:
			return newErrorWithTrace("CaughtError has no property: %s", node, ctx, node.Right.Value)
		}
	}

	instance, ok := leftObj.(*object.Instance)
	if !ok {
		return newErrorWithTrace("type error: %s is not an instance or namespace", node, ctx, getObjectTypeString(leftObj))
	}

	fieldOrMethodName := node.Right.Value

	if val, found := instance.Env.Get(fieldOrMethodName); found {
		return val
	}

	method, ok := instance.Grimoire.Methods[fieldOrMethodName]
	if !ok {
		return newErrorWithTrace("undefined property or method: %s", node, ctx, fieldOrMethodName)
	}

	if method.IsPrivate && !sameClass(env, instance.Grimoire) {
		return newErrorWithTrace(
			"private method '%s' not accessible outside its defining class",
			node, ctx, fieldOrMethodName)
	}

	if method.IsProtected && !sameOrSubclass(env, instance.Grimoire) {
		return newErrorWithTrace("protected method '%s' not accessible here",
			node, ctx, fieldOrMethodName)
	}

	return &object.BoundMethod{
		Instance: instance,
		Method:   method,
		Name:     fieldOrMethodName,
	}
}

// Helper function to check if a node is a literal
func isLiteralNode(node ast.Expression) bool {
	switch node.(type) {
	case *ast.IntegerLiteral, *ast.FloatLiteral, *ast.StringLiteral, *ast.Boolean:
		return true
	default:
		return false
	}
}

// Modify the unwrapReturnValue function to help with method chaining
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func sameClass(env *object.Environment, target *object.Grimoire) bool {
	callerSelf, ok := env.Get("self")
	if !ok || callerSelf == nil {
		return false
	}
	callerInst, ok := callerSelf.(*object.Instance)
	if !ok || callerInst == nil {
		return false
	}
	return callerInst.Grimoire == target
}

func sameOrSubclass(env *object.Environment, target *object.Grimoire) bool {
	callerSelf, ok := env.Get("self")
	if !ok || callerSelf == nil {
		return false
	}
	callerInst, ok := callerSelf.(*object.Instance)
	if !ok || callerInst == nil {
		return false
	}

	sb := callerInst.Grimoire
	for sb != nil {
		if sb == target {
			return true
		}
		sb = sb.Inherits
	}
	return false
}

func evalHashLiteral(
	node *ast.HashLiteral,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)
	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env, ctx)
		if isError(key) {
			return key
		}
		// Unwrap primitives that may have been wrapped in Instance objects
		unwrappedKey := unwrapPrimitive(key)
		hashKey, ok := unwrappedKey.(object.Hashable)
		if !ok {
			return newErrorWithTrace("unusable as hash key: %s", node, ctx, unwrappedKey.Type())
		}
		value := Eval(valueNode, env, ctx)
		if isError(value) {
			return value
		}
		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: unwrappedKey, Value: value}
	}
	return &object.Hash{Pairs: pairs}
}

func evalTupleLiteral(
	tl *ast.TupleLiteral,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	elements := evalExpressions(tl.Elements, env, ctx)
	if len(elements) == 1 && isError(elements[0]) {
		return elements[0]
	}

	return &object.Tuple{Elements: elements}
}

func evalIndexExpression(left, index object.Object, node ast.Node, ctx *CallContext) object.Object {
	// Unwrap instances to get the underlying primitive values
	unwrappedLeft := unwrapPrimitive(left)
	unwrappedIndex := unwrapPrimitive(index)
	
	switch {
	case unwrappedLeft.Type() == object.TUPLE_OBJ:
		return evalTupleIndexExpression(unwrappedLeft, unwrappedIndex, node, ctx)
	case unwrappedLeft.Type() == object.ARRAY_OBJ && unwrappedIndex.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(unwrappedLeft, unwrappedIndex, node, ctx)
	case unwrappedLeft.Type() == object.MAP_OBJ:
		return evalHashIndexExpression(unwrappedLeft, unwrappedIndex, node, ctx)
	case unwrappedLeft.Type() == object.STRING_OBJ && unwrappedIndex.Type() == object.INTEGER_OBJ:
		result := evalStringIndexExpression(unwrappedLeft, unwrappedIndex, node, ctx)
		// If the original left was an instance, wrap the result back to maintain consistency
		if left.Type() == object.INSTANCE_OBJ {
			// Get current environment - use a simple approach
			var currentEnv *object.Environment
			if ctx != nil && ctx.env != nil {
				currentEnv = ctx.env
			} else {
				// Create a temporary environment if none available
				currentEnv = object.NewEnvironment()
			}
			return wrapPrimitive(result, currentEnv, ctx)
		}
		return result
	case left.Type() == object.INSTANCE_OBJ:
		// Handle special case where instance might have custom indexing behavior
		if instance, ok := left.(*object.Instance); ok {
			// Check if the instance has a get method for custom indexing
			if _, exists := instance.Grimoire.Methods["get"]; exists {
				// Get current environment for method call
				var currentEnv *object.Environment
				if ctx != nil && ctx.env != nil {
					currentEnv = ctx.env
				} else {
					currentEnv = object.NewEnvironment()
				}
				return evalGrimoireMethodCall(instance, "get", []object.Object{index}, currentEnv, ctx)
			}
			// If no custom get method, fall back to unwrapped handling
			if unwrappedLeft.Type() == object.STRING_OBJ && unwrappedIndex.Type() == object.INTEGER_OBJ {
				result := evalStringIndexExpression(unwrappedLeft, unwrappedIndex, node, ctx)
				// Wrap result back since original was instance
				var currentEnv *object.Environment
				if ctx != nil && ctx.env != nil {
					currentEnv = ctx.env
				} else {
					currentEnv = object.NewEnvironment()
				}
				return wrapPrimitive(result, currentEnv, ctx)
			}
		}
		return newErrorWithTrace("invalid instance for indexing: %s", node, ctx, left.Type())
	default:
		return newErrorWithTrace("index operator not supported: %s", node, ctx, left.Type())
	}
}

func evalTupleIndexExpression(
	tuple, index object.Object,
	node ast.Node,
	ctx *CallContext,
) object.Object {
	tupleObj, ok := tuple.(*object.Tuple)
	if !ok {
		return newErrorWithTrace("not a tuple: %s", node, ctx, tuple.Type())
	}
	
	indexObj, ok := index.(*object.Integer)
	if !ok {
		return newErrorWithTrace("tuple index must be INTEGER, got %s", node, ctx, index.Type())
	}
	idx := int(indexObj.Value)

	// Handle negative indices
	if idx < 0 {
		idx = len(tupleObj.Elements) + idx
	}

	if idx < 0 || idx >= len(tupleObj.Elements) {
		return newErrorWithTrace("index out of bounds: %d (tuple length: %d)",
			node, ctx, idx, len(tupleObj.Elements))
	}

	return tupleObj.Elements[idx]
}

func evalHashIndexExpression(
	hash, index object.Object,
	node ast.Node,
	ctx *CallContext,
) object.Object {
	hashObject, ok := hash.(*object.Hash)
	if !ok {
		return newErrorWithTrace("not a hash: %s", node, ctx, hash.Type())
	}
	// Unwrap primitives that may have been wrapped in Instance objects
	unwrappedIndex := unwrapPrimitive(index)
	key, ok := unwrappedIndex.(object.Hashable)
	if !ok {
		return newErrorWithTrace("unusable as hash key: %s", node, ctx, unwrappedIndex.Type())
	}
	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return NONE
	}
	return pair.Value
}

func evalArrayIndexExpression(
	array, index object.Object,
	node ast.Node,
	ctx *CallContext,
) object.Object {
	arrayObject, ok := array.(*object.Array)
	if !ok {
		return newErrorWithTrace("index operation not supported on %s", node, ctx, array.Type())
	}

	intIndex, ok := index.(*object.Integer)
	if !ok {
		return newErrorWithTrace("array index must be INTEGER, got %s", node, ctx, index.Type())
	}

	idx := intIndex.Value
	maxIndex := int64(len(arrayObject.Elements) - 1)

	// Handle negative indices like Python
	if idx < 0 {
		idx = int64(len(arrayObject.Elements)) + idx
	}

	if idx < 0 || idx > maxIndex {
		return newErrorWithTrace("index out of bounds: %d (array length: %d)",
			node, ctx, idx, maxIndex+1)
	}

	return arrayObject.Elements[idx]
}

func evalStringIndexExpression(
	str, index object.Object,
	node ast.Node,
	ctx *CallContext,
) object.Object {
	stringObj, ok := str.(*object.String)
	if !ok {
		return newErrorWithTrace("index operation not supported on %s", node, ctx, str.Type())
	}

	intIndex, ok := index.(*object.Integer)
	if !ok {
		return newErrorWithTrace("string index must be INTEGER, got %s", node, ctx, index.Type())
	}

	idx := intIndex.Value
	strLen := int64(len(stringObj.Value))
	maxIndex := strLen - 1

	// Handle negative indices like Python
	if idx < 0 {
		idx = strLen + idx
	}

	if idx < 0 || idx > maxIndex {
		return newErrorWithTrace("string index out of bounds: %d (string length: %d)",
			node, ctx, idx, strLen)
	}

	// Return a single-character string
	return &object.String{Value: string(stringObj.Value[idx])}
}

func evalSliceExpression(left, start, end object.Object, node ast.Node, ctx *CallContext) object.Object {
	// Unwrap instances to get the underlying primitive values
	unwrappedLeft := unwrapPrimitive(left)
	
	switch {
	case unwrappedLeft.Type() == object.STRING_OBJ:
		result := evalStringSliceExpression(unwrappedLeft, start, end, node, ctx)
		// If the original left was an instance, wrap the result back to maintain consistency
		if left.Type() == object.INSTANCE_OBJ {
			var currentEnv *object.Environment
			if ctx != nil && ctx.env != nil {
				currentEnv = ctx.env
			} else {
				currentEnv = object.NewEnvironment()
			}
			return wrapPrimitive(result, currentEnv, ctx)
		}
		return result
	case unwrappedLeft.Type() == object.ARRAY_OBJ:
		return evalArraySliceExpression(unwrappedLeft, start, end, node, ctx)
	default:
		return newErrorWithTrace("slice operator not supported: %s", node, ctx, left.Type())
	}
}

func evalStringSliceExpression(str, start, end object.Object, node ast.Node, ctx *CallContext) object.Object {
	stringObj, ok := str.(*object.String)
	if !ok {
		return newErrorWithTrace("slice operation not supported on %s", node, ctx, str.Type())
	}
	
	strLen := int64(len(stringObj.Value))
	var startIdx, endIdx int64
	
	// Handle start index
	if start != nil {
		startInt, ok := start.(*object.Integer)
		if !ok {
			return newErrorWithTrace("slice start index must be INTEGER, got %s", node, ctx, start.Type())
		}
		startIdx = startInt.Value
		if startIdx < 0 {
			startIdx = strLen + startIdx
		}
	} else {
		startIdx = 0
	}
	
	// Handle end index
	if end != nil {
		endInt, ok := end.(*object.Integer)
		if !ok {
			return newErrorWithTrace("slice end index must be INTEGER, got %s", node, ctx, end.Type())
		}
		endIdx = endInt.Value
		if endIdx < 0 {
			endIdx = strLen + endIdx
		}
	} else {
		endIdx = strLen
	}
	
	// Bounds checking
	if startIdx < 0 {
		startIdx = 0
	}
	if endIdx > strLen {
		endIdx = strLen
	}
	if startIdx > endIdx {
		startIdx = endIdx
	}
	
	// Extract substring
	return &object.String{Value: stringObj.Value[startIdx:endIdx]}
}

func evalArraySliceExpression(arr, start, end object.Object, node ast.Node, ctx *CallContext) object.Object {
	arrayObj, ok := arr.(*object.Array)
	if !ok {
		return newErrorWithTrace("slice operation not supported on %s", node, ctx, arr.Type())
	}
	
	arrLen := int64(len(arrayObj.Elements))
	var startIdx, endIdx int64
	
	// Handle start index
	if start != nil {
		startInt, ok := start.(*object.Integer)
		if !ok {
			return newErrorWithTrace("slice start index must be INTEGER, got %s", node, ctx, start.Type())
		}
		startIdx = startInt.Value
		if startIdx < 0 {
			startIdx = arrLen + startIdx
		}
	} else {
		startIdx = 0
	}
	
	// Handle end index
	if end != nil {
		endInt, ok := end.(*object.Integer)
		if !ok {
			return newErrorWithTrace("slice end index must be INTEGER, got %s", node, ctx, end.Type())
		}
		endIdx = endInt.Value
		if endIdx < 0 {
			endIdx = arrLen + endIdx
		}
	} else {
		endIdx = arrLen
	}
	
	// Bounds checking
	if startIdx < 0 {
		startIdx = 0
	}
	if endIdx > arrLen {
		endIdx = arrLen
	}
	if startIdx > endIdx {
		startIdx = endIdx
	}
	
	// Extract subarray
	newElements := make([]object.Object, endIdx-startIdx)
	copy(newElements, arrayObj.Elements[startIdx:endIdx])
	return &object.Array{Elements: newElements}
}

func extendFunctionEnv(
	fn *object.Function,
	args []object.Object,
	global *object.Environment,
	ctx *CallContext,
) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	// Bind parameters: support ast.Identifier or ast.Parameter nodes
	for i, pExpr := range fn.Parameters {
		switch param := pExpr.(type) {
		case *ast.Identifier:
			name := param.Value
			if i < len(args) {
				env.Set(name, args[i])
			} else {
				env.Set(name, NONE)
			}
		case *ast.Parameter:
			name := param.Name.Value
			if i < len(args) {
				env.Set(name, args[i])
			} else if param.DefaultValue != nil {
				// Default value may refer to an identifier or an expression
				switch dv := param.DefaultValue.(type) {
				case *ast.Identifier:
					if val, ok := global.Get(dv.Value); ok {
						env.Set(name, val)
					} else {
						env.Set(name,
							newErrorWithTrace("identifier not found: %s", dv, ctx, dv.Value))
					}
				default:
					env.Set(name, Eval(param.DefaultValue, fn.Env, ctx))
				}
			} else {
				env.Set(name, NONE)
			}
			
			// Store type hint for parameter if present
			if param.TypeHint != nil {
				if typeHintIdent, ok := param.TypeHint.(*ast.Identifier); ok {
					typeHintKey := "__type_hint__" + name
					env.Set(typeHintKey, &object.String{Value: typeHintIdent.Value})
				}
			}
		default:
			// Unsupported parameter node
		}
	}

	return env
}

func evalIdentifier(node *ast.Identifier, env *object.Environment, ctx *CallContext) object.Object {
	// Handle super keyword
	if node.Value == "super" {
		instance, ok := env.Get("self")
		if !ok || instance == nil {
			return newErrorWithTrace("'super' can only be used in an instance method", node, ctx)
		}

		inst, ok := instance.(*object.Instance)
		if !ok {
			return newErrorWithTrace(
				"'super' must be used in an instance of a grimoire",
				node,
				ctx,
			)
		}

		// Use the method's grimoire context if available, otherwise fall back to instance grimoire
		// This ensures super resolves to the parent of the current method's class, not the instance's class
		// Critical for multi-level inheritance to work correctly
		var currentGrimoire *object.Grimoire
		if ctx.MethodGrimoire != nil {
			currentGrimoire = ctx.MethodGrimoire
		} else {
			currentGrimoire = inst.Grimoire
		}
		
		if currentGrimoire == nil || currentGrimoire.Inherits == nil {
			return newErrorWithTrace("no parent class found for 'super'", node, ctx)
		}

		// Return a super object that can be used for method calls
		return &object.Super{
			Instance: inst,
			Parent:   currentGrimoire.Inherits,
		}
	}

	// First check builtins.
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}
	// Then check the environment.
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	if node.Value == "None" {
		return object.NONE
	}
	return newErrorWithTrace("identifier not found: %s", node, ctx, node.Value)
}

func evalProgram(program *ast.Program, env *object.Environment, ctx *CallContext) object.Object {
	var result object.Object
	var mainStatement *ast.MainStatement
	
	// Create a proper program context if none exists
	if ctx == nil {
		ctx = &CallContext{
			FunctionName:      "<program>",
			Node:              program,
			Parent:            nil,
			IsDirectExecution: true,
			env:               env,
		}
	}
	
	// First pass: check if main statement exists
	for _, statement := range program.Statements {
		if mainStmt, ok := statement.(*ast.MainStatement); ok {
			mainStatement = mainStmt
			break
		}
	}
	
	// Second pass: process statements based on whether main exists
	for _, statement := range program.Statements {
		if _, ok := statement.(*ast.MainStatement); ok {
			// Skip main statement for now
			continue
		}
		
		// If main exists, only execute function definitions, class definitions, and assignments
		// Skip other top-level executable statements
		if mainStatement != nil {
			if _, ok := statement.(*ast.FunctionDefinition); ok {
				// Execute function definitions
				result = Eval(statement, env, ctx)
			} else if _, ok := statement.(*ast.GrimoireDefinition); ok {
				// Execute class definitions
				result = Eval(statement, env, ctx)
			} else if _, ok := statement.(*ast.AssignStatement); ok {
				// Execute assignment statements
				result = Eval(statement, env, ctx)
			} else if exprStmt, ok := statement.(*ast.ExpressionStatement); ok {
				// Skip expression statements like print calls when main exists
				_ = exprStmt // Avoid unused variable warning
			} else {
				// Execute other statements (imports, etc.)
				result = Eval(statement, env, ctx)
			}
		} else {
			// No main block, execute everything normally (backward compatibility)
			result = Eval(statement, env, ctx)
		}
		
		switch result.(type) {
		case *object.ReturnValue:
			return result.(*object.ReturnValue).Value
		case *object.Error, *object.CustomError, *object.ErrorWithTrace:
			return result
		}
	}
	
	// If main statement exists, execute it last (only for direct execution)
	if mainStatement != nil && ctx != nil && ctx.IsDirectExecution {
		mainCtx := &CallContext{
			FunctionName:      "main",
			Node:              mainStatement,
			Parent:            ctx,
			IsDirectExecution: true,
			env:               env,
		}
		result = Eval(mainStatement, env, mainCtx)
		
		switch result.(type) {
		case *object.ReturnValue:
			return result.(*object.ReturnValue).Value
		case *object.Error, *object.CustomError, *object.ErrorWithTrace:
			return result
		}
	}
	
	return result
}

func evalBlockStatement(
	block *ast.BlockStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env, ctx)

		if result != nil {
			rt := result.Type()

			if rt == object.RETURN_VALUE_OBJ ||
				rt == object.ERROR_OBJ ||
				rt == object.CUSTOM_ERROR_OBJ ||
				rt == object.STOP.Type() ||
				rt == object.SKIP.Type() ||
				isErrorWithTrace(result) {
				return result
			}
		}
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(
	operator string,
	node *ast.PrefixExpression,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	switch operator {
	case "!":
		right := Eval(node.Right, env, ctx)
		return evalBangOperatorExpression(right, env, ctx)
	case "not":
		right := Eval(node.Right, env, ctx)
		if isError(right) {
			return right
		}
		return evalBangOperatorExpression(right, env, ctx)
	case "~":
		right := Eval(node.Right, env, ctx)
		if isError(right) {
			return right
		}
		unwrappedRight := unwrapPrimitive(right)
		intOperand, ok := unwrappedRight.(*object.Integer)
		if !ok {
			return newErrorWithTrace("unsupported operand type for ~: %s", node, ctx, unwrappedRight.Type())
		}

		return wrapPrimitive(&object.Integer{Value: ^intOperand.Value}, env, ctx)

	// Prefix increment and decrement operators
	case "++":
		// Only identifiers can be incremented
		ident, ok := node.Right.(*ast.Identifier)
		if !ok {
			return newErrorWithTrace("prefix '++' operator requires an identifier", node, ctx)
		}
		obj := Eval(node.Right, env, ctx)
		if isError(obj) {
			return obj
		}
		intObj, ok := obj.(*object.Integer)
		if !ok {
			return newErrorWithTrace("prefix '++' operator requires an integer", node, ctx)
		}
		newVal := intObj.Value + 1
		env.Set(ident.Value, &object.Integer{Value: newVal})
		return &object.Integer{Value: newVal}
	case "--":
		ident, ok := node.Right.(*ast.Identifier)
		if !ok {
			return newErrorWithTrace("prefix '--' operator requires an identifier", node, ctx)
		}
		obj := Eval(node.Right, env, ctx)
		if isError(obj) {
			return obj
		}
		intObj, ok := obj.(*object.Integer)
		if !ok {
			return newErrorWithTrace("prefix '--' operator requires an integer", node, ctx)
		}
		newValDec := intObj.Value - 1
		env.Set(ident.Value, &object.Integer{Value: newValDec})
		return &object.Integer{Value: newValDec}
	case "-":
		right := Eval(node.Right, env, ctx)
		return evalMinusPrefixOperatorExpression(right, env, ctx)
	default:
		return newErrorWithTrace("unknown operator: %s%s", node, ctx,
			operator, Eval(node.Right, env, ctx).Type())
	}
}

func evalInfixExpression(
	operator string,
	left, right object.Object,
	node ast.Node,
	ctx *CallContext,
) object.Object {
	if debugPrimitiveWrapping && operator == "//" {
		fmt.Fprintf(os.Stderr, "EVAL_INFIX: %s between %T and %T in %s\n", operator, left, right, getContextName(ctx))
	}
	// Unwrap primitive values from instances if needed
	unwrappedLeft := unwrapPrimitive(left)
	unwrappedRight := unwrapPrimitive(right)
	
	// Debug all arithmetic operations in method contexts
	if debugPrimitiveWrapping && ctx != nil && strings.Contains(getContextName(ctx), ".") {
		fmt.Fprintf(os.Stderr, "ARITH: %T %s %T -> unwrapped: %T %s %T in %s\n", 
			left, operator, right, unwrappedLeft, operator, unwrappedRight, getContextName(ctx))
		if leftInt, ok := unwrappedLeft.(*object.Integer); ok {
			if rightInt, ok := unwrappedRight.(*object.Integer); ok {
				fmt.Fprintf(os.Stderr, "ARITH: %d %s %d in %s\n", leftInt.Value, operator, rightInt.Value, getContextName(ctx))
			}
		}
	}
	
	// Handle "in" and "not in" operators before type-specific switches
	if operator == "in" || operator == "not in" {
		result := evalInOperator(unwrappedLeft, unwrappedRight, node, ctx)
		if operator == "not in" {
			// Invert the result for "not in"
			if boolResult, ok := result.(*object.Boolean); ok {
				return nativeBoolToBooleanObject(!boolResult.Value)
			}
		}
		return result
	}
	
	switch {
	case unwrappedLeft.Type() == object.INTEGER_OBJ && unwrappedRight.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, unwrappedLeft, unwrappedRight, node, ctx)
	case unwrappedLeft.Type() == object.BOOLEAN_OBJ && unwrappedRight.Type() == object.BOOLEAN_OBJ:
		return evalBooleanInfixExpression(operator, unwrappedLeft, unwrappedRight, node, ctx)
	case unwrappedLeft.Type() == object.STRING_OBJ && unwrappedRight.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, unwrappedLeft, unwrappedRight, node, ctx)
	case unwrappedLeft.Type() == object.STRING_OBJ && unwrappedRight.Type() == object.INTEGER_OBJ:
		// String multiplication: "hello" * 3
		if operator == "*" {
			strVal := unwrappedLeft.(*object.String).Value
			count := unwrappedRight.(*object.Integer).Value
			if count < 0 {
				return newErrorWithTrace("string multiplication count cannot be negative", node, ctx)
			}
			result := strings.Repeat(strVal, int(count))
			return &object.String{Value: result}
		}
		return newErrorWithTrace("unsupported operation: %s %s %s", node, ctx, unwrappedLeft.Type(), operator, unwrappedRight.Type())
	case unwrappedLeft.Type() == object.INTEGER_OBJ && unwrappedRight.Type() == object.STRING_OBJ:
		// Integer * string: 3 * "hello"
		if operator == "*" {
			count := unwrappedLeft.(*object.Integer).Value
			strVal := unwrappedRight.(*object.String).Value
			if count < 0 {
				return newErrorWithTrace("string multiplication count cannot be negative", node, ctx)
			}
			result := strings.Repeat(strVal, int(count))
			return &object.String{Value: result}
		}
		return newErrorWithTrace("unsupported operation: %s %s %s", node, ctx, unwrappedLeft.Type(), operator, unwrappedRight.Type())
	case unwrappedLeft == object.NONE && unwrappedRight == object.NONE:
		return nativeBoolToBooleanObject(operator == "==")
	case unwrappedLeft.Type() == object.ARRAY_OBJ && unwrappedRight.Type() == object.ARRAY_OBJ:
		if operator == "+" {
			leftArr := unwrappedLeft.(*object.Array)
			rightArr := unwrappedRight.(*object.Array)
			combined := make([]object.Object, len(leftArr.Elements)+len(rightArr.Elements))
			copy(combined, leftArr.Elements)
			copy(combined[len(leftArr.Elements):], rightArr.Elements)
			return &object.Array{Elements: combined}
		}
		return newErrorWithTrace("unknown operator for arrays: %s", node, ctx, operator)
	case unwrappedLeft.Type() == object.ARRAY_OBJ && unwrappedRight.Type() == object.INTEGER_OBJ:
		// Array multiplication: [1, 2] * 3
		if operator == "*" {
			leftArr := unwrappedLeft.(*object.Array)
			count := unwrappedRight.(*object.Integer).Value
			if count < 0 {
				return newErrorWithTrace("array multiplication count cannot be negative", node, ctx)
			}
			
			totalElements := len(leftArr.Elements) * int(count)
			result := make([]object.Object, totalElements)
			
			for i := 0; i < int(count); i++ {
				copy(result[i*len(leftArr.Elements):], leftArr.Elements)
			}
			
			return &object.Array{Elements: result}
		}
		return newErrorWithTrace("unsupported operation: %s %s %s", node, ctx, unwrappedLeft.Type(), operator, unwrappedRight.Type())
	case unwrappedLeft.Type() == object.INTEGER_OBJ && unwrappedRight.Type() == object.ARRAY_OBJ:
		// Integer * array: 3 * [1, 2]
		if operator == "*" {
			count := unwrappedLeft.(*object.Integer).Value
			rightArr := unwrappedRight.(*object.Array)
			if count < 0 {
				return newErrorWithTrace("array multiplication count cannot be negative", node, ctx)
			}
			
			totalElements := len(rightArr.Elements) * int(count)
			result := make([]object.Object, totalElements)
			
			for i := 0; i < int(count); i++ {
				copy(result[i*len(rightArr.Elements):], rightArr.Elements)
			}
			
			return &object.Array{Elements: result}
		}
		return newErrorWithTrace("unsupported operation: %s %s %s", node, ctx, unwrappedLeft.Type(), operator, unwrappedRight.Type())
	case unwrappedLeft == object.NONE || unwrappedRight == object.NONE:
		if operator == "==" {
			return nativeBoolToBooleanObject(false)
		} else if operator == "!=" {
			return nativeBoolToBooleanObject(true)
		}
		return newErrorWithTrace("operation not supported with None: %s", node, ctx, operator)
	case left.Type() == object.INSTANCE_OBJ && right.Type() == object.INSTANCE_OBJ:
		// Handle instance operations specially
		leftInstance := left.(*object.Instance)
		rightInstance := right.(*object.Instance)
		
		// Check if both are Array instances and handle concatenation
		if leftInstance.Grimoire.Name == "Array" && rightInstance.Grimoire.Name == "Array" && operator == "+" {
			// Get the actual arrays from the instances
			var leftArray, rightArray *object.Array
			
			if leftElements, exists := leftInstance.Env.Get("elements"); exists {
				if arr, isArray := leftElements.(*object.Array); isArray {
					leftArray = arr
				} else if elemInstance, isInstance := leftElements.(*object.Instance); isInstance {
					if innerValue, innerExists := elemInstance.Env.Get("value"); innerExists {
						if arr, isArray := innerValue.(*object.Array); isArray {
							leftArray = arr
						}
					}
				}
			}
			
			if rightElements, exists := rightInstance.Env.Get("elements"); exists {
				if arr, isArray := rightElements.(*object.Array); isArray {
					rightArray = arr
				} else if elemInstance, isInstance := rightElements.(*object.Instance); isInstance {
					if innerValue, innerExists := elemInstance.Env.Get("value"); innerExists {
						if arr, isArray := innerValue.(*object.Array); isArray {
							rightArray = arr
						}
					}
				}
			}
			
			if leftArray != nil && rightArray != nil {
				combined := make([]object.Object, len(leftArray.Elements)+len(rightArray.Elements))
				copy(combined, leftArray.Elements)
				copy(combined[len(leftArray.Elements):], rightArray.Elements)
				result := &object.Array{Elements: combined}
				
				// Wrap the result back as an Array instance if we're in an environment with Array grimoire
				if ctx != nil && ctx.env != nil {
					return wrapPrimitive(result, ctx.env, ctx)
				}
				return result
			}
		}
		
		// If not handled above, fall through to type mismatch error
		return newErrorWithTrace("type mismatch: %s %s %s", node, ctx,
			left.Type(), operator, right.Type())
	case unwrappedLeft.Type() != unwrappedRight.Type():
		return newErrorWithTrace("type mismatch: %s %s %s", node, ctx,
			unwrappedLeft.Type(), operator, unwrappedRight.Type())
	case unwrappedLeft.Type() == object.FLOAT_OBJ || unwrappedRight.Type() == object.FLOAT_OBJ:
		leftVal := toFloat(unwrappedLeft)
		rightVal := toFloat(unwrappedRight)
		switch operator {
		case "+":
			return &object.Float{Value: leftVal + rightVal}
		case "-":
			return &object.Float{Value: leftVal - rightVal}
		case "*":
			return &object.Float{Value: leftVal * rightVal}
		case "/":
			if rightVal == 0 {
				return newErrorWithTrace("division by zero", node, ctx)
			}
			return &object.Float{Value: leftVal / rightVal}
		case "**":
			return &object.Float{Value: math.Pow(leftVal, rightVal)}
		case "<":
			return nativeBoolToBooleanObject(leftVal < rightVal)
		case ">":
			return nativeBoolToBooleanObject(leftVal > rightVal)
		case "<=":
			return nativeBoolToBooleanObject(leftVal <= rightVal)
		case ">=":
			return nativeBoolToBooleanObject(leftVal >= rightVal)
		case "==":
			return nativeBoolToBooleanObject(leftVal == rightVal)
		case "!=":
			return nativeBoolToBooleanObject(leftVal != rightVal)
		default:
			return newErrorWithTrace("unknown operator: %s %s %s", node, ctx,
				left.Type(), operator, right.Type())
		}
	}

	return newErrorWithTrace(
		"unknown operator or type mismatch: %s %s %s",
		node, ctx,
		left.Type(),
		operator,
		right.Type(),
	)
}

func toFloat(obj object.Object) float64 {
	switch obj := obj.(type) {
	case *object.Integer:
		return float64(obj.Value)
	case *object.Float:
		return obj.Value
	default:
		return 0.0
	}
}

func evalStringInfixExpression(
	operator string,
	left, right object.Object,
	node ast.Node,
	ctx *CallContext,
) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	
	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	default:
		return newErrorWithTrace("unknown operator: %s %s %s",
			node, ctx, left.Type(), operator, right.Type())
	}
}

func evalBooleanInfixExpression(
	operator string,
	left, right object.Object,
	node ast.Node,
	ctx *CallContext,
) object.Object {
	leftVal := left.(*object.Boolean).Value
	rightVal := right.(*object.Boolean).Value
	switch operator {
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newErrorWithTrace("unknown operator: %s %s %s",
			node, ctx, left.Type(), operator, right.Type())
	}
}

func evalPrefixIncrementDecrement(
	operator string,
	node *ast.PrefixExpression,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	switch operand := node.Right.(type) {
	case *ast.Identifier:
		obj, ok := env.Get(operand.Value)
		if !ok {
			return newErrorWithTrace("undefined variable '%s'", node, ctx, operand.Value)
		}

		// Handle both direct Integer objects and Instance-wrapped integers
		var intObj *object.Integer
		if directInt, ok := obj.(*object.Integer); ok {
			intObj = directInt
		} else if instance, ok := obj.(*object.Instance); ok {
			// Check if this is an Integer wrapper instance
			if instance.Grimoire.Name == "Integer" {
				if value, exists := instance.Env.Get("value"); exists {
					if wrappedInt, ok := value.(*object.Integer); ok {
						intObj = wrappedInt
					}
				}
			}
		}

		if intObj == nil {
			return newErrorWithTrace("prefix '%s' operator requires an integer variable '%s'",
				node, ctx, operator, operand.Value)
		}

		if operator == "++" {
			intObj.Value += 1
		} else if operator == "--" {
			intObj.Value -= 1
		}

		// If it was an instance, update the instance's value
		if instance, ok := obj.(*object.Instance); ok {
			instance.Env.Set("value", intObj)
			env.SetWithGlobalCheck(operand.Value, instance)
		} else {
			env.SetWithGlobalCheck(operand.Value, intObj)
		}
		return intObj

	default:
		return newErrorWithTrace("prefix '%s' operator requires an integer or identifier",
			node, ctx, operator)
	}
}

func evalPostfixIncrementDecrement(
	operator string,
	node *ast.PostfixExpression,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	switch operand := node.Left.(type) {
	case *ast.Identifier:
		obj, ok := env.Get(operand.Value)
		if !ok {
			return newErrorWithTrace("undefined variable '%s'", node, ctx, operand.Value)
		}

		// Handle both direct Integer objects and Instance-wrapped integers
		var intObj *object.Integer
		if directInt, ok := obj.(*object.Integer); ok {
			intObj = directInt
		} else if instance, ok := obj.(*object.Instance); ok {
			// Check if this is an Integer wrapper instance
			if instance.Grimoire.Name == "Integer" {
				if value, exists := instance.Env.Get("value"); exists {
					if wrappedInt, ok := value.(*object.Integer); ok {
						intObj = wrappedInt
					}
				}
			}
		}

		if intObj == nil {
			return newErrorWithTrace("postfix '%s' operator requires an integer variable '%s'",
				node, ctx, operator, operand.Value)
		}

		oldValue := intObj.Value

		var newValue int64
		if operator == "++" {
			newValue = oldValue + 1
		} else if operator == "--" {
			newValue = oldValue - 1
		}

		newIntObj := &object.Integer{Value: newValue}

		// If it was an instance, update the instance's value
		if instance, ok := obj.(*object.Instance); ok {
			instance.Env.Set("value", newIntObj)
			env.SetWithGlobalCheck(operand.Value, instance)
		} else {
			env.SetWithGlobalCheck(operand.Value, newIntObj)
		}

		return &object.Integer{Value: oldValue}
	default:
		return newErrorWithTrace("postfix '%s' operator requires an integer or identifier",
			node, ctx, operator)
	}
}

func evalBangOperatorExpression(
	right object.Object,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NONE:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(
	right object.Object,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	// Unwrap primitive values from instances if needed
	unwrapped := unwrapPrimitive(right)
	
	if unwrapped.Type() != object.INTEGER_OBJ && unwrapped.Type() != object.FLOAT_OBJ {
		// Unknown operand type for prefix minus
		return newError("unknown operator: -%s", right.Type())
	}
	switch unwrapped := unwrapped.(type) {
	case *object.Integer:
		return &object.Integer{Value: -unwrapped.Value}
	case *object.Float:
		return &object.Float{Value: -unwrapped.Value}
	default:
		// Fallback for unexpected types
		return newError("unknown type for minus operator: %s", unwrapped.Type())
	}
}

func evalIntegerInfixExpression(
	operator string,
	left, right object.Object,
	node ast.Node,
	ctx *CallContext,
) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	
	if debugPrimitiveWrapping && operator == "//" {
		fmt.Fprintf(os.Stderr, "INTDIV: %d %s %d in %s\n", leftVal, operator, rightVal, getContextName(ctx))
	}
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newErrorWithTrace("division by zero", node, ctx)
		}
		return &object.Integer{Value: leftVal / rightVal}
	case "%":
		if rightVal == 0 {
			return newErrorWithTrace("modulo by zero", node, ctx)
		}
		return &object.Integer{Value: leftVal % rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "**":
		return &object.Integer{Value: int64(math.Pow(float64(leftVal), float64(rightVal)))}
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case "<<":
		return &object.Integer{Value: leftVal << uint(rightVal)}
	case ">>":
		return &object.Integer{Value: leftVal >> uint(rightVal)}
	case "&":
		return &object.Integer{Value: leftVal & rightVal}
	case "^":
		return &object.Integer{Value: leftVal ^ rightVal}
	case "|":
		return &object.Integer{Value: leftVal | rightVal}
	case "//":
		if rightVal == 0 {
			return newErrorWithTrace("integer division by zero", node, ctx)
		}
		return &object.Integer{Value: leftVal / rightVal}
	default:
		return newErrorWithTrace("unknown operator: %s %s %s",
			node, ctx, left.Type(), operator, right.Type())
	}
}

func evalCompoundAssignment(
	node *ast.InfixExpression,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	rightVal := Eval(node.Right, env, ctx)
	if isError(rightVal) {
		return rightVal
	}

	switch leftNode := node.Left.(type) {
	case *ast.Identifier:
		currVal, ok := env.Get(leftNode.Value)
		if !ok {
			return newErrorWithTrace("undefined variable: %s", node, ctx, leftNode.Value)
		}

		newVal := applyCompoundOperator(node.Operator, currVal, rightVal, node, ctx)
		if isError(newVal) {
			return newVal
		}

		// If the current value is an Instance, update its wrapped value
		if instance, ok := currVal.(*object.Instance); ok {
			instance.Env.Set("value", newVal)
			env.SetWithGlobalCheck(leftNode.Value, instance)
			return newVal
		} else {
			env.SetWithGlobalCheck(leftNode.Value, newVal)
			return newVal
		}

	default:
		return newErrorWithTrace("invalid assignment target: %T", node, ctx, leftNode)
	}
}

func applyCompoundOperator(
	operator string,
	leftVal, rightVal object.Object,
	node ast.Node,
	ctx *CallContext,
) object.Object {
	// Helper function to extract the actual integer/float value from either direct objects or Instance wrappers
	extractInteger := func(obj object.Object) (*object.Integer, bool) {
		if intObj, ok := obj.(*object.Integer); ok {
			return intObj, true
		}
		if instance, ok := obj.(*object.Instance); ok {
			if instance.Grimoire.Name == "Integer" {
				if value, exists := instance.Env.Get("value"); exists {
					if intObj, ok := value.(*object.Integer); ok {
						return intObj, true
					}
				}
			}
		}
		return nil, false
	}

	extractFloat := func(obj object.Object) (*object.Float, bool) {
		if floatObj, ok := obj.(*object.Float); ok {
			return floatObj, true
		}
		if instance, ok := obj.(*object.Instance); ok {
			if instance.Grimoire.Name == "Float" {
				if value, exists := instance.Env.Get("value"); exists {
					if floatObj, ok := value.(*object.Float); ok {
						return floatObj, true
					}
				}
			}
		}
		return nil, false
	}

	// Try to extract integers first
	if lInt, ok := extractInteger(leftVal); ok {
		rInt, ok := extractInteger(rightVal)
		if !ok {
			return newErrorWithTrace("type mismatch: expected INTEGER, got %s",
				node, ctx, rightVal.Type())
		}
		switch operator {
		case "+=":
			return &object.Integer{Value: lInt.Value + rInt.Value}
		case "-=":
			return &object.Integer{Value: lInt.Value - rInt.Value}
		case "*=":
			return &object.Integer{Value: lInt.Value * rInt.Value}
		case "/=":
			if rInt.Value == 0 {
				return newErrorWithTrace("division by zero", node, ctx)
			}
			return &object.Integer{Value: lInt.Value / rInt.Value}
		default:
			return newErrorWithTrace("unknown operator: %s", node, ctx, operator)
		}
	}

	// Try to extract floats 
	if lFloat, ok := extractFloat(leftVal); ok {
		rFloat, ok := extractFloat(rightVal)
		if !ok {
			return newErrorWithTrace("type mismatch: expected FLOAT, got %s",
				node, ctx, rightVal.Type())
		}
		switch operator {
		case "+=":
			return &object.Float{Value: lFloat.Value + rFloat.Value}
		case "-=":
			return &object.Float{Value: lFloat.Value - rFloat.Value}
		case "*=":
			return &object.Float{Value: lFloat.Value * rFloat.Value}
		case "/=":
			if rFloat.Value == 0 {
				return newErrorWithTrace("division by zero", node, ctx)
			}
			return &object.Float{Value: lFloat.Value / rFloat.Value}
		default:
			return newErrorWithTrace("unknown operator: %s", node, ctx, operator)
		}
	}

	// Handle string concatenation with +=
	extractString := func(obj object.Object) (*object.String, bool) {
		if strObj, ok := obj.(*object.String); ok {
			return strObj, true
		}
		if instance, ok := obj.(*object.Instance); ok {
			if instance.Grimoire.Name == "String" {
				if value, exists := instance.Env.Get("value"); exists {
					if strObj, ok := value.(*object.String); ok {
						return strObj, true
					}
				}
			}
		}
		return nil, false
	}
	
	if operator == "+=" {
		if lStr, ok := extractString(leftVal); ok {
			if rStr, ok := extractString(rightVal); ok {
				return &object.String{Value: lStr.Value + rStr.Value}
			}
		}
	}

	return newErrorWithTrace("unsupported type for compound assignment: %s",
		node, ctx, leftVal.Type())
}

func evalIfExpression(
	ie *ast.IfStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	condition := Eval(ie.Condition, env, ctx)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		ifCtx := &CallContext{
			FunctionName: "if_block",
			Node:         ie.Consequence,
			Parent:       ctx,
			env:          env,
		}
		return Eval(ie.Consequence, env, ifCtx)
	}

	for _, branch := range ie.OtherwiseBranches {
		condition = Eval(branch.Condition, env, ctx)
		if isError(condition) {
			return condition
		}
		if isTruthy(condition) {
			otherwiseCtx := &CallContext{
				FunctionName: "otherwise_block",
				Node:         branch.Consequence,
				Parent:       ctx,
				env:          env,
			}
			return Eval(branch.Consequence, env, otherwiseCtx)
		}
	}

	if ie.Alternative != nil {
		elseCtx := &CallContext{
			FunctionName: "else_block",
			Node:         ie.Alternative,
			Parent:       ctx,
			env:          env,
		}
		return Eval(ie.Alternative, env, elseCtx)
	}

	return NONE
}

func isError(obj object.Object) bool {
	if obj == nil {
		return false
	}
	return obj.Type() == object.ERROR_OBJ ||
		obj.Type() == object.CUSTOM_ERROR_OBJ ||
		isErrorWithTrace(obj)
}

// isStopIterationError checks if an object represents a StopIteration error
func isStopIterationError(obj object.Object) bool {
	switch err := obj.(type) {
	case *object.Error:
		return strings.Contains(err.Message, "StopIteration")
	case *object.CustomError:
		// Check the name directly
		if err.Name == "StopIteration" {
			return true
		}
		// Check the message
		if strings.Contains(err.Message, "StopIteration") {
			return true
		}
		// Check details for StopIteration pattern
		if details, ok := err.Details["errorType"]; ok {
			if strType, ok := details.(*object.String); ok && strType.Value == "String" {
				if instance, ok := err.Details["instance"]; ok {
					if strInstance, ok := instance.(*object.String); ok && strInstance.Value == "StopIteration" {
						return true
					}
				}
			}
		}
		return false
	case *object.ErrorWithTrace:
		// Check the message
		if strings.Contains(err.Message, "StopIteration") {
			return true
		}
		// Check CustomDetails for StopIteration pattern
		if err.CustomDetails != nil {
			if errorType, ok := err.CustomDetails["errorType"]; ok {
				if strType, ok := errorType.(*object.String); ok && strType.Value == "String" {
					if instance, ok := err.CustomDetails["instance"]; ok {
						if strInstance, ok := instance.(*object.String); ok && strInstance.Value == "StopIteration" {
							return true
						}
					}
				}
			}
		}
		return false
	default:
		return false
	}
}

func evalWhileStatement(
	node *ast.WhileStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	whileCtx := &CallContext{
		FunctionName: "while_loop",
		Node:         node,
		Parent:       ctx,
		env:          env,
	}

	for {
		condition := Eval(node.Condition, env, whileCtx)
		if isError(condition) {
			return condition
		}
		if !isTruthy(condition) {
			break
		}

		n := len(node.Body.Statements)
		var controlSignal object.Object = nil

		for i := 0; i < n-1; i++ {
			stmtCtx := &CallContext{
				FunctionName: "while_statement",
				Node:         node.Body.Statements[i],
				Parent:       whileCtx,
				env:          env,
			}

			res := Eval(node.Body.Statements[i], env, stmtCtx)

			rt := getObjectType(res)
			if rt == string(object.STOP.Type()) || rt == string(object.SKIP.Type()) ||
				rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ ||
				rt == object.CUSTOM_ERROR_OBJ || isErrorWithTrace(res) {
				controlSignal = res
				break
			}
		}

		if n > 0 {
			lastStmtCtx := &CallContext{
				FunctionName: "while_last_statement",
				Node:         node.Body.Statements[n-1],
				Parent:       whileCtx,
				env:          env,
			}
			_ = Eval(node.Body.Statements[n-1], env, lastStmtCtx)
		}

		if controlSignal != nil {
			rt := getObjectType(controlSignal)
			if rt == string(object.STOP.Type()) {
				break
			}
			if rt == string(object.SKIP.Type()) {
				continue
			}
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ ||
				rt == object.CUSTOM_ERROR_OBJ || isErrorWithTrace(controlSignal) {
				return controlSignal
			}
		}
	}
	return NONE
}

func getObjectType(obj object.Object) string {
	if obj == nil {
		return ""
	}
	return string(obj.Type())
}

func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Boolean:
		return obj.Value
	case *object.String:
		return len(obj.Value) > 0
	case *object.Array:
		return len(obj.Elements) > 0
	case *object.Tuple:
		return len(obj.Elements) > 0
	case *object.Hash:
		return len(obj.Pairs) > 0
	case *object.None:
		return false
	default:
		return true
	}
}

func evalForStatement(
	fs *ast.ForStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	iterable := Eval(fs.Iterable, env, ctx)
	if isError(iterable) {
		return iterable
	}

	forCtx := &CallContext{
		FunctionName: "for_loop",
		Node:         fs,
		Parent:       ctx,
		env:          env,
	}

	switch iter := iterable.(type) {
	case *object.Array:
		result := processArrayIteration(iter.Elements, fs, env, forCtx, ctx)
		if result != NONE {
			return result
		}
	case *object.String:
		// Convert string to array of character strings for iteration
		charElements := make([]object.Object, len(iter.Value))
		for i, char := range iter.Value {
			charElements[i] = &object.String{Value: string(char)}
		}
		result := processArrayIteration(charElements, fs, env, forCtx, ctx)
		if result != NONE {
			return result
		}
	case *object.Hash:
		// Iterate over hash keys by default
		var elements []object.Object
		for _, pair := range iter.Pairs {
			elements = append(elements, pair.Key)
		}
		result := processArrayIteration(elements, fs, env, forCtx, ctx)
		if result != NONE {
			return result
		}
	case *object.Instance:
		// First check if the instance has an iter method
		if _, ok := iter.Grimoire.Methods["iter"]; ok {
			// Call iter to get an iterator
			iteratorObj := evalGrimoireMethodCall(iter, "iter", []object.Object{}, env, forCtx)
			if isError(iteratorObj) {
				return iteratorObj
			}
			
			// Use the iterator to iterate
			if iterator, ok := iteratorObj.(*object.Instance); ok {
				// Process elements one at a time instead of collecting all first
				for {
					if _, hasNext := iterator.Grimoire.Methods["next"]; hasNext {
						nextValue := evalGrimoireMethodCall(iterator, "next", []object.Object{}, env, forCtx)
						
						// Check if it's a StopIteration error
						if isStopIterationError(nextValue) {
							break
						}
						if isError(nextValue) {
							return nextValue
						}
						
						// Process each element immediately
						switch varExpr := fs.Variable.(type) {
						case *ast.Identifier:
							env.Set(varExpr.Value, nextValue)
						case *ast.TupleLiteral:
							var items []object.Object
							if tupObj, ok := nextValue.(*object.Tuple); ok {
								items = tupObj.Elements
							} else if arrObj, ok := nextValue.(*object.Array); ok {
								items = arrObj.Elements
							} else {
								return newErrorWithTrace(fmt.Sprintf("cannot unpack non-iterable element: %s", nextValue.Type()), fs, ctx)
							}
							if len(varExpr.Elements) > len(items) {
								return newErrorWithTrace("not enough values to unpack", fs, ctx)
							}
							for i, elem := range varExpr.Elements {
								ident, ok := elem.(*ast.Identifier)
								if !ok {
									return newErrorWithTrace("invalid assignment target in for loop", fs, ctx)
								}
								env.Set(ident.Value, items[i])
							}
						default:
							env.Set(fs.Variable.String(), nextValue)
						}
						
						if fs.Body != nil {
							loopResult := Eval(fs.Body, env, forCtx)
							if loopResult != nil {
								rt := getObjectType(loopResult)
								if rt == string(object.STOP.Type()) {
									return object.STOP
								}
								if rt == string(object.SKIP.Type()) {
									continue
								}
								if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ ||
									rt == object.CUSTOM_ERROR_OBJ || isErrorWithTrace(loopResult) {
									return loopResult
								}
							}
						}
					} else {
						return newErrorWithTrace("Iterator must have next method", fs, ctx)
					}
				}
				// Iterator exhausted, return NONE
				return NONE
			} else {
				return newErrorWithTrace("iter must return an iterator instance", fs, ctx)
			}
		} else if iter.Grimoire.Name == "Array" {
			// Handle Array instances
			if elementsObj, ok := iter.Env.Get("elements"); ok {
				if arr, ok := elementsObj.(*object.Array); ok {
					result := processArrayIteration(arr.Elements, fs, env, forCtx, ctx)
					if result != NONE {
						return result
					}
				}
			}
		} else if iter.Grimoire.Name == "String" {
			// Handle String instances by getting the value and converting to characters
			if valueObj, ok := iter.Env.Get("value"); ok {
				if str, ok := valueObj.(*object.String); ok {
					charElements := make([]object.Object, len(str.Value))
					for i, char := range str.Value {
						charElements[i] = &object.String{Value: string(char)}
					}
					result := processArrayIteration(charElements, fs, env, forCtx, ctx)
					if result != NONE {
						return result
					}
				}
			}
		} else {
			return newErrorWithTrace("for loop: %s instance is not iterable", fs, ctx, iter.Grimoire.Name)
		}
	default:
		return newErrorWithTrace("for loop requires an iterable, got %s", fs, ctx, iterable.Type())
	}
	return NONE
}

// evalInOperator handles "in" operator for membership testing across different container types
func evalInOperator(
	left, right object.Object,
	node ast.Node,
	ctx *CallContext,
) object.Object {
	switch container := right.(type) {
	case *object.String:
		// Check if left is a substring or character in the string
		if leftStr, ok := left.(*object.String); ok {
			contains := strings.Contains(container.Value, leftStr.Value)
			return nativeBoolToBooleanObject(contains)
		}
		return newErrorWithTrace("'in' operator with string requires string on left side, got %s", node, ctx, left.Type())
		
	case *object.Array:
		// Check if element exists in array
		for _, elem := range container.Elements {
			if isObjectEqual(left, elem) {
				return nativeBoolToBooleanObject(true)
			}
		}
		return nativeBoolToBooleanObject(false)
		
	case *object.Hash:
		// Check if key exists in hash
		hashKey, ok := left.(object.Hashable)
		if !ok {
			return newErrorWithTrace("unusable as hash key: %T", node, ctx, left)
		}
		_, exists := container.Pairs[hashKey.HashKey()]
		return nativeBoolToBooleanObject(exists)
		
	case *object.Instance:
		// Handle wrapped containers
		if container.Grimoire.Name == "String" {
			if valueObj, ok := container.Env.Get("value"); ok {
				if str, ok := valueObj.(*object.String); ok {
					if leftStr, ok := left.(*object.String); ok {
						contains := strings.Contains(str.Value, leftStr.Value)
						return nativeBoolToBooleanObject(contains)
					}
					return newErrorWithTrace("'in' operator with string requires string on left side, got %s", node, ctx, left.Type())
				}
			}
		} else if container.Grimoire.Name == "Array" {
			if elementsObj, ok := container.Env.Get("elements"); ok {
				if arr, ok := elementsObj.(*object.Array); ok {
					for _, elem := range arr.Elements {
						if isObjectEqual(left, elem) {
							return nativeBoolToBooleanObject(true)
						}
					}
					return nativeBoolToBooleanObject(false)
				}
			}
		}
		return newErrorWithTrace("'in' operator not supported for %s instance", node, ctx, container.Grimoire.Name)
		
	default:
		return newErrorWithTrace("'in' operator not supported for %s", node, ctx, right.Type())
	}
}

// Helper function to check if two objects are equal
func isObjectEqual(left, right object.Object) bool {
	// Unwrap primitives to handle instances
	unwrappedLeft := unwrapPrimitive(left)
	unwrappedRight := unwrapPrimitive(right)
	
	if unwrappedLeft.Type() != unwrappedRight.Type() {
		return false
	}
	
	switch leftObj := unwrappedLeft.(type) {
	case *object.Integer:
		rightObj := unwrappedRight.(*object.Integer)
		return leftObj.Value == rightObj.Value
	case *object.Float:
		rightObj := unwrappedRight.(*object.Float)
		return leftObj.Value == rightObj.Value
	case *object.String:
		rightObj := unwrappedRight.(*object.String)
		return leftObj.Value == rightObj.Value
	case *object.Boolean:
		rightObj := unwrappedRight.(*object.Boolean)
		return leftObj.Value == rightObj.Value
	default:
		// For other types, use pointer comparison as fallback
		return unwrappedLeft == unwrappedRight
	}
}

// Helper function to process array iteration with variable unpacking and loop body execution
func processArrayIteration(
	elements []object.Object,
	fs *ast.ForStatement,
	env *object.Environment,
	forCtx *CallContext,
	ctx *CallContext,
) object.Object {
	for _, elem := range elements {
		switch varExpr := fs.Variable.(type) {
		case *ast.Identifier:
			env.Set(varExpr.Value, elem)
		case *ast.TupleLiteral:
			var items []object.Object
			if tupObj, ok := elem.(*object.Tuple); ok {
				items = tupObj.Elements
			} else if arrObj, ok := elem.(*object.Array); ok {
				items = arrObj.Elements
			} else {
				return newErrorWithTrace("cannot unpack non-iterable element: %s",
					fs, ctx, elem.Type())
			}
			if len(varExpr.Elements) != len(items) {
				return newErrorWithTrace("unpacking mismatch: expected %d values, got %d",
					fs, ctx, len(varExpr.Elements), len(items))
			}
			for i, target := range varExpr.Elements {
				ident, ok := target.(*ast.Identifier)
				if !ok {
					return newErrorWithTrace("invalid assignment target in for loop", fs, ctx)
				}
				env.Set(ident.Value, items[i])
			}
		default:
			env.Set(fs.Variable.String(), elem)
		}

		if fs.Body != nil {
			loopResult := Eval(fs.Body, env, forCtx)
			if loopResult != nil {
				rt := getObjectType(loopResult)
				if rt == string(object.STOP.Type()) {
					return object.STOP // Exit the entire for loop
				}
				if rt == string(object.SKIP.Type()) {
					continue // Skip to the next element in the outer loop
				}
				if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ ||
					rt == object.CUSTOM_ERROR_OBJ || isErrorWithTrace(loopResult) {
					return loopResult
				}
			}
		}
	}
	return NONE
}

// resolveImportPath searches for an import file with smart resolution
func resolveImportPath(importPath string) (string, error) {
	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		currentDir = "."
	}
	
	// Handle relative imports explicitly (starts with . or ..)
	if strings.HasPrefix(importPath, "./") || strings.HasPrefix(importPath, "../") {
		return resolveRelativeImport(importPath, currentDir)
	}
	
	// Check if this looks like a package import
	// Pattern analysis:
	// - "packagename.filename" -> package import (look in carrion_modules)
	// - "filename" -> local file (look in current dir first)
	// - "path/filename" -> explicit path
	
	return resolvePackageOrFileImport(importPath, currentDir)
}

// resolveRelativeImport handles relative imports like ../file or ./file
func resolveRelativeImport(importPath, currentDir string) (string, error) {
	// Remove any .crl extension if provided
	cleanPath := strings.TrimSuffix(importPath, ".crl")
	
	// Resolve the relative path
	fullPath := filepath.Join(currentDir, cleanPath+".crl")
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return "", fmt.Errorf("could not resolve relative path %s: %v", importPath, err)
	}
	
	if _, err := os.Stat(absPath); err == nil {
		return absPath, nil
	}
	
	return "", fmt.Errorf("relative import not found: %s", importPath)
}

// resolvePackageOrFileImport handles package and file imports with smart resolution
func resolvePackageOrFileImport(importPath, currentDir string) (string, error) {
	// Define search paths in order of priority
	searchPaths := []string{
		// 1. Current directory (for local files)
		currentDir,
		// 2. Local project modules 
		filepath.Join(currentDir, "carrion_modules"),
		// 3. Global bifrost modules
		"/usr/bin/carrion_modules",
		// 4. User-specific packages (~/.carrion/packages)
		getUserCarrionPackages(),
		// 5. Shared global packages (/usr/local/share/carrion/lib)
		getSharedGlobalPackages(),
	}
	
	// Smart import resolution logic
	if strings.Contains(importPath, "/") {
		// Explicit path provided - try as package/file structure
		return resolveExplicitPath(importPath, searchPaths)
	} else {
		// Simple name - could be local file or package
		return resolveSimpleName(importPath, searchPaths, currentDir)
	}
}

// resolveExplicitPath handles imports with slashes like "package/file" or "path/to/file"
func resolveExplicitPath(importPath string, searchPaths []string) (string, error) {
	parts := strings.Split(importPath, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid explicit path: %s", importPath)
	}
	
	packageName := parts[0]
	subPath := strings.Join(parts[1:], "/")
	
	// Try each search path for package structure
	for _, basePath := range searchPaths {
		if basePath == "" {
			continue
		}
		
		// Try versioned package structure: basePath/package/version/src/file.crl
		packagePath := filepath.Join(basePath, packageName)
		if versions, err := getLatestPackageVersion(packagePath); err == nil && len(versions) > 0 {
			latestVersion := versions[len(versions)-1]
			versionedPath := filepath.Join(packagePath, latestVersion, "src", subPath+".crl")
			if _, err := os.Stat(versionedPath); err == nil {
				return versionedPath, nil
			}
		}
		
		// Try direct path: basePath/package/file.crl
		directPath := filepath.Join(basePath, importPath+".crl")
		if _, err := os.Stat(directPath); err == nil {
			return directPath, nil
		}
	}
	
	return "", fmt.Errorf("explicit path import not found: %s", importPath)
}

// resolveSimpleName handles simple imports like "filename" - looks for local files first, then packages
func resolveSimpleName(importPath string, searchPaths []string, currentDir string) (string, error) {
	// First, try as a local file in current directory
	localPath := filepath.Join(currentDir, importPath+".crl")
	if _, err := os.Stat(localPath); err == nil {
		return localPath, nil
	}
	
	// Then try as a package name in each search location
	for i, basePath := range searchPaths {
		if basePath == "" {
			continue
		}
		
		// Skip current directory since we already tried it
		if i == 0 {
			continue
		}
		
		// Look for package with this name
		packagePath := filepath.Join(basePath, importPath)
		if versions, err := getLatestPackageVersion(packagePath); err == nil && len(versions) > 0 {
			latestVersion := versions[len(versions)-1]
			
			// Try common entry points in package
			possibleEntries := []string{
				filepath.Join(packagePath, latestVersion, "src", importPath+".crl"),
				filepath.Join(packagePath, latestVersion, "src", "main.crl"),
				filepath.Join(packagePath, latestVersion, "src", "index.crl"),
			}
			
			for _, entryPath := range possibleEntries {
				if _, err := os.Stat(entryPath); err == nil {
					return entryPath, nil
				}
			}
		}
		
		// Try direct file path as fallback
		directPath := filepath.Join(basePath, importPath+".crl")
		if _, err := os.Stat(directPath); err == nil {
			return directPath, nil
		}
	}
	
	return "", fmt.Errorf("import not found: %s", importPath)
}

// getUserCarrionPackages returns the user-specific Carrion packages directory
func getUserCarrionPackages() string {
	if home := os.Getenv("CARRION_HOME"); home != "" {
		return filepath.Join(home, "packages")
	}
	
	userHome, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(userHome, ".carrion", "packages")
}

// getSharedGlobalPackages returns the shared global packages directory
func getSharedGlobalPackages() string {
	if runtime.GOOS == "windows" {
		if programData := os.Getenv("ProgramData"); programData != "" {
			return filepath.Join(programData, "Carrion", "lib")
		}
		return filepath.Join("C:", "ProgramData", "Carrion", "lib")
	}
	return "/usr/local/share/carrion/lib"
}

// getLatestPackageVersion returns sorted list of versions for a package
func getLatestPackageVersion(packagePath string) ([]string, error) {
	entries, err := os.ReadDir(packagePath)
	if err != nil {
		return nil, err
	}
	
	var versions []string
	for _, entry := range entries {
		if entry.IsDir() {
			versions = append(versions, entry.Name())
		}
	}
	
	// TODO: Sort by semantic version - for now just return alphabetical
	return versions, nil
}

func evalImportStatement(
	node *ast.ImportStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	// Handle grimoire-only imports (import "GrimoireName")
	if node.FilePath.Value == "" && node.ClassName != nil {
		return evalGrimoireImport(node, env, ctx)
	}
	
	importPath := node.FilePath.Value
	
	// Check if already imported using the import path as key
	if importedFiles[importPath] {
		return object.NONE
	}
	importedFiles[importPath] = true

	// Resolve the import path to an actual file
	resolvedPath, err := resolveImportPath(importPath)
	if err != nil {
		return newErrorWithTrace("could not resolve import: %s", node, ctx, err)
	}

	fileContent, err := os.ReadFile(resolvedPath)
	if err != nil {
		return newErrorWithTrace("could not read import file: %s", node, ctx, err)
	}

	l := lexer.NewWithFilename(string(fileContent), resolvedPath)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		errorDetails := fmt.Sprintf("parsing errors in imported file %s:\n", resolvedPath)
		for _, err := range p.Errors() {
			errorDetails += fmt.Sprintf("- %s\n", err)
		}
		return newErrorWithTrace(errorDetails, node, ctx)
	}

	importEnv := object.NewEnclosedEnvironment(env)
	importCtx := &CallContext{
		FunctionName: "import_" + importPath,
		Node:         program,
		Parent:       ctx,
		env:          importEnv,
		IsDirectExecution: false,  // This is an import, not direct execution
	}

	evalResult := Eval(program, importEnv, importCtx)
	if isError(evalResult) {
		return newErrorWithTrace("error evaluating imported file %s: %s",
			node, ctx, resolvedPath, evalResult.Inspect())
	}

	namespace := &object.Namespace{Env: importEnv}

	if node.ClassName != nil {
		// Selective import: import only the specified grimoire
		className := node.ClassName.Value
		val, exists := importEnv.Get(className)
		if !exists {
			return newErrorWithTrace("grimoire '%s' not found in module '%s'", 
				node, ctx, className, importPath)
		}
		if val.Type() != object.GRIMOIRE_OBJ {
			return newErrorWithTrace("'%s' is not a grimoire in module '%s'", 
				node, ctx, className, importPath)
		}
		
		if node.Alias != nil {
			// Selective import with alias: bind the specific grimoire to the alias
			env.Set(node.Alias.Value, val)
		} else {
			// Selective import without alias: bind to original name
			env.Set(className, val)
		}
	} else if node.Alias != nil {
		// Module import with alias: create a namespace
		env.Set(node.Alias.Value, namespace)
	} else {
		// Import all grimoires from the module
		for _, name := range importEnv.GetNames() {
			val, _ := importEnv.Get(name)
			if val.Type() == object.GRIMOIRE_OBJ {
				env.Set(name, val)
			}
		}
	}

	return object.NONE
}

// evalGrimoireImport handles grimoire-only imports like import "HelloWorld" as hw
func evalGrimoireImport(
	node *ast.ImportStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	grimoireName := node.ClassName.Value
	
	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		currentDir = "."
	}
	
	// Define search paths for grimoire files
	searchPaths := []string{
		// 1. Current directory
		currentDir,
		// 2. Local project modules 
		filepath.Join(currentDir, "carrion_modules"),
		// 3. Global carrion modules
		"/usr/bin/carrion_modules",
		// 4. User-specific packages
		getUserCarrionPackages(),
		// 5. Shared global packages
		getSharedGlobalPackages(),
	}
	
	// Search for files containing the grimoire
	var foundGrimoire object.Object
	var foundInFile string
	
	for _, searchPath := range searchPaths {
		if searchPath == "" {
			continue
		}
		
		// Look for .crl files in the search path
		files, err := findCarrionFiles(searchPath)
		if err != nil {
			continue
		}
		
		// Check each file for the grimoire
		for _, filePath := range files {
			// Skip if already imported
			if importedFiles[filePath] {
				continue
			}
			
			// Read and parse the file
			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				continue
			}
			
			l := lexer.NewWithFilename(string(fileContent), filePath)
			p := parser.New(l)
			program := p.ParseProgram()
			
			if len(p.Errors()) > 0 {
				continue
			}
			
			// Evaluate the file in a temporary environment
			tempEnv := object.NewEnclosedEnvironment(env)
			tempCtx := &CallContext{
				FunctionName: "import_search",
				Node:         program,
				Parent:       ctx,
				env:          tempEnv,
				IsDirectExecution: false,
			}
			
			evalResult := Eval(program, tempEnv, tempCtx)
			if isError(evalResult) {
				continue
			}
			
			// Check if the grimoire exists in this file
			if val, exists := tempEnv.Get(grimoireName); exists && val.Type() == object.GRIMOIRE_OBJ {
				foundGrimoire = val
				foundInFile = filePath
				break
			}
		}
		
		if foundGrimoire != nil {
			break
		}
		
		// Also check in bifrost package structure
		packagePaths, _ := findBifrostPackages(searchPath)
		for _, pkgPath := range packagePaths {
			mainFile := filepath.Join(pkgPath, "src", "main.crl")
			if _, err := os.Stat(mainFile); err != nil {
				continue
			}
			
			// Skip if already imported
			if importedFiles[mainFile] {
				continue
			}
			
			fileContent, err := os.ReadFile(mainFile)
			if err != nil {
				continue
			}
			
			l := lexer.NewWithFilename(string(fileContent), mainFile)
			p := parser.New(l)
			program := p.ParseProgram()
			
			if len(p.Errors()) > 0 {
				continue
			}
			
			tempEnv := object.NewEnclosedEnvironment(env)
			tempCtx := &CallContext{
				FunctionName: "import_search",
				Node:         program,
				Parent:       ctx,
				env:          tempEnv,
				IsDirectExecution: false,
			}
			
			evalResult := Eval(program, tempEnv, tempCtx)
			if isError(evalResult) {
				continue
			}
			
			if val, exists := tempEnv.Get(grimoireName); exists && val.Type() == object.GRIMOIRE_OBJ {
				foundGrimoire = val
				foundInFile = mainFile
				break
			}
		}
		
		if foundGrimoire != nil {
			break
		}
	}
	
	if foundGrimoire == nil {
		return newErrorWithTrace("grimoire '%s' not found in any available module", 
			node, ctx, grimoireName)
	}
	
	// Mark the file as imported
	importedFiles[foundInFile] = true
	
	// Bind the grimoire to the environment
	if node.Alias != nil {
		env.Set(node.Alias.Value, foundGrimoire)
	} else {
		env.Set(grimoireName, foundGrimoire)
	}
	
	return object.NONE
}

// findCarrionFiles finds all .crl files in a directory (non-recursive for current dir, limited recursion for modules)
func findCarrionFiles(dir string) ([]string, error) {
	var files []string
	
	// For current directory, only check top level
	currentDir, _ := os.Getwd()
	if dir == currentDir {
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, err
		}
		
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".crl") {
				files = append(files, filepath.Join(dir, entry.Name()))
			}
		}
		return files, nil
	}
	
	// For other directories, do limited recursive search
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // Skip errors
		}
		
		// Skip hidden directories
		if info.IsDir() && strings.HasPrefix(filepath.Base(path), ".") && path != dir {
			return filepath.SkipDir
		}
		
		if !info.IsDir() && strings.HasSuffix(path, ".crl") {
			files = append(files, path)
		}
		
		// Don't recurse too deep
		if info.IsDir() && strings.Count(path, string(os.PathSeparator)) - strings.Count(dir, string(os.PathSeparator)) > 2 {
			return filepath.SkipDir
		}
		
		return nil
	})
	
	return files, err
}

// findBifrostPackages finds all bifrost package directories (pattern: package/version)
func findBifrostPackages(basePath string) ([]string, error) {
	var packages []string
	
	entries, err := os.ReadDir(basePath)
	if err != nil {
		return nil, err
	}
	
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		
		packagePath := filepath.Join(basePath, entry.Name())
		versions, err := os.ReadDir(packagePath)
		if err != nil {
			continue
		}
		
		for _, version := range versions {
			if version.IsDir() {
				versionPath := filepath.Join(packagePath, version.Name())
				// Check if it has src directory
				if _, err := os.Stat(filepath.Join(versionPath, "src")); err == nil {
					packages = append(packages, versionPath)
				}
			}
		}
	}
	
	return packages, nil
}

func evalGlobalStatement(
	node *ast.GlobalStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	// Mark each variable name as global in the current environment
	for _, name := range node.Names {
		env.MarkGlobal(name.Value)
	}
	return object.NONE
}

func evalUnpackStatement(
	node *ast.UnpackStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	// Evaluate the value being unpacked
	val := Eval(node.Value, env, ctx)
	if isError(val) {
		return val
	}
	
	// Handle different types of unpacking
	switch value := val.(type) {
	case *object.Instance:
		// Handle instances of Array, Tuple, Map grimoires
		switch value.Grimoire.Name {
		case "Array":
			// Extract the internal array from the instance
			if elements, ok := value.Env.Get("elements"); ok {
				if arr, ok := elements.(*object.Array); ok {
					return unpackArray(node.Variables, arr, env, ctx, node)
				}
			}
		case "Tuple":
			// Extract the internal tuple from the instance
			if elements, ok := value.Env.Get("elements"); ok {
				if tup, ok := elements.(*object.Tuple); ok {
					return unpackTuple(node.Variables, tup, env, ctx, node)
				}
			}
		case "Map", "Dict":
			// Extract the internal hash from the instance
			if pairs, ok := value.Env.Get("pairs"); ok {
				if hash, ok := pairs.(*object.Hash); ok {
					return unpackMap(node.Variables, hash, env, ctx, node)
				}
			}
		}
		return newErrorWithTrace(
			fmt.Sprintf("cannot unpack instance of %s", value.Grimoire.Name),
			node,
			ctx,
		)
	case *object.Array:
		return unpackArray(node.Variables, value, env, ctx, node)
	case *object.Tuple:
		return unpackTuple(node.Variables, value, env, ctx, node)
	case *object.Hash:
		return unpackMap(node.Variables, value, env, ctx, node)
	default:
		return newErrorWithTrace(
			fmt.Sprintf("cannot unpack object of type %T", val),
			node,
			ctx,
		)
	}
}

func unpackArray(
	variables []ast.Expression,
	arr *object.Array,
	env *object.Environment,
	ctx *CallContext,
	node ast.Node,
) object.Object {
	if len(variables) == 2 && len(arr.Elements) > 2 {
		// Special case: k, v <- [10, 20, 30]
		// k gets indices [0, 1, 2], v gets values [10, 20, 30]
		
		// Create indices array
		indices := &object.Array{Elements: []object.Object{}}
		for i := range arr.Elements {
			indices.Elements = append(indices.Elements, &object.Integer{Value: int64(i)})
		}
		
		// Assign to variables
		if ident, ok := variables[0].(*ast.Identifier); ok {
			env.Set(ident.Value, indices)
		}
		if ident, ok := variables[1].(*ast.Identifier); ok {
			env.Set(ident.Value, arr)
		}
		
		return object.NONE
	}
	
	// Regular unpacking: a, b, c <- [1, 2, 3]
	if len(variables) != len(arr.Elements) {
		return newErrorWithTrace(
			"cannot unpack %d values into %d variables",
			node, ctx, len(arr.Elements), len(variables))
	}
	
	for i, variable := range variables {
		if ident, ok := variable.(*ast.Identifier); ok {
			env.Set(ident.Value, arr.Elements[i])
		}
	}
	
	return object.NONE
}

func unpackTuple(
	variables []ast.Expression,
	tuple *object.Tuple,
	env *object.Environment,
	ctx *CallContext,
	node ast.Node,
) object.Object {
	// Regular unpacking
	if len(variables) != len(tuple.Elements) {
		return newErrorWithTrace(
			"cannot unpack %d values into %d variables",
			node, ctx, len(tuple.Elements), len(variables))
	}
	
	for i, variable := range variables {
		if ident, ok := variable.(*ast.Identifier); ok {
			env.Set(ident.Value, tuple.Elements[i])
		}
	}
	
	return object.NONE
}

func unpackMap(
	variables []ast.Expression,
	hash *object.Hash,
	env *object.Environment,
	ctx *CallContext,
	node ast.Node,
) object.Object {
	if len(variables) != 2 {
		return newErrorWithTrace(
			"map unpacking requires exactly 2 variables (keys, values)",
			node, ctx)
	}
	
	// Extract keys and values
	keys := &object.Array{Elements: []object.Object{}}
	values := &object.Array{Elements: []object.Object{}}
	
	for _, pair := range hash.Pairs {
		keys.Elements = append(keys.Elements, pair.Key)
		values.Elements = append(values.Elements, pair.Value)
	}
	
	// Assign to variables
	if ident, ok := variables[0].(*ast.Identifier); ok {
		env.Set(ident.Value, keys)
	}
	if ident, ok := variables[1].(*ast.Identifier); ok {
		env.Set(ident.Value, values)
	}
	
	return object.NONE
}

// Type checking helper functions
func getTypeString(expr ast.Expression) string {
	if ident, ok := expr.(*ast.Identifier); ok {
		return ident.Value
	}
	return "Unknown"
}

func getObjectTypeString(obj object.Object) string {
	switch o := obj.(type) {
	case *object.Instance:
		// For instances, return the grimoire name (which is the type)
		return o.Grimoire.Name
	case *object.Integer:
		return "Integer"
	case *object.Float:
		return "Float"
	case *object.String:
		return "String"
	case *object.Boolean:
		return "Boolean"
	case *object.Array:
		return "Array"
	case *object.Hash:
		return "Map"
	case *object.Tuple:
		return "Tuple"
	case *object.Function:
		return "Function"
	case *object.None:
		return "None"
	default:
		return "Unknown"
	}
}

func isTypeCompatible(expected, actual string) bool {
	// Exact match
	if expected == actual {
		return true
	}
	
	// Handle type hint aliases (int -> Integer, str -> String, etc.)
	expectedNormalized := normalizeTypeName(expected)
	actualNormalized := normalizeTypeName(actual)
	
	if expectedNormalized == actualNormalized {
		return true
	}
	
	// Special cases for numeric types
	if (expectedNormalized == "Integer" && actualNormalized == "Float") || 
	   (expectedNormalized == "Float" && actualNormalized == "Integer") {
		return true
	}
	
	// None can be assigned to any type
	if actualNormalized == "None" {
		return true
	}
	
	return false
}

// normalizeTypeName converts type hint names to their grimoire equivalents
func normalizeTypeName(typeName string) string {
	switch typeName {
	case "int":
		return "Integer"
	case "str":
		return "String"
	case "float":
		return "Float"
	case "bool":
		return "Boolean"
	case "list":
		return "Array"
	case "dict":
		return "Map"
	default:
		return typeName
	}
}

func checkParameterTypes(fn *object.Function, args []object.Object, ctx *CallContext) object.Object {
	for i, pExpr := range fn.Parameters {
		if param, ok := pExpr.(*ast.Parameter); ok && param.TypeHint != nil {
			if i < len(args) {
				expectedType := getTypeString(param.TypeHint)
				actualType := getObjectTypeString(args[i])
				if !isTypeCompatible(expectedType, actualType) {
					return newErrorWithTrace(
						"Type error: parameter '%s' expects %s but got %s",
						ctx.Node, ctx, param.Name.Value, expectedType, actualType)
				}
			}
		}
	}
	return nil
}

// Global goroutine manager
var globalGoroutineManager = object.NewGoroutineManager()

func evalDivergeStatement(
	node *ast.DivergeStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	// Create a new goroutine
	goroutine := &object.Goroutine{
		Done:      make(chan bool, 1),
		IsRunning: true,
	}
	
	// Set name if provided
	if node.Name != nil {
		goroutine.Name = node.Name.Value
		globalGoroutineManager.AddNamedGoroutine(goroutine.Name, goroutine)
	} else {
		globalGoroutineManager.AddAnonymousGoroutine(goroutine)
	}
	
	// Start the goroutine
	go func() {
		defer func() {
			// Recover from any panic to prevent deadlock
			if r := recover(); r != nil {
				// Convert panic to error and store it
				goroutine.Error = &object.Error{
					Message: fmt.Sprintf("Goroutine panic: %v", r),
				}
			}
			
			// Always ensure Done channel receives a value
			goroutine.IsRunning = false
			goroutine.Done <- true
		}()
		
		// Create a new environment for the goroutine
		goroutineEnv := object.NewEnclosedEnvironment(env)
		
		// Create a new context for the goroutine
		goroutineCtx := &CallContext{
			FunctionName: "diverge",
			Node:         node.Body,
			Parent:       ctx,
			env:          goroutineEnv,
		}
		
		// Execute the body
		result := Eval(node.Body, goroutineEnv, goroutineCtx)
		
		// Store the result or error
		if isError(result) {
			goroutine.Error = result
		} else {
			goroutine.Result = result
		}
	}()
	
	return goroutine
}

func evalConvergeStatement(
	node *ast.ConvergeStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	if len(node.Names) == 0 {
		// Wait for all goroutines
		
		// Wait for named goroutines
		namedGoroutines := globalGoroutineManager.GetAllNamedGoroutines()
		for _, goroutine := range namedGoroutines {
			// Wait for goroutine completion with race condition protection
			select {
			case <-goroutine.Done:
				// Goroutine completed
			default:
				// If Done channel doesn't have a value yet, wait for it
				if goroutine.IsRunning {
					<-goroutine.Done
				}
			}
		}
		
		// Wait for anonymous goroutines
		anonymousGoroutines := globalGoroutineManager.GetAllAnonymousGoroutines()
		for _, goroutine := range anonymousGoroutines {
			// Wait for goroutine completion with race condition protection
			select {
			case <-goroutine.Done:
				// Goroutine completed
			default:
				// If Done channel doesn't have a value yet, wait for it
				if goroutine.IsRunning {
					<-goroutine.Done
				}
			}
		}
		
		// Clear all goroutines
		globalGoroutineManager.ClearAll()
		
	} else {
		// Wait for specific named goroutines
		for _, nameExpr := range node.Names {
			nameIdent, ok := nameExpr.(*ast.Identifier)
			if !ok {
				return newErrorWithTrace("converge expects goroutine names", node, ctx)
			}
			
			goroutine, exists := globalGoroutineManager.GetNamedGoroutine(nameIdent.Value)
			if !exists {
				return newErrorWithTrace("goroutine '%s' not found", node, ctx, nameIdent.Value)
			}
			
			// Wait for goroutine completion regardless of IsRunning state
			// to handle race conditions where IsRunning might have changed
			select {
			case <-goroutine.Done:
				// Goroutine completed
			default:
				// If Done channel doesn't have a value yet, wait for it
				if goroutine.IsRunning {
					<-goroutine.Done
				}
			}
			
			// Remove from manager with proper cleanup
			globalGoroutineManager.RemoveAndCleanupNamed(nameIdent.Value)
		}
	}
	
	return object.NONE
}
