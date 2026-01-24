package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/javanhut/TheCarrionLanguage/src/debug"
	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/repl"
)

var versionNum string = "0.1.9"

const CROW_IMAGE = `
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⡟⠋⢻⣷⣄⡀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣾⣿⣷⣿⣿⣿⣿⣿⣶⣾⣿⣿⠿⠿⠿⠶⠄⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠉⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⡟⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⣿⣿⣿⣿⣿⠟⠻⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⣿⣿⣿⣿⣿⣆⣤⠿⢶⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠀⠑⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠸⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠉⠙⠛⠋⠉⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
  `

func main() {
	// Define command line flags
	version := flag.Bool("version", false, "Prints out the Current version of the Carrion Language")
	shortVersion := flag.Bool("v", false, "Prints out the current Carrion Version (short from)")
	idebug := flag.Bool("idebug", false, "Enable interpreter debugging")
	id := flag.Bool("id", false, "Enable interpreter debugging (short form)")
	lexerDebug := flag.Bool("lexer", false, "Enable lexer debugging (use with --idebug)")
	parserDebug := flag.Bool("parser", false, "Enable parser debugging (use with --idebug)")
	evaluatorDebug := flag.Bool("evaluator", false, "Enable evaluator debugging (use with --idebug)")
	allDebug := flag.Bool("all", false, "Enable all debugging outputs (use with --idebug)")

	flag.Parse()

	// Create debug configuration
	debugConfig := debug.NewConfig()
	if *idebug || *id {
		debugConfig.Enabled = true
		if *allDebug {
			debugConfig.EnableAll()
		} else {
			debugConfig.Lexer = *lexerDebug
			debugConfig.Parser = *parserDebug
			debugConfig.Evaluator = *evaluatorDebug
			// If no specific debug flag is set, enable all
			if !*lexerDebug && !*parserDebug && !*evaluatorDebug {
				debugConfig.EnableAll()
			}
		}
	}

	// Create a global environment
	env := object.NewEnvironment()
	env.SetDebugConfig(debugConfig)

	// Print out the version of Carrion Lang
	if *version || *shortVersion {
		print := fmt.Println
		versionInfo := fmt.Sprintf("Carrion Language version %v", versionNum)
		print(versionInfo)
		return
	}

	// Attempt to load the standard library
	if err := evaluator.LoadMuninStdlib(env); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load stdlib: %v\n", err)
		os.Exit(1)
	}

	// Get non-flag arguments
	args := flag.Args()

	if len(args) > 0 {
		filePath := args[0]
		if strings.HasSuffix(filePath, ".crl") {
			err := repl.ProcessFileWithDebug(filePath, os.Stdout, env, debugConfig)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
		} else if filePath == "version" || filePath == "v" {
			fmt.Fprintln(os.Stderr, "Use --version or -v flag to display version information.")
			os.Exit(1)

		} else {
			fmt.Fprintln(os.Stderr, "Unsupported file type. Only .crl files are allowed.")
			os.Exit(1)
		}
	} else {
		fmt.Printf("%s\n", CROW_IMAGE)
		repl.StartWithDebug(os.Stdin, os.Stdout, env, debugConfig)
	}
}
