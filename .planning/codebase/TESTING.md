# Testing Patterns

**Analysis Date:** 2026-01-18

## Test Framework

**Runner:**
- Go built-in `testing` package
- Custom test framework: Sindri (`cmd/sindri/main.go`) for Carrion tests

**Assertion Library:**
- Go standard library comparisons
- Manual assertions with `t.Errorf()`, `t.Fatalf()`

**Run Commands:**
```bash
go test ./...                           # Run all tests
go test ./src/evaluator/...             # Run evaluator tests
go test -v ./...                        # Verbose output
make test                               # Via Makefile (if configured)
sindri appraise examples/               # Run Carrion tests with Sindri
sindri appraise -d examples/            # Detailed output
sindri appraise -r examples/            # Generate HTML report
```

## Test File Organization

**Location:**
- `*_test.go` alongside source files (Go standard)
- `src/tests/` for integration tests

**Naming:**
- `<module>_test.go` for unit tests (e.g., `lexer_test.go`, `evaluator_test.go`)
- `integration_test.go` for integration tests
- `bug_repro_test.go` for regression tests

**Structure:**
```
src/
  ast/
    ast.go
    ast_test.go
  evaluator/
    evaluator.go
    evaluator_test.go
  lexer/
    lexer.go
    lexer_test.go
  parser/
    parser.go
    parse_test.go
  object/
    object.go
    object_test.go
  tests/
    integration_test.go
    bug_repro_test.go
bifrost/internal/
  auth/
    auth.go
    auth_test.go
  config/
    config_test.go
  manifest/
    manifest_test.go
  archive/
    archive_test.go
  version/
    version_test.go
```

## Test Structure

**Suite Organization:**
```go
func TestFeatureName(t *testing.T) {
    // Table-driven tests preferred
    tests := []struct {
        name     string
        input    string
        expected string
        wantErr  bool
    }{
        {"success case", "valid input", "expected output", false},
        {"error case", "invalid input", "", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // arrange
            // act
            result, err := functionUnderTest(tt.input)
            // assert
            if (err != nil) != tt.wantErr {
                t.Errorf("unexpected error: %v", err)
            }
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}
```

**Patterns:**
- Table-driven tests with `t.Run()` for subtests
- `t.TempDir()` for temporary file isolation
- HTTP mocking with `httptest.NewServer()`
- Cleanup with `defer` or `t.Cleanup()`

## Mocking

**Framework:**
- Go standard library (`httptest` for HTTP mocking)
- Manual mock implementations

**Patterns:**
```go
// HTTP mocking (from bifrost/internal/auth/auth_test.go)
server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // Check request
    if r.Method != "POST" {
        t.Errorf("expected POST, got %s", r.Method)
    }
    // Return response
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}))
defer server.Close()
```

**What to Mock:**
- HTTP servers for API testing
- File system via temporary directories

**What NOT to Mock:**
- Core interpreter components (test as integration)
- Pure functions

## Fixtures and Factories

**Test Data:**
```go
// Inline test data in table-driven tests
tests := []struct {
    input    string
    expected int64
}{
    {"5", 5},
    {"10", 10},
    {"999", 999},
}
```

**Location:**
- Inline in test files for simple cases
- `examples/` directory for Carrion test programs
- `t.TempDir()` for temporary test files

## Coverage

**Requirements:**
- No enforced coverage target
- Coverage tracked for awareness

**Configuration:**
- Go built-in coverage via `-cover` flag

**View Coverage:**
```bash
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Test Types

**Unit Tests:**
- Test single functions/types in isolation
- Located in `*_test.go` files
- Examples: `src/lexer/lexer_test.go`, `src/parser/parse_test.go`

**Integration Tests:**
- Test full Carrion programs
- Located in `src/tests/integration_test.go`
- Runs `.crl` files from `examples/` directory
- Tracks pass/fail/skip rates

**Sindri Tests (Carrion-level):**
- Test functions with "appraise" in name
- Written in Carrion language
- Executed by Sindri framework
- Generates HTML reports

## Common Patterns

**Async Testing:**
```go
func TestAsyncOperation(t *testing.T) {
    // Use channels or sync.WaitGroup
    done := make(chan bool)
    go func() {
        // async operation
        done <- true
    }()
    select {
    case <-done:
        // success
    case <-time.After(time.Second):
        t.Fatal("timeout")
    }
}
```

**Error Testing:**
```go
func TestErrorCase(t *testing.T) {
    _, err := functionThatShouldFail(invalidInput)
    if err == nil {
        t.Error("expected error, got nil")
    }
}
```

**Lexer/Parser Testing:**
```go
func TestLexer(t *testing.T) {
    input := `let x = 5`
    l := lexer.New(input)

    tests := []struct {
        expectedType    token.TokenType
        expectedLiteral string
    }{
        {token.LET, "let"},
        {token.IDENT, "x"},
        {token.ASSIGN, "="},
        {token.INT, "5"},
    }

    for i, tt := range tests {
        tok := l.NextToken()
        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - wrong type. expected=%q, got=%q",
                i, tt.expectedType, tok.Type)
        }
    }
}
```

**Snapshot Testing:**
- Not used in this codebase
- Prefer explicit assertions

---

*Testing analysis: 2026-01-18*
*Update when test patterns change*
