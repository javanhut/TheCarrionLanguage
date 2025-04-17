// src/utils/error_printer.go
package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/javanhut/TheCarrionLanguage/src/object"
)

// ANSI color codes for terminal output
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Bold   = "\033[1m"
	Cyan   = "\033[36m"
)

// Store REPL history for error context
var replHistory = make(map[int]string)

// RegisterReplLine stores a line from REPL history for error context
func RegisterReplLine(lineNum int, content string) {
	replHistory[lineNum] = content
}

// ClearReplHistory clears the stored REPL history
func ClearReplHistory() {
	replHistory = make(map[int]string)
}

// PrintError formats and prints a Carrion error with source context
func PrintError(err *object.ErrorWithTrace) {
	// Print main error header with clear separation
	fmt.Printf(
		"\n%s%s══════════════════════════════════════════════════════════════════%s\n",
		Bold,
		Red,
		Reset,
	)
	fmt.Printf("%s%sERROR: %s%s\n", Bold, Red, err.Message, Reset)

	// Print location with clearer formatting
	filename := formatFilename(err.Position.Filename)
	if err.Position.Line > 0 {
		fmt.Printf(
			"%sLocation:%s %s at line %d, column %d\n",
			Bold,
			Reset,
			filename,
			err.Position.Line,
			err.Position.Column,
		)
	} else {
		fmt.Printf("%sLocation:%s %s\n", Bold, Reset, filename)
	}

	// Print source context with better formatting
	printSourceContext(err.Position, true)

	// Print stack trace with clearer formatting
	if len(err.Stack) > 0 {
		fmt.Printf("\n%sCall Stack:%s\n", Bold, Reset)
		for i, entry := range err.Stack {
			funcName := formatFunctionName(entry.FunctionName)
			locationInfo := formatLocationInfo(entry.Position)
			fmt.Printf("  %d: %s%s%s at %s\n", i+1, Bold, funcName, Reset, locationInfo)
		}
	}

	// Print custom error details with better formatting
	if err.Type() == object.CUSTOM_ERROR_OBJ && len(err.CustomDetails) > 0 {
		fmt.Printf("\n%sAdditional Information:%s\n", Bold, Reset)
		for key, value := range err.CustomDetails {
			fmt.Printf("  • %s%s:%s %s\n", Bold, key, Reset, value.Inspect())
		}
	}

	// Print fix suggestions based on error type
	printFixSuggestions(err)

	fmt.Printf(
		"%s%s══════════════════════════════════════════════════════════════════%s\n\n",
		Bold,
		Red,
		Reset,
	)
}

// Format filename for display
func formatFilename(filename string) string {
	if filename == "" || filename == "unknown" {
		return fmt.Sprintf("%s<unknown source>%s", Yellow, Reset)
	} else if filename == "<input>" || filename == "<repl>" {
		return fmt.Sprintf("%sREPL Input%s", Cyan, Reset)
	} else {
		return fmt.Sprintf("%s%s%s", Bold, filename, Reset)
	}
}

// Format function name for display
func formatFunctionName(name string) string {
	if name == "<anonymous function>" {
		return "anonymous function"
	} else if name == "while_loop" || name == "for_loop" || name == "if_block" {
		return name + " block"
	}
	return name + "()"
}

// Format location info for display
func formatLocationInfo(pos object.SourcePosition) string {
	if pos.Filename == "unknown" || pos.Line <= 0 {
		return fmt.Sprintf("%s<unknown location>%s", Yellow, Reset)
	} else if pos.Filename == "<input>" || pos.Filename == "<repl>" {
		return fmt.Sprintf("%sREPL Input line %d%s", Cyan, pos.Line, Reset)
	} else {
		return fmt.Sprintf("%s:%d", pos.Filename, pos.Line)
	}
}

// Helper to print the source code context for an error
func printSourceContext(pos object.SourcePosition, isMainError bool) {
	// Handle <repl> or <input> source specially
	if pos.Filename == "<input>" || pos.Filename == "<repl>" {
		printReplContext(pos.Line, pos.Column, isMainError)
		return
	}

	// Only try to print context if we have a valid filename and line number
	if pos.Filename == "" || pos.Filename == "unknown" || pos.Line <= 0 {
		fmt.Printf("  %s(source code not available)%s\n", Yellow, Reset)
		return
	}

	// Try to read the source file
	file, err := os.Open(pos.Filename)
	if err != nil {
		fmt.Printf("  %s(could not open source file: %s)%s\n", Yellow, err, Reset)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentLine := 1

	// Find the error line and surrounding context
	context := []string{}
	lineStart := max(1, pos.Line-2)
	lineEnd := pos.Line + 2

	for scanner.Scan() {
		if currentLine >= lineStart && currentLine <= lineEnd {
			context = append(context, scanner.Text())
		}
		if currentLine > lineEnd {
			break
		}
		currentLine++
	}

	if scanErr := scanner.Err(); scanErr != nil {
		fmt.Printf("  %s(error reading source file: %s)%s\n", Yellow, scanErr, Reset)
	}

	// If no context could be found
	if len(context) == 0 {
		fmt.Printf("  %s(source code not available for line %d)%s\n", Yellow, pos.Line, Reset)
		return
	}

	// Print the context with clear formatting
	fmt.Println()
	contextLineIndex := pos.Line - lineStart
	if contextLineIndex >= 0 && contextLineIndex < len(context) {
		// Print a header for the code section
		if isMainError {
			fmt.Printf("%s%sSource Code:%s\n", Bold, Blue, Reset)
		}

		for i, line := range context {
			lineNum := lineStart + i

			// Calculate prefix based on maximum line number width
			maxLineNum := lineStart + len(context) - 1
			lineNumWidth := len(fmt.Sprintf("%d", maxLineNum))

			if lineNum == pos.Line {
				// The error line - highlighted
				fmt.Printf("  %s%*d |%s %s\n", Bold, lineNumWidth, lineNum, Reset, line)

				// Print the error indicator if we have column info
				if pos.Column > 0 {
					pointerCol := min(pos.Column, len(line)+1)
					// padding := strings.Repeat(" ", pointerCol+lineNumWidth+3)
					fmt.Printf("  %s |%s%s%s%s^ %sError occurs here%s\n",
						strings.Repeat(" ", lineNumWidth),
						Reset,
						strings.Repeat(" ", pointerCol),
						Bold+Red,
						Reset+Yellow,
						Reset)
				}
			} else {
				// Context line
				fmt.Printf("  %*d | %s\n", lineNumWidth, lineNum, line)
			}
		}
		fmt.Println()
	}
}

// Print context for REPL input
func printReplContext(line int, column int, isMainError bool) {
	if len(replHistory) == 0 || replHistory[line] == "" {
		fmt.Printf("  %s(REPL source code not available for line %d)%s\n", Yellow, line, Reset)
		return
	}

	// Get the line from history
	codeLine := replHistory[line]

	// Print the context with clear formatting
	fmt.Println()
	if isMainError {
		fmt.Printf("%s%sREPL Input:%s\n", Bold, Blue, Reset)
	}

	fmt.Printf("  %s%3d |%s %s\n", Bold, line, Reset, codeLine)

	// Print the error indicator if we have column info
	if column > 0 {
		pointerCol := min(column, len(codeLine)+1)
		padding := strings.Repeat(" ", pointerCol)
		fmt.Printf("      |%s%s%s%s^ %sError occurs here%s\n",
			Reset,
			padding,
			Bold+Red,
			Reset+Yellow,
			Reset)
	}
	fmt.Println()
}

// Print suggestions based on error type
func printFixSuggestions(err *object.ErrorWithTrace) {
	// Extract error type from the message
	errorType := ""
	if strings.Contains(err.Message, "invalid assignment target") {
		errorType = "invalid_assignment"
	} else if strings.Contains(err.Message, "index operator not supported") {
		errorType = "invalid_index"
	} else if strings.Contains(err.Message, "identifier not found") {
		errorType = "undefined_variable"
	} else if strings.Contains(err.Message, "type mismatch") {
		errorType = "type_mismatch"
	}

	// Print suggestion based on error type
	if errorType != "" {
		fmt.Printf("\n%sSuggestion:%s ", Bold+Green, Reset)

		switch errorType {
		case "invalid_assignment":
			fmt.Println(
				"Carrion doesn't support direct array element assignment like arr[i] = value.",
			)
            fmt.Println(
                "  Try creating an Array grim with a 'set(index, value)' method instead:",
            )
            fmt.Println("    grim Array:")
			fmt.Println("      spell set(index, value):")
			fmt.Println("        new_elements = []")
			fmt.Println("        for i in range(len(self.elements)):")
			fmt.Println("          if i == index:")
			fmt.Println("            new_elements = new_elements + [value]")
			fmt.Println("          else:")
			fmt.Println("            new_elements = new_elements + [self.elements[i]]")
			fmt.Println("        self.elements = new_elements")

		case "invalid_index":
			fmt.Println("You're trying to use [] indexing on a type that doesn't support it.")
			fmt.Println("  If this is a custom object, try implementing a 'get(index)' method:")
            fmt.Println("    grim YourType:")
			fmt.Println("      spell get(index):")
			fmt.Println("        // Your implementation here")

		case "undefined_variable":
			fmt.Println("The variable name you're trying to use hasn't been defined.")
			fmt.Println("  Make sure to declare variables before using them:")
			fmt.Println("    my_var = initial_value")

		case "type_mismatch":
			fmt.Println("You're trying to perform an operation on incompatible types.")
			fmt.Println("  Try using type conversion functions like int(), str(), or float():")
			fmt.Println("    converted_value = int(other_value)")
		}
	}
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// PrintParseFail formats parser error information
func PrintParseFail(filename string, content string, errors []string) {
	fmt.Printf(
		"\n%s%s══════════════════════════════════════════════════════════════════%s\n",
		Bold,
		Red,
		Reset,
	)
	fmt.Printf("%s%sPARSE ERROR in %s%s\n", Bold, Red, filename, Reset)

	lines := strings.Split(content, "\n")
	for i, errMsg := range errors {
		if i > 0 {
			fmt.Println() // Add spacing between multiple errors
		}

		// Try to extract line and column information from error messages
		var lineNum, colNum int
		if n, _ := fmt.Sscanf(errMsg, "at line %d, column %d", &lineNum, &colNum); n < 2 {
			// Try another common format
			fmt.Sscanf(errMsg, "line %d:%d", &lineNum, &colNum)
		}

		fmt.Printf("%s%sError:%s %s\n", Bold, Red, Reset, formatErrorMessage(errMsg))

		// If we could extract line info, show the code snippet
		if lineNum > 0 && lineNum <= len(lines) {
			fmt.Printf("%s%sSource:%s\n", Bold, Blue, Reset)
			lineStart := max(1, lineNum-1)
			lineEnd := min(len(lines), lineNum+1)

			maxLineNum := lineEnd
			lineNumWidth := len(fmt.Sprintf("%d", maxLineNum))

			for i := lineStart; i <= lineEnd; i++ {
				if i > 0 && i <= len(lines) {
					if i == lineNum {
						fmt.Printf("  %s%*d |%s %s\n", Bold, lineNumWidth, i, Reset, lines[i-1])
						if colNum > 0 {
							pointerCol := min(colNum, len(lines[i-1])+1)
							padding := strings.Repeat(" ", pointerCol)
							fmt.Printf("  %s |%s%s%s%s^ %sError occurs here%s\n",
								strings.Repeat(" ", lineNumWidth),
								Reset,
								padding,
								Bold+Red,
								Reset+Yellow,
								Reset)
						}
					} else {
						fmt.Printf("  %*d | %s\n", lineNumWidth, i, lines[i-1])
					}
				}
			}

			// Add suggestions for common parsing errors
			printParseErrorSuggestion(errMsg)
		}
	}
	fmt.Printf(
		"%s%s══════════════════════════════════════════════════════════════════%s\n\n",
		Bold,
		Red,
		Reset,
	)
}

// Format common error messages to be more user-friendly
func formatErrorMessage(msg string) string {
	// Replace common internal terms with more user-friendly ones
	msg = strings.ReplaceAll(msg, "expected next token to be", "expected")
	msg = strings.ReplaceAll(msg, "got=", "got ")

	// Extract the main message without technical details
	if idx := strings.Index(msg, ", got"); idx > 0 {
		return msg[:idx] + msg[idx:]
	}

	return msg
}

// Print suggestions for common parsing errors
func printParseErrorSuggestion(errMsg string) {
	fmt.Printf("\n%sSuggestion:%s ", Bold+Green, Reset)

	switch {
	case strings.Contains(errMsg, "expected COLON"):
		fmt.Println(
			"You may be missing a colon (:) at the end of a statement or block declaration.",
		)
		fmt.Println("  Example: if condition:")
		fmt.Println("           spell function_name():")

	case strings.Contains(errMsg, "expected RPAREN"):
		fmt.Println(
			"You may be missing a closing parenthesis ')' in an expression or function call.",
		)
		fmt.Println("  Example: function_name(arg1, arg2)")

	case strings.Contains(errMsg, "expected INDENT"):
		fmt.Println("You need to indent the block after a colon (:).")
		fmt.Println("  Example: if condition:")
		fmt.Println("             indented_statement")

	case strings.Contains(errMsg, "unexpected DEDENT"):
		fmt.Println("You may have incorrect indentation. Check that all lines in a block")
		fmt.Println("  have consistent indentation levels (use spaces, not tabs).")

	case strings.Contains(errMsg, "unexpected token"):
		fmt.Println("There's an unexpected symbol or keyword in your code.")
		fmt.Println("  Check for typos or misplaced operators/keywords.")

	default:
		fmt.Println("Check your syntax carefully. Make sure:")
		fmt.Println("  - All blocks (if, spell, etc.) end with a colon (:)")
		fmt.Println("  - All parentheses, brackets, and braces are properly closed")
		fmt.Println("  - Indentation is consistent (use 4 spaces per level)")
		fmt.Println("  - Variable names follow the language conventions")
	}
}
