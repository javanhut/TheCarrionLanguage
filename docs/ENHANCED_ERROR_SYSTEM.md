# Enhanced Error System for Carrion Language

## Overview

The Carrion language now features a comprehensive enhanced error system that provides detailed, contextual error messages with helpful suggestions and fix recommendations, similar to the Rust compiler's error reporting system.

## Key Features

### 1. **Detailed Error Context**
- **Source Code Spans**: Precise location information with start/end positions
- **Multiple Labels**: Point to different parts of code relevant to the error
- **Stack Traces**: Complete call stack information
- **Variable Context**: Current variable states and values

### 2. **Rich Error Types**
- **Error Levels**: Error, Warning, Note, Help
- **Error Categories**: Syntax, Type, Runtime, Semantic, Import, I/O, Custom
- **Error Codes**: Unique identifiers for each error type (e.g., "E0001")
- **Error Titles**: Human-readable error descriptions

### 3. **Intelligent Suggestions**
- **Fix Recommendations**: Concrete steps to resolve errors
- **Code Replacements**: Suggested code changes
- **Pattern Matching**: Automatic suggestions based on common error patterns
- **Language-Specific Help**: Carrion-specific syntax guidance

### 4. **Visual Error Formatting**
- **Syntax Highlighting**: Color-coded error output
- **Source Code Display**: Shows relevant code with error indicators
- **Caret Indicators**: Points to exact error locations
- **Multi-line Support**: Handles errors spanning multiple lines

## Error System Architecture

### Core Components

1. **`EnhancedError`** - The main error type with comprehensive context
2. **`ErrorSpan`** - Represents a span of code in the source
3. **`ErrorLabel`** - Points to specific parts of code with messages
4. **`ErrorSuggestion`** - Contains fix recommendations
5. **`ErrorFix`** - Specific fix instructions
6. **`ErrorNote`** - Additional context and help information

### Error Categories

- **`ERROR_CATEGORY_SYNTAX`** - Parsing and syntax errors
- **`ERROR_CATEGORY_TYPE`** - Type mismatch and type-related errors
- **`ERROR_CATEGORY_RUNTIME`** - Runtime execution errors
- **`ERROR_CATEGORY_SEMANTIC`** - Semantic analysis errors
- **`ERROR_CATEGORY_IMPORT`** - Module import errors
- **`ERROR_CATEGORY_IO`** - File and I/O errors
- **`ERROR_CATEGORY_CUSTOM`** - User-defined errors

### Error Levels

- **`ERROR_LEVEL_ERROR`** - Critical errors that prevent execution
- **`ERROR_LEVEL_WARNING`** - Warnings about potential issues
- **`ERROR_LEVEL_NOTE`** - Additional information
- **`ERROR_LEVEL_HELP`** - Helpful suggestions and guidance

## Usage Examples

### Basic Error Creation

```go
// Create a syntax error
span := object.ErrorSpan{
    Start: object.SourcePosition{
        Filename: "example.crl",
        Line:     1,
        Column:   10,
    },
    End: object.SourcePosition{
        Filename: "example.crl",
        Line:     1,
        Column:   11,
    },
    Source: "spell greet(name)",
}

err := object.NewSyntaxError("expected COLON, got NEWLINE", span).
    WithCode("E0001").
    WithTitle("Missing Colon")
```

### Adding Suggestions

```go
err.AddSuggestion(
    "Add missing colon",
    "Function definitions require a colon (:) at the end",
    object.ErrorFix{
        Description: "Add ':' after the function signature",
        Replacement: ":",
    },
)
```

### Adding Context Labels

```go
err.AddLabel(
    typeSpan,
    "this is a STRING",
    object.ERROR_LEVEL_NOTE,
)
```

### Adding Notes

```go
err.AddNote(
    "In Carrion, all function definitions must end with a colon (:)",
    object.ERROR_LEVEL_HELP,
    nil,
)
```

## Integration with Existing Code

### Error Context Management

```go
// Create an error context
context := object.NewErrorContext("example.crl", sourceCode)
context.SetPosition(line, column)

// Create enhanced errors with context
err := context.NewEnhancedRuntimeError("identifier not found: x")
```

### Error Propagation

```go
// Propagate errors up the call stack
enhanced := object.PropagateError(err, context)

// Wrap errors with additional context
wrapped := object.WrapError(err, "function call failed", context)
```

### Migration from Old Errors

```go
// Upgrade existing errors to enhanced errors
enhanced := object.UpgradeToEnhancedError(oldError, span)

// Add suggestions based on error patterns
enhanced = utils.EnhanceErrorWithSuggestions(enhanced, sourceCode)
```

## Error Formatting and Display

### Enhanced Error Printer

```go
// Print detailed error with full context
utils.PrintEnhancedError(enhancedError)

// Print simplified error
utils.PrintEnhancedErrorSimple(enhancedError)
```

### Example Error Output

```
error[syntax] E0001
Missing Colon: expected COLON, got NEWLINE
  --> example.crl:1:17
   |
 1 | spell greet(name)
   |                 ^ expected ':'
   |
help: Add missing colon
    Function definitions require a colon (:) at the end
    • Add ':' after the function signature

note: In Carrion, all function definitions must end with a colon (:)
```

## Common Error Patterns

### Undefined Variable Error

```go
err := context.CreateUndefinedVariableError("x")
// Automatically includes suggestions for:
// - Defining the variable
// - Checking for typos
// - Similar variable names
```

### Type Mismatch Error

```go
err := context.NewEnhancedTypeError(
    "cannot add STRING and INTEGER",
    "STRING",
    "INTEGER",
)
// Automatically includes suggestions for:
// - Type conversion functions
// - Compatible operations
```

### Index Out of Bounds Error

```go
err := context.NewEnhancedIndexError(index, length)
// Automatically includes suggestions for:
// - Bounds checking
// - Safe access methods
// - Array length validation
```

## Configuration Options

### Error Configuration

```go
config := &object.ErrorConfig{
    UseEnhancedErrors: true,
    ShowSuggestions:   true,
    ShowStackTrace:    true,
    VerboseMode:       false,
}
```

### Customization

- **Enable/Disable Enhanced Errors**: Fall back to basic errors if needed
- **Show/Hide Suggestions**: Control suggestion display
- **Stack Trace Control**: Enable/disable stack trace output
- **Verbose Mode**: Show additional debugging information

## Language-Specific Features

### Carrion Syntax Guidance

The error system provides specific guidance for Carrion language features:

- **Function Definitions**: Use `spell` instead of `def`
- **Class Definitions**: Use `grim` instead of `class`
- **Control Flow**: Use `otherwise` instead of `elif`
- **Loop Control**: Use `stop`/`skip` instead of `break`/`continue`
- **Error Handling**: Use `attempt`/`ensnare`/`resolve` blocks

### Common Mistakes Detection

- **Python-style syntax**: Detects and suggests Carrion equivalents
- **Missing colons**: Automatic detection and suggestions
- **Indentation errors**: Helpful guidance for proper indentation
- **Type confusion**: Clear explanations of type-related issues

## Performance Considerations

- **Lazy Evaluation**: Suggestions are generated only when needed
- **Memory Management**: Bounded error history to prevent memory leaks
- **Efficient Formatting**: Optimized string formatting for error display
- **Context Cleanup**: Automatic cleanup of error contexts

## Testing and Validation

### Error System Testing

```go
// Test error creation
func TestEnhancedErrorCreation(t *testing.T) {
    span := createTestSpan()
    err := object.NewSyntaxError("test error", span)
    assert.Equal(t, "test error", err.Message)
    assert.Equal(t, object.ERROR_CATEGORY_SYNTAX, err.Category)
}
```

### Suggestion Testing

```go
// Test suggestion generation
func TestErrorSuggestions(t *testing.T) {
    err := createTestError()
    enhanced := utils.EnhanceErrorWithSuggestions(err, sourceCode)
    assert.True(t, len(enhanced.Suggestions) > 0)
}
```

## Migration Guide

### From Basic Errors

1. **Replace `newError()` calls** with `NewEnhancedError()`
2. **Add position information** using `ErrorSpan`
3. **Include suggestions** using `AddSuggestion()`
4. **Update error printing** to use `PrintEnhancedError()`

### From Error With Trace

1. **Convert to EnhancedError** using `UpgradeToEnhancedError()`
2. **Add structured suggestions** instead of plain text
3. **Enhance with labels** for better context
4. **Update formatting** for better visual presentation

## Best Practices

### Error Creation

- **Be Specific**: Use precise error messages and spans
- **Add Context**: Include relevant variable states and call stack
- **Provide Suggestions**: Always include helpful fix recommendations
- **Use Appropriate Levels**: Choose the right error level (Error, Warning, Note, Help)

### Error Handling

- **Propagate Context**: Maintain error context through the call stack
- **Chain Related Errors**: Use cause chains for related errors
- **Collect Multiple Errors**: Group related errors together
- **Clean Up Resources**: Ensure proper cleanup of error contexts

### Performance

- **Minimize Allocations**: Reuse error objects when possible
- **Lazy Generation**: Generate expensive content only when needed
- **Bounded History**: Limit error history to prevent memory issues
- **Efficient Formatting**: Use optimized string operations

## Future Enhancements

### Planned Features

- **Interactive Error Fixing**: IDE integration for automatic fixes
- **Error Analytics**: Collect error patterns for improvement
- **Localization**: Multi-language error messages
- **Custom Error Types**: User-defined error categories
- **Advanced Suggestions**: AI-powered fix recommendations

### Extensibility

- **Plugin System**: Allow custom error handlers
- **Custom Formatters**: Support for different output formats
- **Integration APIs**: APIs for IDE and tool integration
- **Metrics Collection**: Performance and usage metrics

## Documentation and Resources

- **API Reference**: Complete API documentation
- **Examples**: Comprehensive usage examples
- **Migration Guide**: Step-by-step migration instructions
- **Best Practices**: Recommended patterns and practices
- **Performance Guide**: Optimization tips and tricks

## Contributing

To contribute to the enhanced error system:

1. **Follow the error creation patterns** established in the codebase
2. **Add appropriate test coverage** for new error types
3. **Update documentation** for new features
4. **Ensure performance** doesn't degrade with changes
5. **Maintain backward compatibility** where possible

The enhanced error system significantly improves the developer experience by providing clear, actionable error messages that help developers quickly identify and fix issues in their Carrion code.