# Sindri Testing Framework

Sindri is the built-in testing framework for Carrion language. It provides a simple and efficient way to write and run tests for your Carrion code.

## Overview

Sindri allows you to write test functions using the `appraise` naming convention and run them using the `sindri appraise` command. The framework automatically discovers test functions and executes them, providing colored output and comprehensive test summaries.

## Installation

Sindri is built alongside the main Carrion installation. To build the Sindri testing tool:

```bash
make sindri
```

This creates the `sindri` executable in the `cmd/sindri/` directory.

## Usage

### Basic Commands

```bash
sindri appraise <test_file.crl>           # Run specific test file
sindri appraise [directory]               # Run all appraise files in directory
sindri appraise                           # Run all appraise files in current directory
sindri appraise -d <path>                 # Run with detailed output
sindri appraise --detailed <path>         # Run with detailed output
```

### Examples

```bash
sindri appraise examples/sindri_demo.crl  # Run specific file
sindri appraise ./tests                   # Run all test files in tests directory
sindri appraise -d                        # Run all tests with detailed output
sindri appraise --detailed ./tests        # Run tests in directory with details
```

## Writing Tests

### Test Function Naming

Test functions must contain the word "appraise" in their name. Common patterns include:

- `appraise_<feature_name>()`
- `test_appraise_<feature_name>()`
- `<feature_name>_appraise()`

### Test Structure

Tests are written as regular Carrion spells (functions) that use the built-in `check()` function to make assertions:

```carrion
spell appraise_arithmetic():
    # Test basic arithmetic operations
    check(2 + 2 == 4)
    check(10 - 3 == 7)
    check(6 * 7 == 42)

spell test_appraise_strings():
    # Test string operations
    greeting = "Hello, " + "World!"
    check(greeting == "Hello, World!")
```

### The check Statement

The `check` statement is the core assertion mechanism in Sindri. It supports two formats:

#### Boolean Assertion Format
```carrion
check(boolean_expression)
```

#### Value Comparison Format
```carrion
check(actual_value, expected_value)
```

- If the values are equal, the test passes
- If the values differ, the test fails with a descriptive error message
- Type mismatches are automatically detected and reported

#### Error Message Format
When a test fails, you'll see detailed error messages like:
```
Value 9 didn't match Value 10, Expected 9 to Equal 10 got 9 instead
```

## Test Examples

### Basic Assertions

```carrion
spell appraise_boolean_operations():
    check(True and True, True)
    check(True and False, False)
    check(not False, True)

spell appraise_comparison():
    check(5 > 3, True)
    check(10 >= 10, True)
    check(2 < 8, True)
```

### Mathematical Operations

```carrion
spell math_appraise_division():
    result = 15 / 3
    check(result, 5)
```

### Intentional Failures

```carrion
spell appraise_intentional_failure():
    # This test will fail for demonstration
    check(1, 2)  # Will output: Value 1 didn't match Value 2, Expected 1 to Equal 2 got 1 instead
```

## Test Output

Sindri provides colored terminal output for easy test result identification:

- **Green "PASSED"**: Test executed successfully
- **Red "FAILED"**: Test failed with error message
- **Test Summary**: Shows total passed/failed counts

### Output Modes

#### Summary Mode (Default)
When running multiple test files, Sindri shows a concise summary:
```
=== Test Summary ===
tests/math_appraise.crl                            - FAILED
tests/string_test_appraise.crl                     - PASSED

Passed: 7/10
Failed: 3/10

Some tests failed! Use -d or --detailed for more information.
```

#### Detailed Mode (-d or --detailed)
Shows full test output with specific error messages:
```
Running tests in tests/math_appraise.crl...
Found 3 test function(s)

Running appraise_addition... PASSED
Running appraise_multiplication_fail... FAILED - Assertion Check Failed: Value 12 didn't match Value 15, Expected 12 to Equal 15 got 12 instead

=== Test Summary for tests/math_appraise.crl ===
Passed: 2
Failed: 1
Total:  3
```

### Sample Output

```
Running tests in sindri_demo.crl...
Found 5 test function(s)

Running appraise_arithmetic... PASSED
Running test_appraise_strings... PASSED
Running appraise_boolean_operations... PASSED
Running math_appraise_division... PASSED
Running appraise_intentional_failure... FAILED - Type mismatch: expected INTEGER, got INTEGER

=== Test Summary ===
Passed: 4
Failed: 1
Total:  5

Some tests failed!
```

## Supported Features

### Function Types

- **Standalone Functions**: Regular Carrion spells containing "appraise"
- **Grimoire Methods**: Instance methods within grimoires (limited support)

### Test Discovery

Sindri automatically discovers and runs tests in multiple ways:

#### Function Discovery
Within each test file, Sindri scans for:
1. Function definitions containing "appraise"
2. Grimoire method definitions containing "appraise"

#### File Discovery
Sindri can automatically find test files containing "appraise" in their filename:
- `test_appraise.crl`
- `appraise_math.crl`
- `string_appraise_tests.crl`

#### Directory Scanning
- `sindri appraise` - Runs all appraise files in current directory
- `sindri appraise ./tests` - Runs all appraise files in specified directory
- `sindri appraise file.crl` - Runs specific test file

#### Path Display
- **Relative paths**: Used when files are in the same directory or subdirectories
- **Absolute paths**: Used when files are in parent directories or different branches

### Error Handling

- Parser errors are caught and reported before test execution
- Runtime errors in tests are captured and displayed
- Type mismatches in assertions are automatically detected

## Limitations

- Grimoire method testing has limited support
- Test files must have `.crl` extension
- Tests must be written as functions (not inline code)

## File Structure

The Sindri framework consists of:

- `cmd/sindri/main.go`: Main executable and test runner
- `cmd/sindri/go.mod`: Go module definition
- Built-in `check()` function in evaluator builtins

## Best Practices

1. Use descriptive test function names that clearly indicate what is being tested
2. Group related assertions within the same test function
3. Include both positive and negative test cases
4. Use meaningful variable names in your tests
5. Keep tests simple and focused on specific functionality

## Integration

Sindri integrates seamlessly with the Carrion language ecosystem:

- Uses the same lexer, parser, and evaluator as the main interpreter
- Supports all Carrion language features within tests
- Provides the same error handling and type checking