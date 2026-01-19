// src/utils/enhanced_error_formatter.go
package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/javanhut/TheCarrionLanguage/src/object"
)

// Enhanced ANSI color codes
const (
	// Base colors
	ColorReset     = "\033[0m"
	ColorRed       = "\033[31m"
	ColorGreen     = "\033[32m"
	ColorYellow    = "\033[33m"
	ColorBlue      = "\033[34m"
	ColorMagenta   = "\033[35m"
	ColorCyan      = "\033[36m"
	ColorWhite     = "\033[37m"
	
	// Bright colors
	ColorBrightRed    = "\033[91m"
	ColorBrightGreen  = "\033[92m"
	ColorBrightYellow = "\033[93m"
	ColorBrightBlue   = "\033[94m"
	ColorBrightMagenta = "\033[95m"
	ColorBrightCyan   = "\033[96m"
	
	// Styles
	StyleBold      = "\033[1m"
	StyleDim       = "\033[2m"
	StyleItalic    = "\033[3m"
	StyleUnderline = "\033[4m"
	
	// Background colors
	BgRed    = "\033[41m"
	BgGreen  = "\033[42m"
	BgYellow = "\033[43m"
	BgBlue   = "\033[44m"
)

// PrintEnhancedError formats and prints an enhanced error with comprehensive context
func PrintEnhancedError(err *object.EnhancedError) {
	// Header with error level and category
	levelColor := getLevelColor(err.Level)
	categoryColor := getCategoryColor(err.Category)
	
	fmt.Printf("\n%s%s%s %s[%s] %s%s%s\n", 
		StyleBold, levelColor, err.Level.String(), 
		categoryColor, err.Category.String(), 
		ColorReset, StyleBold, err.Code)
	
	// Title and main message
	if err.Title != "" {
		fmt.Printf("%s%s%s: %s\n", StyleBold, ColorRed, err.Title, ColorReset)
	}
	
	// Main error span with source code
	printErrorSpan(err.MainSpan, err.Message, err.Level, true)
	
	// Additional labels
	for _, label := range err.Labels {
		fmt.Printf("\n%s%s%s:%s %s\n", 
			StyleBold, getLevelColor(label.Level), label.Level.String(), 
			ColorReset, label.Message)
		printErrorSpan(label.Span, label.Message, label.Level, false)
	}
	
	// Suggestions
	if len(err.Suggestions) > 0 {
		fmt.Printf("\n%s%s%s:%s\n", StyleBold, ColorBrightGreen, "help", ColorReset)
		for _, suggestion := range err.Suggestions {
			fmt.Printf("  %s%s%s\n", StyleBold, suggestion.Title, ColorReset)
			if suggestion.Description != "" {
				fmt.Printf("    %s\n", suggestion.Description)
			}
			
			// Print fixes
			for _, fix := range suggestion.Fixes {
				if fix.Description != "" {
					fmt.Printf("    %s• %s%s\n", ColorBrightGreen, fix.Description, ColorReset)
				}
				if fix.Replacement != "" {
					fmt.Printf("      %sReplace with:%s %s\n", 
						StyleBold, ColorReset, fix.Replacement)
				}
			}
		}
	}
	
	// Notes
	if len(err.Notes) > 0 {
		fmt.Printf("\n%s%s%s:%s\n", StyleBold, ColorBrightCyan, "note", ColorReset)
		for _, note := range err.Notes {
			fmt.Printf("  %s\n", note.Message)
			if note.Span != nil {
				printErrorSpan(*note.Span, note.Message, note.Level, false)
			}
		}
	}
	
	// Stack trace
	if len(err.Stack) > 0 {
		fmt.Printf("\n%s%s%s:%s\n", StyleBold, ColorBrightBlue, "stack backtrace", ColorReset)
		for i, entry := range err.Stack {
			fmt.Printf("  %2d: %s%s%s\n", i, StyleBold, entry.FunctionName, ColorReset)
			fmt.Printf("       at %s\n", formatPosition(entry.Position))
		}
	}
	
	// Related errors
	if len(err.RelatedErrors) > 0 {
		fmt.Printf("\n%s%s%s:%s\n", StyleBold, ColorYellow, "related errors", ColorReset)
		for _, relatedErr := range err.RelatedErrors {
			fmt.Printf("  %s%s%s: %s\n", 
				StyleBold, getLevelColor(relatedErr.Level), relatedErr.Level.String(), 
				relatedErr.Message)
			fmt.Printf("    --> %s\n", formatErrorSpan(relatedErr.MainSpan))
		}
	}
	
	// Cause chain
	if err.Cause != nil {
		fmt.Printf("\n%s%s%s:%s\n", StyleBold, ColorMagenta, "caused by", ColorReset)
		PrintEnhancedError(err.Cause)
	}
	
	fmt.Printf("\n")
}

// printErrorSpan prints a source code span with highlighting
func printErrorSpan(span object.ErrorSpan, message string, level object.ErrorLevel, isMain bool) {
	// Position info
	fmt.Printf("  %s--> %s%s\n", ColorBrightBlue, formatErrorSpan(span), ColorReset)
	
	// Source code context
	if span.Start.Filename == "<repl>" || span.Start.Filename == "<input>" {
		printReplSpanContext(span, message, level, isMain)
	} else {
		printFileSpanContext(span, message, level, isMain)
	}
}

// printFileSpanContext prints file-based source context
func printFileSpanContext(span object.ErrorSpan, message string, level object.ErrorLevel, isMain bool) {
	if span.Start.Filename == "" || span.Start.Line <= 0 {
		return
	}
	
	file, err := os.Open(span.Start.Filename)
	if err != nil {
		fmt.Printf("   %s%s%s = note: could not read source file\n", 
			ColorYellow, StyleBold, ColorReset)
		return
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	lines := []string{}
	currentLine := 1
	
	// Read relevant lines
	contextStart := max(1, span.Start.Line-2)
	contextEnd := span.End.Line + 2
	
	for scanner.Scan() {
		if currentLine >= contextStart && currentLine <= contextEnd {
			lines = append(lines, scanner.Text())
		}
		if currentLine > contextEnd {
			break
		}
		currentLine++
	}
	
	if len(lines) == 0 {
		return
	}
	
	// Print line numbers and source
	maxLineNum := contextEnd
	lineNumWidth := len(fmt.Sprintf("%d", maxLineNum))
	
	for i, line := range lines {
		lineNum := contextStart + i
		
		// Line number gutter
		if lineNum >= span.Start.Line && lineNum <= span.End.Line {
			// Error line
			fmt.Printf("   %s%*d%s %s|%s ", 
				ColorBrightRed, lineNumWidth, lineNum, ColorReset, 
				ColorBrightRed, ColorReset)
		} else {
			// Context line
			fmt.Printf("   %s%*d%s %s|%s ", 
				ColorBrightBlue, lineNumWidth, lineNum, ColorReset, 
				ColorBrightBlue, ColorReset)
		}
		
		// Print line with highlighting
		if lineNum >= span.Start.Line && lineNum <= span.End.Line {
			printHighlightedLine(line, span, lineNum, level)
		} else {
			fmt.Printf("%s\n", line)
		}
	}
	
	// Print error indicator
	if span.Start.Line == span.End.Line {
		printErrorIndicator(span, lineNumWidth, level, message)
	}
}

// printReplSpanContext prints REPL-based source context
func printReplSpanContext(span object.ErrorSpan, message string, level object.ErrorLevel, isMain bool) {
	if span.Start.Line <= 0 {
		return
	}
	
	line, exists := replHistory[span.Start.Line]
	if !exists {
		fmt.Printf("   %s%s%s = note: REPL history not available\n", 
			ColorYellow, StyleBold, ColorReset)
		return
	}
	
	// Print with context
	lineNumWidth := len(fmt.Sprintf("%d", span.Start.Line))
	
	// Print the line
	fmt.Printf("   %s%*d%s %s|%s ", 
		ColorBrightRed, lineNumWidth, span.Start.Line, ColorReset, 
		ColorBrightRed, ColorReset)
	
	printHighlightedLine(line, span, span.Start.Line, level)
	
	// Print error indicator
	printErrorIndicator(span, lineNumWidth, level, message)
}

// printHighlightedLine prints a line with error highlighting
func printHighlightedLine(line string, span object.ErrorSpan, lineNum int, level object.ErrorLevel) {
	if lineNum != span.Start.Line {
		fmt.Printf("%s\n", line)
		return
	}
	
	startCol := max(1, span.Start.Column) - 1
	endCol := span.End.Column
	if span.Start.Line != span.End.Line {
		endCol = len(line)
	}
	
	// Print line with highlighting
	if startCol < len(line) {
		// Before error
		fmt.Printf("%s", line[:startCol])
		
		// Error part
		errorPart := line[startCol:min(endCol, len(line))]
		levelColor := getLevelColor(level)
		fmt.Printf("%s%s%s%s", levelColor, StyleBold, errorPart, ColorReset)
		
		// After error
		if endCol < len(line) {
			fmt.Printf("%s", line[endCol:])
		}
	} else {
		fmt.Printf("%s", line)
	}
	fmt.Printf("\n")
}

// printErrorIndicator prints the error indicator with carets
func printErrorIndicator(span object.ErrorSpan, lineNumWidth int, level object.ErrorLevel, message string) {
	if span.Start.Line != span.End.Line {
		return
	}
	
	startCol := max(1, span.Start.Column) - 1
	endCol := span.End.Column - 1
	if endCol <= startCol {
		endCol = startCol + 1
	}
	
	// Padding for line number
	padding := strings.Repeat(" ", lineNumWidth+3)
	
	// Spaces before caret
	spaces := strings.Repeat(" ", startCol)
	
	// Carets
	caretLen := endCol - startCol
	if caretLen <= 0 {
		caretLen = 1
	}
	carets := strings.Repeat("^", caretLen)
	
	levelColor := getLevelColor(level)
	fmt.Printf("   %s%s|%s%s%s%s%s %s\n", 
		padding[:lineNumWidth], ColorBrightRed, ColorReset, 
		spaces, levelColor, carets, ColorReset, message)
}

// getLevelColor returns the appropriate color for an error level
func getLevelColor(level object.ErrorLevel) string {
	switch level {
	case object.ERROR_LEVEL_ERROR:
		return ColorBrightRed
	case object.ERROR_LEVEL_WARNING:
		return ColorBrightYellow
	case object.ERROR_LEVEL_NOTE:
		return ColorBrightCyan
	case object.ERROR_LEVEL_HELP:
		return ColorBrightGreen
	default:
		return ColorReset
	}
}

// getCategoryColor returns the appropriate color for an error category
func getCategoryColor(category object.ErrorCategory) string {
	switch category {
	case object.ERROR_CATEGORY_SYNTAX:
		return ColorBrightRed
	case object.ERROR_CATEGORY_TYPE:
		return ColorBrightMagenta
	case object.ERROR_CATEGORY_RUNTIME:
		return ColorBrightYellow
	case object.ERROR_CATEGORY_SEMANTIC:
		return ColorBrightCyan
	case object.ERROR_CATEGORY_IMPORT:
		return ColorBrightBlue
	case object.ERROR_CATEGORY_IO:
		return ColorBrightGreen
	case object.ERROR_CATEGORY_CUSTOM:
		return ColorWhite
	default:
		return ColorReset
	}
}

// formatErrorSpan formats an error span for display
func formatErrorSpan(span object.ErrorSpan) string {
	if span.Start.Filename == span.End.Filename {
		if span.Start.Line == span.End.Line {
			if span.Start.Column == span.End.Column {
				return fmt.Sprintf("%s:%d:%d", span.Start.Filename, span.Start.Line, span.Start.Column)
			}
			return fmt.Sprintf("%s:%d:%d-%d", span.Start.Filename, span.Start.Line, span.Start.Column, span.End.Column)
		}
		return fmt.Sprintf("%s:%d:%d-%d:%d", span.Start.Filename, span.Start.Line, span.Start.Column, span.End.Line, span.End.Column)
	}
	return fmt.Sprintf("%s:%d:%d to %s:%d:%d", 
		span.Start.Filename, span.Start.Line, span.Start.Column,
		span.End.Filename, span.End.Line, span.End.Column)
}

// formatPosition formats a source position for display
func formatPosition(pos object.SourcePosition) string {
	if pos.Filename == "" || pos.Line <= 0 {
		return "<unknown>"
	}
	return fmt.Sprintf("%s:%d:%d", pos.Filename, pos.Line, pos.Column)
}

// PrintEnhancedParseError prints enhanced parser errors
func PrintEnhancedParseError(filename string, content string, errors []object.EnhancedError) {
	fmt.Printf("\n%s%s%s[%s] %s%s%s\n", 
		StyleBold, ColorBrightRed, "error", 
		"syntax", ColorReset, StyleBold, "PARSE_ERROR")
	
	for i, err := range errors {
		if i > 0 {
			fmt.Printf("\n")
		}
		PrintEnhancedError(&err)
	}
}

// ConvertParseErrorsToEnhanced converts traditional parser errors to enhanced errors
func ConvertParseErrorsToEnhanced(filename string, content string, errors []string) []object.EnhancedError {
	var enhancedErrors []object.EnhancedError
	lines := strings.Split(content, "\n")
	
	for _, errMsg := range errors {
		// Try to extract position information
		var lineNum, colNum int
		
		// Try different formats
		if n, _ := fmt.Sscanf(errMsg, "at line %d, column %d", &lineNum, &colNum); n < 2 {
			fmt.Sscanf(errMsg, "line %d:%d", &lineNum, &colNum)
		}
		
		// Create span
		span := object.ErrorSpan{
			Start: object.SourcePosition{
				Filename: filename,
				Line:     lineNum,
				Column:   colNum,
			},
			End: object.SourcePosition{
				Filename: filename,
				Line:     lineNum,
				Column:   colNum + 1,
			},
		}
		
		// Add source if available
		if lineNum > 0 && lineNum <= len(lines) {
			span.Source = lines[lineNum-1]
		}
		
		// Create enhanced error
		enhanced := object.NewSyntaxError(errMsg, span)
		
		// Add context-specific suggestions
		addParseErrorSuggestions(enhanced, errMsg)
		
		enhancedErrors = append(enhancedErrors, *enhanced)
	}
	
	return enhancedErrors
}

// addParseErrorSuggestions adds suggestions based on parse error patterns
func addParseErrorSuggestions(err *object.EnhancedError, errMsg string) {
	switch {
	case strings.Contains(errMsg, "expected COLON"):
		err.AddSuggestion(
			"Add missing colon",
			"Control structures and function definitions require a colon (:) at the end",
			object.ErrorFix{
				Description: "Add ':' at the end of the line",
				Replacement: ":",
			},
		)
		
	case strings.Contains(errMsg, "expected RPAREN"):
		err.AddSuggestion(
			"Add missing closing parenthesis",
			"Every opening parenthesis '(' must have a matching closing parenthesis ')'",
			object.ErrorFix{
				Description: "Add ')' to close the parentheses",
				Replacement: ")",
			},
		)
		
	case strings.Contains(errMsg, "expected INDENT"):
		err.AddSuggestion(
			"Add proper indentation",
			"Code blocks after ':' must be indented (use 4 spaces)",
			object.ErrorFix{
				Description: "Indent the next line with 4 spaces",
			},
		)
		
	case strings.Contains(errMsg, "no prefix parse function"):
		err.AddSuggestion(
			"Check expression syntax",
			"The token cannot be used at the beginning of an expression",
			object.ErrorFix{
				Description: "Review the expression syntax and fix any typos",
			},
		)
		
	case strings.Contains(errMsg, "expected next token"):
		err.AddSuggestion(
			"Check token sequence",
			"The parser expected a different token at this position",
			object.ErrorFix{
				Description: "Review the syntax and ensure proper token ordering",
			},
		)
	}
}

// Helper function to create a span from a single position
func CreateSpanFromPosition(pos object.SourcePosition) object.ErrorSpan {
	return object.ErrorSpan{
		Start: pos,
		End: object.SourcePosition{
			Filename: pos.Filename,
			Line:     pos.Line,
			Column:   pos.Column + 1,
		},
	}
}

// Helper function to create a span from token positions
func CreateSpanFromTokens(start, end object.SourcePosition) object.ErrorSpan {
	return object.ErrorSpan{
		Start: start,
		End:   end,
	}
}