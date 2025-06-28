package evaluator

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

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
		if len(n.Statements) > 0 {
			return getNodeToken(n.Statements[0])
		}
		return nil
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
			entry := object.StackTraceEntry{
				FunctionName: currentCtx.FunctionName,
				Position:     nodePos,
			}
			err.Stack = append(err.Stack, entry)
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
		// Don't wrap arguments to builtin functions
		if isBuiltinFunction(ctx.FunctionName) {
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

			// For now, just set the value directly to avoid init method issues
			instance.Env.Set("self", instance)
			instance.Env.Set("value", obj)

			return instance
		}
	}

	// If grimoire not found, return the original object
	return obj
}

// unwrapPrimitive extracts the primitive value from a wrapped instance if applicable
func unwrapPrimitive(obj object.Object) object.Object {
	if instance, ok := obj.(*object.Instance); ok {
		if value, exists := instance.Env.Get("value"); exists {
			return value
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

	case *ast.StopStatement:
		return object.STOP
	case *ast.SkipStatement:
		return object.SKIP
	case *ast.CheckStatement:
		cond := Eval(node.Condition, env, ctx)
		if isError(cond) {
			return cond
		}
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
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env, ctx)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.Boolean:
		primitive := nativeBoolToBooleanObject(node.Value)
		return wrapPrimitive(primitive, env, ctx)
	case *ast.AssignStatement:
		return evalAssignStatement(node, env, ctx)
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
	case *ast.GrimoireDefinition:
		return evalGrimoireDefinition(node, env, ctx)
	case *ast.AttemptStatement:
		return evalAttemptStatement(node, env, ctx)
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

		details := make(map[string]object.Object)
		details["errorType"] = &object.String{Value: instance.Grimoire.Name}
		details["instance"] = instance

		return newCustomErrorWithTrace(instance.Grimoire.Name, message, node, ctx, details)
	}

	if str, ok := errObj.(*object.String); ok {
		return newCustomErrorWithTrace("Error", str.Value, node, ctx, nil)
	}

	return newErrorWithTrace("cannot raise non-error object: %s", node, ctx, errObj.Type())
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
	}
	tryResult := Eval(node.TryBlock, env, tryCtx)

	if isError(tryResult) {
		if customErr, ok := tryResult.(*object.CustomError); ok {
			for _, ensnare := range node.EnsnareClauses {
				condition := Eval(ensnare.Condition, env, ctx)
				if isError(condition) {
					result = condition
					break
				}

				if grimoire, ok := condition.(*object.Grimoire); ok {
					if customErr.ErrorType == grimoire {
						ensnareCtx := &CallContext{
							FunctionName: "ensnare",
							Node:         ensnare.Consequence,
							Parent:       ctx,
							env:          env,
						}
						result = Eval(ensnare.Consequence, env, ensnareCtx)
						break
					}
				} else if str, ok := condition.(*object.String); ok {
					if customErr.Name == str.Value {
						ensnareCtx := &CallContext{
							FunctionName: "ensnare",
							Node:         ensnare.Consequence,
							Parent:       ctx,
							env:          env,
						}
						result = Eval(ensnare.Consequence, env, ensnareCtx)
						break
					}
				}
			}
		} else if errWithTrace, ok := tryResult.(*object.ErrorWithTrace); ok {
			// Similar handling for our new error type
			for _, ensnare := range node.EnsnareClauses {
				condition := Eval(ensnare.Condition, env, ctx)
				if isError(condition) {
					result = condition
					break
				}

				if str, ok := condition.(*object.String); ok {
					if strings.HasPrefix(errWithTrace.Message, str.Value) {
						ensnareCtx := &CallContext{
							FunctionName: "ensnare",
							Node:         ensnare.Consequence,
							Parent:       ctx,
							env:          env,
						}
						result = Eval(ensnare.Consequence, env, ensnareCtx)
						break
					}
				}
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
	switch obj1 := obj1.(type) {
	case *object.Integer:
		if obj2, ok := obj2.(*object.Integer); ok {
			return obj1.Value == obj2.Value
		}
	case *object.String:
		if obj2, ok := obj2.(*object.String); ok {
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

		// Don't wrap primitives - this breaks arithmetic operations
		// Wrapping should only happen for explicit method calls on literals

		env.Set(target.Value, val)
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
		key, ok := index.(object.Hashable)
		if !ok {
			return newErrorWithTrace("unusable as hash key: %s", node, ctx, index.Type())
		}

		pair := object.HashPair{Key: index, Value: value}
		array.Pairs[key.HashKey()] = pair
		return value

	default:
		return newErrorWithTrace("index assignment not supported: %s", node, ctx, array.Type())
	}
}

func checkType(val object.Object, expectedType string) bool {
	switch expectedType {
	case "str":
		return val.Type() == object.STRING_OBJ
	case "int":
		return val.Type() == object.INTEGER_OBJ
	case "float":
		return val.Type() == object.FLOAT_OBJ
	case "bool":
		return val.Type() == object.BOOLEAN_OBJ
	default:
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

	// Create method context
	methodCtx := &CallContext{
		FunctionName: instance.Grimoire.Name + "." + methodName,
		Node:         ctx.Node,
		Parent:       ctx,
		env:          methodEnv,
	}

	// Bind arguments: support simple identifiers or full Parameter nodes
	for i, pExpr := range method.Parameters {
		switch param := pExpr.(type) {
		case *ast.Identifier:
			name := param.Value
			if i < len(args) {
				methodEnv.Set(name, args[i])
			} else {
				methodEnv.Set(name, NONE)
			}
		case *ast.Parameter:
			name := param.Name.Value
			if i < len(args) {
				methodEnv.Set(name, args[i])
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
		return evalGrimoireMethodCall(fnTyped.Instance, fnTyped.Name, args, env, ctx)

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

		parentMethod, ok := inst.Grimoire.Inherits.Methods[node.Right.Value]
		if !ok {
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

	instance, ok := leftObj.(*object.Instance)
	if !ok {
		return newErrorWithTrace("type error: %s is not an instance", node, ctx, leftObj.Type())
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
		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newErrorWithTrace("unusable as hash key: %s", node, ctx, key.Type())
		}
		value := Eval(valueNode, env, ctx)
		if isError(value) {
			return value
		}
		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
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
	switch {
	case left.Type() == object.TUPLE_OBJ:
		return evalTupleIndexExpression(left, index, node, ctx)
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index, node, ctx)
	case left.Type() == object.HASH_OBJ:
		return evalHashIndexExpression(left, index, node, ctx)
	case left.Type() == object.STRING_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalStringIndexExpression(left, index, node, ctx)
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
	key, ok := index.(object.Hashable)
	if !ok {
		return newErrorWithTrace("unusable as hash key: %s", node, ctx, index.Type())
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
		default:
			// Unsupported parameter node
		}
	}

	return env
}

func evalIdentifier(node *ast.Identifier, env *object.Environment, ctx *CallContext) object.Object {
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

	for _, statement := range program.Statements {
		result = Eval(statement, env, ctx)

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
	
	switch {
	case unwrappedLeft.Type() == object.INTEGER_OBJ && unwrappedRight.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, unwrappedLeft, unwrappedRight, node, ctx)
	case unwrappedLeft.Type() == object.BOOLEAN_OBJ && unwrappedRight.Type() == object.BOOLEAN_OBJ:
		return evalBooleanInfixExpression(operator, unwrappedLeft, unwrappedRight, node, ctx)
	case unwrappedLeft.Type() == object.STRING_OBJ && unwrappedRight.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, unwrappedLeft, unwrappedRight, node, ctx)
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
	case unwrappedLeft == object.NONE || unwrappedRight == object.NONE:
		if operator == "==" {
			return nativeBoolToBooleanObject(false)
		} else if operator == "!=" {
			return nativeBoolToBooleanObject(true)
		}
		return newErrorWithTrace("operation not supported with None: %s", node, ctx, operator)
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
	if operator != "+" {
		return newErrorWithTrace("unknown operator: %s %s %s",
			node, ctx, left.Type(), operator, right.Type())
	}
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	return &object.String{Value: leftVal + rightVal}
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

		intObj, ok := obj.(*object.Integer)
		if !ok {
			return newErrorWithTrace("prefix '%s' operator requires an integer variable '%s'",
				node, ctx, operator, operand.Value)
		}

		if operator == "++" {
			intObj.Value += 1
		} else if operator == "--" {
			intObj.Value -= 1
		}

		env.Set(operand.Value, intObj)
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

		intObj, ok := obj.(*object.Integer)
		if !ok {
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

		newObj := &object.Integer{Value: newValue}

		env.Set(operand.Value, newObj)

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
	if right.Type() != object.INTEGER_OBJ && right.Type() != object.FLOAT_OBJ {
		// Unknown operand type for prefix minus
		return newError("unknown operator: -%s", right.Type())
	}
	switch right := right.(type) {
	case *object.Integer:
		return &object.Integer{Value: -right.Value}
	case *object.Float:
		return &object.Float{Value: -right.Value}
	default:
		// Fallback for unexpected types
		return newError("unknown type for minus operator: %s", right.Type())
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

		env.Set(leftNode.Value, newVal)
		return newVal

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
	switch l := leftVal.(type) {
	case *object.Integer:
		rInt, ok := rightVal.(*object.Integer)
		if !ok {
			return newErrorWithTrace("type mismatch: expected INTEGER, got %s",
				node, ctx, rightVal.Type())
		}
		switch operator {
		case "+=":
			return &object.Integer{Value: l.Value + rInt.Value}
		case "-=":
			return &object.Integer{Value: l.Value - rInt.Value}
		case "*=":
			return &object.Integer{Value: l.Value * rInt.Value}
		case "/=":
			if rInt.Value == 0 {
				return newErrorWithTrace("division by zero", node, ctx)
			}
			return &object.Integer{Value: l.Value / rInt.Value}
		default:
			return newErrorWithTrace("unknown operator: %s", node, ctx, operator)
		}

	case *object.Float:
		rFloat, ok := rightVal.(*object.Float)
		if !ok {
			return newErrorWithTrace("type mismatch: expected FLOAT, got %s",
				node, ctx, rightVal.Type())
		}
		switch operator {
		case "+=":
			return &object.Float{Value: l.Value + rFloat.Value}
		case "-=":
			return &object.Float{Value: l.Value - rFloat.Value}
		case "*=":
			return &object.Float{Value: l.Value * rFloat.Value}
		case "/=":
			if rFloat.Value == 0 {
				return newErrorWithTrace("division by zero", node, ctx)
			}
			return &object.Float{Value: l.Value / rFloat.Value}
		default:
			return newErrorWithTrace("unknown operator: %s", node, ctx, operator)
		}

	default:
		return newErrorWithTrace("unsupported type for compound assignment: %s",
			node, ctx, leftVal.Type())
	}
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
		for _, elem := range iter.Elements {
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
	}
	return NONE
}

func evalImportStatement(
	node *ast.ImportStatement,
	env *object.Environment,
	ctx *CallContext,
) object.Object {
	filePath := node.FilePath.Value + ".crl"

	if importedFiles[filePath] {
		return object.NONE
	}
	importedFiles[filePath] = true

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return newErrorWithTrace("could not import file: %s", node, ctx, err)
	}

	l := lexer.NewWithFilename(string(fileContent), filePath)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		errorDetails := fmt.Sprintf("parsing errors in imported file %s:\n", filePath)
		for _, err := range p.Errors() {
			errorDetails += fmt.Sprintf("- %s\n", err)
		}
		return newErrorWithTrace(errorDetails, node, ctx)
	}

	importEnv := object.NewEnclosedEnvironment(env)
	importCtx := &CallContext{
		FunctionName: "import_" + filePath,
		Node:         program,
		Parent:       ctx,
		env:          importEnv,
	}

	evalResult := Eval(program, importEnv, importCtx)
	if isError(evalResult) {
		return newErrorWithTrace("error evaluating imported file %s: %s",
			node, ctx, filePath, evalResult.Inspect())
	}

	namespace := &object.Namespace{Env: importEnv}

	if node.Alias != nil {
		env.Set(node.Alias.Value, namespace)
	} else {
		for _, name := range importEnv.GetNames() {
			val, _ := importEnv.Get(name)
			if val.Type() == object.GRIMOIRE_OBJ {
				env.Set(name, val)
			}
		}
	}

	return object.NONE
}
