# Codebase Concerns

**Analysis Date:** 2026-01-18

## Tech Debt

**Large monolithic evaluator file:**
- Issue: `src/evaluator/evaluator.go` is 5,257 lines with high cyclomatic complexity
- Why: Organic growth during development, all evaluation logic in one file
- Impact: Hard to navigate, test, and maintain; increases cognitive load
- Fix approach: Split into multiple files by concern (expressions, statements, control flow, imports)

**Large parser file:**
- Issue: `src/parser/parser.go` is 2,626 lines
- Why: Full language grammar in single file
- Impact: Difficult to locate specific parsing logic
- Fix approach: Extract expression parsing, statement parsing into separate files

**Large sockets module:**
- Issue: `src/modules/sockets.go` is 1,601 lines mixing TCP, UDP, WebSocket, HTTP
- Why: All network functionality evolved in one place
- Impact: Hard to test individual socket types; poor separation of concerns
- Fix approach: Split into `tcp.go`, `udp.go`, `websocket.go`, `http_server.go`

## Known Bugs

**No bugs explicitly documented**
- No FIXME or BUG comments found in codebase
- Integration tests track expected failures in skip list

## Security Considerations

**Path traversal validation incomplete:**
- Risk: `src/modules/sockets.go:268-271` checks for `..` but doesn't canonicalize paths
- Current mitigation: Basic string check for `..` in path
- Recommendations: Use `filepath.Clean()` and validate against base directory

**Hardcoded file permissions:**
- Risk: Files created with hardcoded `0644` or `0755` permissions
- Files: `src/modules/file.go:110-118`, `bifrost/cmd/bifrost/main.go:77,89`
- Current mitigation: None
- Recommendations: Make permissions configurable or respect umask

**Missing error checks on file operations:**
- Risk: Silent failures in Bifrost init command
- Files: `bifrost/cmd/bifrost/main.go:77,89` - `os.MkdirAll` and `os.WriteFile` without error checking
- Current mitigation: None
- Recommendations: Check and report all file operation errors

## Performance Bottlenecks

**No HTTP connection pooling:**
- Problem: New HTTP client created for every request
- Files: `src/modules/httpmodule.go:39-40`
- Measurement: Not benchmarked
- Cause: Client instantiated per-request instead of reused
- Improvement path: Create singleton HTTP client with connection pooling

**MIME type map recreated per call:**
- Problem: `getMimeType` creates new map on every call
- Files: `src/modules/sockets.go:232-258`
- Measurement: Not benchmarked
- Cause: Map literal instead of package-level constant
- Improvement path: Move MIME type map to package-level variable

**Version list not semantically sorted:**
- Problem: Package versions sorted alphabetically, not semantically
- Files: `src/evaluator/evaluator.go:4462`, `bifrost/internal/integration/carrion.go:76`
- Measurement: Not applicable
- Cause: TODO left in code: "Sort by semantic version - for now just return alphabetical"
- Improvement path: Implement semantic version comparison and sorting

## Fragile Areas

**Global state in evaluator:**
- Files: `src/evaluator/evaluator.go:28-48`
- Why fragile: Multiple global maps (`importedFiles`, `callStack`, `recursionDepths`, `globalGoroutineManager`)
- Common failures: State not cleaned up between runs, memory leaks in long-running processes
- Safe modification: Always call `CleanupGlobalState()` after evaluation
- Test coverage: Integration tests exist but don't verify cleanup

**File handle registry:**
- Files: `src/modules/file.go:12-15`
- Why fragile: Global `fileHandles` map with no automatic cleanup
- Common failures: Handle exhaustion if files not explicitly closed
- Safe modification: Always pair fileOpen with fileClose
- Test coverage: No dedicated tests for file module

**Socket handle registry:**
- Files: `src/modules/sockets.go:20-32`
- Why fragile: Same pattern as file handles; global state with manual cleanup
- Common failures: Socket/port exhaustion
- Safe modification: Close all sockets before program exit
- Test coverage: No dedicated tests for socket module

## Scaling Limits

**Recursion depth:**
- Current capacity: MAX_CALL_DEPTH = 1000 function calls
- Limit: Stack overflow protection kicks in at 1000 depth
- Symptoms at limit: Error returned, evaluation stops
- Scaling path: Increase limit or implement tail-call optimization

**Goroutine cleanup timeout:**
- Current capacity: 100ms wait for goroutine cleanup
- Limit: Slow goroutines may be orphaned
- Files: `src/evaluator/evaluator.go:86-90`
- Scaling path: Make timeout configurable or implement graceful shutdown

## Dependencies at Risk

**liner library:**
- Risk: `github.com/peterh/liner` is sole dependency for REPL line editing
- Impact: REPL would break without it; no graceful fallback
- Migration plan: Could implement basic readline fallback

## Missing Critical Features

**Registry integration incomplete:**
- Problem: Bifrost package manager has stubs for registry operations
- Files: `bifrost/cmd/bifrost/main.go:112`, `bifrost/internal/integration/carrion.go:76`
- Current workaround: Local package installation only
- Blocks: Cannot download packages from remote registry
- Implementation complexity: High (requires server-side registry)

## Test Coverage Gaps

**Socket module untested:**
- What's not tested: `src/modules/sockets.go` (1,601 lines) has no dedicated tests
- Risk: Network functionality could break silently
- Priority: High - complex I/O code
- Difficulty to test: Need mock servers, port management

**HTTP module untested:**
- What's not tested: `src/modules/httpmodule.go` (756 lines) has no dedicated tests
- Risk: HTTP client functionality could regress
- Priority: High - external API interactions
- Difficulty to test: Need HTTP mocking

**File module untested:**
- What's not tested: `src/modules/file.go` has no dedicated tests
- Risk: File operations could fail in edge cases
- Priority: Medium - relies on Go stdlib
- Difficulty to test: Need temporary file handling

**Time module untested:**
- What's not tested: `src/modules/timemodule.go` has no dedicated tests
- Risk: Time operations could have edge cases
- Priority: Low - mostly wraps Go time package
- Difficulty to test: Need time mocking for some tests

---

*Concerns audit: 2026-01-18*
*Update as issues are fixed or new ones discovered*
