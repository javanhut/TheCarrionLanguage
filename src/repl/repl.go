// src/repl/repl.go
package repl

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/peterh/liner"

	"github.com/javanhut/TheCarrionLanguage/src/debug"
	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
	"github.com/javanhut/TheCarrionLanguage/src/token"
	"github.com/javanhut/TheCarrionLanguage/src/utils"
)

const ODINS_EYE = `

  ███████████████████████████████████████████████████████████████████
  ███████████████████████████████████████████████████████████████████
  ███████████████████████████████████████████████████████████████████
  █████████████████████████████████  █████████  ██  ▒████████████████
  ███████████████████████  █████████ █████████     ██████████████████
  ████████████████████  █  █████████     ██████       ███████████████
  █████████████████   █   █████████  ██   ████    ███████████████████
  █████████████████████   ██   █████    █████  ██ ███████████████████
  ████████████████████████     █████  █████  █  █  ██████████████████
  █████████████████████████  ███████  ████  ███ █████████████████████
  ███████████████████████████  █          ░██████████████████████████
  ███████████████████████████  ████   ███  ██████████████████████████
  ████████████████    ███████ ██████  ████  █████████████████████████
  ███████████████  ██   ████                             ████████████
  ████████████                    ██  █████ ███      ████████████████
  ████████████████████████████ █████  ████  █████  ██████████████████
  ███████████████████████ █████  ███  ██    █████████████████████████
  █████████████████████ █  ██  ██      █████  ███████████████████████
  █████████████████████  █    ██████  █████     ███    ██████████████
  ██████████████████████    ███████   ████  ██     ██████████████████
  █████████████████████    ██████     █████   ██     ████████████████
  █████████████████        █████   █  ███████████  █ ████████████████
  ███████████████████   █  ██████     ███████████  ██████████████████
  ███████████████████ ███  █████████  ███████████████████████████████
  ██████████████████████████████████  ███████████████████████████████
  ███████████████████████████████████████████████████████████████████
  ███████████████████████████████████████████████████████████████████

  `

// Start begins the REPL
func Start(in io.Reader, out io.Writer, env *object.Environment) {
	line := liner.NewLiner()
	evaluator.LineReader = line
	print := fmt.Fprintln
	defer func() {
		ok := line.Close()
		if ok != nil {
			log.Fatal("Unable to close the file: ", ok)
		}
		evaluator.LineReader = nil
		// Clean up global state to prevent memory leaks
		evaluator.CleanupGlobalState()
		utils.ClearReplHistory()
	}()

	if env == nil {
		env = object.NewEnvironment()
	}

	// Optional: Set a custom tab completion function
	line.SetCompleter(func(input string) []string {
		keywords := []string{
			// Control flow
			"if", "else", "otherwise", "for", "in", "while", "match", "case", "skip", "stop", "return",
			// Literals and constants
			"True", "False", "None",
			// Object-oriented
			"spell", "grim", "init", "self", "super", "arcane", "arcanespell",
			// Error handling
			"attempt", "resolve", "ensnare", "raise", "check",
			// Module system
			"import", "as",
			// Built-in functions
			"print", "input", "len", "type", "range", "max", "abs", "ord", "chr",
			"int", "float", "str", "bool", "list", "tuple", "enumerate", "pairs", "is_sametype",
			// Standard library constructors
			"Array", "String", "Integer", "Float", "Boolean", "File", "OS",
			// Standard library functions
			"help", "version", "modules",
			// REPL commands
			"clear", "quit", "exit",
		}

		// Built-in method suggestions for common patterns
		methodSuggestions := []string{
			// Array methods
			".append", ".sort", ".reverse", ".contains", ".length", ".get", ".set", ".clear",
			".index_of", ".remove", ".first", ".last", ".slice", ".is_empty", ".to_string",
			// String methods
			".upper", ".lower", ".find", ".char_at", ".reverse",
			// Integer methods
			".to_bin", ".to_oct", ".to_hex", ".abs", ".pow", ".gcd", ".lcm",
			".is_even", ".is_odd", ".is_prime", ".to_float",
			// Float methods
			".round", ".floor", ".ceil", ".sqrt", ".sin", ".cos", ".abs",
			".is_integer", ".is_positive", ".is_negative", ".is_zero", ".to_int",
			// Boolean methods
			".to_int", ".negate", ".and_with", ".or_with", ".xor_with",
			// File methods
			".read", ".write", ".append", ".exists",
			// OS methods
			".cwd", ".chdir", ".listdir", ".mkdir", ".remove", ".getenv", ".setenv", ".run", ".sleep",
		}

		// Only suggest keywords at the beginning of input
		if strings.TrimSpace(input) == "" {
			return keywords
		}

		prefix := strings.ToLower(input)
		var suggestions []string

		// Check for method completions (when input contains a dot)
		if strings.Contains(input, ".") {
			for _, method := range methodSuggestions {
				if strings.HasPrefix(strings.ToLower(method), "."+strings.ToLower(strings.Split(input, ".")[len(strings.Split(input, "."))-1])) {
					// Reconstruct the full suggestion
					parts := strings.Split(input, ".")
					if len(parts) > 1 {
						base := strings.Join(parts[:len(parts)-1], ".")
						suggestions = append(suggestions, base+method)
					}
				}
			}
		} else {
			// Regular keyword completion
			for _, kw := range keywords {
				if strings.HasPrefix(strings.ToLower(kw), prefix) {
					suggestions = append(suggestions, kw)
				}
			}
		}

		return suggestions
	})

	// Optional: Load history from a file
	historyFile := ".carrion_history"
	if f, err := os.Open(historyFile); err == nil {
		_, err := line.ReadHistory(f)
		if err != nil {
			log.Fatal("Error occured", err)
		}

		closed := f.Close()
		if closed != nil {
			log.Fatal("Unable to close file. Error: ", closed)
		}
	}

	// Save history on exit
	defer func() {
		if f, err := os.Create(historyFile); err == nil {
			line.WriteHistory(f)
			f.Close()
		}
	}()

	// Clear REPL history for error tracking
	utils.ClearReplHistory()

	if len(os.Args) > 1 {
		filePath := os.Args[1]
		if strings.HasSuffix(filePath, ".crl") {
			err := ProcessFile(filePath, out, env)
			if err != nil {
				fmt.Fprintf(out, "Error processing file: %v\n", err)
				return
			}
			return
		} else {
			print(out, "Unsupported file type. Only .crl files are allowed.")
			return
		}
	}

	var inputBuffer strings.Builder
	isMultiline := false
	currentIndentLevel := 0
	baseIndentLevel := 0
	inIfBlock := false
	lineNumber := 1            // Track line numbers for error context
	consecutiveEmptyLines := 0 // Track consecutive empty lines for double-enter detection

	print(out, "Welcome to the Carrion Programming Language REPL!")
	print(out, "")
	print(out, "Quick Help:")
	print(out, "   • Type 'help()' for basic information")
	print(out, "   • Type 'version()' to see current version")
	print(out, "   • Type 'modules()' to list standard library modules")
	print(out, "")
	print(out, "REPL Commands:")
	print(out, "   • 'clear' - Clear screen")
	print(out, "   • 'quit', 'exit', 'q', 'qa' - Exit REPL")
	print(out, "   • Use Tab for auto-completion")
	print(out, "")
	print(out, "May Mimir guide your coding journey! Type commands below:")

	for {
		var prompt string
		if !isMultiline {
			prompt = ">>> "
		} else {
			prompt = "... "
		}

		// Get input from the user
		input, err := line.Prompt(prompt)
		if err != nil {
			if err == io.EOF {
				print(out, "\nFarewell, May the All Father bless your travels!")
				return
			}
			fmt.Fprintf(out, "Error reading line: %v\n", err)
			continue
		}

		// Register this line for error context
		utils.RegisterReplLine(lineNumber, input)
		lineNumber++

		trimmedLine := strings.ToLower(strings.TrimSpace(input))

		// Handle special commands only at the primary prompt
		if !isMultiline {
			switch trimmedLine {
			case "exit", "quit", "q", "qa", "qa!":
				print(out, "Farewell, May the All Father bless your travels!")
				return
			case "clear":
				clearScreen(out)
				utils.ClearReplHistory() // Clear history on screen clear
				lineNumber = 1           // Reset line counter
				continue
			case "mimir", "MIMIR", "Mimir", "help":
				cmd := exec.Command("mimir")
				cmd.Stdout = out
				cmd.Stderr = out
				cmd.Stdin = os.Stdin
				err := cmd.Run()
				if err != nil {
					fmt.Fprintf(out, "Error running mimir: %v\n", err)
				}
				continue
			case "":
				continue
			}
		}

		// Add input to history if not empty
		if trimmedLine != "" {
			line.AppendHistory(input)
		}

		// Count leading spaces to determine indentation level
		indentSpaces := len(input) - len(strings.TrimLeft(input, " "))
		currentIndentLevel = indentSpaces / 4 // Assuming 4 spaces per indent level

		// Handle empty lines
		if trimmedLine == "" {
			consecutiveEmptyLines++
			if isMultiline {
				inputBuffer.WriteString(input)
				inputBuffer.WriteString("\n")
				// Double-enter detection: if user presses enter twice on empty lines
				// and we're in a multi-line block, force evaluation
				if consecutiveEmptyLines >= 2 {
					shouldEvaluate := true
					if shouldEvaluate {
						input := inputBuffer.String()
						if strings.TrimSpace(input) != "" {
							evaluated, complete := tryParseAndEval(input, out, env)
							if complete {
								if evaluated != nil && evaluated.Type() != object.NONE_OBJ {
									fmt.Fprintf(out, "%s\n", evaluated.Inspect())
								}
							}
						}
						// Reset all state
						inputBuffer.Reset()
						isMultiline = false
						baseIndentLevel = 0
						inIfBlock = false
						consecutiveEmptyLines = 0
					}
				}
			}
			continue
		}

		// Reset consecutive empty lines counter when we get actual content
		consecutiveEmptyLines = 0

		// Check if this is the start of an if block
		if strings.HasPrefix(trimmedLine, "if ") && strings.HasSuffix(trimmedLine, ":") {
			inIfBlock = true
			isMultiline = true
			baseIndentLevel = currentIndentLevel
		}

		// Check for otherwise or else clauses
		if inIfBlock && currentIndentLevel <= baseIndentLevel &&
			(strings.HasPrefix(trimmedLine, "otherwise") || strings.HasPrefix(trimmedLine, "else")) {
			isMultiline = true
		}

		// Append the line to our input buffer
		inputBuffer.WriteString(input)
		inputBuffer.WriteString("\n")

		// Determine if we should evaluate
		shouldEvaluate := false

		// Check for a complete block
		if isMultiline {
			if currentIndentLevel <= baseIndentLevel && !strings.HasSuffix(trimmedLine, ":") &&
				!strings.HasPrefix(trimmedLine, "otherwise") &&
				!strings.HasPrefix(trimmedLine, "else") {
				shouldEvaluate = true
				inIfBlock = false
			}
		} else if !strings.HasSuffix(trimmedLine, ":") {
			shouldEvaluate = true
		}

		if shouldEvaluate {
			input := inputBuffer.String()
			if strings.TrimSpace(input) == "" {
				inputBuffer.Reset()
				continue
			}

			evaluated, complete := tryParseAndEval(input, out, env)
			if complete {
				if evaluated != nil && evaluated.Type() != object.NONE_OBJ {
					fmt.Fprintf(out, "%s\n", evaluated.Inspect())
				}
				inputBuffer.Reset()
				isMultiline = false
				baseIndentLevel = 0
				inIfBlock = false
			} else {
				isMultiline = true
			}
		}
	}
}

// clearScreen clears the terminal screen
func clearScreen(out io.Writer) {
	// ANSI escape sequence to clear screen and move cursor to home position
	fmt.Fprint(out, "\033[H\033[2J")
	// Additional sequence to clear scroll-back buffer on some terminals
	fmt.Fprint(out, "\033[3J")
}

// ProcessFile reads, parses, and evaluates a Carrion source file
func ProcessFile(filePath string, out io.Writer, env *object.Environment) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	// Tokenize, parse, and evaluate the file contents
	l := lexer.NewWithFilename(string(content), filePath)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		utils.PrintParseFail(filePath, string(content), p.Errors())
		return fmt.Errorf("file %s contains syntax errors", filePath)
	}

	evaluated := evaluator.Eval(program, env, nil)

	// Handle errors with improved formatting
	if errObj, ok := evaluated.(*object.ErrorWithTrace); ok {
		utils.PrintError(errObj)
		return fmt.Errorf("runtime error in file %s", filePath)
	}

	if errObj, ok := evaluated.(*object.Error); ok {
		// Convert simple errors to error with trace for consistent formatting
		traceError := &object.ErrorWithTrace{
			ErrorType: object.ERROR_OBJ,
			Message:   errObj.Message,
			Position: object.SourcePosition{
				Filename: filePath,
				Line:     1,
				Column:   1,
			},
		}
		utils.PrintError(traceError)
		return fmt.Errorf("runtime error in file %s", filePath)
	}

	if evaluated != nil && evaluated.Type() != object.NONE_OBJ {
		fmt.Fprintf(out, "%s\n", evaluated.Inspect())
	}
	return nil
}

// tryParseAndEval attempts to parse and evaluate the input
func tryParseAndEval(input string, out io.Writer, env *object.Environment) (object.Object, bool) {
	if out == nil {
	}
	l := lexer.NewWithFilename(
		input,
		"<repl>",
	) // Use <repl> as the filename for better error reporting
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		if isIncompleteParse(p.Errors()) {
			return nil, false
		}
		utils.PrintParseFail("<repl>", input, p.Errors())
		return nil, true
	}

	evaluated := evaluator.Eval(program, env, nil)
	if evaluated == nil {
		return nil, true
	}

	// Use custom error printer for all errors
	if errObj, ok := evaluated.(*object.ErrorWithTrace); ok {
		utils.PrintError(errObj)
		return nil, true
	}

	if errObj, ok := evaluated.(*object.Error); ok {
		// Convert simple errors to error with trace for consistent formatting
		traceError := &object.ErrorWithTrace{
			ErrorType: object.ERROR_OBJ,
			Message:   errObj.Message,
			Position: object.SourcePosition{
				Filename: "<repl>",
				Line:     1,
				Column:   1,
			},
		}
		utils.PrintError(traceError)
		return nil, true
	}

	if returnValue, ok := evaluated.(*object.ReturnValue); ok {
		evaluated = returnValue.Value
	}

	return evaluated, true
}

// isIncompleteParse checks if the parser errors indicate incomplete input
func isIncompleteParse(errs []string) bool {
	for _, err := range errs {
		if strings.Contains(strings.ToLower(err), "unexpected end") ||
			strings.Contains(strings.ToLower(err), "unexpected eof") ||
			strings.Contains(strings.ToLower(err), "incomplete") ||
			strings.Contains(strings.ToLower(err), "expected next token") {
			return true
		}
	}
	return false
}

// StartWithDebug begins the REPL with debug configuration
func StartWithDebug(in io.Reader, out io.Writer, env *object.Environment, debugConfig *debug.Config) {
	Start(in, out, env)
}

// ProcessFileWithDebug reads, parses, and evaluates a Carrion source file with debug output
func ProcessFileWithDebug(filePath string, out io.Writer, env *object.Environment, debugConfig *debug.Config) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	// Tokenize with debug output
	l := lexer.NewWithFilename(string(content), filePath)

	if debugConfig.ShouldDebugLexer() {
		fmt.Fprintf(os.Stderr, "\n=== LEXER OUTPUT ===\n")
		// Create a copy of the lexer for debug output
		debugLexer := lexer.NewWithFilename(string(content), filePath)
		for {
			tok := debugLexer.NextToken()
			fmt.Fprintf(os.Stderr, "lexer: Token{Type: %s, Literal: %q, Line: %d, Column: %d}\n",
				tok.Type, tok.Literal, tok.Line, tok.Column)
			if tok.Type == token.EOF {
				break
			}
		}
		fmt.Fprintf(os.Stderr, "===================\n\n")
	}

	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		utils.PrintParseFail(filePath, string(content), p.Errors())
		return fmt.Errorf("file %s contains syntax errors", filePath)
	}

	if debugConfig.ShouldDebugParser() {
		fmt.Fprintf(os.Stderr, "\n=== PARSER OUTPUT ===\n")
		fmt.Fprintf(os.Stderr, "parser: Program with %d statements\n", len(program.Statements))
		for i, stmt := range program.Statements {
			fmt.Fprintf(os.Stderr, "parser: Statement[%d]: %T - %s\n", i, stmt, stmt.String())
		}
		fmt.Fprintf(os.Stderr, "====================\n\n")
	}

	if debugConfig.ShouldDebugEvaluator() {
		fmt.Fprintf(os.Stderr, "\n=== EVALUATOR OUTPUT ===\n")
	}

	evaluated := evaluator.EvalWithDebug(program, env, nil, debugConfig)

	if debugConfig.ShouldDebugEvaluator() {
		fmt.Fprintf(os.Stderr, "=====================\n\n")
	}

	// Handle errors with improved formatting
	if errObj, ok := evaluated.(*object.ErrorWithTrace); ok {
		utils.PrintError(errObj)
		return fmt.Errorf("runtime error in file %s", filePath)
	}

	if errObj, ok := evaluated.(*object.Error); ok {
		// Convert simple errors to error with trace for consistent formatting
		traceError := &object.ErrorWithTrace{
			ErrorType: object.ERROR_OBJ,
			Message:   errObj.Message,
			Position: object.SourcePosition{
				Filename: filePath,
				Line:     1,
				Column:   1,
			},
		}
		utils.PrintError(traceError)
		return fmt.Errorf("runtime error in file %s", filePath)
	}

	if evaluated != nil && evaluated.Type() != object.NONE_OBJ {
		fmt.Fprintf(out, "%s\n", evaluated.Inspect())
	}
	return nil
}
