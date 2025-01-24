// object/object.go
package object

import (
	"fmt"
	"strings"
)

const (
	CUSTOM_ERROR_OBJ = "USER DEFINED ERROR"
)

type CustomError struct {
	Name    string            // Name of the error type (e.g., "ValueError")
	Message string            // Error message
	Details map[string]Object // Additional details (optional)
}

func (e *CustomError) Type() ObjectType { return CUSTOM_ERROR_OBJ }
func (e *CustomError) Inspect() string {
	return fmt.Sprintf("%s: %s", e.Name, e.Message)
}

func (e *CustomError) ErrorDetails() string {
	details := []string{}
	for key, value := range e.Details {
		details = append(details, fmt.Sprintf("%s: %s", key, value.Inspect()))
	}
	return fmt.Sprintf("%s: %s\nDetails:\n%s", e.Name, e.Message, strings.Join(details, "\n"))
}
