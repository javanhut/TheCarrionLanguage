package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"thecarrionlang/evaluator"
	"thecarrionlang/lexer"
	"thecarrionlang/object"
	"thecarrionlang/parser"
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

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	var inputBuffer strings.Builder
	isMultiline := false
	currentIndentLevel := 0
	expectedIndentLevel := 0

	fmt.Fprintln(out, "Welcome to the Carrion Programming Language REPL!")
	fmt.Fprintln(out, "Type 'exit' or 'quit' to exit, 'clear' to clear the screen.")

	for {
		if !isMultiline {
			fmt.Fprint(out, ">>> ")
		} else {
			fmt.Fprint(out, "... ")
		}

		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)

		// Handle special commands only at the primary prompt
		if !isMultiline {
			switch trimmedLine {
			case "exit", "quit":
				fmt.Fprintln(out, "Farewell, brave coder!")
				return
			case "clear":
				clearScreen(out)
				continue
			case "":
				continue
			}
		}

		// Count leading spaces to determine indentation level
		indentSpaces := len(line) - len(strings.TrimLeft(line, " "))
		currentIndentLevel = indentSpaces / 4 // Assuming 4 spaces per indent level

		// Skip empty lines in multiline mode
		if isMultiline && trimmedLine == "" {
			inputBuffer.WriteString(line)
			inputBuffer.WriteString("\n")
			continue
		}

		// Validate indentation if we're in multiline mode
		if isMultiline && trimmedLine != "" {
			if currentIndentLevel < expectedIndentLevel {
				// Check if this is an 'otherwise' or 'else' statement
				if strings.HasPrefix(trimmedLine, "otherwise") ||
					strings.HasPrefix(trimmedLine, "else") {
					expectedIndentLevel = currentIndentLevel
				} else if currentIndentLevel == 0 {
					// If we're back to no indentation, try to evaluate
					isMultiline = false
				} else {
					fmt.Fprintln(out, "IndentationError: expected an indented block")
					inputBuffer.Reset()
					isMultiline = false
					continue
				}
			}
		}

		// Check if this line starts a new block
		if strings.HasSuffix(trimmedLine, ":") {
			isMultiline = true
			expectedIndentLevel = currentIndentLevel + 1
		}

		// Append the line to our input buffer
		inputBuffer.WriteString(line)
		inputBuffer.WriteString("\n")

		// Try to evaluate if we're not in multiline mode or we've returned to base indentation
		if !isMultiline || (currentIndentLevel == 0 && !strings.HasSuffix(trimmedLine, ":")) {
			input := inputBuffer.String()
			if strings.TrimSpace(input) == "" {
				inputBuffer.Reset()
				continue
			}

			evaluated, complete := tryParseAndEval(input, out)
			if complete {
				if evaluated != nil {
					fmt.Fprintf(out, "%s\n", evaluated.Inspect())
				}
				inputBuffer.Reset()
				isMultiline = false
				expectedIndentLevel = 0
			} else {
				// If parsing is incomplete, continue in multiline mode
				isMultiline = true
			}
		}
	}
}

func tryParseAndEval(input string, out io.Writer) (object.Object, bool) {
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

	evaluated := evaluator.Eval(program)
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
	fmt.Fprint(out, ODINS_EYE)
	io.WriteString(out, "Sorry Friend! Odin's eye sees all and you seem to have errors.\n")
	io.WriteString(out, "Parser Errors:\n")
	for _, msg := range errors {
		fmt.Fprintf(out, "\t%s\n", msg)
	}
}

func clearScreen(out io.Writer) {
	fmt.Fprint(out, "\033[H\033[2J")
}

