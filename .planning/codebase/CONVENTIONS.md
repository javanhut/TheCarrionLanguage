# Coding Conventions

**Analysis Date:** 2026-01-18

## Naming Patterns

**Files:**
- `snake_case.go` for all Go source files (e.g., `enhanced_error_integration.go`, `bound_method.go`)
- `*_test.go` alongside source files for tests
- `*module.go` suffix for module implementations (e.g., `httpmodule.go`, `timemodule.go`)

**Functions:**
- PascalCase for exported functions (e.g., `New()`, `NextToken()`, `Eval()`)
- camelCase for unexported functions (e.g., `readString()`, `peekChar()`)
- `Eval*()` prefix for evaluation functions (e.g., `EvalProgram`, `evalStatement`)
- `is*()` prefix for predicates (e.g., `isTruthy`, `isBuiltinFunction`)

**Variables:**
- camelCase for local variables (e.g., `indentStack`, `currLine`, `sourceFile`)
- UPPER_CASE for constants (e.g., `INDENT`, `DEDENT`, `MAX_CALL_DEPTH`)
- No underscore prefix for private members

**Types:**
- PascalCase for interfaces and structs (e.g., `Lexer`, `Parser`, `Object`)
- No `I` prefix for interfaces (e.g., `Object`, not `IObject`)
- `*Statement`/`*Expression` suffix for AST nodes
- `*Obj` suffix optional for object type constants

## Code Style

**Formatting:**
- Go standard: tabs for indentation
- Consistent spacing around operators
- Run `go fmt ./...` or `make fmt`

**Linting:**
- `go vet` for static analysis (`make vet`)
- `golangci-lint` when available (`make lint`)

## Import Organization

**Order:**
1. Standard library packages
2. Third-party packages
3. Internal project packages

**Grouping:**
- Blank line between groups
- Alphabetical within each group

**Path Aliases:**
- Full import paths used (e.g., `github.com/javanhut/TheCarrionLanguage/src/object`)

## Error Handling

**Patterns:**
- Return `*object.Error` for runtime errors in evaluator
- Collect errors in slice for parser (non-fatal until parsing complete)
- Use Go error type for internal Go errors
- Include context in error messages

**Error Types:**
- `*object.Error` - Carrion runtime errors
- `*object.ErrorWithTrace` - Enhanced errors with stack traces
- Go `error` interface - Internal tooling errors

**Logging:**
- `fmt.Printf` for debug output
- Console output via Carrion's print() builtin
- No structured logging library

## Logging

**Framework:**
- Standard fmt package for output
- No external logging framework

**Patterns:**
- Debug mode controlled by environment variable (`CARRION_DEBUG_WRAPPING`)
- Command-line flags for verbose output (`--idebug`, `--lexer`, `--parser`, `--evaluator`)

## Comments

**When to Comment:**
- Explain complex algorithms or non-obvious logic
- Document public API functions
- Mark TODO items for future work

**GoDoc:**
- Brief comments above exported functions/types
- Example: `// New creates a new Lexer instance`

**TODO Comments:**
- Format: `// TODO: description`
- Example: `// TODO: Sort by semantic version - for now just return alphabetical`

## Function Design

**Size:**
- Most functions under 50 lines
- Notable exception: `src/evaluator/evaluator.go` contains large eval functions

**Parameters:**
- Use variadic `...object.Object` for builtin functions
- Pointer receivers for struct methods

**Return Values:**
- Return `object.Object` from evaluation functions
- Return error as second value for Go-style error handling
- Use named returns sparingly

## Module Design

**Exports:**
- PascalCase for exported items
- Organize by functional domain (lexer, parser, evaluator, etc.)

**Package Structure:**
- One main type per package (e.g., Lexer in lexer package)
- Helper types in same package
- Tests in same package with `_test` suffix

**Barrel Files:**
- Not used (Go doesn't have barrel file convention)

## Carrion Language Conventions

**Keywords:**
- `spell` for function definition
- `grim` for class (grimoire) definition
- `attempt/resolve/ensnare` for try/catch/finally
- `diverge/converge` for goroutine spawn/sync

**Naming in Carrion:**
- snake_case for variables and functions
- PascalCase for grimoire (class) names
- `self` for instance reference
- `init` for constructor

---

*Convention analysis: 2026-01-18*
*Update when patterns change*
