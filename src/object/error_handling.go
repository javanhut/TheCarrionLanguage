// object/object.go
package object

import (
	"fmt"
	"strings"
)

const (
	CUSTOM_ERROR_OBJ = "USER DEFINED ERROR"
)

// CustomError represents a user-defined error in the language.
type CustomError struct {
	Name      string            // Name of the error type (e.g., "ValueError")
	Message   string            // Error message
	Details   map[string]Object // Additional details (optional)
	ErrorType *Grimoire         // The grimoire (class) this error belongs to (renamed to avoid conflict)
	Instance  *Instance         // Instance of the error (if applicable)
}

// Type returns the type of the object (implements the Object interface).
func (ce *CustomError) Type() ObjectType { return CUSTOM_ERROR_OBJ }

// Inspect returns a string representation of the error (implements the Object interface).
func (ce *CustomError) Inspect() string {
	var details []string
	for key, value := range ce.Details {
		details = append(details, fmt.Sprintf("%s: %s", key, value.Inspect()))
	}

	if len(details) > 0 {
		return fmt.Sprintf("%s: %s (%s)", ce.Name, ce.Message, strings.Join(details, ", "))
	}
	return fmt.Sprintf("%s: %s", ce.Name, ce.Message)
}

// NewCustomError creates a new CustomError object.
func NewCustomError(name, message string) *CustomError {
	return &CustomError{
		Name:    name,
		Message: message,
		Details: make(map[string]Object),
	}
}

// AddDetail adds a key-value pair to the error's details.
func (ce *CustomError) AddDetail(key string, value Object) {
	ce.Details[key] = value
}
