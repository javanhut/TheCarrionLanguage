# Carrion Language Changelog

## Version 0.1.9 - Developer Experience Improvements

### Release Date: January 2026

### Overview
Carrion Language v0.1.9 focuses on **developer experience improvements** with intelligent error suggestions, enhanced HTTP server capabilities for building web applications, improved REPL behavior, and data structure refinements.

### Major Features

#### Intelligent Error Suggestions
- **"Did You Mean?" System**: Provides suggestions when you mistype variable names, method calls, or function names
- **Levenshtein Distance Matching**: Finds similar names based on edit distance
- **Context-Aware Suggestions**: Suggests from available methods, variables, and builtins in scope
- **Type-Specific Guidance**: Different suggestions for string/number mismatches, boolean operations, and type hint violations

#### HTTP Server Enhancements
- **Full Request Object Fields**: Access method, path, headers, query parameters, and body
- **Request Headers Access**: All incoming HTTP headers are captured and accessible
- **Query Parameter Parsing**: Automatic URL query string parsing
- **Request Body Support**: Full request body capture for POST/PUT/PATCH requests
- **Response Headers**: Set custom headers via `http_response()` third parameter

#### Cleaner REPL Experience
- **Silent Assignments**: Assignment statements no longer echo values
- **No Function Reference Noise**: Function/grimoire references without calls don't print
- **Unified Error Handling**: All error types use consistent formatting with suggestions
- **Non-Fatal Warnings**: History file errors log warnings instead of crashing

#### BTree Enhancements
- **New `display_tree()` Method**: Visualize binary search tree structure
- **Renamed `size()` to `get_size()`**: Consistent naming with Stack, Queue, and Heap
- **None Value Handling**: `insert()` now gracefully skips `None` values

### Bug Fixes
- Fixed indentation handling for nested blocks
- Fixed `input()` function behavior in various contexts
- Fixed float instance wrapping
- Fixed module import resolution
- Fixed tuple operations and unpacking
- Fixed HTTP server socket operations
- Fixed response headers to be set before status code (HTTP spec compliance)

### Documentation Updates
- **HTTP Server Enhancement Guide**: Complete guide to building web applications
- **Indentation Guide**: Indentation rules and best practices
- **Enhanced Error System**: Understanding the new suggestion system
- **Updated Data Structures**: Documentation for `get_size()` and `display_tree()`

### Migration Guide
```python
# BTree method rename:
# Old:
count = my_tree.size()

# New:
count = my_tree.get_size()
```

See [RELEASE_NOTES_0.1.9.md](../RELEASE_NOTES_0.1.9.md) for complete details.

---

## Version 0.1.8 - Complete Tooling Ecosystem

### Release Date: July 20, 2025

### Overview
Carrion Language v0.1.8 introduces a comprehensive tooling ecosystem that transforms the development experience with three major new tools: **Sindri Testing Framework**, **Mimir Documentation System**, and **Bifrost Package Manager**. This release focuses on developer productivity, testing capabilities, and documentation accessibility while maintaining full backward compatibility.

### Major Features

#### Sindri Testing Framework
- **Automatic Test Discovery**: Finds test functions using the "appraise" naming convention
- **Colored Terminal Output**: Green for passed tests, red for failures
- **Flexible Assertions**: Support for both boolean and value comparison assertions
- **Multiple Output Modes**: Summary and detailed reporting modes
- **Directory Testing**: Run all tests in a directory or specific files
- **Built-in `check()` function**: Core assertion mechanism for all tests

#### Mimir Documentation System
- **Interactive Documentation Browser**: Menu-driven exploration of language features
- **Command-Line Lookup**: Quick help for specific functions (`mimir scry <function>`)
- **Comprehensive Coverage**: Built-in functions, standard library, language features
- **Search Functionality**: Find functions by name or purpose
- **REPL Integration**: Seamless integration with the Carrion REPL
- **Category Browsing**: Organized documentation by topic

#### Bifrost Package Manager
- **Git Submodule Integration**: Integrated directly into the Carrion repository
- **Package Management**: Install, manage, and distribute Carrion packages
- **Build System Updates**: Enhanced Makefile and installation scripts
- **Future-Ready**: Foundation for growing Carrion ecosystem

### Build System Improvements
- New Makefile targets: `make sindri`, `make mimir`, `make install`
- Updated installation scripts with Sindri support
- Enhanced setup with complete development environment
- Clean uninstall support for all tools

### Documentation Updates
- **Sindri.md**: Complete testing framework guide with examples
- **Mimir.md**: Documentation system reference and usage
- **Updated Module Documentation**: Enhanced for new tooling ecosystem
- **README Updates**: Tool ecosystem overview
- **Installation Guides**: Updated for new tool installation

### Backward Compatibility
- **Zero Breaking Changes**: All existing Carrion code continues to work unchanged
- **Optional Tooling**: New tools enhance but don't replace existing workflows
- **Existing APIs**: All functions and modules remain the same

### API Improvements & Breaking Changes

#### Static Method Support for Grimoires
- **Implemented static method calls on grimoire classes**
  - Added support for `Grimoire.method()` syntax
  - Created `StaticMethod` object type for handling static calls
  - **Location**: `src/evaluator/evaluator.go`, `src/object/static_method.go`

#### File & OS API Unification (Breaking Change)
- **Refactored file and OS operations to use consistent grimoire API**

**New File grimoire static methods**:
- `File.read(path)` - Read entire file content
- `File.write(path, content)` - Write content to file (overwrites)
- `File.append(path, content)` - Append content to file
- `File.exists(path)` - Check if file exists
- `File.open(path, mode)` - Create File object for complex operations

**New OS grimoire static methods**:
- `OS.cwd()` - Get current working directory
- `OS.chdir(path)` - Change directory
- `OS.listdir(path)` - List directory contents
- `OS.getenv(key)` - Get environment variable
- `OS.setenv(key, value)` - Set environment variable
- `OS.remove(path)` - Remove file/directory
- `OS.mkdir(path, perm)` - Create directory
- `OS.run(cmd, args, capture)` - Execute system commands
- `OS.sleep(seconds)` - Sleep for specified time
- `OS.expandEnv(string)` - Expand environment variables

### Critical Bug Fixes

#### String Concatenation Type Fix
- **Fixed critical bug where string concatenation operations returned BUILTIN type objects instead of proper String instances**
  - Previously, long string concatenations or concatenations involving triple-quoted strings could result in incorrect object types
  - Socket operations and other modules expecting string types would fail with "data must be a string" errors
  - Now all string concatenation operations return properly wrapped String instances with method access
  - **Location**: `src/evaluator/evaluator.go`

#### Multi-Level Inheritance Fix
- **Fixed critical bug where `super.init()` caused infinite recursion in 3+ level inheritance chains**
  - Previously, Level2's `super.init()` would call itself instead of Level1's init
  - Now correctly resolves to immediate parent class at each level
  - Supports inheritance hierarchies of any depth
  - Added `MethodGrimoire` field to `CallContext` to track which class owns the current method
  - **Location**: `src/evaluator/evaluator.go`

#### Fixed Variable Resolution Order
- **Updated `evalIdentifier` to check environment variables before builtin functions**
  - Prevents variable name conflicts with builtin function names
  - Ensures user-defined variables take precedence over system functions
  - **Location**: `src/evaluator/evaluator.go:2766-2773`

### Migration Guide

#### Updating File/OS Operations
```python
# Old API (deprecated)
content = fileRead("data.txt")
fileWrite("output.txt", "hello")
current_dir = osGetCwd()

# New API (v0.1.8+)
content = File.read("data.txt")
File.write("output.txt", "hello")
current_dir = OS.cwd()
```

### Enhanced Import System

#### Selective Imports
- **Added support for selective grimoire and spell imports**
  - Import specific grimoires: `import "module.GrimoireName"`
  - Import specific spells: `import "module.spell_name"`
  - Enhanced parser to validate identifiers in import paths
  - Supports both uppercase (grimoires) and lowercase (spells) selective imports
  - **Location**: `src/parser/parser.go`, `src/evaluator/evaluator.go`

### REPL Improvements

#### Cleaner Output Display
- **REPL no longer displays output for assignments and definitions**
  - Assignment statements don't print values
  - Function/grimoire definitions don't clutter output
  - Only expression results and function call returns are displayed
  - Function, grimoire, and builtin identifiers without calls are not printed
  - **Location**: `src/repl/repl.go`

### Tuple Handling Fix

#### Removed Automatic Tuple Unpacking
- **Fixed tuple handling in function calls**
  - Removed automatic single-tuple argument unpacking
  - More predictable tuple behavior in function parameters
  - Prevents unexpected function signature mismatches
  - **Location**: `src/evaluator/evaluator.go`

### Files Changed
- Core evaluator: Added `check()` builtin function, fixed inheritance, fixed string concatenation, tuple handling
- REPL: Enhanced output display logic, cleaner interactive experience
- Parser: Improved import statement parsing for selective imports
- Build system: Updated Makefile, setup.sh, install scripts
- New tools: cmd/sindri/, cmd/mimir/, bifrost/ submodule
- Module system: src/modules/file.go, src/modules/os.go refactored
- Documentation: Added Sindri.md, Mimir.md, Modules.md enhanced, Grimoires.md enhanced, updated guides

See [RELEASE_NOTES_0.1.8.md](RELEASE_NOTES_0.1.8.md) for complete details.

---

## Version 0.1.6 - String Indexing & Standard Library Enhancement

### 🎉 Major Features

#### String Indexing Support
- **Implemented string indexing for primitive strings**
  - Supports positive indices: `s[0]`, `s[1]`, `s[6]`
  - Supports negative indices: `s[-1]`, `s[-2]` (Python-style)
  - Proper bounds checking with descriptive error messages
  - Returns single-character strings
  - **Location**: `src/evaluator/evaluator.go:1498-1616`

#### String Grimoire (Standard Library)
- **Created comprehensive String grimoire** in Munin standard library
  - `String(value)` - constructor for string objects
  - `length()` - get string length
  - `upper()` - convert to uppercase
  - `lower()` - convert to lowercase  
  - `reverse()` - reverse string order
  - `find(substring)` - find substring position (returns -1 if not found)
  - `contains(substring)` - check if string contains substring
  - `char_at(index)` - safe character access with bounds checking
  - **Location**: `src/munin/string.crl`

### 🔧 New Builtin Functions

#### Character/ASCII Functions
- **`ord(char)`** - Convert single character to ASCII code
  - Example: `ord("A")` returns `65`
  - **Location**: `src/evaluator/builtins.go:767-780`

- **`chr(code)`** - Convert ASCII code to character  
  - Example: `chr(65)` returns `"A"`
  - Supports range 0-255
  - **Location**: `src/evaluator/builtins.go:782-795`

### 🗂️ Project Organization

#### File Structure Improvements
- **Moved all test files to proper locations**
  - All `test_*.crl` files → `src/examples/` (66 files)
  - All debug files → `debug/` directory
  - Cleaned up root directory structure

#### Code Quality
- **Enhanced error handling** for string operations
- **Improved bounds checking** with clear error messages
- **Maintained backward compatibility** with existing code

### 🧪 Testing & Verification

#### Functionality Verified
- ✅ String indexing with positive indices
- ✅ String indexing with negative indices  
- ✅ Bounds checking and error handling
- ✅ Integration with existing recursion system
- ✅ String grimoire instantiation and basic operations
- ✅ New builtin functions (`ord`, `chr`)
- ✅ Existing functionality preserved

#### Example Usage
```carrion
// String indexing
s = "hello world"
print(s[0])    // "h"
print(s[6])    // "w" 
print(s[-1])   // "d"
print(s[-2])   // "l"

// String grimoire
sg = String("Hello World")
print(sg.length())           // 11
print(sg.upper())           // "HELLO WORLD"
print(sg.find("World"))     // 6
print(sg.contains("Hello")) // True

// New builtins
print(ord("A"))  // 65
print(chr(65))   // "A"
```

### 🔄 Recursive String Operations
- **Enhanced recursive function support** with string indexing
- **Example**: Recursive string reversal now possible
```carrion
spell reverse_string(s, index):
    if index < 0:
        return ""
    return s[index] + reverse_string(s, index - 1)

spell reverse(s):
    return reverse_string(s, len(s) - 1)

print(reverse("Carrion"))  // "noirraC"
```

### 🏗️ Technical Implementation

#### Core Changes
- **Modified `evalIndexExpression`** to handle `STRING_OBJ` type
- **Added `evalStringIndexExpression`** for string-specific indexing logic
- **Enhanced builtin function registry** with character conversion functions
- **Implemented String grimoire** following Carrion's grimoire patterns

#### Error Handling
- **Descriptive error messages** for out-of-bounds access
- **Type safety** for index operations (must be INTEGER)
- **Graceful handling** of negative indices
- **Clear stack traces** for debugging

### 📝 Notes
- **Backward Compatible**: All existing code continues to work
- **Performance**: String indexing is O(1) operation
- **Memory Safe**: Proper bounds checking prevents crashes
- **Consistent**: Follows Python-style negative indexing conventions

---

**Contributors**: Claude Code Assistant  
**Date**: June 7, 2025  
**Commit Range**: Latest development commits  
**Files Changed**: 2 core files, 1 new grimoire, 66+ files reorganized