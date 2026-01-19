// src/utils/example_enhanced_errors.go
package utils

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/object"
)

// ExampleEnhancedErrorSystem demonstrates the enhanced error system
func ExampleEnhancedErrorSystem() {
	fmt.Println("=== Enhanced Error System Examples ===")
	
	// Example 1: Basic syntax error
	fmt.Println("1. Syntax Error Example:")
	syntaxError := createSyntaxErrorExample()
	PrintEnhancedError(syntaxError)
	
	// Example 2: Type error with suggestions
	fmt.Println("2. Type Error Example:")
	typeError := createTypeErrorExample()
	PrintEnhancedError(typeError)
	
	// Example 3: Runtime error with context
	fmt.Println("3. Runtime Error Example:")
	runtimeError := createRuntimeErrorExample()
	PrintEnhancedError(runtimeError)
	
	// Example 4: Complex error with multiple labels and suggestions
	fmt.Println("4. Complex Error Example:")
	complexError := createComplexErrorExample()
	PrintEnhancedError(complexError)
	
	// Example 5: Chained errors
	fmt.Println("5. Chained Error Example:")
	chainedError := createChainedErrorExample()
	PrintEnhancedError(chainedError)
}

// createSyntaxErrorExample creates a syntax error example
func createSyntaxErrorExample() *object.EnhancedError {
	
	span := object.ErrorSpan{
		Start: object.SourcePosition{
			Filename: "example.crl",
			Line:     1,
			Column:   17,
		},
		End: object.SourcePosition{
			Filename: "example.crl",
			Line:     1,
			Column:   18,
		},
		Source: "spell greet(name)",
	}
	
	err := object.NewSyntaxError("expected COLON, got NEWLINE", span).
		WithCode("E0001").
		WithTitle("Missing Colon")
	
	err.AddSuggestion(
		"Add missing colon",
		"Function definitions require a colon (:) at the end",
		object.ErrorFix{
			Span:        span,
			Replacement: ":",
			Description: "Add ':' after the function signature",
		},
	)
	
	err.AddNote(
		"In Carrion, all function definitions must end with a colon (:)",
		object.ERROR_LEVEL_HELP,
		nil,
	)
	
	return err
}

// createTypeErrorExample creates a type error example
func createTypeErrorExample() *object.EnhancedError {
	
	span := object.ErrorSpan{
		Start: object.SourcePosition{
			Filename: "example.crl",
			Line:     3,
			Column:   10,
		},
		End: object.SourcePosition{
			Filename: "example.crl",
			Line:     3,
			Column:   15,
		},
		Source: "result = x + y",
	}
	
	err := object.NewTypeError("cannot add STRING and INTEGER", span).
		WithCode("E0002").
		WithTitle("Type Mismatch")
	
	// Add labels for the operands
	err.AddLabel(
		object.ErrorSpan{
			Start: object.SourcePosition{
				Filename: "example.crl",
				Line:     3,
				Column:   10,
			},
			End: object.SourcePosition{
				Filename: "example.crl",
				Line:     3,
				Column:   11,
			},
			Source: "x",
		},
		"this is a STRING",
		object.ERROR_LEVEL_NOTE,
	)
	
	err.AddLabel(
		object.ErrorSpan{
			Start: object.SourcePosition{
				Filename: "example.crl",
				Line:     3,
				Column:   14,
			},
			End: object.SourcePosition{
				Filename: "example.crl",
				Line:     3,
				Column:   15,
			},
			Source: "y",
		},
		"this is an INTEGER",
		object.ERROR_LEVEL_NOTE,
	)
	
	err.AddSuggestion(
		"Convert types",
		"Convert one operand to match the other's type",
		object.ErrorFix{
			Description: "Convert string to integer: int(x) + y",
		},
		object.ErrorFix{
			Description: "Convert integer to string: x + str(y)",
		},
	)
	
	err.AddContext("left_operand", &object.String{Value: "hello"})
	err.AddContext("right_operand", &object.Integer{Value: 42})
	
	return err
}

// createRuntimeErrorExample creates a runtime error example
func createRuntimeErrorExample() *object.EnhancedError {
	
	span := object.ErrorSpan{
		Start: object.SourcePosition{
			Filename: "example.crl",
			Line:     3,
			Column:   9,
		},
		End: object.SourcePosition{
			Filename: "example.crl",
			Line:     3,
			Column:   19,
		},
		Source: "value = arr[index]",
	}
	
	err := object.NewRuntimeError("index 5 out of bounds for array of length 3", span).
		WithCode("E0003").
		WithTitle("Index Out of Bounds")
	
	err.AddContext("array_length", &object.Integer{Value: 3})
	err.AddContext("requested_index", &object.Integer{Value: 5})
	
	err.AddSuggestion(
		"Check array bounds",
		"Always verify that the index is within the array bounds",
		object.ErrorFix{
			Description: "Use: if index >= 0 and index < len(arr):",
		},
		object.ErrorFix{
			Description: "Use safe access: arr.get(index) or arr.get(index, default_value)",
		},
	)
	
	err.AddNote(
		"Array indices in Carrion start at 0 and go up to length - 1",
		object.ERROR_LEVEL_HELP,
		nil,
	)
	
	// Add stack trace
	err.AddStackEntry("main", object.SourcePosition{
		Filename: "example.crl",
		Line:     3,
		Column:   9,
	})
	
	return err
}

// createComplexErrorExample creates a complex error with multiple components
func createComplexErrorExample() *object.EnhancedError {
	
	span := object.ErrorSpan{
		Start: object.SourcePosition{
			Filename: "example.crl",
			Line:     7,
			Column:   9,
		},
		End: object.SourcePosition{
			Filename: "example.crl",
			Line:     7,
			Column:   27,
		},
		Source: "value = calculate(-5, 0)",
	}
	
	err := object.NewRuntimeError("ValueError: x must be positive", span).
		WithCode("E0004").
		WithTitle("Value Error")
	
	// Add multiple labels
	err.AddLabel(
		object.ErrorSpan{
			Start: object.SourcePosition{
				Filename: "example.crl",
				Line:     7,
				Column:   19,
			},
			End: object.SourcePosition{
				Filename: "example.crl",
				Line:     7,
				Column:   21,
			},
			Source: "-5",
		},
		"negative value not allowed",
		object.ERROR_LEVEL_ERROR,
	)
	
	err.AddLabel(
		object.ErrorSpan{
			Start: object.SourcePosition{
				Filename: "example.crl",
				Line:     7,
				Column:   23,
			},
			End: object.SourcePosition{
				Filename: "example.crl",
				Line:     7,
				Column:   24,
			},
			Source: "0",
		},
		"division by zero would occur",
		object.ERROR_LEVEL_WARNING,
	)
	
	// Add multiple suggestions
	err.AddSuggestion(
		"Fix input validation",
		"Ensure input values meet the function requirements",
		object.ErrorFix{
			Description: "Use a positive value for x: calculate(5, 0)",
		},
		object.ErrorFix{
			Description: "Use a non-zero value for y: calculate(-5, 1)",
		},
	)
	
	err.AddSuggestion(
		"Add error handling",
		"Handle errors gracefully in your code",
		object.ErrorFix{
			Description: "Use attempt/ensnare blocks to handle exceptions",
		},
	)
	
	// Add stack trace
	err.AddStackEntry("calculate", object.SourcePosition{
		Filename: "example.crl",
		Line:     3,
		Column:   9,
	})
	err.AddStackEntry("main", object.SourcePosition{
		Filename: "example.crl",
		Line:     7,
		Column:   9,
	})
	
	// Add context
	err.AddContext("function_name", &object.String{Value: "calculate"})
	err.AddContext("argument_x", &object.Integer{Value: -5})
	err.AddContext("argument_y", &object.Integer{Value: 0})
	
	return err
}

// createChainedErrorExample creates a chained error example
func createChainedErrorExample() *object.EnhancedError {
	
	// Create the primary error (division by zero)
	primarySpan := object.ErrorSpan{
		Start: object.SourcePosition{
			Filename: "example.crl",
			Line:     2,
			Column:   12,
		},
		End: object.SourcePosition{
			Filename: "example.crl",
			Line:     2,
			Column:   17,
		},
		Source: "return a / b",
	}
	
	primaryErr := object.NewRuntimeError("division by zero", primarySpan).
		WithCode("E0005").
		WithTitle("Division by Zero")
	
	primaryErr.AddContext("dividend", &object.Integer{Value: 10})
	primaryErr.AddContext("divisor", &object.Integer{Value: 0})
	
	// Create the secondary error (propagation)
	secondarySpan := object.ErrorSpan{
		Start: object.SourcePosition{
			Filename: "example.crl",
			Line:     5,
			Column:   14,
		},
		End: object.SourcePosition{
			Filename: "example.crl",
			Line:     5,
			Column:   35,
		},
		Source: "result = divide(data[0], data[1])",
	}
	
	secondaryErr := object.NewRuntimeError("error in process_data function", secondarySpan).
		WithCode("E0006").
		WithTitle("Function Call Error")
	
	// Chain the errors
	secondaryErr.WithCause(primaryErr)
	
	// Add stack trace
	secondaryErr.AddStackEntry("divide", object.SourcePosition{
		Filename: "example.crl",
		Line:     2,
		Column:   12,
	})
	secondaryErr.AddStackEntry("process_data", object.SourcePosition{
		Filename: "example.crl",
		Line:     5,
		Column:   14,
	})
	secondaryErr.AddStackEntry("main", object.SourcePosition{
		Filename: "example.crl",
		Line:     9,
		Column:   10,
	})
	
	// Add suggestions
	secondaryErr.AddSuggestion(
		"Add input validation",
		"Check for zero divisors before calling divide",
		object.ErrorFix{
			Description: "Add: if data[1] != 0: before calling divide",
		},
	)
	
	return secondaryErr
}

// DemonstrateErrorMigration shows how to migrate from old errors to enhanced errors
func DemonstrateErrorMigration() {
	fmt.Println("\n=== Error Migration Example ===")
	
	// Create an old-style error
	oldError := &object.Error{Message: "identifier not found: x"}
	
	// Create a span for the upgrade
	span := object.ErrorSpan{
		Start: object.SourcePosition{
			Filename: "migration.crl",
			Line:     1,
			Column:   10,
		},
		End: object.SourcePosition{
			Filename: "migration.crl",
			Line:     1,
			Column:   11,
		},
		Source: "result = x + y",
	}
	
	// Upgrade to enhanced error
	enhanced := object.UpgradeToEnhancedError(oldError, span)
	
	// Add suggestions
	enhanced = EnhanceErrorWithSuggestions(enhanced, "result = x + y")
	
	fmt.Println("Old error: ", oldError.Inspect())
	fmt.Println("\nUpgraded to enhanced error:")
	PrintEnhancedError(enhanced)
}

// ShowErrorComparison compares old and new error output
func ShowErrorComparison() {
	fmt.Println("\n=== Error Output Comparison ===")
	
	// Old style error
	fmt.Println("OLD STYLE ERROR:")
	fmt.Println("ERROR: identifier not found: x")
	fmt.Println("  at line 1, column 10")
	
	fmt.Println("\nNEW ENHANCED ERROR:")
	
	// Enhanced error
	span := object.ErrorSpan{
		Start: object.SourcePosition{
			Filename: "comparison.crl",
			Line:     1,
			Column:   10,
		},
		End: object.SourcePosition{
			Filename: "comparison.crl",
			Line:     1,
			Column:   11,
		},
		Source: "result = x + y",
	}
	
	enhanced := object.NewRuntimeError("identifier not found: x", span).
		WithCode("E0007").
		WithTitle("Undefined Variable")
	
	enhanced.AddSuggestion(
		"Define the variable",
		"Variables must be defined before use",
		object.ErrorFix{
			Description: "Add: x = value",
		},
	)
	
	PrintEnhancedError(enhanced)
}

// RunAllExamples runs all error system examples
func RunAllExamples() {
	ExampleEnhancedErrorSystem()
	DemonstrateErrorMigration()
	ShowErrorComparison()
}