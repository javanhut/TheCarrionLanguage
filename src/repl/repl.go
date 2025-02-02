package repl

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/parser"

	"github.com/peterh/liner"
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

func Start(in io.Reader, out io.Writer, env *object.Environment) {
	line := liner.NewLiner()
	evaluator.LineReader = line

	defer func() {
		line.Close()
		evaluator.LineReader = nil
	}()

	if env == nil {
		env = object.NewEnvironment()
	}

	// Optional: Set a custom tab completion function
	// line.SetCompleter(func(line string) []string {
	// 	// Implement auto-completion logic here
	// 	return nil
	// })

	// Optional: Load history from a file
	// if f, err := os.Open(historyFile); err == nil {
	// 	line.ReadHistory(f)
	// 	f.Close()
	// }
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

	fmt.Fprintln(out, "Welcome to the Carrion Programming Language REPL!")
	fmt.Fprintln(out, "Type 'exit' or 'quit' to exit, 'clear' to clear the screen.")
	fmt.Fprintln(out, "Type any commands you like may Mimir guide your hand.")

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

		trimmedLine := strings.TrimSpace(input)

		// Handle special commands only at the primary prompt
		if !isMultiline {
			switch trimmedLine {
			case "exit", "quit":
				fmt.Fprintln(out, "Farewell, May the All Father bless your travels!")
				return
			case "clear":
				clearScreen(out)
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
			if isMultiline {
				inputBuffer.WriteString(input)
				inputBuffer.WriteString("\n")
			}
			continue
		}

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

func tryParseAndEval(input string, out io.Writer, env *object.Environment) (object.Object, bool) {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		if isIncompleteParse(p.Errors()) {
			return nil, false
		}
		printParserErrors(out, p.Errors())
		return nil, true
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated == nil {
		return nil, true
	}

	if returnValue, ok := evaluated.(*object.ReturnValue); ok {
		evaluated = returnValue.Value
	}

	return evaluated, true
}

func isIncompleteParse(errs []string) bool {
	for _, err := range errs {
		if strings.Contains(strings.ToLower(err), "unexpected end") ||
			strings.Contains(strings.ToLower(err), "unexpected eof") ||
			strings.Contains(strings.ToLower(err), "incomplete") {
			return true
		}
	}
	return false
}

func printParserErrors(out io.Writer, errors []string) {
	// fmt.Fprint(out, ODINS_EYE)
	io.WriteString(out, "Sorry Friend! Odin's eye sees all and you seem to have errors.\n")
	io.WriteString(out, "Parser Errors:\n")
	for _, msg := range errors {
		fmt.Fprintf(out, "\t%s\n", msg)
	}
}

func clearScreen(out io.Writer) {
	fmt.Fprint(out, "\033[H\033[2J")
}

func ProcessFile(filePath string, out io.Writer, env *object.Environment) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	// Tokenize, parse, and evaluate the file contents
	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		printParserErrors(out, p.Errors())
		return fmt.Errorf("file %s contains syntax errors", filePath)
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil && evaluated.Type() != object.NONE_OBJ {
		fmt.Fprintf(out, "%s\n", evaluated.Inspect())
	}
	return nil
}
