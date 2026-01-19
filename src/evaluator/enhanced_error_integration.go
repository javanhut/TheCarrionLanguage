// src/evaluator/enhanced_error_integration.go
package evaluator

import (
	"fmt"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/utils"
)

// EnhancedEvaluator wraps the standard evaluator with enhanced error handling
type EnhancedEvaluator struct {
	Context *object.ErrorContext
}

// NewEnhancedEvaluator creates a new enhanced evaluator
func NewEnhancedEvaluator(filename, sourceCode string) *EnhancedEvaluator {
	return &EnhancedEvaluator{
		Context: object.NewErrorContext(filename, sourceCode),
	}
}

// Enhanced error creation functions for the evaluator

// newEnhancedError creates an enhanced error with current context
func (ee *EnhancedEvaluator) newEnhancedError(message string, args ...interface{}) *object.EnhancedError {
	formattedMessage := fmt.Sprintf(message, args...)
	return ee.Context.NewEnhancedRuntimeError(formattedMessage)
}

// newEnhancedTypeError creates an enhanced type error
func (ee *EnhancedEvaluator) newEnhancedTypeError(message string, expected, actual string) *object.EnhancedError {
	return ee.Context.NewEnhancedTypeError(message, expected, actual)
}

// newEnhancedIndexError creates an enhanced index error
func (ee *EnhancedEvaluator) newEnhancedIndexError(index, length int) *object.EnhancedError {
	return ee.Context.NewEnhancedIndexError(index, length)
}

// newEnhancedArgumentError creates an enhanced argument error
func (ee *EnhancedEvaluator) newEnhancedArgumentError(functionName string, expected, actual int) *object.EnhancedError {
	return ee.Context.NewEnhancedArgumentError(functionName, expected, actual)
}

// Enhanced error checking functions

// isEnhancedError checks if an object is an enhanced error
func isEnhancedError(obj object.Object) bool {
	return object.IsEnhancedError(obj)
}

// extractEnhancedError extracts an enhanced error from an object
func extractEnhancedError(obj object.Object) *object.EnhancedError {
	return object.ExtractEnhancedError(obj)
}

// Error propagation helpers

// propagateEnhancedError propagates an error with additional context
func (ee *EnhancedEvaluator) propagateEnhancedError(err object.Object) *object.EnhancedError {
	return object.PropagateError(err, ee.Context)
}

// wrapEnhancedError wraps an error with additional context
func (ee *EnhancedEvaluator) wrapEnhancedError(err object.Object, wrapMessage string) *object.EnhancedError {
	return object.WrapError(err, wrapMessage, ee.Context)
}

// Context management

// pushFunction pushes a function context onto the stack
func (ee *EnhancedEvaluator) pushFunction(functionName string, line, column int) {
	ee.Context.SetPosition(line, column)
	pos := ee.Context.CreatePosition()
	ee.Context.PushCallContext(functionName, pos)
}

// popFunction pops the current function context from the stack
func (ee *EnhancedEvaluator) popFunction() {
	ee.Context.PopCallContext()
}

// setPosition sets the current position in the source code
func (ee *EnhancedEvaluator) setPosition(line, column int) {
	ee.Context.SetPosition(line, column)
}

// addVariable adds a variable to the error context
func (ee *EnhancedEvaluator) addVariable(name string, value object.Object) {
	ee.Context.AddVariable(name, value.Inspect())
}

// Integration with existing evaluator functions

// EnhancedNewError creates an enhanced error with suggestions
func EnhancedNewError(message string, sourceCode string, filename string, line, column int) *object.EnhancedError {
	context := object.NewErrorContext(filename, sourceCode)
	context.SetPosition(line, column)
	
	err := context.NewEnhancedRuntimeError(message)
	
	// Add suggestions based on error pattern
	enhanced := utils.EnhanceErrorWithSuggestions(err, sourceCode)
	
	return enhanced
}

// EnhancedNewErrorf creates an enhanced error with formatted message
func EnhancedNewErrorf(sourceCode string, filename string, line, column int, format string, args ...interface{}) *object.EnhancedError {
	message := fmt.Sprintf(format, args...)
	return EnhancedNewError(message, sourceCode, filename, line, column)
}

// UpgradeExistingError upgrades an existing error to enhanced error
func UpgradeExistingError(err object.Object, sourceCode string, filename string, line, column int) *object.EnhancedError {
	if err == nil {
		return nil
	}
	
	// If it's already an enhanced error, return it
	if enhanced := extractEnhancedError(err); enhanced != nil {
		return enhanced
	}
	
	// Create context
	context := object.NewErrorContext(filename, sourceCode)
	context.SetPosition(line, column)
	
	// Upgrade the error
	enhanced := context.UpgradeError(err)
	
	// Add suggestions
	enhanced = utils.EnhanceErrorWithSuggestions(enhanced, sourceCode)
	
	return enhanced
}

// Helper functions for common error patterns

// CreateUndefinedIdentifierError creates an error for undefined identifiers
func CreateUndefinedIdentifierError(identifier string, sourceCode string, filename string, line, column int) *object.EnhancedError {
	context := object.NewErrorContext(filename, sourceCode)
	context.SetPosition(line, column)
	
	return context.CreateUndefinedVariableError(identifier)
}

// CreateTypeMismatchError creates a type mismatch error
func CreateTypeMismatchError(operation string, leftType, rightType object.ObjectType, sourceCode string, filename string, line, column int) *object.EnhancedError {
	context := object.NewErrorContext(filename, sourceCode)
	context.SetPosition(line, column)
	
	message := fmt.Sprintf("type mismatch: %s %s %s", leftType, operation, rightType)
	return context.NewEnhancedTypeError(message, string(leftType), string(rightType))
}

// CreateIndexOutOfBoundsError creates an index out of bounds error
func CreateIndexOutOfBoundsError(index, length int, sourceCode string, filename string, line, column int) *object.EnhancedError {
	context := object.NewErrorContext(filename, sourceCode)
	context.SetPosition(line, column)
	
	return context.NewEnhancedIndexError(index, length)
}

// CreateArgumentCountError creates an argument count error
func CreateArgumentCountError(functionName string, expected, actual int, sourceCode string, filename string, line, column int) *object.EnhancedError {
	context := object.NewErrorContext(filename, sourceCode)
	context.SetPosition(line, column)
	
	return context.NewEnhancedArgumentError(functionName, expected, actual)
}

// Enhanced error printing functions

// PrintEnhancedError prints an enhanced error with full context
func PrintEnhancedError(err *object.EnhancedError) {
	utils.PrintEnhancedError(err)
}

// PrintEnhancedErrorSimple prints an enhanced error with minimal context
func PrintEnhancedErrorSimple(err *object.EnhancedError) {
	// Create a simplified version for basic output
	fmt.Printf("Error: %s\n", err.Message)
	fmt.Printf("Location: %s\n", err.MainSpan.String())
	
	if len(err.Suggestions) > 0 {
		fmt.Printf("Suggestion: %s\n", err.Suggestions[0].Title)
	}
}

// Error detection and handling helpers

// HasEnhancedErrors checks if a slice of objects contains enhanced errors
func HasEnhancedErrors(objects []object.Object) bool {
	for _, obj := range objects {
		if isEnhancedError(obj) {
			return true
		}
	}
	return false
}

// CollectEnhancedErrors collects all enhanced errors from a slice of objects
func CollectEnhancedErrors(objects []object.Object) []*object.EnhancedError {
	var errors []*object.EnhancedError
	
	for _, obj := range objects {
		if enhanced := extractEnhancedError(obj); enhanced != nil {
			errors = append(errors, enhanced)
		}
	}
	
	return errors
}

// MergeEnhancedErrors merges multiple enhanced errors into a single error
func MergeEnhancedErrors(errors []*object.EnhancedError) *object.EnhancedError {
	return object.MergeErrors(errors)
}

// Configuration and setup

// SetupEnhancedErrorHandling configures enhanced error handling for the evaluator
func SetupEnhancedErrorHandling(evaluator *EnhancedEvaluator, config *object.ErrorConfig) {
	evaluator.Context.Config = config
}

// EnableVerboseErrors enables verbose error output
func EnableVerboseErrors(evaluator *EnhancedEvaluator) {
	evaluator.Context.Config.VerboseMode = true
}

// DisableEnhancedErrors disables enhanced error handling (fallback to basic errors)
func DisableEnhancedErrors(evaluator *EnhancedEvaluator) {
	evaluator.Context.Config.UseEnhancedErrors = false
}

// Example usage functions (for demonstration)

// ExampleEnhancedErrorUsage demonstrates how to use enhanced errors
func ExampleEnhancedErrorUsage() {
	// Create an enhanced evaluator
	evaluator := NewEnhancedEvaluator("example.crl", "x = y + z")
	
	// Set position
	evaluator.setPosition(1, 5)
	
	// Create an enhanced error
	err := evaluator.newEnhancedError("identifier not found: %s", "y")
	
	// Add suggestions
	err.AddSuggestion(
		"Define the variable",
		"Variables must be defined before use",
		object.ErrorFix{
			Description: "Add: y = value",
		},
	)
	
	// Print the error
	PrintEnhancedError(err)
}

// Migration helpers for existing code

// MigrateToEnhancedError migrates existing error handling code to use enhanced errors
func MigrateToEnhancedError(obj object.Object, sourceCode, filename string, line, column int) object.Object {
	// If it's already an enhanced error, return as-is
	if enhanced := extractEnhancedError(obj); enhanced != nil {
		return enhanced
	}
	
	// If it's a basic error, upgrade it
	if basicErr, ok := obj.(*object.Error); ok {
		return UpgradeExistingError(basicErr, sourceCode, filename, line, column)
	}
	
	// If it's an error with trace, upgrade it
	if traceErr, ok := obj.(*object.ErrorWithTrace); ok {
		return UpgradeExistingError(traceErr, sourceCode, filename, line, column)
	}
	
	// Not an error, return as-is
	return obj
}