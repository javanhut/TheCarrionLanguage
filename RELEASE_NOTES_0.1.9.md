# Carrion Language v0.1.9 Release Notes

**Release Date**: January 2026
**Version**: 0.1.9
**Previous Version**: 0.1.8

## Overview

Carrion Language v0.1.9 focuses on **developer experience improvements** with intelligent error suggestions, enhanced HTTP server capabilities for building web applications, improved REPL behavior, and data structure refinements. This release makes debugging easier, web development more powerful, and the interactive experience cleaner.

## 🎯 Intelligent Error Suggestions

### "Did You Mean?" System
The error system now provides intelligent suggestions when you mistype variable names, method calls, or function names.

### Key Features
- **Levenshtein Distance Matching**: Finds similar names based on edit distance
- **Context-Aware Suggestions**: Suggests from available methods, variables, and builtins in scope
- **Type-Specific Guidance**: Different suggestions for string/number mismatches, boolean operations, and type hint violations

### Example
```carrion
name = "Carrion"
print(nme)  # Typo!

# Error output now includes:
# error: identifier not found: nme
#   --> example.crl:2:7
#    |
#  2 | print(nme)
#    |       ^^^ unknown identifier
#    |
# help: Did you mean 'name'?
```

### Enhanced Type Mismatch Messages
```carrion
x = "42" + 10  # String + Integer

# Now shows specific guidance:
# error: String/Number type mismatch
#   You're mixing strings and numbers in an operation.
#
# help:
#   • Convert string to number: int("123") or float("3.14")
#   • Convert number to string: str(123) or str(3.14)
#   • Use string concatenation (+) only with strings
```

## 🌐 HTTP Server Enhancements

### Full-Featured Web Applications
Build complete REST APIs and web applications with comprehensive request/response handling.

### New Request Object Fields
```carrion
spell handler(request):
    # All fields now available:
    method = request["method"]      # GET, POST, PUT, DELETE, etc.
    path = request["path"]          # /api/users
    headers = request["headers"]    # All HTTP headers
    query = request["query"]        # Query string parameters
    body = request["body"]          # Request body (POST/PUT)

    return http_response(200, "OK", {"Content-Type": "text/plain"})
```

### Request Headers Access
```carrion
spell authenticated_handler(request):
    if "Authorization" not in request["headers"]:
        return http_response(401, "Unauthorized", {})

    auth = request["headers"]["Authorization"]
    # Validate token...

    return http_response(200, "Success", {
        "Content-Type": "application/json",
        "X-Request-ID": generate_id()
    })
```

### Query Parameter Parsing
```carrion
spell list_items(request):
    # URL: /items?page=1&limit=10&filter=active
    page = request["query"]["page"]      # "1"
    limit = request["query"]["limit"]    # "10"
    filter = request["query"]["filter"]  # "active"

    # Process with pagination...
    return http_response(200, items_json, {"Content-Type": "application/json"})
```

### Request Body Support
```carrion
spell create_user(request):
    if request["method"] != "POST":
        return http_response(405, "Method Not Allowed", {})

    body = request["body"]  # Raw JSON string
    # Parse and process...

    return http_response(201, "{\"id\": 123}", {
        "Content-Type": "application/json"
    })
```

## 🖥️ REPL Improvements

### Cleaner Output
The REPL no longer displays function definitions when you reference them without calling:

```carrion
# Before (0.1.8):
>>> print
<builtin function: print>

# After (0.1.9):
>>> print
>>>              # Clean, no output for function references

>>> print("Hello")
Hello            # Only actual results are shown
```

### Assignment Silence
Assignment statements no longer echo values:

```carrion
# Before (0.1.8):
>>> x = 42
42

# After (0.1.9):
>>> x = 42
>>>              # Silent assignment

>>> x            # Explicit reference shows value
42
```

### Unified Error Handling
All error types (EnhancedError, ErrorWithTrace, Error) now use consistent formatting with the new suggestion system.

### Non-Fatal Warnings
History file errors and liner issues now log warnings instead of crashing the REPL.

## 🌳 Data Structure Updates

### BTree Enhancements

#### New `display_tree()` Method
Visualize your binary search tree structure:

```carrion
bst = BTree()
for value in [10, 5, 15, 3, 7, 12, 18]:
    bst.insert(value)

bst.display_tree()
# Output:
# 10
# |-- 5
# |   |-- 3
# |   |-- 7
# |-- 15
#     |-- 12
#     |-- 18
```

#### Renamed `size()` to `get_size()`
Consistent naming with Stack, Queue, and Heap:

```carrion
# Before (0.1.8):
count = bst.size()

# After (0.1.9):
count = bst.get_size()
```

#### None Value Handling
`insert()` now gracefully skips `None` values:

```carrion
bst = BTree()
bst.insert(10)
bst.insert(None)  # Silently skipped
bst.insert(5)
print(bst.get_size())  # 2
```

## 🐛 Bug Fixes

### Core Language
- **Indentation Handling**: Fixed issues with nested block indentation
- **Input Function**: Fixed `input()` behavior in various contexts
- **Float Instance Wrapping**: Fixed instance wrapping for float values
- **Import Logic**: Fixed module import resolution
- **Tuple Handling**: Fixed tuple operations and unpacking

### REPL
- **Evaluator Output**: Fixed to only display appropriate results
- **Function Display**: Functions, grimoires, and builtins no longer print definitions on reference

### Networking
- **Server Sockets**: Fixed issues with HTTP server operations
- **Response Headers**: Headers now properly set before status code (HTTP spec compliance)

## 📖 Documentation Updates

### New Documentation
- **[HTTP Server Enhancement](docs/HTTP-Server-Enhancement.md)**: Complete guide to building web applications
- **[Indentation](docs/Indentation.md)**: Indentation rules and best practices
- **[Changelog](docs/Changelog/README.md)**: Version history tracking

### Updated Documentation
- **[Data Structures](docs/Data-Structures.md)**: Updated with `get_size()` and `display_tree()`
- **[README](docs/README.md)**: Added references to new documentation
- **[Language Reference](docs/Language-Reference.md)**: Updated for 0.1.9 changes

## 🔄 Migration Guide

### From 0.1.8 to 0.1.9

#### BTree Method Rename
```carrion
# Update your code:
# Old:
count = my_tree.size()

# New:
count = my_tree.get_size()
```

#### HTTP Handler Updates (Optional Enhancement)
```carrion
# You can now access more request data:
spell handler(request):
    # New fields available (backward compatible):
    headers = request["headers"]  # New!
    query = request["query"]      # New!
    body = request["body"]        # New!

    # Existing fields still work:
    method = request["method"]
    path = request["path"]
```

### Backward Compatibility
- All existing code works unchanged (except `BTree.size()` → `BTree.get_size()`)
- New HTTP request fields are additive
- Error messages are enhanced but don't change behavior

## 🚀 Getting Started with v0.1.9

### For Existing Users
1. **Update Installation**: Run `make install` to get the latest version
2. **Update BTree Code**: Change `size()` calls to `get_size()`
3. **Enjoy Better Errors**: Typos now get helpful suggestions
4. **Build Web Apps**: Use enhanced HTTP features for REST APIs

### For New Users
1. **Install Carrion**: Follow standard installation process
2. **Explore with Mimir**: Run `mimir` to explore language features
3. **Try HTTP Server**: Check `examples/http_rest_api_demo.crl`
4. **Write Tests**: Use Sindri for testing your code

## 🎯 What's Next

### Planned for Future Releases
- **List Comprehensions**: Python-like collection processing
- **Async/Await**: Modern concurrency patterns
- **JSON Module**: Built-in JSON parsing and generation
- **WebSocket Support**: Real-time communication
- **Enhanced IDE Integration**: Language server improvements

## 📋 Summary

Carrion Language v0.1.9 enhances the developer experience with:

- 🎯 **Intelligent Error Suggestions**: "Did you mean?" for typos and mismatches
- 🌐 **Full HTTP Server**: Headers, query params, body, and response headers
- 🖥️ **Cleaner REPL**: No noise from function references or assignments
- 🌳 **Better Data Structures**: Visual tree display and consistent naming
- 🐛 **Bug Fixes**: Indentation, input, floats, imports, tuples, and more
- 📖 **Updated Documentation**: New guides and comprehensive changelog

This release maintains backward compatibility (with one minor rename) while significantly improving the development and debugging experience.

**Download Carrion v0.1.9 and enjoy smarter error messages and powerful web development!**

---

**Contributors**: Carrion Language Team
**Documentation**: Complete guides available in `docs/`
**Support**: [GitHub Issues](https://github.com/javanhut/TheCarrionLanguage/issues)
**License**: Same as Carrion Language
