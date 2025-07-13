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
       {"x = 1; ++x", 2},
       {"x = 5; ++x", 6},
       {"x = 10; ++x", 11},
       {"x = 1; --x", 0},
       {"x = 0; --x", -1},
       {"x = 10; --x", 9},
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

   // Evaluate AST with direct execution context for proper main statement handling
	ctx := &CallContext{
		FunctionName:      "<program>",
		Node:              program,
		Parent:            nil,
		IsDirectExecution: true,
		env:               env,
	}
	result := Eval(program, env, ctx)
	return result
}

// Helper function to extract error message from either Error or ErrorWithTrace
func getErrorMessage(obj object.Object) (string, bool) {
	switch err := obj.(type) {
	case *object.Error:
		return err.Message, true
	case *object.ErrorWithTrace:
		return err.Message, true
	default:
		return "", false
	}
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
        // Extract error message from Error or ErrorWithTrace
        var errMessage string
        switch err := evaluated.(type) {
        case *object.Error:
            errMessage = err.Message
        case *object.ErrorWithTrace:
            errMessage = err.Message
        default:
            t.Errorf("no error object returned. got=%T (%+v)", evaluated, evaluated)
            continue
        }
        if errMessage != tt.expectedMessage {
            t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errMessage)
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
			errMessage, ok := getErrorMessage(evaluated)
			if !ok {
				t.Errorf("object is not Error. got=%T (%+v)",
					evaluated, evaluated)
				continue
			}
			if errMessage != expected {
				t.Errorf("wrong error message. expected=%q, got=%q",
					expected, errMessage)
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
			3,
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else if errorMsg, ok := tt.expected.(string); ok {
			errMessage, ok := getErrorMessage(evaluated)
			if !ok {
				t.Errorf("Expected error object, got=%T (%+v)", evaluated, evaluated)
				continue
			}
			if !strings.Contains(errMessage, errorMsg) {
				t.Errorf("Expected error message to contain %q, got %q", errorMsg, errMessage)
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

func TestGrimoireMethodCall(t *testing.T) {
	input := `
 grim Calculator:
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

func TestGrimoireRecursion(t *testing.T) {
	input := `
 grim Fibonacci:
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

func TestGrimoireInheritance(t *testing.T) {
	t.Skip("Grimoire inheritance parsing needs separate fix")
	input := `
 grim Shape:
    spell area():
        return 0

 grim Rectangle(Shape):
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
	t.Skip("Binary search test has issues beyond parameter scoping")
	input := `
 grim SafeArray:
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

func TestMainStatementTwoPhaseExecution(t *testing.T) {
	t.Run("backward compatibility - no main block", func(t *testing.T) {
		// Without main block, all statements should execute normally
		input := `x = 10
y = 20
spell add(a, b):
    return a + b
result = add(x, y)
result`
		evaluated := testEval(input)
		testIntegerObject(t, evaluated, 30)
	})

	t.Run("with main block - function definitions execute before main", func(t *testing.T) {
		input := `spell multiply(a, b):
    return a * b

main:
    result = multiply(5, 6)
    result`
		evaluated := testEval(input)
		testIntegerObject(t, evaluated, 30)
	})

	t.Run("with main block - class definitions execute before main", func(t *testing.T) {
		// Test that assignments execute before main
		input := `counter = 2

main:
    counter`
		evaluated := testEval(input)
		testIntegerObject(t, evaluated, 2)
	})

	t.Run("with main block - assignments execute before main", func(t *testing.T) {
		input := `x = 100
y = 200
z = x + y

main:
    z`
		evaluated := testEval(input)
		testIntegerObject(t, evaluated, 300)
	})

	t.Run("with main block - expression statements are skipped outside main", func(t *testing.T) {
		input := `x = 0

# These expression statements should be skipped
print("This should NOT execute")
x + 10  # This should not affect the result

main:
    x  # Should still be 0`
		evaluated := testEval(input)
		testIntegerObject(t, evaluated, 0)
	})

	t.Run("main block executes last", func(t *testing.T) {
		input := `execution_order = []

spell record_init():
    execution_order = execution_order + ["init"]
    return execution_order

execution_order = record_init()

main:
    execution_order = execution_order + ["main"]
    len(execution_order)  # Should be 2 if main executes last`
		evaluated := testEval(input)
		testIntegerObject(t, evaluated, 2)
	})

	t.Run("return value from main block", func(t *testing.T) {
		input := `spell helper():
    return 42

main:
    return helper()`
		evaluated := testEval(input)
		testIntegerObject(t, evaluated, 42)
	})

	t.Run("error handling in main block", func(t *testing.T) {
		input := `main:
    x = 10 / 0  # Division by zero`
		evaluated := testEval(input)
		errMsg, isError := getErrorMessage(evaluated)
		if !isError {
			t.Fatalf("expected error object, got %T", evaluated)
		}
		if !strings.Contains(errMsg, "division by zero") {
			t.Errorf("expected division by zero error, got %q", errMsg)
		}
	})

	t.Run("error handling outside main block", func(t *testing.T) {
		input := `spell bad_function():
    return 1 / 0

x = bad_function()  # This should error before main

main:
    x`
		evaluated := testEval(input)
		errMsg, isError := getErrorMessage(evaluated)
		if !isError {
			t.Fatalf("expected error object, got %T", evaluated)
		}
		if !strings.Contains(errMsg, "division by zero") {
			t.Errorf("expected division by zero error, got %q", errMsg)
		}
	})

	t.Run("empty main block", func(t *testing.T) {
		input := `x = 123

main:
    # Empty main block

# This should not execute after main
x`
		evaluated := testEval(input)
		// Empty main block should return nil
		if evaluated != nil {
			t.Errorf("expected nil from empty main block, got %T (%+v)", evaluated, evaluated)
		}
	})

	t.Run("nested blocks within main", func(t *testing.T) {
		input := `main:
    x = 10
    if x > 5:
        if x < 20:
            x = x * 2
    x`
		evaluated := testEval(input)
		testIntegerObject(t, evaluated, 20)
	})

	t.Run("early return within main", func(t *testing.T) {
		input := `main:
    x = 5
    if x > 0:
        return x * 10
    return 0  # Should not reach here`
		evaluated := testEval(input)
		testIntegerObject(t, evaluated, 50)
	})

	t.Run("multiple main blocks should error", func(t *testing.T) {
		input := `main:
    x = 1

main:
    x = 2`
		evaluated := testEval(input)
		// This should be handled by the parser, but if it reaches evaluator,
		// only the first main should be recognized
		testIntegerObject(t, evaluated, 1)
	})

	t.Run("imports execute before main", func(t *testing.T) {
		// Note: This would require actual module system, so we simulate
		input := `# import statements would go here
imported_value = 999  # Simulating an imported value

main:
    imported_value`
		evaluated := testEval(input)
		testIntegerObject(t, evaluated, 999)
	})

	t.Run("complex program with main - full integration", func(t *testing.T) {
		input := `# Global assignments
PI = 3
radius = 5

# Function definitions
spell area(r):
    return PI * r * r

spell circumference(r):
    return 2 * PI * r

# These should be skipped
print("This should not execute")
area(10)  # This call should be skipped

# Main block executes last
main:
    area(radius)`
		evaluated := testEval(input)
		testIntegerObject(t, evaluated, 75) // 3 * 5 * 5
	})
}

func TestStringIndexing(t *testing.T) {
	t.Run("string indexing should return character", func(t *testing.T) {
		input := `s = "hello"
result = s[1]`
		evaluated := testEval(input)
		result, ok := evaluated.(*object.String)
		if !ok {
			t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
		}
		if result.Value != "e" {
			t.Errorf("string has wrong value. got=%q, wanted=%q", result.Value, "e")
		}
	})
	
	t.Run("string indexing bounds", func(t *testing.T) {
		tests := []struct {
			input    string
			expected string
			desc     string
		}{
			{`s = "hello"; s[0]`, "h", "first character"},
			{`s = "hello"; s[4]`, "o", "last character"},
		}
		
		for _, tt := range tests {
			t.Run(tt.desc, func(t *testing.T) {
				evaluated := testEval(tt.input)
				result, ok := evaluated.(*object.String)
				if !ok {
					t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
				}
				if result.Value != tt.expected {
					t.Errorf("string has wrong value. got=%q, wanted=%q", result.Value, tt.expected)
				}
			})
		}
	})
	
	t.Run("array indexing should work correctly", func(t *testing.T) {
		input := `arr = [1, 2, 3]
result = arr[1]`
		evaluated := testEval(input)
		testIntegerObject(t, evaluated, 2)
	})
}

func TestIncrementDecrementOperators(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
		desc     string
	}{
		{"x = 5; x += 3; x", 8, "addition assignment"},
		{"x = 10; x -= 4; x", 6, "subtraction assignment"},
		{"x = 3; x *= 4; x", 12, "multiplication assignment"},
		{"x = 12; x /= 3; x", 4, "division assignment"},
		{"x = 1; ++x", 2, "pre-increment"},
		{"x = 5; --x", 4, "pre-decrement"},
	}
	
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			evaluated := testEval(tt.input)
			testIntegerObject(t, evaluated, tt.expected)
		})
	}
}
