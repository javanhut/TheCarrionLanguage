package evaluator

import (
	"fmt"
	"path/filepath"

	"thecarrionlanguage/src/lexer"
	"thecarrionlanguage/src/munin"
	"thecarrionlanguage/src/object"
	"thecarrionlanguage/src/parser"
)

// LoadMuninStdlib loads all .crl files from the given directory into the provided env.
func LoadMuninStdlib(env *object.Environment) error {
	// Read all items in the muninDir
	entries, err := munin.MuninFs.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read munin dir: %w", err)
	}

	// Iterate over each item, looking for .crl files
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".crl" {
			// Read the .crl file
			content, err := munin.MuninFs.ReadFile(entry.Name())
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", entry.Name(), err)
			}

			// Lex & parse
			l := lexer.New(string(content))
			p := parser.New(l)
			program := p.ParseProgram()

			// Check parse errors
			if len(p.Errors()) > 0 {
				return fmt.Errorf("parse errors in %s: %v", entry.Name(), p.Errors())
			}

			// Evaluate the program in the global environment
			result := Eval(program, env)
			// If evaluation produced an error object, return that
			if isError(result) {
				return fmt.Errorf("runtime error in %s: %s", entry.Name(), result.Inspect())
			}
		}
	}

	return nil
}
