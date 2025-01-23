package evaluator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"thecarrionlanguage/lexer"
	"thecarrionlanguage/object"
	"thecarrionlanguage/parser"
)

// LoadMuninStdlib loads all .crl files from the given directory into the provided env.
func LoadMuninStdlib(env *object.Environment, muninDir string) error {
	// Read all items in the muninDir
	entries, err := ioutil.ReadDir(muninDir)
	if err != nil {
		return fmt.Errorf("failed to read munin dir: %w", err)
	}

	// Iterate over each item, looking for .crl files
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".crl" {
			filePath := filepath.Join(muninDir, entry.Name())
			// Read the .crl file
			content, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", filePath, err)
			}

			// Lex & parse
			l := lexer.New(string(content))
			p := parser.New(l)
			program := p.ParseProgram()

			// Check parse errors
			if len(p.Errors()) > 0 {
				return fmt.Errorf("parse errors in %s: %v", filePath, p.Errors())
			}

			// Evaluate the program in the global environment
			result := Eval(program, env)
			// If evaluation produced an error object, return that
			if isError(result) {
				return fmt.Errorf("runtime error in %s: %s", filePath, result.Inspect())
			}
		}
	}

	return nil
}
