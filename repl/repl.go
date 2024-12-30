package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"thecarrionlang/evaluator"
	"thecarrionlang/lexer"
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
const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := strings.TrimSpace(scanner.Text())

		// Handle special commands
		switch line {
		case "clear":
			clearScreen()
			continue
		case "exit":
			fmt.Fprintln(out, "Terminated!")
			return
		case "quit":
			fmt.Fprintln(out, "Quit Application!")
			return
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, ODINS_EYE)
	io.WriteString(out, "Sorry Friend! Odin's eye sees all and you seem to have errors.\n")
	io.WriteString(out, "Error in the parser: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
