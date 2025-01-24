package evaluator

import (
	"fmt"
	"math"
	"os"

	"thecarrionlanguage/src/ast"
	"thecarrionlanguage/src/lexer"
	"thecarrionlanguage/src/object"
	"thecarrionlanguage/src/parser"
)

var (
	NONE          = &object.None{Value: "None"}
	TRUE          = &object.Boolean{Value: true}
	FALSE         = &object.Boolean{Value: false}
	importedFiles = map[string]bool{}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.IfStatement:
		return evalIfExpression(node, env)
	case *ast.PrefixExpression:
		if node.Operator == "++" || node.Operator == "--" {
			return evalPrefixIncrementDecrement(node.Operator, node, env)
		}
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, node, env)

	case *ast.InfixExpression:
		if node.Operator == "+=" || node.Operator == "-=" ||
			node.Operator == "*=" || node.Operator == "/=" {
			return evalCompoundAssignment(node, env)
		}
		right := Eval(node.Right, env)
		if isError(right) {
			// fmt.Printf("Error in right operand: %v\n", right)
			return right
		}
		left := Eval(node.Left, env)
		if isError(left) {
			// fmt.Printf("Error in left operand: %v\n", left)
			return left
		}
		result := evalInfixExpression(node.Operator, left, right)
		// fmt.Printf("InfixExpression result: %v\n", result)
		return result
	case *ast.PostfixExpression:
		return evalPostfixIncrementDecrement(node.Operator, node, env)
		// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}
	case *ast.NoneLiteral:
		return object.NONE
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.AssignStatement:
		return evalAssignStatement(node, env)
	case *ast.WhileStatement:
		return evalWhileStatement(node, env)
	case *ast.ForStatement:
		return evalForStatement(node, env)
	case *ast.ImportStatement:
		return evalImportStatement(node, env)
	case *ast.MatchStatement:
		return evalMatchStatement(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.ArrayLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.TupleLiteral:
		return evalTupleLiteral(node, env)
	case *ast.HashLiteral:
		return evalHashLiteral(node, env)
	case *ast.FunctionDefinition:
		fnObj := &object.Function{
			Parameters: node.Parameters,
			Body:       node.Body,
			Env:        env,
		}
		env.Set(node.Name.Value, fnObj)
		return fnObj
	case *ast.DotExpression:
		return evalDotExpression(node, env)
	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)
	case *ast.SpellbookDefinition:
		return evalSpellbookDefinition(node, env)
	case *ast.AttemptStatement:
		return evalAttemptStatement(node, env)
	case *ast.CallExpression:
		return evalCallExpression(Eval(node.Function, env), evalExpressions(node.Arguments, env))
	}
	return NONE
}

func evalAttemptStatement(node *ast.AttemptStatement, env *object.Environment) object.Object {
	// Evaluate the try block
	result := Eval(node.TryBlock, env)

	// If an error occurs and a resolve block exists, execute the resolve block
	if isError(result) && node.ResolveBlock != nil {
		return Eval(node.ResolveBlock, env)
	}

	return result
}

func evalMatchStatement(ms *ast.MatchStatement, env *object.Environment) object.Object {
	matchValue := Eval(ms.MatchValue, env)
	if isError(matchValue) {
		return matchValue
	}

	for _, caseClause := range ms.Cases {
		caseCondition := Eval(caseClause.Condition, env)
		if isError(caseCondition) {
			return caseCondition
		}

		// Compare with isEqual(...) or however you prefer
		if isEqual(matchValue, caseCondition) {
			// As soon as we find a match, evaluate and return
			return Eval(caseClause.Body, env)
		}
	}

	// If no case matched, look for default case
	if ms.Default != nil {
		return Eval(ms.Default.Body, env)
	}

	return NONE // no matches, no default => return None
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
	// Add more cases for other types (e.g., Float, Boolean, etc.)
	default:
		return false
	}
	return false
}

func evalAssignStatement(node *ast.AssignStatement, env *object.Environment) object.Object {
	switch name := node.Name.(type) {
	case *ast.Identifier:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(name.Value, val)
		return val

	case *ast.DotExpression:
		// Handle assignments to object properties
		left := Eval(name.Left, env)
		if isError(left) {
			return left
		}

		instance, ok := left.(*object.Instance)
		if !ok {
			return newError("invalid assignment target: %s", left.Type())
		}

		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}

		instance.Env.Set(name.Right.Value, val)
		return val

	default:
		return newError("invalid assignment target: %T", node.Name)
	}
}

func evalSpellbookDefinition(node *ast.SpellbookDefinition, env *object.Environment) object.Object {
	methods := map[string]*object.Function{}

	for _, method := range node.Methods {
		fn := &object.Function{
			Parameters: method.Parameters,
			Body:       method.Body,
			Env:        env,
		}
		methods[method.Name.Value] = fn
	}

	spellbook := &object.Spellbook{
		Name:    node.Name.Value,
		Methods: methods,
		Env:     env, // Assign the environment
	}

	if node.InitMethod != nil {
		initFn := &object.Function{
			Parameters: node.InitMethod.Parameters,
			Body:       node.InitMethod.Body,
			Env:        env,
		}
		spellbook.InitMethod = initFn
	}

	env.Set(node.Name.Value, spellbook)
	return spellbook
}

func evalCallExpression(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Spellbook:
		// This is the case for calling the spellbook constructor, e.g. Person("Alice", 30)
		instance := &object.Instance{
			Spellbook: fn,
			Env:       object.NewEnclosedEnvironment(fn.Env),
		}
		// If there's an init method, call it
		if fn.InitMethod != nil {
			extendedEnv := extendFunctionEnv(fn.InitMethod, args)
			extendedEnv.Set("self", instance)
			Eval(fn.InitMethod.Body, extendedEnv)
		}
		return instance

	case *object.Function:
		// Normal function call (no bound instance)
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)

	case *object.BoundMethod:
		// *** THIS is how you set self in the method environment ***
		extendedEnv := extendFunctionEnv(fn.Method, args)
		extendedEnv.Set("self", fn.Instance)

		evaluated := Eval(fn.Method.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)

	default:
		return newError("not a function: %s", fn.Type())
	}
}

func evalDotExpression(node *ast.DotExpression, env *object.Environment) object.Object {
	leftObj := Eval(node.Left, env)
	if isError(leftObj) {
		return leftObj
	}

	instance, ok := leftObj.(*object.Instance)
	if !ok {
		return newError("type error: %s is not an instance", leftObj.Type())
	}

	fieldOrMethodName := node.Right.Value

	// If it's a field in instance.Env (like self.name)
	if val, found := instance.Env.Get(fieldOrMethodName); found {
		return val
	}

	// Otherwise, check if it's a method in instance.Spellbook.Methods
	method, ok := instance.Spellbook.Methods[fieldOrMethodName]
	if !ok {
		return newError("undefined property or method: %s", fieldOrMethodName)
	}

	// Return a BoundMethod, DO NOT CALL THE METHOD:
	return &object.BoundMethod{
		Instance: instance,
		Method:   method,
	}
}

func evalHashLiteral(
	node *ast.HashLiteral,
	env *object.Environment,
) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)
	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}
		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type())
		}
		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}
		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}
	return &object.Hash{Pairs: pairs}
}

func evalTupleLiteral(tl *ast.TupleLiteral, env *object.Environment) object.Object {
	elements := evalExpressions(tl.Elements, env)
	if len(elements) == 1 && isError(elements[0]) {
		return elements[0]
	}

	return &object.Tuple{Elements: elements}
}

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.TUPLE_OBJ:
		return evalTupleIndexExpression(left, index)
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.HASH_OBJ:
		return evalHashIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

func evalTupleIndexExpression(tuple, index object.Object) object.Object {
	tupleObj := tuple.(*object.Tuple)
	idx := int(index.(*object.Integer).Value)
	if idx < 0 || idx >= len(tupleObj.Elements) {
		return NONE
	}
	return tupleObj.Elements[idx]
}

func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)
	key, ok := index.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", index.Type())
	}
	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return NONE
	}
	return pair.Value
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	maxIndex := int64(len(arrayObject.Elements) - 1)
	if idx < 0 || idx > maxIndex {
		return NONE
	}
	return arrayObject.Elements[idx]
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newError("not a function: %s", fn.Type())
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for i, param := range fn.Parameters {
		// If an argument is provided, use it
		if i < len(args) {
			env.Set(param.Name.Value, args[i])
		} else if param.DefaultValue != nil {
			// If no argument is provided, use the default value
			defaultVal := Eval(param.DefaultValue, env)
			env.Set(param.Name.Value, defaultVal)
		} else {
			env.Set(param.Name.Value, NONE) // No default value
		}
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	// if the function returned via `return`, unwrap the *object.ReturnValue
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}
	if node.Value == "None" {
		return object.NONE
	}
	return newError("identifier not found: " + node.Value)
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		// fmt.Printf("Evaluating statement: %T\n", statement)
		result = Eval(statement, env)
		// fmt.Printf("Statement result: %v\n", result)
		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			// fmt.Printf("Error found: %v\n", result)
			return result
		}
	}
	return result
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result // Propagate return/error up
			}
		}
	}

	return result // Return the last evaluated value (or nil if empty block)
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
) object.Object {
	switch operator {
	case "!":
		right := Eval(node.Right, env)
		return evalBangOperatorExpression(right, env)
	case "-":
		right := Eval(node.Right, env)
		return evalMinusPrefixOperatorExpression(right, env)
	default:
		return newError("unknown operator: %s%s", operator, Eval(node.Right, env).Type())
	}
}

func evalInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	// fmt.Printf("InfixExpression operator: %s, left: %v, right: %v\n", operator, left, right)
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ:
		return evalBooleanInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case left == object.NONE && right == object.NONE:
		return nativeBoolToBooleanObject(operator == "==")
	case left == object.NONE || right == object.NONE:
		if operator == "==" {
			return nativeBoolToBooleanObject(false)
		} else if operator == "!=" {
			return nativeBoolToBooleanObject(true)
		}
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), operator, right.Type())
	case left.Type() == object.FLOAT_OBJ || right.Type() == object.FLOAT_OBJ:
		leftVal := toFloat(left)
		rightVal := toFloat(right)
		switch operator {
		case "+":
			return &object.Float{Value: leftVal + rightVal}
		case "-":
			return &object.Float{Value: leftVal - rightVal}
		case "*":
			return &object.Float{Value: leftVal * rightVal}
		case "/":
			return &object.Float{Value: leftVal / rightVal}
		case "**":
			return &object.Float{Value: math.Pow(leftVal, rightVal)}
		default:
			return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
		}
	}

	// Fallback return for unmatched cases
	return newError(
		"unknown operator or type mismatch: %s %s %s",
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
		return 0.0 // Shouldn't reach here
	}
}

func evalStringInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	if operator != "+" {
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	return &object.String{Value: leftVal + rightVal}
}

func evalBooleanInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Boolean).Value
	rightVal := right.(*object.Boolean).Value
	switch operator {
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalPrefixIncrementDecrement(
	operator string,
	node *ast.PrefixExpression,
	env *object.Environment,
) object.Object {
	switch operand := node.Right.(type) {
	case *ast.Identifier:
		obj, ok := env.Get(operand.Value)
		if !ok {
			return newError("undefined variable '%s'", operand.Value)
		}

		intObj, ok := obj.(*object.Integer)
		if !ok {
			return newError("prefix '%s' operator requires an integer variable '%s'", operator, operand.Value)
		}

		if operator == "++" {
			intObj.Value += 1
		} else if operator == "--" {
			intObj.Value -= 1
		}

		env.Set(operand.Value, intObj)
		return intObj

	default:
		return newError("prefix '%s' operator requires an integer or identifier", operator)
	}
}

func evalPostfixIncrementDecrement(
	operator string,
	node *ast.PostfixExpression,
	env *object.Environment,
) object.Object {
	switch operand := node.Left.(type) {
	case *ast.Identifier:
		// Get the current value from environment
		obj, ok := env.Get(operand.Value)
		if !ok {
			return newError("undefined variable '%s'", operand.Value)
		}

		// Make sure it's an integer
		intObj, ok := obj.(*object.Integer)
		if !ok {
			return newError("postfix '%s' operator requires an integer variable '%s'", operator, operand.Value)
		}

		// Store old value to return (postfix behavior)
		oldValue := intObj.Value

		// Create new value
		var newValue int64
		if operator == "++" {
			newValue = oldValue + 1
		} else if operator == "--" {
			newValue = oldValue - 1
		}

		// Create new integer object and set it
		newObj := &object.Integer{Value: newValue}

		// Update in current environment
		env.Set(operand.Value, newObj)

		// Return the original value (postfix behavior)
		return &object.Integer{Value: oldValue}
	default:
		return newError("postfix '%s' operator requires an integer or identifier", operator)
	}
}

func evalPostfixExpression(
	operator string,
	node *ast.PostfixExpression,
	env *object.Environment,
) object.Object {
	switch operator {
	case "++", "--":
		return evalPostfixIncrementDecrement(operator, node, env)
	default:
		return newError("unknown operator: %s", operator)
	}
}

func evalBangOperatorExpression(right object.Object, env *object.Environment) object.Object {
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

func evalMinusPrefixOperatorExpression(right object.Object, env *object.Environment) object.Object {
	if right.Type() != object.INTEGER_OBJ && right.Type() != object.FLOAT_OBJ {
		return newError("unknown operator: -%s", right.Type())
	}
	switch right := right.(type) {
	case *object.Integer:
		return &object.Integer{Value: -right.Value}
	case *object.Float:
		return &object.Float{Value: -right.Value}
	default:
		return newError("unknown type for minus operator: %s", right.Type())
	}
}

func evalIncrementOperatorExpression(side object.Object) object.Object {
	if side.Type() != object.INTEGER_OBJ {
		return NONE
	}
	value := side.(*object.Integer).Value
	return &object.Integer{Value: value + 1}
}

func evalDecrementOperatorExpression(side object.Object) object.Object {
	if side.Type() != object.INTEGER_OBJ {
		return NONE
	}
	value := side.(*object.Integer).Value
	return &object.Integer{Value: value - 1}
}

func evalIntegerInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "%":
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
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalCompoundAssignment(node *ast.InfixExpression, env *object.Environment) object.Object {
	// Evaluate the right-hand side
	rightVal := Eval(node.Right, env)
	if isError(rightVal) {
		return rightVal
	}

	// The left side must be an AST expression (often an Identifier, Dot, or Index).
	switch leftNode := node.Left.(type) {
	case *ast.Identifier:
		// 1. Look up the current value of x in environment
		currVal, ok := env.Get(leftNode.Value)
		if !ok {
			return newError("undefined variable: %s", leftNode.Value)
		}
		// 2. Apply the arithmetic
		newVal := applyCompoundOperator(node.Operator, currVal, rightVal)
		if isError(newVal) {
			return newVal
		}
		// 3. Store back into environment
		env.Set(leftNode.Value, newVal)
		return newVal

	// (Optional) If you want to allow `arr[0] += 5` or `obj.field += 5`, handle
	// *ast.IndexExpression or *ast.DotExpression similarly here.

	default:
		return newError("invalid assignment target: %T", leftNode)
	}
}

func applyCompoundOperator(operator string, leftVal, rightVal object.Object) object.Object {
	switch l := leftVal.(type) {
	case *object.Integer:
		rInt, ok := rightVal.(*object.Integer)
		if !ok {
			return newError("type mismatch: expected INTEGER, got %s", rightVal.Type())
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
				return newError("division by zero")
			}
			return &object.Integer{Value: l.Value / rInt.Value}
		default:
			return newError("unknown operator: %s", operator)
		}

	case *object.Float:
		rFloat, ok := rightVal.(*object.Float)
		if !ok {
			return newError("type mismatch: expected FLOAT, got %s", rightVal.Type())
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
				return newError("division by zero")
			}
			return &object.Float{Value: l.Value / rFloat.Value}
		default:
			return newError("unknown operator: %s", operator)
		}

	default:
		return newError("unsupported type for compound assignment: %s", leftVal.Type())
	}
}

func evalIfExpression(ie *ast.IfStatement, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)
	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	}

	for _, branch := range ie.OtherwiseBranches {
		condition = Eval(branch.Condition, env)
		if isError(condition) {
			return condition
		}
		if isTruthy(condition) {
			return Eval(branch.Consequence, env)
		}
	}

	if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	}

	return NONE
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	return obj != nil && obj.Type() == object.ERROR_OBJ
}

func evalWhileStatement(node *ast.WhileStatement, env *object.Environment) object.Object {
	for {
		// Re-evaluate the condition in the current environment
		condition := Eval(node.Condition, env)
		if isError(condition) {
			return condition
		}

		// Stop loop if condition evaluates to FALSE
		if !isTruthy(condition) {
			break
		}

		// Evaluate the body
		result := Eval(node.Body, env)
		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return NONE
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case TRUE:
		return true
	case FALSE, NONE:
		return false
	default:
		return true
	}
}

func evalForStatement(fs *ast.ForStatement, env *object.Environment) object.Object {
	iterable := Eval(fs.Iterable, env)
	if isError(iterable) {
		return iterable
	}

	var result object.Object = NONE

	switch iter := iterable.(type) {
	case *object.Array:
		for _, elem := range iter.Elements {
			env.Set(fs.Variable.Value, elem)
			result = Eval(fs.Body, env)
			if result != nil {
				rt := result.Type()
				if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
					return result
				}
			}
		}
	default:
		return newError("unsupported iterable type: %s", iterable.Type())
	}

	// Handle optional else block
	if fs.Alternative != nil {
		result = Eval(fs.Alternative, env)
	}

	return result
}

func evalImportStatement(node *ast.ImportStatement, env *object.Environment) object.Object {
	filePath := node.FilePath.Value + ".crl" // Append .crl extension
	className := ""
	if node.ClassName != nil {
		className = node.ClassName.Value
	}

	// Check if the file is already imported
	if importedFiles[filePath] {
		return NONE
	}
	importedFiles[filePath] = true

	// Load and parse the file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return newError("could not import file: %s", err)
	}

	l := lexer.New(string(fileContent))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		return newError("parsing errors in imported file: %v", p.Errors())
	}

	// Evaluate the program in a new environment
	importEnv := object.NewEnclosedEnvironment(env)
	Eval(program, importEnv)

	// If a class name is specified, import only that class
	if className != "" {
		val, ok := importEnv.Get(className)
		if !ok || val.Type() != object.SPELLBOOK_OBJ {
			return newError("class '%s' not found in file '%s'", className, filePath)
		}
		env.Set(className, val)
	} else {
		// Import all spellbooks from the file
		for _, name := range importEnv.GetNames() {
			val, _ := importEnv.Get(name)
			if val.Type() == object.SPELLBOOK_OBJ {
				env.Set(name, val)
			}
		}
	}

	return NONE
}
