# Carrion Language Changelog

## Version 0.1.8 - String Concatenation Bug Fix

### Critical Bug Fix

#### Fixed String Concatenation Type Issues
- **Fixed critical bug where string concatenation operations returned BUILTIN type objects instead of proper String instances**
  - Previously, long string concatenations or concatenations involving triple-quoted strings could result in incorrect object types
  - Socket operations and other modules expecting string types would fail with "data must be a string" errors
  - Now all string concatenation operations return properly wrapped String instances with method access
  - **Location**: `src/evaluator/evaluator.go`

### Technical Implementation

#### Enhanced Infix Expression Evaluation
- **Updated `evalStringInfixExpression` to use `wrapPrimitive()` for string concatenation results**
  - Ensures concatenated strings are properly wrapped String instances
  - Maintains access to String grimoire methods like `.upper()`, `.lower()`, `.length()`
  - **Location**: `src/evaluator/evaluator.go:3224`

#### Fixed Variable Resolution Order
- **Updated `evalIdentifier` to check environment variables before builtin functions**
  - Prevents variable name conflicts with builtin function names
  - Ensures user-defined variables take precedence over system functions
  - **Location**: `src/evaluator/evaluator.go:2766-2773`

#### Integer and Float Operations
- **Applied same fix to `evalIntegerInfixExpression` and `evalFloatInfixExpression`**
  - All arithmetic operations now return properly wrapped instances
  - Maintains consistency across all primitive type operations
  - **Locations**: Multiple in `src/evaluator/evaluator.go`

### Verification & Testing

#### Issues Resolved
- **HTTP server socket operations** - `socket_send` now works correctly with concatenated strings
- **Triple-quoted string concatenation** - Works properly with regular string concatenation
- **Long string concatenations** - No longer result in BUILTIN type objects
- **String method access** - Concatenated strings maintain access to String methods

#### Example of Fixed Behavior
```carrion
# Previously failed with "socket_send: data must be a string"
html_content = """<!DOCTYPE html>
<html><body><h1>Hello World</h1></body></html>"""

headers = "HTTP/1.1 200 OK\r\n\r\n"
response = headers + html_content  # Now returns proper String instance

# Works correctly with socket operations
socket_send(client_id, response)

# String methods now accessible on concatenated results
print(response.length())  # Works correctly
```

### Documentation Updates

#### Updated Type System Documentation
- **Added String Concatenation and Type Consistency section** to `docs/Type-System.md`
- **Documented concatenation behavior** with examples
- **Explained previous issues and fixes** for reference

### Performance Impact

- **No performance regression** - Fix maintains existing operation speed
- **Improved type safety** - Prevents runtime type errors
- **Memory consistent** - Proper object lifecycle management

### Migration Notes

- **No breaking changes** - This is a bug fix that makes existing code work correctly
- **Socket operations** - Code that failed due to string type issues will now work
- **String methods** - Concatenated strings now have proper method access

## Version 0.1.8 - Multi-Level Inheritance Fix & File/OS Grimoire Refactoring

### Major Bug Fix

#### Fixed Infinite Recursion in Multi-Level Inheritance
- **Fixed critical bug where `super.init()` caused infinite recursion in 3+ level inheritance chains**
  - Previously, Level2's `super.init()` would call itself instead of Level1's init
  - Now correctly resolves to immediate parent class at each level
  - Supports inheritance hierarchies of any depth
  - **Location**: `src/evaluator/evaluator.go`

### Technical Implementation

#### Method Grimoire Tracking
- **Added `MethodGrimoire` field to `CallContext`** to track which class owns the current method
  - Enables proper super resolution based on method's declaring class
  - Ensures context inheritance through evaluation chains
  - **Location**: `src/evaluator/evaluator.go:40`

#### Fixed Context Inheritance
- **Updated all `CallContext` creations to inherit `MethodGrimoire` from parent context**
  - Ensures method context is preserved through nested evaluations
  - Critical for proper super resolution in deep call stacks
  - **Locations**: Multiple in `src/evaluator/evaluator.go`

#### Enhanced Method Call Handling
- **Updated method call evaluators to set proper grimoire context**:
  - `evalBoundMethodCall` - uses `findMethodOwner()` helper
  - `evalGrimoireMethodCall` - sets method's owning grimoire
  - `evalStaticMethodCall` - sets static method's grimoire
  - Init method calls - sets constructor's grimoire

### Verification & Testing

#### Test Cases Passing
- **2-level inheritance** - Still works correctly (no regression)
- **3-level inheritance** - Now works correctly (was broken)
- **4-level inheritance** - Works correctly
- **5-level inheritance** - Works correctly
- **Multiple independent inheritance chains** - All work correctly

#### Example of Fixed Behavior
```python
grim Level1:
    init(x):
        self.l1 = x

grim Level2(Level1):
    init(x):
        super.init(x)  # Now correctly calls Level1.init()
        self.l2 = x

grim Level3(Level2):
    init(x):
        super.init(x)  # Correctly calls Level2.init()
        self.l3 = x

obj = Level3("test")
# Before fix: Stack overflow due to infinite recursion
# After fix: Works perfectly, all constructors called in order
```

### Documentation Updates

#### Updated Grimoires Documentation
- **Added Multi-Level Inheritance section** with comprehensive examples
- **Added Deep Inheritance Hierarchies examples** showing 5+ levels
- **Enhanced Best Practices section** with inheritance guidelines:
  - Always call `super.init()` in constructors
  - Limit inheritance depth for maintainability
  - Use abstract classes for contracts
  - Prefer composition for complex behaviors
  - Override methods consistently
  - Document inheritance hierarchies

### Migration Notes

- **No breaking changes** - This is a bug fix that makes existing code work correctly
- **No migration needed** - Code that worked around the bug will continue to work
- **Improved stability** - Deep inheritance hierarchies now work as expected

### Performance Impact

- **No performance regression** - Fix adds minimal overhead
- **Prevents stack overflow** - Eliminates infinite recursion
- **Memory safe** - Proper cleanup of method contexts

### Major Features

#### Static Method Support for Grimoires
- **Implemented static method calls on grimoire classes**
  - Added support for `Grimoire.method()` syntax
  - Created `StaticMethod` object type for handling static calls
  - Enhanced evaluator to support grimoire static methods
  - **Location**: `src/evaluator/evaluator.go`, `src/object/static_method.go`

#### File & OS API Unification
- **Refactored file and OS operations to use consistent grimoire API**
  - All file operations now use `File.method()` syntax
  - All OS operations now use `OS.method()` syntax
  - Moved builtin functions to dedicated modules
  - **Locations**: `src/modules/file.go`, `src/modules/os.go`

### API Changes

#### File Operations (Breaking Change)
- **New File grimoire static methods**:
  - `File.read(path)` - Read entire file content
  - `File.write(path, content)` - Write content to file (overwrites)
  - `File.append(path, content)` - Append content to file
  - `File.exists(path)` - Check if file exists
  - `File.open(path, mode)` - Create File object for complex operations

#### OS Operations (Breaking Change)
- **New OS grimoire static methods**:
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

### Architecture Improvements

#### Module System
- **Created dedicated modules for system operations**
  - `src/modules/file.go` - File operation implementations
  - `src/modules/os.go` - OS operation implementations
  - Enhanced module loading in `src/evaluator/stdlib.go`

#### Builtin Function Cleanup
- **Removed system-level functions from core builtins**
  - Removed 14 file and OS functions from `src/evaluator/builtins.go`
  - Kept only core language functions as builtins
  - Improved separation of concerns

#### Argument Handling
- **Enhanced argument processing for wrapped primitives**
  - Added helper functions to handle both direct and instance-wrapped arguments
  - Improved compatibility with automatic primitive wrapping
  - **Location**: `src/modules/file.go`, `src/modules/os.go`

### Documentation Updates

#### Updated Documentation
- **Standard Library documentation** - Reflect new File and OS APIs
- **Builtin Functions documentation** - Remove deprecated functions, add grimoire methods
- **Version numbers** - Updated to 0.1.8 throughout documentation

### Migration Guide

#### Updating Existing Code
```python
# Old API (deprecated)
content = fileRead("data.txt")
fileWrite("output.txt", "hello")
current_dir = osGetCwd()

# New API (recommended)
content = File.read("data.txt")
File.write("output.txt", "hello")
current_dir = OS.cwd()
```

### Backward Compatibility
- **File object operations** remain unchanged (`file.read_content()`, `file.write_content()`)
- **Autoclose statement** works with both `open()` and `File.open()`
- **Munin standard library** maintains existing grimoire APIs

## Version 0.1.6 - String Indexing & Standard Library Enhancement

### Major Features

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

### New Builtin Functions

#### Character/ASCII Functions
- **`ord(char)`** - Convert single character to ASCII code
  - Example: `ord("A")` returns `65`
  - **Location**: `src/evaluator/builtins.go:767-780`

- **`chr(code)`** - Convert ASCII code to character  
  - Example: `chr(65)` returns `"A"`
  - Supports range 0-255
  - **Location**: `src/evaluator/builtins.go:782-795`

### Project Organization

#### File Structure Improvements
- **Moved all test files to proper locations**
  - All `test_*.crl` files → `src/examples/` (66 files)
  - All debug files → `debug/` directory
  - Cleaned up root directory structure

#### Code Quality
- **Enhanced error handling** for string operations
- **Improved bounds checking** with clear error messages
- **Maintained backward compatibility** with existing code

### Testing & Verification

#### Functionality Verified
- String indexing with positive indices
- String indexing with negative indices  
- Bounds checking and error handling
- Integration with existing recursion system
- String grimoire instantiation and basic operations
- New builtin functions (`ord`, `chr`)
- Existing functionality preserved

#### Example Usage
```carrion
# String indexing
s = "hello world"
print(s[0])    # "h"
print(s[6])    # "w" 
print(s[-1])   # "d"
print(s[-2])   # "l"

# String grimoire
sg = String("Hello World")
print(sg.length())           # 11
print(sg.upper())           # "HELLO WORLD"
print(sg.find("World"))     # 6
print(sg.contains("Hello")) # True

# New builtins
print(ord("A"))  # 65
print(chr(65))   # "A"
```

### Recursive String Operations
- **Enhanced recursive function support** with string indexing
- **Example**: Recursive string reversal now possible
```carrion
spell reverse_string(s, index):
    if index < 0:
        return ""
    return s[index] + reverse_string(s, index - 1)

spell reverse(s):
    return reverse_string(s, len(s) - 1)

print(reverse("Carrion"))  # "noirraC"
```

### Technical Implementation

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

### Notes
- **Backward Compatible**: All existing code continues to work
- **Performance**: String indexing is O(1) operation
- **Memory Safe**: Proper bounds checking prevents crashes
- **Consistent**: Follows Python-style negative indexing conventions

---

**Contributors**: Claude Code Assistant  
**Date**: June 7, 2025 - July 15, 2025  
**Major Changes**: Multi-level inheritance fix, File & OS Grimoire refactoring, String indexing  
**Files Changed**: `src/evaluator/evaluator.go`, `src/object/object.go`, `docs/Grimoires.md`, `src/modules/file.go`, `src/modules/os.go`, `src/munin/string.crl`, and more
