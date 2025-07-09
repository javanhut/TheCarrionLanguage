// src/object/enhanced_errors.go
package object

import (
	"fmt"
	"strings"
)

// ErrorLevel defines the severity of an error
type ErrorLevel int

const (
	ERROR_LEVEL_ERROR ErrorLevel = iota
	ERROR_LEVEL_WARNING
	ERROR_LEVEL_NOTE
	ERROR_LEVEL_HELP
)

func (el ErrorLevel) String() string {
	switch el {
	case ERROR_LEVEL_ERROR:
		return "error"
	case ERROR_LEVEL_WARNING:
		return "warning"
	case ERROR_LEVEL_NOTE:
		return "note"
	case ERROR_LEVEL_HELP:
		return "help"
	default:
		return "unknown"
	}
}

// ErrorCategory defines the category of error for better organization
type ErrorCategory int

const (
	ERROR_CATEGORY_SYNTAX ErrorCategory = iota
	ERROR_CATEGORY_TYPE
	ERROR_CATEGORY_RUNTIME
	ERROR_CATEGORY_SEMANTIC
	ERROR_CATEGORY_IMPORT
	ERROR_CATEGORY_IO
	ERROR_CATEGORY_CUSTOM
)

func (ec ErrorCategory) String() string {
	switch ec {
	case ERROR_CATEGORY_SYNTAX:
		return "syntax"
	case ERROR_CATEGORY_TYPE:
		return "type"
	case ERROR_CATEGORY_RUNTIME:
		return "runtime"
	case ERROR_CATEGORY_SEMANTIC:
		return "semantic"
	case ERROR_CATEGORY_IMPORT:
		return "import"
	case ERROR_CATEGORY_IO:
		return "io"
	case ERROR_CATEGORY_CUSTOM:
		return "custom"
	default:
		return "unknown"
	}
}

// ErrorSpan represents a span of code in the source
type ErrorSpan struct {
	Start  SourcePosition
	End    SourcePosition
	Source string // The actual source code that caused the error
}

func (es ErrorSpan) String() string {
	if es.Start.Filename == es.End.Filename {
		if es.Start.Line == es.End.Line {
			return fmt.Sprintf("%s:%d:%d-%d", es.Start.Filename, es.Start.Line, es.Start.Column, es.End.Column)
		}
		return fmt.Sprintf("%s:%d:%d-%d:%d", es.Start.Filename, es.Start.Line, es.Start.Column, es.End.Line, es.End.Column)
	}
	return fmt.Sprintf("%s:%d:%d to %s:%d:%d", es.Start.Filename, es.Start.Line, es.Start.Column, es.End.Filename, es.End.Line, es.End.Column)
}

// ErrorLabel represents a label pointing to a specific part of code
type ErrorLabel struct {
	Span    ErrorSpan
	Message string
	Level   ErrorLevel
}

// ErrorSuggestion represents a suggestion for fixing the error
type ErrorSuggestion struct {
	Title       string
	Description string
	Fixes       []ErrorFix
}

// ErrorFix represents a concrete fix suggestion
type ErrorFix struct {
	Span        ErrorSpan
	Replacement string
	Description string
}

// ErrorNote represents additional information about an error
type ErrorNote struct {
	Message string
	Level   ErrorLevel
	Span    *ErrorSpan // Optional span for the note
}

// EnhancedError represents a comprehensive error with detailed context and suggestions
type EnhancedError struct {
	ErrorType     ObjectType
	Code          string          // Error code (e.g., "E0001", "SYNTAX_ERROR")
	Title         string          // Main error title
	Message       string          // Detailed error message
	Level         ErrorLevel      // Error severity level
	Category      ErrorCategory   // Error category
	MainSpan      ErrorSpan       // Primary error location
	Labels        []ErrorLabel    // Additional labels pointing to relevant code
	Suggestions   []ErrorSuggestion // Suggestions for fixing the error
	Notes         []ErrorNote     // Additional notes and context
	Cause         *EnhancedError  // Underlying cause
	Stack         []StackTraceEntry // Call stack
	Context       map[string]Object // Additional context information
	RelatedErrors []*EnhancedError  // Related errors
}

// Implement Object interface
func (e *EnhancedError) Type() ObjectType {
	return e.ErrorType
}

func (e *EnhancedError) Inspect() string {
	return e.Message
}

func (e *EnhancedError) String() string {
	return e.Message
}

// Builder methods for fluent API
func (e *EnhancedError) WithCode(code string) *EnhancedError {
	e.Code = code
	return e
}

func (e *EnhancedError) WithTitle(title string) *EnhancedError {
	e.Title = title
	return e
}

func (e *EnhancedError) WithLevel(level ErrorLevel) *EnhancedError {
	e.Level = level
	return e
}

func (e *EnhancedError) WithCategory(category ErrorCategory) *EnhancedError {
	e.Category = category
	return e
}

func (e *EnhancedError) WithSpan(span ErrorSpan) *EnhancedError {
	e.MainSpan = span
	return e
}

func (e *EnhancedError) WithCause(cause *EnhancedError) *EnhancedError {
	e.Cause = cause
	return e
}

func (e *EnhancedError) AddLabel(span ErrorSpan, message string, level ErrorLevel) *EnhancedError {
	e.Labels = append(e.Labels, ErrorLabel{
		Span:    span,
		Message: message,
		Level:   level,
	})
	return e
}

func (e *EnhancedError) AddSuggestion(title, description string, fixes ...ErrorFix) *EnhancedError {
	e.Suggestions = append(e.Suggestions, ErrorSuggestion{
		Title:       title,
		Description: description,
		Fixes:       fixes,
	})
	return e
}

func (e *EnhancedError) AddNote(message string, level ErrorLevel, span *ErrorSpan) *EnhancedError {
	e.Notes = append(e.Notes, ErrorNote{
		Message: message,
		Level:   level,
		Span:    span,
	})
	return e
}

func (e *EnhancedError) AddStackEntry(functionName string, pos SourcePosition) *EnhancedError {
	e.Stack = append(e.Stack, StackTraceEntry{
		FunctionName: functionName,
		Position:     pos,
	})
	return e
}

func (e *EnhancedError) AddContext(key string, value Object) *EnhancedError {
	if e.Context == nil {
		e.Context = make(map[string]Object)
	}
	e.Context[key] = value
	return e
}

func (e *EnhancedError) AddRelatedError(related *EnhancedError) *EnhancedError {
	e.RelatedErrors = append(e.RelatedErrors, related)
	return e
}

// Constructor functions
func NewEnhancedError(errorType ObjectType, message string, span ErrorSpan) *EnhancedError {
	return &EnhancedError{
		ErrorType:   errorType,
		Message:     message,
		Level:       ERROR_LEVEL_ERROR,
		Category:    ERROR_CATEGORY_RUNTIME,
		MainSpan:    span,
		Labels:      []ErrorLabel{},
		Suggestions: []ErrorSuggestion{},
		Notes:       []ErrorNote{},
		Context:     make(map[string]Object),
	}
}

func NewSyntaxError(message string, span ErrorSpan) *EnhancedError {
	return NewEnhancedError(ERROR_OBJ, message, span).
		WithCode("SYNTAX_ERROR").
		WithCategory(ERROR_CATEGORY_SYNTAX).
		WithTitle("Syntax Error")
}

func NewTypeError(message string, span ErrorSpan) *EnhancedError {
	return NewEnhancedError(ERROR_OBJ, message, span).
		WithCode("TYPE_ERROR").
		WithCategory(ERROR_CATEGORY_TYPE).
		WithTitle("Type Error")
}

func NewRuntimeError(message string, span ErrorSpan) *EnhancedError {
	return NewEnhancedError(ERROR_OBJ, message, span).
		WithCode("RUNTIME_ERROR").
		WithCategory(ERROR_CATEGORY_RUNTIME).
		WithTitle("Runtime Error")
}

func NewSemanticError(message string, span ErrorSpan) *EnhancedError {
	return NewEnhancedError(ERROR_OBJ, message, span).
		WithCode("SEMANTIC_ERROR").
		WithCategory(ERROR_CATEGORY_SEMANTIC).
		WithTitle("Semantic Error")
}

func NewImportError(message string, span ErrorSpan) *EnhancedError {
	return NewEnhancedError(ERROR_OBJ, message, span).
		WithCode("IMPORT_ERROR").
		WithCategory(ERROR_CATEGORY_IMPORT).
		WithTitle("Import Error")
}

func NewIOError(message string, span ErrorSpan) *EnhancedError {
	return NewEnhancedError(ERROR_OBJ, message, span).
		WithCode("IO_ERROR").
		WithCategory(ERROR_CATEGORY_IO).
		WithTitle("I/O Error")
}

// Predefined error suggestions for common cases
var CommonErrorSuggestions = map[string]ErrorSuggestion{
	"UNDEFINED_VARIABLE": {
		Title:       "Variable not defined",
		Description: "The variable you're trying to use hasn't been defined yet.",
		Fixes: []ErrorFix{
			{
				Description: "Define the variable before using it",
			},
		},
	},
	"UNDEFINED_FUNCTION": {
		Title:       "Function not defined",
		Description: "The function you're trying to call hasn't been defined yet.",
		Fixes: []ErrorFix{
			{
				Description: "Define the function before calling it",
			},
		},
	},
	"WRONG_ARGUMENT_COUNT": {
		Title:       "Wrong number of arguments",
		Description: "The function call has the wrong number of arguments.",
		Fixes: []ErrorFix{
			{
				Description: "Check the function definition and provide the correct number of arguments",
			},
		},
	},
	"INVALID_ASSIGNMENT": {
		Title:       "Invalid assignment target",
		Description: "You can only assign to variables, not to expressions.",
		Fixes: []ErrorFix{
			{
				Description: "Make sure the left side of the assignment is a variable",
			},
		},
	},
	"DIVISION_BY_ZERO": {
		Title:       "Division by zero",
		Description: "You cannot divide by zero.",
		Fixes: []ErrorFix{
			{
				Description: "Check that the divisor is not zero before dividing",
			},
		},
	},
	"TYPE_MISMATCH": {
		Title:       "Type mismatch",
		Description: "The operation cannot be performed on values of these types.",
		Fixes: []ErrorFix{
			{
				Description: "Convert the values to compatible types before performing the operation",
			},
		},
	},
	"EXPECTED_TOKEN": {
		Title:       "Expected token",
		Description: "The parser expected a specific token but found something else.",
		Fixes: []ErrorFix{
			{
				Description: "Check the syntax and add the missing token",
			},
		},
	},
	"UNEXPECTED_TOKEN": {
		Title:       "Unexpected token",
		Description: "The parser found a token that doesn't belong here.",
		Fixes: []ErrorFix{
			{
				Description: "Remove the unexpected token or check the syntax",
			},
		},
	},
	"MISSING_COLON": {
		Title:       "Missing colon",
		Description: "Control structures like if, for, while, etc. require a colon at the end.",
		Fixes: []ErrorFix{
			{
				Description: "Add a colon (:) at the end of the statement",
			},
		},
	},
	"INDENTATION_ERROR": {
		Title:       "Indentation error",
		Description: "The code is not properly indented.",
		Fixes: []ErrorFix{
			{
				Description: "Use consistent indentation (4 spaces per level)",
			},
		},
	},
}

// Helper function to get suggestions based on error patterns
func GetSuggestionForError(errorMessage string) *ErrorSuggestion {
	// Simple pattern matching for common errors
	lowerMessage := strings.ToLower(errorMessage)
	
	if strings.Contains(lowerMessage, "identifier not found") || strings.Contains(lowerMessage, "undefined") {
		if strings.Contains(lowerMessage, "function") {
			suggestion := CommonErrorSuggestions["UNDEFINED_FUNCTION"]
			return &suggestion
		}
		suggestion := CommonErrorSuggestions["UNDEFINED_VARIABLE"]
		return &suggestion
	}
	
	if strings.Contains(lowerMessage, "wrong number of arguments") {
		suggestion := CommonErrorSuggestions["WRONG_ARGUMENT_COUNT"]
		return &suggestion
	}
	
	if strings.Contains(lowerMessage, "invalid assignment") {
		suggestion := CommonErrorSuggestions["INVALID_ASSIGNMENT"]
		return &suggestion
	}
	
	if strings.Contains(lowerMessage, "division by zero") {
		suggestion := CommonErrorSuggestions["DIVISION_BY_ZERO"]
		return &suggestion
	}
	
	if strings.Contains(lowerMessage, "type mismatch") || strings.Contains(lowerMessage, "cannot be used") {
		suggestion := CommonErrorSuggestions["TYPE_MISMATCH"]
		return &suggestion
	}
	
	if strings.Contains(lowerMessage, "expected") && strings.Contains(lowerMessage, "token") {
		suggestion := CommonErrorSuggestions["EXPECTED_TOKEN"]
		return &suggestion
	}
	
	if strings.Contains(lowerMessage, "unexpected") && strings.Contains(lowerMessage, "token") {
		suggestion := CommonErrorSuggestions["UNEXPECTED_TOKEN"]
		return &suggestion
	}
	
	if strings.Contains(lowerMessage, "missing") && strings.Contains(lowerMessage, "colon") {
		suggestion := CommonErrorSuggestions["MISSING_COLON"]
		return &suggestion
	}
	
	if strings.Contains(lowerMessage, "indent") {
		suggestion := CommonErrorSuggestions["INDENTATION_ERROR"]
		return &suggestion
	}
	
	return nil
}

// Helper function to convert old errors to enhanced errors
func UpgradeToEnhancedError(oldError Object, span ErrorSpan) *EnhancedError {
	if oldError == nil {
		return nil
	}
	
	switch err := oldError.(type) {
	case *Error:
		enhanced := NewRuntimeError(err.Message, span)
		if suggestion := GetSuggestionForError(err.Message); suggestion != nil {
			enhanced.AddSuggestion(suggestion.Title, suggestion.Description, suggestion.Fixes...)
		}
		return enhanced
		
	case *ErrorWithTrace:
		enhanced := NewRuntimeError(err.Message, span)
		enhanced.Stack = err.Stack
		if err.Cause != nil {
			enhanced.Cause = UpgradeToEnhancedError(err.Cause, span)
		}
		for key, value := range err.CustomDetails {
			enhanced.AddContext(key, value)
		}
		if suggestion := GetSuggestionForError(err.Message); suggestion != nil {
			enhanced.AddSuggestion(suggestion.Title, suggestion.Description, suggestion.Fixes...)
		}
		return enhanced
		
	case *CustomError:
		enhanced := NewEnhancedError(CUSTOM_ERROR_OBJ, err.Message, span).
			WithCode("CUSTOM_ERROR").
			WithCategory(ERROR_CATEGORY_CUSTOM).
			WithTitle(err.Name)
		for key, value := range err.Details {
			enhanced.AddContext(key, value)
		}
		return enhanced
		
	default:
		return NewRuntimeError(oldError.Inspect(), span)
	}
}