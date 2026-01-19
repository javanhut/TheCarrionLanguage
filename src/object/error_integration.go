// src/object/error_integration.go
package object

import (
	"fmt"
	"strings"
)

// ErrorConfig holds configuration for error handling
type ErrorConfig struct {
	UseEnhancedErrors bool
	ShowSuggestions   bool
	ShowStackTrace    bool
	VerboseMode       bool
}

// DefaultErrorConfig returns a default error configuration
func DefaultErrorConfig() *ErrorConfig {
	return &ErrorConfig{
		UseEnhancedErrors: true,
		ShowSuggestions:   true,
		ShowStackTrace:    true,
		VerboseMode:       false,
	}
}

// ErrorContext holds context information for error creation
type ErrorContext struct {
	CurrentFile     string
	CurrentFunction string
	CurrentLine     int
	CurrentColumn   int
	SourceCode      string
	CallStack       []StackTraceEntry
	Environment     map[string]string // Variable state
	Config          *ErrorConfig
}

// NewErrorContext creates a new error context
func NewErrorContext(filename string, sourceCode string) *ErrorContext {
	return &ErrorContext{
		CurrentFile: filename,
		SourceCode:  sourceCode,
		Config:      DefaultErrorConfig(),
		Environment: make(map[string]string),
	}
}

// SetPosition updates the current position in the context
func (ctx *ErrorContext) SetPosition(line, column int) {
	ctx.CurrentLine = line
	ctx.CurrentColumn = column
}

// SetFunction updates the current function in the context
func (ctx *ErrorContext) SetFunction(functionName string) {
	ctx.CurrentFunction = functionName
}

// AddVariable adds a variable to the context environment
func (ctx *ErrorContext) AddVariable(name, value string) {
	ctx.Environment[name] = value
}

// CreatePosition creates a SourcePosition from the current context
func (ctx *ErrorContext) CreatePosition() SourcePosition {
	return SourcePosition{
		Filename: ctx.CurrentFile,
		Line:     ctx.CurrentLine,
		Column:   ctx.CurrentColumn,
	}
}

// CreateSpan creates an ErrorSpan from the current context
func (ctx *ErrorContext) CreateSpan() ErrorSpan {
	pos := ctx.CreatePosition()
	return ErrorSpan{
		Start:  pos,
		End:    pos,
		Source: ctx.getSourceLine(ctx.CurrentLine),
	}
}

// getSourceLine extracts a specific line from source code
func (ctx *ErrorContext) getSourceLine(lineNum int) string {
	if ctx.SourceCode == "" || lineNum <= 0 {
		return ""
	}
	
	lines := strings.Split(ctx.SourceCode, "\n")
	if lineNum > len(lines) {
		return ""
	}
	
	return lines[lineNum-1]
}

// Enhanced error creation functions that integrate with existing code

// NewEnhancedRuntimeError creates a runtime error with context
func (ctx *ErrorContext) NewEnhancedRuntimeError(message string) *EnhancedError {
	if !ctx.Config.UseEnhancedErrors {
		// Fall back to basic error for compatibility
		return UpgradeToEnhancedError(&Error{Message: message}, ctx.CreateSpan())
	}
	
	span := ctx.CreateSpan()
	err := NewRuntimeError(message, span)
	
	// Add context information
	if ctx.CurrentFunction != "" {
		err.AddContext("function", &String{Value: ctx.CurrentFunction})
	}
	
	// Add variable context
	for name, value := range ctx.Environment {
		err.AddContext(fmt.Sprintf("variable_%s", name), &String{Value: value})
	}
	
	// Add stack trace
	for _, entry := range ctx.CallStack {
		err.AddStackEntry(entry.FunctionName, entry.Position)
	}
	
	return err
}

// NewEnhancedTypeError creates a type error with context
func (ctx *ErrorContext) NewEnhancedTypeError(message string, expectedType, actualType string) *EnhancedError {
	span := ctx.CreateSpan()
	err := NewTypeError(message, span)
	
	// Add type-specific context
	err.AddContext("expected_type", &String{Value: expectedType})
	err.AddContext("actual_type", &String{Value: actualType})
	
	// Add type conversion suggestion
	err.AddSuggestion(
		"Type conversion",
		fmt.Sprintf("Convert %s to %s", actualType, expectedType),
		ErrorFix{
			Description: fmt.Sprintf("Use %s() to convert the value", expectedType),
		},
	)
	
	return err
}

// NewEnhancedSyntaxError creates a syntax error with context
func (ctx *ErrorContext) NewEnhancedSyntaxError(message string, token string) *EnhancedError {
	span := ctx.CreateSpan()
	err := NewSyntaxError(message, span)
	
	// Add token context
	err.AddContext("token", &String{Value: token})
	
	return err
}

// NewEnhancedIndexError creates an index error with context
func (ctx *ErrorContext) NewEnhancedIndexError(index int, length int) *EnhancedError {
	span := ctx.CreateSpan()
	message := fmt.Sprintf("index %d out of bounds for length %d", index, length)
	err := NewRuntimeError(message, span)
	
	// Add index-specific context
	err.AddContext("index", &Integer{Value: int64(index)})
	err.AddContext("length", &Integer{Value: int64(length)})
	
	// Add bounds checking suggestion
	err.AddSuggestion(
		"Bounds checking",
		"Always check array bounds before accessing elements",
		ErrorFix{
			Description: "Use: if index >= 0 and index < len(array):",
		},
		ErrorFix{
			Description: "Or use safe access methods that return None for invalid indices",
		},
	)
	
	return err
}

// NewEnhancedArgumentError creates an argument error with context
func (ctx *ErrorContext) NewEnhancedArgumentError(functionName string, expected, actual int) *EnhancedError {
	span := ctx.CreateSpan()
	message := fmt.Sprintf("function %s expects %d arguments, got %d", functionName, expected, actual)
	err := NewRuntimeError(message, span)
	
	// Add function-specific context
	err.AddContext("function", &String{Value: functionName})
	err.AddContext("expected_args", &Integer{Value: int64(expected)})
	err.AddContext("actual_args", &Integer{Value: int64(actual)})
	
	// Add argument suggestion
	if actual < expected {
		err.AddSuggestion(
			"Missing arguments",
			"The function call is missing required arguments",
			ErrorFix{
				Description: fmt.Sprintf("Add %d more argument(s) to the function call", expected-actual),
			},
		)
	} else {
		err.AddSuggestion(
			"Too many arguments",
			"The function call has too many arguments",
			ErrorFix{
				Description: fmt.Sprintf("Remove %d argument(s) from the function call", actual-expected),
			},
		)
	}
	
	return err
}

// Integration functions for existing error types

// UpgradeError upgrades an existing error to enhanced error with context
func (ctx *ErrorContext) UpgradeError(err Object) *EnhancedError {
	span := ctx.CreateSpan()
	enhanced := UpgradeToEnhancedError(err, span)
	
	// Add current context
	if ctx.CurrentFunction != "" {
		enhanced.AddContext("function", &String{Value: ctx.CurrentFunction})
	}
	
	// Add stack trace
	for _, entry := range ctx.CallStack {
		enhanced.AddStackEntry(entry.FunctionName, entry.Position)
	}
	
	return enhanced
}

// IsEnhancedError checks if an object is an enhanced error
func IsEnhancedError(obj Object) bool {
	_, ok := obj.(*EnhancedError)
	return ok
}

// ExtractEnhancedError extracts an enhanced error from an object
func ExtractEnhancedError(obj Object) *EnhancedError {
	if enhanced, ok := obj.(*EnhancedError); ok {
		return enhanced
	}
	return nil
}

// Utility functions for error handling

// ChainErrors creates a chain of related errors
func ChainErrors(primary *EnhancedError, secondary *EnhancedError) *EnhancedError {
	if primary == nil {
		return secondary
	}
	if secondary == nil {
		return primary
	}
	
	return primary.WithCause(secondary)
}

// MergeErrors merges multiple errors into a single enhanced error
func MergeErrors(errors []*EnhancedError) *EnhancedError {
	if len(errors) == 0 {
		return nil
	}
	
	if len(errors) == 1 {
		return errors[0]
	}
	
	// Use the first error as the primary
	primary := errors[0]
	
	// Add others as related errors
	for i := 1; i < len(errors); i++ {
		primary.AddRelatedError(errors[i])
	}
	
	return primary
}

// CreateCompoundError creates a compound error from multiple individual errors
func CreateCompoundError(errors []Object, context *ErrorContext) *EnhancedError {
	if len(errors) == 0 {
		return nil
	}
	
	var enhancedErrors []*EnhancedError
	
	// Convert all errors to enhanced errors
	for _, err := range errors {
		if enhanced := ExtractEnhancedError(err); enhanced != nil {
			enhancedErrors = append(enhancedErrors, enhanced)
		} else {
			// Upgrade basic errors
			enhanced := context.UpgradeError(err)
			enhancedErrors = append(enhancedErrors, enhanced)
		}
	}
	
	return MergeErrors(enhancedErrors)
}

// Error propagation helpers

// PropagateError propagates an error up the call stack
func PropagateError(err Object, context *ErrorContext) *EnhancedError {
	if err == nil {
		return nil
	}
	
	// If it's already an enhanced error, just add current context
	if enhanced := ExtractEnhancedError(err); enhanced != nil {
		// Add current function to stack if not already there
		if context.CurrentFunction != "" {
			enhanced.AddStackEntry(context.CurrentFunction, context.CreatePosition())
		}
		return enhanced
	}
	
	// Upgrade to enhanced error
	return context.UpgradeError(err)
}

// WrapError wraps an error with additional context
func WrapError(err Object, wrapMessage string, context *ErrorContext) *EnhancedError {
	if err == nil {
		return nil
	}
	
	enhanced := context.UpgradeError(err)
	
	// Create a new error that wraps the original
	span := context.CreateSpan()
	wrapper := NewRuntimeError(wrapMessage, span)
	wrapper.WithCause(enhanced)
	
	return wrapper
}

// Context management for error handling

// PushCallContext pushes a new call context onto the stack
func (ctx *ErrorContext) PushCallContext(functionName string, position SourcePosition) {
	ctx.CallStack = append(ctx.CallStack, StackTraceEntry{
		FunctionName: functionName,
		Position:     position,
	})
	ctx.SetFunction(functionName)
}

// PopCallContext pops the current call context from the stack
func (ctx *ErrorContext) PopCallContext() {
	if len(ctx.CallStack) > 0 {
		ctx.CallStack = ctx.CallStack[:len(ctx.CallStack)-1]
	}
	
	// Update current function
	if len(ctx.CallStack) > 0 {
		ctx.CurrentFunction = ctx.CallStack[len(ctx.CallStack)-1].FunctionName
	} else {
		ctx.CurrentFunction = ""
	}
}

// ClearCallStack clears the call stack
func (ctx *ErrorContext) ClearCallStack() {
	ctx.CallStack = []StackTraceEntry{}
	ctx.CurrentFunction = ""
}

// Helper functions for specific error types

// CreateUndefinedVariableError creates an error for undefined variables
func (ctx *ErrorContext) CreateUndefinedVariableError(variableName string) *EnhancedError {
	span := ctx.CreateSpan()
	message := fmt.Sprintf("identifier not found: %s", variableName)
	err := NewRuntimeError(message, span)
	
	err.AddContext("variable_name", &String{Value: variableName})
	
	// Add helpful suggestions
	err.AddSuggestion(
		"Define the variable",
		"Variables must be defined before use",
		ErrorFix{
			Description: fmt.Sprintf("Add: %s = value", variableName),
		},
	)
	
	// Check for similar variable names
	ctx.suggestSimilarVariables(err, variableName)
	
	return err
}

// suggestSimilarVariables suggests similar variable names
func (ctx *ErrorContext) suggestSimilarVariables(err *EnhancedError, target string) {
	var suggestions []string
	
	// Simple similarity check (can be enhanced with better algorithms)
	for name := range ctx.Environment {
		if strings.Contains(name, target) || strings.Contains(target, name) {
			suggestions = append(suggestions, name)
		}
	}
	
	if len(suggestions) > 0 {
		err.AddNote(
			fmt.Sprintf("Did you mean: %s", strings.Join(suggestions, ", ")),
			ERROR_LEVEL_HELP,
			nil,
		)
	}
}