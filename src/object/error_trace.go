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

	// headline
	sb.WriteString("Error: ")
	sb.WriteString(e.Message)
	sb.WriteString("\n  at ")
	sb.WriteString(e.Position.Filename)
	sb.WriteString(fmt.Sprintf(", Line: %d, Column: %d\n",
		e.Position.Line, e.Position.Column))

	// stack (most‑recent first)
	if n := len(e.Stack); n > 0 {
		sb.WriteString("Stack trace:\n")
		for i := n - 1; i >= 0; i-- {
			f := e.Stack[i]
			name := f.FunctionName
			if name == "" {
				name = "<anon>"
			}
			sb.WriteString(fmt.Sprintf("  %d: %s (%s:Line: %d, Column: %d)\n",
				n-1-i, name, f.Position.Filename, f.Position.Line, f.Position.Column))
		}
	}

	// custom details
	if e.ErrorType == CUSTOM_ERROR_OBJ && len(e.CustomDetails) > 0 {
		sb.WriteString("Details:\n")
		for k, v := range e.CustomDetails {
			sb.WriteString("  ")
			sb.WriteString(k)
			sb.WriteString(": ")
			sb.WriteString(v.Inspect())
			sb.WriteString("\n")
		}
	}

	// chained cause
	if e.Cause != nil {
		sb.WriteString("\nCaused by:\n  ")
		sb.WriteString(strings.ReplaceAll(e.Cause.Inspect(), "\n", "\n  "))
		sb.WriteString("\n")
	}

	return sb.String()
}

// Optional, so fmt.Println(err) prints same view.
func (e *ErrorWithTrace) String() string { return e.Inspect() }

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
