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
)

func PrintError(err *object.ErrorWithTrace) {
	fmt.Printf("%s%sError: %s%s\n", Bold, Red, err.Message, Reset)

	// Print location
	fmt.Printf("  at Positon: %s%s:Line: %d, Column: %d%s\n",
		Bold, err.Position.Filename, err.Position.Line, err.Position.Column, Reset)

	// Try to print source context
	printSourceContext(err.Position)

	// Print stack trace
	if len(err.Stack) > 0 {
		fmt.Printf("\n%sStack trace:%s\n", Bold, Reset)
		for i, entry := range err.Stack {
			fmt.Printf("  %d: %s%s%s (%s: Line: %d: Column: %d)\n",
				i, Bold, entry.FunctionName, Reset,
				entry.Position.Filename, entry.Position.Line, entry.Position.Column)
			printSourceContext(entry.Position)
		}
	}

	// Print custom error details
	if err.ErrorType == object.CUSTOM_ERROR_OBJ && len(err.CustomDetails) > 0 {
		fmt.Printf("\n%sDetails:%s\n", Bold, Reset)
		for key, value := range err.CustomDetails {
			fmt.Printf("  %s%s:%s %s\n", Bold, key, Reset, value.Inspect())
		}
	}
}

func printSourceContext(pos object.SourcePosition) {
	// Skip for REPL input
	if pos.Filename == "<repl>" || pos.Filename == "" || pos.Line <= 0 {
		return
	}

	file, err := os.Open(pos.Filename)
	if err != nil {
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

	if len(context) == 0 {
		return
	}

	// Print the context
	fmt.Println()
	contextLineIndex := pos.Line - lineStart
	if contextLineIndex >= 0 && contextLineIndex < len(context) {
		for i, line := range context {
			lineNum := lineStart + i

			if lineNum == pos.Line {
				// The error line
				fmt.Printf("  %s%3d |%s %s\n", Bold, lineNum, Reset, line)

				// Print the error indicator
				if pos.Column > 0 {
					padding := strings.Repeat(" ", min(pos.Column, len(line)+1))
					fmt.Printf("      |%s%s%s^\n", Reset, padding, Red)
				}
			} else {
				// Context line
				fmt.Printf("  %3d | %s\n", lineNum, line)
			}
		}
		fmt.Println()
	}
}

// Helper function
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func PrintErrorToWriter(out *os.File, err *object.ErrorWithTrace) {
	fmt.Fprintf(out, "%s%sError: %s%s\n", Bold, Red, err.Message, Reset)
	fmt.Fprintf(out, "  at %s%s:Line: %d, Column: %d%s\n",
		Bold, err.Position.Filename, err.Position.Line, err.Position.Column, Reset)

	// Print stack trace
	if len(err.Stack) > 0 {
		fmt.Fprintf(out, "\n%sStack trace:%s\n", Bold, Reset)
		for i, entry := range err.Stack {
			fmt.Fprintf(out, "  %d: %s%s%s (%s:%d:%d)\n",
				i, Bold, entry.FunctionName, Reset,
				entry.Position.Filename, entry.Position.Line, entry.Position.Column)
		}
	}
}

// PrintParseFail formats parser error information
func PrintParseFail(filename string, content string, errors []string) {
	fmt.Printf("%s%sParse Error in %s%s\n", Bold, Red, filename, Reset)

	lines := strings.Split(content, "\n")
	for _, errMsg := range errors {
		// Try to extract line and column information from error messages
		// Assuming error messages have some pattern like "expected X at line Y, column Z"
		var lineNum, colNum int
		fmt.Sscanf(errMsg, "at line %d, column %d", &lineNum, &colNum)

		fmt.Printf("  %s\n", errMsg)

		// If we could extract line info, show the code snippet
		if lineNum > 0 && lineNum <= len(lines) {
			fmt.Println()
			lineStart := max(1, lineNum-1)
			lineEnd := min(len(lines), lineNum+1)

			for i := lineStart; i <= lineEnd; i++ {
				if i > 0 && i <= len(lines) {
					if i == lineNum {
						fmt.Printf("  %s%3d |%s %s\n", Bold, i, Reset, lines[i-1])
						if colNum > 0 {
							padding := strings.Repeat(" ", colNum+5)
							indicatorPos := min(colNum+1, len(padding))
							fmt.Printf("      |%s%s%s^\n", Reset, padding[:indicatorPos], Red)
						}
					} else {
						fmt.Printf("  %3d | %s\n", i, lines[i-1])
					}
				}
			}
			fmt.Println()
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
