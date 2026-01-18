# Carrion Language Changelog

## Version {{VERSION}}

### Enhanced Error System
- **Intelligent "Did You Mean?" Suggestions**: The error system now suggests similar variable/method names when encountering unknown identifiers using Levenshtein distance matching
- **Improved Type Mismatch Messages**: More specific error messages for string/number, boolean, and type hint violations
- **Context-Aware Suggestions**: Error suggestions now include available methods, variables, and builtins from the current scope

### Data Structures
- **BTree `display_tree()` Method**: New visual tree representation showing hierarchical structure with branch characters
- **BTree `size()` Renamed to `get_size()`**: Consistent naming with other data structures (Stack, Queue, Heap)
- **BTree None Handling**: `insert()` now gracefully skips `None` values instead of failing

### HTTP Server Enhancements
- **Full Request Headers Access**: All incoming HTTP headers are now available in `request["headers"]`
- **Query Parameter Parsing**: Automatic URL query string parsing available in `request["query"]`
- **Request Body Support**: Full request body capture for POST/PUT/PATCH in `request["body"]`
- **Response Headers**: Set custom response headers via `http_response()` third parameter

### REPL Improvements
- **Cleaner Output**: Functions, grimoires, and builtins no longer display their definitions when referenced without calling
- **Assignment Silence**: Assignment statements no longer echo the assigned value
- **Better Error Handling**: Unified error formatting for all error types (EnhancedError, ErrorWithTrace, Error)
- **Non-Fatal Warnings**: History file and liner errors now log warnings instead of crashing

### Bug Fixes
- Fixed indentation handling issues
- Fixed `input()` function behavior
- Fixed instance wrapping for float values
- Fixed import logic for modules
- Fixed tuple handling
- Fixed REPL evaluator to only display results appropriately
- Fixed server socket issues

### Documentation
- Added HTTP Server Enhancement documentation
- Added Indentation documentation
- Updated Data Structures documentation with new BTree methods
- Updated Language Reference

---

## Version 0.1.8

### Testing Framework
- **Sindri Testing Framework**: Comprehensive testing and benchmarking tool for Carrion
- HTML report generation for test results
- ASCII art branding for test runner

### Concurrency
- **Converge/Diverge Functions**: Goroutine-based concurrency support
- Time module enhancements

### Documentation System
- **Mimir Documentation Tool**: Interactive documentation and help system
- Docstring support for functions and classes

### Type System
- **Type Hints**: Optional static type checking with type annotations
- Return type hints for functions

### Standard Library
- Socket module enhancements
- New array and string methods
- Enhanced OS operations

### Error Handling
- Improved error messages with source locations
- Stack trace support for runtime errors

---

## Version 0.1.7 and Earlier

See [GitHub Releases](https://github.com/javanhut/TheCarrionLanguage/releases) for earlier version history.
