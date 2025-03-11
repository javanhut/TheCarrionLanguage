// src/object/error_trace.go
package object

import (
	"fmt"
	"strings"
)

// SourcePosition tracks the location of code in source files
type SourcePosition struct {
	Filename string
	Line     int
	Column   int
}

func (sp SourcePosition) String() string {
	if sp.Filename == "" || sp.Line == 0 {
		return "unknown position"
	}
	return fmt.Sprintf("%s:%d:%d", sp.Filename, sp.Line, sp.Column)
}

// StackTraceEntry represents a single frame in the error stack trace
type StackTraceEntry struct {
	FunctionName string
	Position     SourcePosition
}

func (ste StackTraceEntry) String() string {
	return fmt.Sprintf("at %s (%s)", ste.FunctionName, ste.Position)
}

// ErrorWithTrace extends the basic error with stack trace and source position information
type ErrorWithTrace struct {
	ErrorType     ObjectType // Changed from Type to ErrorType to avoid conflict
	Message       string
	Position      SourcePosition
	Cause         *ErrorWithTrace
	Stack         []StackTraceEntry
	CustomDetails map[string]Object
}

// To satisfy the Object interface
func (e *ErrorWithTrace) Type() ObjectType {
	return e.ErrorType // Return the stored error type
}

func (e *ErrorWithTrace) Inspect() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Error: %s\n", e.Message))
	sb.WriteString(
		fmt.Sprintf(
			"  at %s, Line: %d, Column: %d\n",
			e.Position.Filename,
			e.Position.Line,
			e.Position.Column,
		),
	)

	// Add the stack trace
	if len(e.Stack) > 0 {
		sb.WriteString("Stack trace:\n")
		for i, entry := range e.Stack {
			sb.WriteString(fmt.Sprintf(
				"  %d: %s (%s:Line: %d, Column: %d)\n",
				i,
				entry.FunctionName,
				entry.Position.Filename,
				entry.Position.Line,
				entry.Position.Column,
			))
		}
	}

	if e.ErrorType == CUSTOM_ERROR_OBJ && len(e.CustomDetails) > 0 {
		sb.WriteString("Details:\n")
		for key, value := range e.CustomDetails {
			sb.WriteString(fmt.Sprintf("  %s: %s\n", key, value.Inspect()))
		}
	}

	// Add the cause if present
	if e.Cause != nil {
		sb.WriteString("\nCaused by:\n")
		causeLines := strings.Split(e.Cause.Inspect(), "\n")
		for _, line := range causeLines {
			sb.WriteString("  " + line + "\n")
		}
	}

	return sb.String()
}

// Added setter methods for fluent API
func (e *ErrorWithTrace) WithCause(cause *ErrorWithTrace) *ErrorWithTrace {
	e.Cause = cause
	return e
}

func (e *ErrorWithTrace) AddDetail(key string, value Object) *ErrorWithTrace {
	if e.CustomDetails == nil {
		e.CustomDetails = make(map[string]Object)
	}
	e.CustomDetails[key] = value
	return e
}

func (e *ErrorWithTrace) AddStackEntry(functionName string, pos SourcePosition) *ErrorWithTrace {
	e.Stack = append(e.Stack, StackTraceEntry{
		FunctionName: functionName,
		Position:     pos,
	})
	return e
}
