package evaluator

import (
	"strings"
	"testing"

	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{`x = 1 
      ++x`, 2},
		{`x = 5 
      ++x`, 6},
		{`x = 10
      ++x`, 11},
		{`x = 1 
      --x`, 0},
		{`x = 0 
      --x`, -1},
		{`x = 10 
      --x`, 9},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		{"10 % 3", 1},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	if len(p.Errors()) > 0 {
		// Return error to help with debugging
		return &object.Error{Message: strings.Join(p.Errors(), ", ")}
	}

	result := Eval(program, env)
	return result
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, wanted=%d", result.Value, expected)
		return false
	}
	return true
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"True", true},
		{"False", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"10 >= 10", true},
		{"10 <= 9", false},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t",
			result.Value, expected)
		return false
	}
	return true
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!True", false},
		{"!False", true},
		{"!5", false},
		{"!!True", true},
		{"!!False", false},
		{"!!5", true},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if(True): 10", 10},
		{"if(False): 10", nil},
		{"if(1): 10", 10},
		{"if(1 < 2): 10", 10},
		{"if(1 > 10): 10", nil},
		{"if(1 < 2): 10 else: 20", 10},
		{"if(1 > 2): 10 else: 20", 20},
		{`
      if (1<0): 
        return 0
      otherwise (1 > 0):
        return 1 
      else:
        return -1`, 1},
		{`if 10 > 1:
        if 10 > 1:
              return 10
        return 1`, 10},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNoneObject(t, evaluated)
		}
	}
}

func testNoneObject(t *testing.T, obj object.Object) bool {
	if obj == nil {
		t.Errorf("object is nil, expected NONE")
		return false
	}
	if obj.Type() != object.NONE_OBJ {
		t.Errorf("object is not NONE. got=%s (%+v)", obj.Type(), obj)
		return false
	}
	return true
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + True",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + True 5",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-True",
			"unknown operator: -BOOLEAN",
		},
		{
			"True + False",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5 True + False 5",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1): True + False ",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
      if (10 > 1):
          if (10 > 1):
              return True + False
      return 1
      `,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{"foobar", "identifier not found: foobar"},
		{`"Hello" - "World"`, "unknown operator: STRING - STRING"},
		//{`{"name": "Carrion"}[spell add(x,y): return x + y]`, "unusable as hash key: SPELL"},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)",
				evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q",
				tt.expectedMessage, errObj.Message)
		}
	}
}

func TestAssignmentStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"x = 5 x", 5},
		{"x = 5 * 5 x", 25},
		{"a = 5 b= 5 a b ", 5},
		{"a = 5 b = a c = a + b + 5 c", 15},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionDefinitionAndCall(t *testing.T) {
	input := `
spell add(x, y):
    return x + y

result = add(2, 3)
result
`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 5)
}

func TestFunctionDefinitionInline(t *testing.T) {
	// demonstrates a single-line function body
	input := `
spell identity(x): return x
identity(42)
`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 42)
}

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`
	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestStringConcatenation(t *testing.T) {
	input := `"Hello" + " " + "World!"`
	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
	}
	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. got=%q", str.Value)
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument to `len` not supported, got INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		switch expected := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(expected))
		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)",
					evaluated, evaluated)
				continue
			}
			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, got=%q",
					expected, errObj.Message)
			}
		}
	}
}

func TestArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"
	evaluated := testEval(input)
	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object is not Array. got=%T (%+v)", evaluated, evaluated)
	}
	if len(result.Elements) != 3 {
		t.Fatalf("array has wrong num of elements. got=%d",
			len(result.Elements))
	}
	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 4)
	testIntegerObject(t, result.Elements[2], 6)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		//{"i=0 [1][i]",1,},
		{
			"[1, 2, 3][1 + 1]",
			3,
		},
		{
			"myArray = [1, 2, 3] myArray[2]",
			3,
		},
		{
			"myArray = [1, 2, 3] myArray[0] + myArray[1] + myArray[2]",
			6,
		},
		{
			"myArray = [1, 2, 3] i = myArray[0] myArray[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			"index out of bounds: 3 (array length: 3)",
		},
		{
			"[1, 2, 3][-1]",
			"index out of bounds: -1 (array length: 3)",
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else if errorMsg, ok := tt.expected.(string); ok {
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("Expected error object, got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if !strings.Contains(errObj.Message, errorMsg) {
				t.Errorf("Expected error message to contain %q, got %q", errorMsg, errObj.Message)
			}
		} else {
			testNoneObject(t, evaluated)
		}
	}
}

func TestRecursiveFunction(t *testing.T) {
	input := `
spell factorial(n):
    if n == 0:
        return 1
    return n * factorial(n - 1)

factorial(5)
`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 120)
}

func TestHashLiterals(t *testing.T) {
	// Fix the hash syntax - make it a single line to avoid formatting issues
	input := `two = "two"
{"one": 10 - 9, two: 1 + 1, "thr" + "ee": 6 / 2, 4: 4, True: 5, False: 6}`

	evaluated := testEval(input)
	result, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("Eval didn't return Hash. got=%T (%+v)", evaluated, evaluated)
	}
	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		TRUE.HashKey():                             5,
		FALSE.HashKey():                            6,
	}
	if len(result.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong num of pairs. got=%d", len(result.Pairs))
	}
	for expectedKey, expectedValue := range expected {
		pair, ok := result.Pairs[expectedKey]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}
		testIntegerObject(t, pair.Value, expectedValue)
	}
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 5}["foo"]`,
			5,
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`key = "foo" {"foo": 5}[key]`,
			5,
		},
		{
			`{}["foo"]`,
			nil,
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{True: 5}[True]`,
			5,
		},
		{
			`{False: 5}[False]`,
			5,
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNoneObject(t, evaluated)
		}
	}
}

func TestSpellbookMethodCall(t *testing.T) {
	input := `
spellbook Calculator:
    spell add(x, y):
        return x + y
    
    spell multiply(x, y):
        return x * y

calc = Calculator()
result = calc.add(5, 3)
result
`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 8)
}

func TestSpellbookRecursion(t *testing.T) {
	input := `
spellbook Fibonacci:
    spell calc(n):
        if n <= 1:
            return n
        return self.calc(n-1) + self.calc(n-2)

fib = Fibonacci()
fib.calc(10)
`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 55)
}

func TestSpellbookInheritance(t *testing.T) {
	input := `
spellbook Shape:
    spell area():
        return 0

spellbook Rectangle(Shape):
    init(width, height):
        self.width = width
        self.height = height
    
    spell area():
        return self.width * self.height

rect = Rectangle(5, 10)
rect.area()
`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 50)
}

func TestBinarySearch(t *testing.T) {
	input := `
spellbook SafeArray:
    init(elements):
        self.elements = elements
        
    spell get(index):
        if index < 0 or index >= len(self.elements):
            return None
        return self.elements[index]
        
    spell binary_search(target):
        return self._binary_search_helper(target, 0, len(self.elements) - 1)
        
    spell _binary_search_helper(target, left, right):
        if left > right:
            return -1
            
        mid = (left + right) / 2
        mid_val = self.get(mid)
        
        if mid_val == target:
            return mid
        otherwise mid_val < target:
            return self._binary_search_helper(target, mid + 1, right)
        else:
            return self._binary_search_helper(target, left, mid - 1)

// Test the binary search with a sorted array
arr = SafeArray([1, 3, 5, 7, 9, 11, 13, 15])
result = arr.binary_search(7)
result
`
	evaluated := testEval(input)
	testIntegerObject(t, evaluated, 3) // 7 is at index 3
}
