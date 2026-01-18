// src/utils/error_suggestions.go
package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/javanhut/TheCarrionLanguage/src/object"
)

// ErrorPattern represents a pattern for matching error messages
type ErrorPattern struct {
	Pattern     *regexp.Regexp
	Suggestion  object.ErrorSuggestion
	ContextHelp string
}

// Common error patterns with suggestions
var ErrorPatterns = []ErrorPattern{
	// Undefined variable/function errors
	{
		Pattern: regexp.MustCompile(`identifier not found: (.+)`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Undefined identifier",
			Description: "The identifier you're trying to use hasn't been defined in the current scope.",
			Fixes: []object.ErrorFix{
				{
					Description: "Define the identifier before using it",
				},
				{
					Description: "Check for typos in the identifier name",
				},
				{
					Description: "Import the module containing this identifier",
				},
			},
		},
		ContextHelp: "Variables and functions must be defined before they can be used. In Carrion, use 'spell' for functions and simple assignment for variables.",
	},
	
	// Type mismatch with STRING - suggest conversion
	{
		Pattern: regexp.MustCompile(`type mismatch:.*STRING.*(INTEGER|FLOAT)|type mismatch:.*(INTEGER|FLOAT).*STRING`),
		Suggestion: object.ErrorSuggestion{
			Title:       "String/Number type mismatch",
			Description: "You're mixing strings and numbers in an operation.",
			Fixes: []object.ErrorFix{
				{
					Description: "Convert string to number: int(\"123\") or float(\"3.14\")",
				},
				{
					Description: "Convert number to string: str(123) or str(3.14)",
				},
				{
					Description: "Use string concatenation (+) only with strings, arithmetic only with numbers",
				},
			},
		},
		ContextHelp: "In Carrion, + means concatenation for strings and addition for numbers. Convert types explicitly to avoid ambiguity.",
	},

	// Type mismatch with BOOLEAN - suggest comparison or conversion
	{
		Pattern: regexp.MustCompile(`type mismatch:.*(BOOLEAN).*(INTEGER|FLOAT|STRING)|type mismatch:.*(INTEGER|FLOAT|STRING).*(BOOLEAN)`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Boolean type mismatch",
			Description: "You're mixing boolean with other types in an operation.",
			Fixes: []object.ErrorFix{
				{
					Description: "Use comparison operators (==, !=, <, >) to get boolean results",
				},
				{
					Description: "Convert to boolean with bool(): bool(0) is False, bool(1) is True",
				},
				{
					Description: "Convert boolean to int: int(True) is 1, int(False) is 0",
				},
			},
		},
		ContextHelp: "Booleans represent True/False values. Use comparisons to create booleans, or convert explicitly.",
	},

	// Type mismatch for assignment with type hints
	{
		Pattern: regexp.MustCompile(`type mismatch: cannot assign (\w+) to variable '(\w+)' with type hint (\w+)`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Type hint violation",
			Description: "You're assigning a value that doesn't match the variable's declared type.",
			Fixes: []object.ErrorFix{
				{
					Description: "Convert the value to match the type hint before assignment",
				},
				{
					Description: "Remove or update the type hint if you need flexibility",
				},
				{
					Description: "Check that your data source provides the expected type",
				},
			},
		},
		ContextHelp: "Type hints in Carrion enforce type safety. Use conversion functions to match types.",
	},

	// Generic type mismatch (fallback)
	{
		Pattern: regexp.MustCompile(`type mismatch: (\w+) (.+) (\w+)`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Type mismatch",
			Description: "You're trying to perform an operation on incompatible types.",
			Fixes: []object.ErrorFix{
				{
					Description: "Use type conversion: int(), float(), str(), bool(), list(), tuple()",
				},
				{
					Description: "Check that both operands are compatible for this operation",
				},
				{
					Description: "For arithmetic: both values must be numbers (INTEGER or FLOAT)",
				},
			},
		},
		ContextHelp: "Carrion has dynamic typing but operations require compatible types. Use conversion functions when needed.",
	},
	
	// Wrong number of arguments
	{
		Pattern: regexp.MustCompile(`wrong number of arguments: want=(\d+), got=(\d+)`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Incorrect argument count",
			Description: "The function call has the wrong number of arguments.",
			Fixes: []object.ErrorFix{
				{
					Description: "Check the function definition and provide the correct number of arguments",
				},
				{
					Description: "Use default parameters if the function supports them",
				},
				{
					Description: "Review the function signature for required vs optional parameters",
				},
			},
		},
		ContextHelp: "Function calls must match the number of parameters defined in the function signature.",
	},
	
	// Division by zero
	{
		Pattern: regexp.MustCompile(`division by zero`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Division by zero",
			Description: "Division by zero is undefined and not allowed.",
			Fixes: []object.ErrorFix{
				{
					Description: "Check if the divisor is zero before performing division",
				},
				{
					Description: "Use conditional logic: if divisor != 0:",
				},
				{
					Description: "Handle the zero case with appropriate error handling",
				},
			},
		},
		ContextHelp: "Always validate that divisors are non-zero before performing division operations.",
	},
	
	// Invalid assignment target
	{
		Pattern: regexp.MustCompile(`invalid assignment target`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Invalid assignment target",
			Description: "You can only assign values to variables, not to expressions or literals.",
			Fixes: []object.ErrorFix{
				{
					Description: "Ensure the left side of '=' is a variable name",
				},
				{
					Description: "Use method calls for object property assignment",
				},
				{
					Description: "For arrays, use methods like append() or set() instead of direct indexing",
				},
			},
		},
		ContextHelp: "In Carrion, you cannot assign to expressions like function calls or array indices directly.",
	},
	
	// Index out of bounds (with valid range info)
	{
		Pattern: regexp.MustCompile(`index out of bounds: (-?\d+).*(?:length: (\d+)|is empty)`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Index out of bounds",
			Description: "You're trying to access an index outside the valid range.",
			Fixes: []object.ErrorFix{
				{
					Description: "Indices start at 0, not 1 - if you meant the last element, use index length-1",
				},
				{
					Description: "Check bounds first: if index >= 0 and index < len(collection):",
				},
				{
					Description: "Use negative indices for end-relative access: -1 for last, -2 for second-to-last",
				},
			},
		},
		ContextHelp: "Array, tuple, and string indices start at 0. For a collection of length N, valid indices are 0 to N-1. Negative indices count from the end (-1 is the last element).",
	},

	// Empty collection index error
	{
		Pattern: regexp.MustCompile(`index out of bounds: (-?\d+) \((array|tuple|string) is empty\)`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Cannot index empty collection",
			Description: "The collection has no elements to access.",
			Fixes: []object.ErrorFix{
				{
					Description: "Check if collection is empty before indexing: if len(collection) > 0:",
				},
				{
					Description: "Ensure the collection was properly initialized with data",
				},
				{
					Description: "Handle the empty case separately in your logic",
				},
			},
		},
		ContextHelp: "An empty array, tuple, or string has no valid indices. Always check length before accessing.",
	},
	
	// Parsing errors
	{
		Pattern: regexp.MustCompile(`expected next token to be (.+), got (.+)`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Unexpected token",
			Description: "The parser expected a different token at this position.",
			Fixes: []object.ErrorFix{
				{
					Description: "Check the syntax and add the expected token",
				},
				{
					Description: "Review language syntax rules for this construct",
				},
				{
					Description: "Look for missing punctuation or keywords",
				},
			},
		},
		ContextHelp: "Syntax errors occur when the code doesn't follow Carrion's grammar rules.",
	},
	
	// Missing colon
	{
		Pattern: regexp.MustCompile(`expected.*COLON`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Missing colon",
			Description: "Control structures and function definitions require a colon (:) at the end.",
			Fixes: []object.ErrorFix{
				{
					Description: "Add ':' at the end of if, for, while, spell, or grim statements",
				},
				{
					Description: "Example: if condition: or spell function_name():",
				},
			},
		},
		ContextHelp: "Colons are required after control flow keywords and function/class definitions in Carrion.",
	},
	
	// Indentation errors
	{
		Pattern: regexp.MustCompile(`expected.*INDENT`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Missing indentation",
			Description: "Code blocks after ':' must be indented.",
			Fixes: []object.ErrorFix{
				{
					Description: "Indent the code block with 4 spaces",
				},
				{
					Description: "Ensure consistent indentation throughout the block",
				},
				{
					Description: "Use spaces, not tabs, for indentation",
				},
			},
		},
		ContextHelp: "Carrion uses indentation to define code blocks, similar to Python. Use 4 spaces per indentation level.",
	},

	// Mixed tabs and spaces in indentation
	{
		Pattern: regexp.MustCompile(`mixed tabs and spaces in indentation`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Mixed indentation",
			Description: "You have both tabs and spaces on the same indented line.",
			Fixes: []object.ErrorFix{
				{
					Description: "Use only spaces: Replace all tabs with 4 spaces",
				},
				{
					Description: "Use only tabs: Replace all spaces with tabs",
				},
				{
					Description: "Configure your editor to convert tabs to spaces on save",
				},
			},
		},
		ContextHelp: "Carrion requires consistent indentation. Pick either tabs or spaces and use them throughout your file.",
	},

	// Inconsistent indentation style (tabs vs spaces between lines)
	{
		Pattern: regexp.MustCompile(`inconsistent indentation: expected (\w+), got (\w+)`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Inconsistent indentation style",
			Description: "Your file started with one indentation style but switched to another.",
			Fixes: []object.ErrorFix{
				{
					Description: "Convert this line to match the file's established style",
				},
				{
					Description: "Use find-and-replace to convert all indentation in the file",
				},
				{
					Description: "Configure your editor's indentation settings for .crl files",
				},
			},
		},
		ContextHelp: "Once you indent a line with tabs or spaces, stick with that style for the entire file.",
	},

	// Import errors
	{
		Pattern: regexp.MustCompile(`module (.+) not found`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Module not found",
			Description: "The module you're trying to import doesn't exist or isn't in the import path.",
			Fixes: []object.ErrorFix{
				{
					Description: "Check the module name for typos",
				},
				{
					Description: "Ensure the module file exists in the correct location",
				},
				{
					Description: "Check that the module is in the Carrion standard library",
				},
			},
		},
		ContextHelp: "Modules must be available in the current directory or standard library path.",
	},
	
	// Attribute errors
	{
		Pattern: regexp.MustCompile(`(.+) has no attribute (.+)`),
		Suggestion: object.ErrorSuggestion{
			Title:       "Attribute not found",
			Description: "The object doesn't have the attribute or method you're trying to access.",
			Fixes: []object.ErrorFix{
				{
					Description: "Check the object type and available methods",
				},
				{
					Description: "Use type() to verify the object type",
				},
				{
					Description: "Check for typos in the attribute name",
				},
			},
		},
		ContextHelp: "Objects only have attributes and methods that are defined for their type.",
	},
	
	// Null/None errors
	{
		Pattern: regexp.MustCompile(`.*None.*has no.*`),
		Suggestion: object.ErrorSuggestion{
			Title:       "None value error",
			Description: "You're trying to use None as if it were a different type.",
			Fixes: []object.ErrorFix{
				{
					Description: "Check if the value is None before using it",
				},
				{
					Description: "Use conditional logic: if value is not None:",
				},
				{
					Description: "Provide default values for potentially None results",
				},
			},
		},
		ContextHelp: "None represents the absence of a value. Always check for None before using values that might be None.",
	},
}

// Language-specific suggestions for common patterns
var LanguageSpecificSuggestions = map[string]object.ErrorSuggestion{
	"spell_syntax": {
		Title:       "Function definition syntax",
		Description: "Functions in Carrion are defined using the 'spell' keyword.",
		Fixes: []object.ErrorFix{
			{
				Description: "Use 'spell function_name():' for function definitions",
			},
			{
				Description: "Example: spell greet(name): return f\"Hello, {name}!\"",
			},
		},
	},
	
	"grim_syntax": {
		Title:       "Class definition syntax",
		Description: "Classes in Carrion are defined using the 'grim' keyword.",
		Fixes: []object.ErrorFix{
			{
				Description: "Use 'grim ClassName:' for class definitions",
			},
			{
				Description: "Example: grim Person: spell init(name): self.name = name",
			},
		},
	},
	
	"array_operations": {
		Title:       "Array operations",
		Description: "Carrion arrays have special methods for manipulation.",
		Fixes: []object.ErrorFix{
			{
				Description: "Use array.append(item) to add elements",
			},
			{
				Description: "Use array[index] to access elements",
			},
			{
				Description: "Use len(array) to get array length",
			},
		},
	},
	
	"string_operations": {
		Title:       "String operations",
		Description: "Carrion strings support various manipulation methods.",
		Fixes: []object.ErrorFix{
			{
				Description: "Use string.upper() and string.lower() for case conversion",
			},
			{
				Description: "Use f\"Hello {name}\" for string interpolation",
			},
				{
				Description: "Use string.find(substring) to search within strings",
			},
		},
	},
	
	"control_flow": {
		Title:       "Control flow",
		Description: "Carrion uses 'otherwise' instead of 'elif' and specific loop keywords.",
		Fixes: []object.ErrorFix{
			{
				Description: "Use 'otherwise' instead of 'elif' in conditional statements",
			},
			{
				Description: "Use 'stop' instead of 'break' and 'skip' instead of 'continue'",
			},
			{
				Description: "Use 'match/case' for pattern matching",
			},
		},
	},
	
	"error_handling": {
		Title:       "Error handling",
		Description: "Carrion uses 'attempt/ensnare/resolve' for error handling.",
		Fixes: []object.ErrorFix{
			{
				Description: "Use 'attempt:' instead of 'try:'",
			},
			{
				Description: "Use 'ensnare ErrorType:' instead of 'except:'",
			},
			{
				Description: "Use 'resolve:' instead of 'finally:'",
			},
		},
	},
}

// EnhanceErrorWithSuggestions adds contextual suggestions to an error
func EnhanceErrorWithSuggestions(err *object.EnhancedError, sourceCode string) *object.EnhancedError {
	// Try to match against known error patterns
	for _, pattern := range ErrorPatterns {
		if pattern.Pattern.MatchString(err.Message) {
			// Add the suggestion
			err.AddSuggestion(
				pattern.Suggestion.Title,
				pattern.Suggestion.Description,
				pattern.Suggestion.Fixes...,
			)
			
			// Add context help as a note
			if pattern.ContextHelp != "" {
				err.AddNote(pattern.ContextHelp, object.ERROR_LEVEL_HELP, nil)
			}
			
			// Add language-specific suggestions based on context
			addLanguageSpecificSuggestions(err, sourceCode)
			
			return err
		}
	}
	
	// If no pattern matched, try to infer suggestions from context
	addInferredSuggestions(err, sourceCode)
	
	return err
}

// addLanguageSpecificSuggestions adds suggestions based on language context
func addLanguageSpecificSuggestions(err *object.EnhancedError, sourceCode string) {
	// Check for common language confusion patterns
	if strings.Contains(sourceCode, "def ") || strings.Contains(err.Message, "def") {
		suggestion := LanguageSpecificSuggestions["spell_syntax"]
		err.AddSuggestion(suggestion.Title, suggestion.Description, suggestion.Fixes...)
	}
	
	if strings.Contains(sourceCode, "class ") || strings.Contains(err.Message, "class") {
		suggestion := LanguageSpecificSuggestions["grim_syntax"]
		err.AddSuggestion(suggestion.Title, suggestion.Description, suggestion.Fixes...)
	}
	
	if strings.Contains(sourceCode, "elif ") || strings.Contains(err.Message, "elif") {
		suggestion := LanguageSpecificSuggestions["control_flow"]
		err.AddSuggestion(suggestion.Title, suggestion.Description, suggestion.Fixes...)
	}
	
	if strings.Contains(sourceCode, "try:") || strings.Contains(err.Message, "try") {
		suggestion := LanguageSpecificSuggestions["error_handling"]
		err.AddSuggestion(suggestion.Title, suggestion.Description, suggestion.Fixes...)
	}
}

// addInferredSuggestions adds suggestions based on error message analysis
func addInferredSuggestions(err *object.EnhancedError, sourceCode string) {
	message := strings.ToLower(err.Message)
	
	// Common typos and alternatives
	if strings.Contains(message, "unexpected") && strings.Contains(message, "token") {
		err.AddSuggestion(
			"Check syntax",
			"Review the syntax around the error location",
			object.ErrorFix{
				Description: "Look for missing punctuation, keywords, or typos",
			},
		)
	}
	
	// Memory and performance suggestions
	if strings.Contains(message, "recursion") || strings.Contains(message, "stack") {
		err.AddSuggestion(
			"Recursion limit",
			"You may have infinite recursion or very deep recursive calls",
			object.ErrorFix{
				Description: "Add a base case to stop recursion",
			},
			object.ErrorFix{
				Description: "Consider using iterative solutions instead of recursion",
			},
		)
	}
	
	// Generic programming advice
	if err.Category == object.ERROR_CATEGORY_RUNTIME {
		err.AddNote(
			"Runtime errors occur during program execution. Review the logic and data flow.",
			object.ERROR_LEVEL_HELP,
			nil,
		)
	}
}

// GetCodeSuggestions returns suggestions for improving code quality
func GetCodeSuggestions(errorType object.ErrorCategory, sourceCode string) []object.ErrorSuggestion {
	var suggestions []object.ErrorSuggestion
	
	switch errorType {
	case object.ERROR_CATEGORY_SYNTAX:
		suggestions = append(suggestions, object.ErrorSuggestion{
			Title:       "Syntax best practices",
			Description: "Follow Carrion's syntax conventions for better code",
			Fixes: []object.ErrorFix{
				{Description: "Use consistent indentation (4 spaces)"},
				{Description: "End control structures with colons (:)"},
				{Description: "Use meaningful variable names"},
			},
		})
		
	case object.ERROR_CATEGORY_TYPE:
		suggestions = append(suggestions, object.ErrorSuggestion{
			Title:       "Type handling",
			Description: "Use proper type checking and conversion",
			Fixes: []object.ErrorFix{
				{Description: "Use type() function to check types"},
				{Description: "Use conversion functions: int(), str(), float()"},
				{Description: "Handle None values explicitly"},
			},
		})
		
	case object.ERROR_CATEGORY_RUNTIME:
		suggestions = append(suggestions, object.ErrorSuggestion{
			Title:       "Runtime safety",
			Description: "Add defensive programming practices",
			Fixes: []object.ErrorFix{
				{Description: "Validate inputs before processing"},
				{Description: "Use error handling with attempt/ensnare blocks"},
				{Description: "Check bounds before accessing arrays/strings"},
			},
		})
	}
	
	return suggestions
}

// FormatSuggestionForDisplay formats a suggestion for console output
func FormatSuggestionForDisplay(suggestion object.ErrorSuggestion) string {
	var result strings.Builder
	
	result.WriteString(fmt.Sprintf("💡 %s\n", suggestion.Title))
	if suggestion.Description != "" {
		result.WriteString(fmt.Sprintf("   %s\n", suggestion.Description))
	}
	
	for i, fix := range suggestion.Fixes {
		result.WriteString(fmt.Sprintf("   %d. %s\n", i+1, fix.Description))
		if fix.Replacement != "" {
			result.WriteString(fmt.Sprintf("      → %s\n", fix.Replacement))
		}
	}
	
	return result.String()
}

// CreateContextualErrorMessage creates a more helpful error message with context
func CreateContextualErrorMessage(originalMessage string, context map[string]string) string {
	var enhanced strings.Builder
	
	enhanced.WriteString(originalMessage)
	
	// Add context information
	if len(context) > 0 {
		enhanced.WriteString("\n\nContext:")
		for key, value := range context {
			enhanced.WriteString(fmt.Sprintf("\n  %s: %s", key, value))
		}
	}
	
	return enhanced.String()
}

// GetErrorDocumentation returns documentation links for error types
func GetErrorDocumentation(errorCode string) string {
	docs := map[string]string{
		"SYNTAX_ERROR":    "https://carrion-lang.org/docs/syntax",
		"TYPE_ERROR":      "https://carrion-lang.org/docs/types",
		"RUNTIME_ERROR":   "https://carrion-lang.org/docs/runtime",
		"IMPORT_ERROR":    "https://carrion-lang.org/docs/modules",
		"SEMANTIC_ERROR":  "https://carrion-lang.org/docs/semantics",
	}
	
	if url, exists := docs[errorCode]; exists {
		return fmt.Sprintf("For more information, see: %s", url)
	}
	
	return "For more information, see: https://carrion-lang.org/docs/errors"
}