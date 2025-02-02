package evaluator

import (
	"fmt"
	"path/filepath"

	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/munin"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
)

func LoadMuninStdlib(env *object.Environment) error {
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
				return fmt.Errorf("parse errors in %s: %v", entry.Name(), p.Errors())
			}

			// 6. Evaluate in the global environment
			result := Eval(program, env)
			if isError(result) {
				return fmt.Errorf("runtime error in %s: %s", entry.Name(), result.Inspect())
			}
		}
	}

	return nil
}
