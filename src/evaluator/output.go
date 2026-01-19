package evaluator

import (
	"io"
	"os"
)

// outputWriter is the writer used for print output
// Defaults to os.Stdout but can be changed for WASM
var outputWriter io.Writer = os.Stdout

// SetOutputWriter sets the output writer for print statements
// This is used by WASM to capture output
func SetOutputWriter(w io.Writer) {
	outputWriter = w
}

// GetOutputWriter returns the current output writer
func GetOutputWriter() io.Writer {
	return outputWriter
}

// ResetOutputWriter resets the output writer to stdout
func ResetOutputWriter() {
	outputWriter = os.Stdout
}
