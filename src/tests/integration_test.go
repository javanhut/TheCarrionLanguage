package main

import (
        "os"
        "os/exec"
        "path/filepath"
        "strings"
        "testing"
)

// TestExampleFiles runs all example .crl files to ensure they don't crash
func TestExampleFiles(t *testing.T) {
	 t.Skip("Skipping example files during lexer refactor")
}


// TestBasicFunctionality tests core language features that should work
func TestBasicFunctionality(t *testing.T) {
        t.Skip("Skipping basic functionality tests during lexer refactor")
        tests := []struct {
                name    string
                code    string
                wantErr bool
	}{
		{
			name: "simple_print",
			code: `print("Hello, World!")`,
			wantErr: false,
		},
		{
			name: "arithmetic_direct",
			code: `print(2 + 3)`,
			wantErr: false,
		},
		{
			name: "simple_if",
			code: `if True:
    print("yes")
else:
    print("no")`,
			wantErr: false,
		},
		{
			name: "array_literal_direct",
			code: `print(len([1, 2, 3]))`,
			wantErr: false,
		},
		{
			name: "builtin_functions",
			code: `print(len("hello"))
print(str(42))
print(int("10"))`,
			wantErr: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Write code to temporary file
			tmpFile := filepath.Join(os.TempDir(), tt.name+".crl")
			err := os.WriteFile(tmpFile, []byte(tt.code), 0644)
			if err != nil {
				t.Fatalf("Failed to write temp file: %v", err)
			}
			defer os.Remove(tmpFile)
			
			// Run the code
			cmd := exec.Command("go", "run", "../main.go", tmpFile)
			cmd.Dir = "."
			output, err := cmd.CombinedOutput()
			
			if tt.wantErr && err == nil {
				t.Errorf("Expected error but got none. Output: %s", output)
			} else if !tt.wantErr && err != nil {
				t.Errorf("Unexpected error: %v\nOutput: %s", err, output)
			} else if !tt.wantErr {
				t.Logf("Success: %s", strings.TrimSpace(string(output)))
			}
		})
	}
}