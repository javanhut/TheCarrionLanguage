package evaluator

import (
	"fmt"
	"strings"

	"thecarrionlanguage/ast"
	"thecarrionlanguage/object"
)

var (
	NONE  = &object.None{Value: "None"}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
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
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.AssignStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
	case *ast.WhileStatement:
		return evalWhileStatement(node, env)
	case *ast.ForStatement:
		return evalForStatement(node, env)

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
		methods := make(map[string]*object.Function)
		for _, spell := range node.Spells {
			fn := &object.Function{
				Parameters: spell.Parameters,
				Body:       spell.Body,
				Env:        env,
			}
			methods[spell.Name.Value] = fn
		}
		spellbook := &object.Spellbook{
			Name:    node.Name.Value,
			Methods: methods,
		}
		env.Set(node.Name.Value, spellbook)
		return spellbook

	case *ast.CallExpression:
		if identifier, ok := node.Function.(*ast.Identifier); ok && strings.Contains(identifier.Value, ".") {
			parts := strings.Split(identifier.Value, ".")
			if len(parts) == 2 {
				spellbookName := parts[0]
				spellName := parts[1]

				// Get the spellbook
				spellbookObj, ok := env.Get(spellbookName)
				if !ok {
					return newError("undefined spellbook: %s", spellbookName)
				}
				spellbook, ok := spellbookObj.(*object.Spellbook)
				if !ok {
					return newError("%s is not a spellbook", spellbookName)
				}

				// Get the spell
				spell, ok := spellbook.Methods[spellName]
				if !ok {
					return newError("undefined spell: %s in spellbook %s", spellName, spellbookName)
				}

				// Apply the spell
				args := evalExpressions(node.Arguments, env)
				if len(args) == 1 && isError(args[0]) {
					return args[0]
				}
				return applyFunction(spell, args)
			}
		}
	}
	return NONE
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
	// Create a new child environment so function variables don’t pollute outer env
	env := object.NewEnclosedEnvironment(fn.Env)

	for i, param := range fn.Parameters {
		env.Set(param.Value, args[i])
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

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	// fmt.Printf("InfixExpression operator: %s, left: %v, right: %v\n", operator, left, right)
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ:
		return evalBooleanInfixExpression(operator, left, right)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)

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
		default:
			return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
		}
	default:
		// fmt.Printf("Error: type mismatch or unknown operator\n")
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
	}
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
		// Fetch the variable from the environment
		obj, ok := env.Get(operand.Value)
		if !ok {
			return newError("undefined variable '%s'", operand.Value)
		}

		intObj, ok := obj.(*object.Integer)
		if !ok {
			return newError("postfix '%s' operator requires an integer variable '%s'", operator, operand.Value)
		}

		// Store the old value to return
		oldValue := intObj.Value

		// Increment or decrement the value
		if operator == "++" {
			intObj.Value += 1
		} else if operator == "--" {
			intObj.Value -= 1
		}

		// Update the environment with the new value
		env.Set(operand.Value, intObj)

		// Return the old value (postfix behavior)
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
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	case "+=":
		return &object.Integer{Value: leftVal + rightVal}
	case "*=":
		return &object.Integer{Value: leftVal * rightVal}
	case "-=":
		return &object.Integer{Value: leftVal - rightVal}
	case "/=":
		return &object.Integer{Value: leftVal / rightVal}
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), operator, right.Type())
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

/*func isTruthy(obj object.Object) bool {
	switch obj {
	case NONE:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}*/

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
			if returnValue, ok := result.(*object.ReturnValue); ok {
				return returnValue
			}
			if result.Type() == object.ERROR_OBJ {
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
