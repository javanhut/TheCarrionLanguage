package evaluator

import (
	"fmt"
	"path/filepath"

	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/modules"
	"github.com/javanhut/TheCarrionLanguage/src/munin"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
)

// LoadModules loads built-in functions from the time module into the provided environment.
func LoadModules(env *object.Environment) {
	// Load time module functions into the environment
	for name, builtin := range modules.TimeModule {
		env.Set(name, builtin)
	}
}

// LoadMuninStdlib loads the standard library into the provided environment.
//
// It first loads Go-based module functions, then reads and evaluates all embedded `.crl` standard library files into the environment. If any file fails to parse or evaluate, an error is returned. On success, the environment is set as the global standard library environment for built-in functions.
//
// Returns an error if reading, parsing, or evaluating any standard library file fails.
func LoadMuninStdlib(env *object.Environment) error {
	// Load Go modules first
	LoadModules(env)
	// 1. List embedded files in the current directory (".")
	//    if you used //go:embed *.crl with no subdirectory
	entries, err := munin.MuninFs.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read embedded stdlib: %w", err)
	}

	// 2. Load each .crl file
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".crl" {
			// 3. Read the file’s content
			content, err := munin.MuninFs.ReadFile(entry.Name())
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", entry.Name(), err)
			}

			// 4. Lex & parse the content
			l := lexer.New(string(content))
			p := parser.New(l)
			program := p.ParseProgram()

			// 5. Check for parse errors
			if len(p.Errors()) > 0 {
				fmt.Printf("Parse errors in %s:\n", entry.Name())
				for i, err := range p.Errors() {
					fmt.Printf("  Error %d: %s\n", i+1, err)
				}
				return fmt.Errorf("parse errors in %s: %v", entry.Name(), p.Errors())
			}

			// 6. Evaluate in the global environment
			result := Eval(program, env, nil)
			if isError(result) {
				return fmt.Errorf("runtime error in %s: %s", entry.Name(), result.Inspect())
			}
		}
	}

	// Set the global reference to the stdlib environment for builtin functions
	SetStdlibEnv(env)

	return nil
}
