// src/repl/repl.go
package repl

import (
	"fmt"
	"io"
	"log"
	"os"
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
			"mimir", "clear", "quit", "exit",
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
			fmt.Fprintln(out, "Unsupported file type. Only .crl files are allowed.")
			return
		}
	}

	var inputBuffer strings.Builder
	isMultiline := false
	currentIndentLevel := 0
	baseIndentLevel := 0
	inIfBlock := false
	lineNumber := 1 // Track line numbers for error context
	consecutiveEmptyLines := 0 // Track consecutive empty lines for double-enter detection

	fmt.Fprintln(out, "🦅 Welcome to the Carrion Programming Language REPL! 🦅")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "📖 Quick Help:")
	fmt.Fprintln(out, "   • Type 'mimir' for comprehensive function help")
	fmt.Fprintln(out, "   • Type 'help()' for basic information")
	fmt.Fprintln(out, "   • Type 'version()' to see current version")
	fmt.Fprintln(out, "   • Type 'modules()' to list standard library modules")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "⚡ REPL Commands:")
	fmt.Fprintln(out, "   • 'clear' - Clear screen")
	fmt.Fprintln(out, "   • 'quit', 'exit', 'q', 'qa' - Exit REPL")
	fmt.Fprintln(out, "   • Use Tab for auto-completion")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "✨ Quick Examples:")
	fmt.Fprintln(out, "   print(\"Hello, World!\")     // Basic output")
	fmt.Fprintln(out, "   x, y = (10, 20)            // Tuple unpacking")
	fmt.Fprintln(out, "   42.to_bin()                // \"0b101010\"")
	fmt.Fprintln(out, "   \"hello\".upper()            // \"HELLO\"")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "May Mimir guide your coding journey! Type commands below:")

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
				fmt.Fprintln(out, "\nFarewell, May the All Father bless your travels!")
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
			case "exit", "quit", "q", "qa":
				fmt.Fprintln(out, "Farewell, May the All Father bless your travels!")
				return
			case "clear":
				clearScreen(out)
				utils.ClearReplHistory() // Clear history on screen clear
				lineNumber = 1           // Reset line counter
				continue
			case "mimir":
				startInteractiveHelp(line, out)
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

// startInteractiveHelp launches the interactive help system
func startInteractiveHelp(line *liner.State, out io.Writer) {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "═══════════════════════════════════════════════════════════════════")
	fmt.Fprintln(out, "🧙‍♂️ Welcome to Mimir's Interactive Help System 🧙‍♂️")
	fmt.Fprintln(out, "═══════════════════════════════════════════════════════════════════")
	fmt.Fprintln(out, "")

	for {
		showHelpMenu(out)
		
		input, err := line.Prompt("help> ")
		if err != nil {
			if err == io.EOF {
				fmt.Fprintln(out, "\nReturning to main REPL...")
				return
			}
			fmt.Fprintf(out, "Error reading input: %v\n", err)
			continue
		}

		choice := strings.ToLower(strings.TrimSpace(input))
		
		switch choice {
		case "1", "builtins", "builtin":
			showBuiltinFunctions(line, out)
		case "2", "stdlib", "standard", "munin":
			showStandardLibrary(line, out)
		case "3", "syntax", "language":
			showLanguageFeatures(line, out)
		case "4", "examples", "demo":
			showExamples(line, out)
		case "5", "search", "find":
			searchFunctions(line, out)
		case "6", "tips", "tricks":
			showTipsAndTricks(out)
		case "h", "help", "menu":
			// Will show menu again on next iteration
			continue
		case "q", "quit", "exit", "back":
			fmt.Fprintln(out, "\n🦅 Returning to main REPL...")
			return
		case "clear":
			clearScreen(out)
		default:
			if choice != "" {
				// Try to find function by name
				if found := searchSpecificFunction(choice, out); !found {
					fmt.Fprintf(out, "❌ Unknown command '%s'. Type 'h' for menu or 'q' to quit.\n\n", input)
				}
			}
		}
	}
}

// showHelpMenu displays the main help menu
func showHelpMenu(out io.Writer) {
	fmt.Fprintln(out, "📚 What would you like help with?")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "  1️⃣  Built-in Functions    - Core language functions (print, len, type, etc.)")
	fmt.Fprintln(out, "  2️⃣  Standard Library      - Munin modules (Array, String, File, OS, etc.)")
	fmt.Fprintln(out, "  3️⃣  Language Features     - Syntax, control flow, OOP, error handling")
	fmt.Fprintln(out, "  4️⃣  Examples & Demos      - Working code examples and tutorials")
	fmt.Fprintln(out, "  5️⃣  Search Functions      - Find specific functions by name or purpose")
	fmt.Fprintln(out, "  6️⃣  Tips & Tricks         - REPL shortcuts and advanced features")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "💡 Commands: Type number/name (e.g., '1' or 'builtins'), 'h' for menu, 'q' to quit")
	fmt.Fprintln(out, "🔍 Quick search: Type any function name directly (e.g., 'print', 'Array')")
	fmt.Fprintln(out, "")
}

// showBuiltinFunctions shows the built-in functions menu
func showBuiltinFunctions(line *liner.State, out io.Writer) {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "🔧 BUILT-IN FUNCTIONS")
	fmt.Fprintln(out, "═══════════════════════")
	
	categories := map[string][]string{
		"1": {"Type Conversion", "int, float, str, bool, list, tuple"},
		"2": {"Utility Functions", "len, type, print, input, range"},
		"3": {"Mathematical", "max, abs, ord, chr"},
		"4": {"Collections", "enumerate, pairs, is_sametype"},
		"5": {"System Functions", "help, version, modules"},
	}
	
	for {
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "Select a category:")
		for k, v := range categories {
			fmt.Fprintf(out, "  %s. %s - %s\n", k, v[0], v[1])
		}
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "Commands: Enter number, function name, 'all' for everything, or 'b' to go back")
		
		input, err := line.Prompt("builtins> ")
		if err != nil || strings.ToLower(strings.TrimSpace(input)) == "b" {
			return
		}
		
		choice := strings.ToLower(strings.TrimSpace(input))
		switch choice {
		case "1":
			showTypeConversionFunctions(out)
		case "2":
			showUtilityFunctions(out)
		case "3":
			showMathFunctions(out)
		case "4":
			showCollectionFunctions(out)
		case "5":
			showSystemFunctions(out)
		case "all":
			showAllBuiltinFunctions(out)
		case "":
			continue
		default:
			if !searchSpecificFunction(choice, out) {
				fmt.Fprintf(out, "❌ Unknown category '%s'\n", input)
			}
		}
		
		fmt.Fprintln(out, "\nPress Enter to continue...")
		line.Prompt("")
	}
}

// showStandardLibrary shows the standard library menu
func showStandardLibrary(line *liner.State, out io.Writer) {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "🏛️  MUNIN STANDARD LIBRARY")
	fmt.Fprintln(out, "═════════════════════════")
	
	modules := map[string]string{
		"1": "Array - Enhanced array operations and manipulation",
		"2": "String - String processing and text manipulation",
		"3": "Integer - Integer utilities and number base conversion",
		"4": "Float - Floating-point operations and math functions",
		"5": "Boolean - Boolean logic and operations",
		"6": "File - File I/O and filesystem operations",
		"7": "OS - Operating system interface and process management",
		"8": "Math - Mathematical functions and constants",
	}
	
	for {
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "Select a module:")
		for k, v := range modules {
			fmt.Fprintf(out, "  %s. %s\n", k, v)
		}
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "Commands: Enter number, module name, 'all' for everything, or 'b' to go back")
		
		input, err := line.Prompt("stdlib> ")
		if err != nil || strings.ToLower(strings.TrimSpace(input)) == "b" {
			return
		}
		
		choice := strings.ToLower(strings.TrimSpace(input))
		switch choice {
		case "1", "array":
			showArrayModule(out)
		case "2", "string":
			showStringModule(out)
		case "3", "integer", "int":
			showIntegerModule(out)
		case "4", "float":
			showFloatModule(out)
		case "5", "boolean", "bool":
			showBooleanModule(out)
		case "6", "file":
			showFileModule(out)
		case "7", "os":
			showOSModule(out)
		case "8", "math":
			showMathModule(out)
		case "all":
			showAllStandardLibrary(out)
		case "":
			continue
		default:
			if !searchSpecificFunction(choice, out) {
				fmt.Fprintf(out, "❌ Unknown module '%s'\n", input)
			}
		}
		
		fmt.Fprintln(out, "\nPress Enter to continue...")
		line.Prompt("")
	}
}

// showLanguageFeatures shows language syntax and features
func showLanguageFeatures(line *liner.State, out io.Writer) {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "⚡ CARRION LANGUAGE FEATURES")
	fmt.Fprintln(out, "═══════════════════════════")
	
	features := map[string]string{
		"1": "Variables & Assignment - Basic assignment, tuple unpacking, operators",
		"2": "Control Flow - if/otherwise/else, for/while loops, match/case",
		"3": "Functions (Spells) - Function definition, parameters, return values",
		"4": "Classes (Grimoires) - OOP, inheritance, methods, properties",
		"5": "Error Handling - attempt/ensnare/resolve, raising errors",
		"6": "Modules & Imports - Code organization, importing files",
		"7": "Data Types - Primitives, collections, type checking",
		"8": "Operators - Arithmetic, logical, comparison, bitwise",
	}
	
	for {
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "Select a topic:")
		for k, v := range features {
			fmt.Fprintf(out, "  %s. %s\n", k, v)
		}
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "Commands: Enter number, topic name, 'all' for everything, or 'b' to go back")
		
		input, err := line.Prompt("syntax> ")
		if err != nil || strings.ToLower(strings.TrimSpace(input)) == "b" {
			return
		}
		
		choice := strings.ToLower(strings.TrimSpace(input))
		switch choice {
		case "1", "variables", "assignment":
			showVariablesAndAssignment(out)
		case "2", "control", "flow", "if", "for", "while":
			showControlFlow(out)
		case "3", "functions", "spells", "function":
			showFunctions(out)
		case "4", "classes", "grimoires", "oop", "class":
			showClasses(out)
		case "5", "errors", "error", "exceptions":
			showErrorHandling(out)
		case "6", "modules", "imports", "import":
			showModules(out)
		case "7", "types", "data":
			showDataTypes(out)
		case "8", "operators", "operator":
			showOperators(out)
		case "all":
			showAllLanguageFeatures(out)
		case "":
			continue
		default:
			fmt.Fprintf(out, "❌ Unknown topic '%s'\n", input)
		}
		
		fmt.Fprintln(out, "\nPress Enter to continue...")
		line.Prompt("")
	}
}

// showExamples shows code examples and tutorials
func showExamples(line *liner.State, out io.Writer) {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "🎯 EXAMPLES & TUTORIALS")
	fmt.Fprintln(out, "══════════════════════")
	
	examples := map[string]string{
		"1": "Hello World & Basics - Getting started with Carrion",
		"2": "Working with Arrays - Manipulation, sorting, searching",
		"3": "String Processing - Text manipulation and formatting",
		"4": "File Operations - Reading, writing, file management",
		"5": "Mathematical Calculations - Numbers, formulas, algorithms",
		"6": "Object-Oriented Programming - Classes, inheritance, methods",
		"7": "Error Handling Examples - Robust error management",
		"8": "Complete Mini Programs - Full working applications",
	}
	
	for {
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "Select an example category:")
		for k, v := range examples {
			fmt.Fprintf(out, "  %s. %s\n", k, v)
		}
		fmt.Fprintln(out, "")
		fmt.Fprintln(out, "Commands: Enter number, 'all' for everything, or 'b' to go back")
		
		input, err := line.Prompt("examples> ")
		if err != nil || strings.ToLower(strings.TrimSpace(input)) == "b" {
			return
		}
		
		choice := strings.ToLower(strings.TrimSpace(input))
		switch choice {
		case "1":
			showBasicExamples(out)
		case "2":
			showArrayExamples(out)
		case "3":
			showStringExamples(out)
		case "4":
			showFileExamples(out)
		case "5":
			showMathExamples(out)
		case "6":
			showOOPExamples(out)
		case "7":
			showErrorExamples(out)
		case "8":
			showMiniPrograms(out)
		case "all":
			showAllExamples(out)
		case "":
			continue
		default:
			fmt.Fprintf(out, "❌ Unknown example category '%s'\n", input)
		}
		
		fmt.Fprintln(out, "\nPress Enter to continue...")
		line.Prompt("")
	}
}

// searchFunctions provides interactive search functionality
func searchFunctions(line *liner.State, out io.Writer) {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "🔍 FUNCTION SEARCH")
	fmt.Fprintln(out, "═════════════════")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "Enter keywords to search for functions, or type 'categories' to browse by category.")
	fmt.Fprintln(out, "Examples: 'array', 'string upper', 'file read', 'math', 'convert'")
	fmt.Fprintln(out, "Commands: 'b' to go back, 'q' to quit, 'h' for help, 'categories' to browse")
	fmt.Fprintln(out, "")
	
	for {
		input, err := line.Prompt("search> ")
		if err != nil {
			if err == io.EOF {
				fmt.Fprintln(out, "\n🦅 Returning to help menu...")
			}
			return
		}
		
		query := strings.ToLower(strings.TrimSpace(input))
		
		// Handle exit commands
		switch query {
		case "b", "back":
			fmt.Fprintln(out, "\n🦅 Returning to help menu...")
			return
		case "q", "quit", "exit":
			fmt.Fprintln(out, "\n🦅 Returning to help menu...")
			return
		case "h", "help", "?":
			fmt.Fprintln(out, "")
			fmt.Fprintln(out, "🔍 SEARCH HELP:")
			fmt.Fprintln(out, "   • Type keywords to search: 'array', 'string upper', 'file read'")
			fmt.Fprintln(out, "   • 'categories' - Browse function categories")
			fmt.Fprintln(out, "   • 'b' or 'back' - Return to help menu")
			fmt.Fprintln(out, "   • 'q' or 'quit' - Return to help menu")
			fmt.Fprintln(out, "")
			continue
		case "":
			continue
		case "categories":
			showSearchCategories(out)
			continue
		}
		
		results := performFunctionSearch(query)
		if len(results) == 0 {
			fmt.Fprintf(out, "❌ No functions found matching '%s'\n", input)
			fmt.Fprintln(out, "💡 Try broader terms like 'array', 'string', 'file', or 'math'")
		} else {
			fmt.Fprintf(out, "\n🎯 Found %d function(s) matching '%s':\n\n", len(results), input)
			for i, result := range results {
				fmt.Fprintf(out, "%d. %s\n", i+1, result)
			}
		}
		fmt.Fprintln(out, "")
	}
}

// showTipsAndTricks shows REPL tips and advanced features
func showTipsAndTricks(out io.Writer) {
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "💡 TIPS & TRICKS")
	fmt.Fprintln(out, "═══════════════")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "🚀 REPL Shortcuts:")
	fmt.Fprintln(out, "   • Tab - Auto-complete functions and keywords")
	fmt.Fprintln(out, "   • ↑/↓ - Navigate command history")
	fmt.Fprintln(out, "   • 'clear' - Clear the screen")
	fmt.Fprintln(out, "   • 'mimir' - Open this interactive help")
	fmt.Fprintln(out, "   • Double Enter - Execute multi-line blocks")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "⚡ Language Features:")
	fmt.Fprintln(out, "   • All primitives auto-wrap: 42.to_bin(), \"hello\".upper()")
	fmt.Fprintln(out, "   • Tuple unpacking: x, y = (10, 20)")
	fmt.Fprintln(out, "   • F-strings: f\"Value is {x}\"")
	fmt.Fprintln(out, "   • Method chaining: arr.sort().reverse().to_string()")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "🎯 Quick Testing:")
	fmt.Fprintln(out, "   • Test expressions: type(42), len(\"hello\")")
	fmt.Fprintln(out, "   • Explore objects: dir(Array([1,2,3]))")
	fmt.Fprintln(out, "   • Check versions: version(), modules()")
	fmt.Fprintln(out, "")
	fmt.Fprintln(out, "🐛 Debugging:")
	fmt.Fprintln(out, "   • Print types: print(type(variable))")
	fmt.Fprintln(out, "   • Inspect values: print(repr(value))")
	fmt.Fprintln(out, "   • Use check() for assertions")
	fmt.Fprintln(out, "")
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
