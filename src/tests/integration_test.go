package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// findProjectRoot finds the project root by looking for go.mod
func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found")
		}
		dir = parent
	}
}

// TestExampleFiles runs all example .crl files to ensure they don't crash
func TestExampleFiles(t *testing.T) {
	projectRoot, err := findProjectRoot()
	if err != nil {
		t.Fatalf("Failed to find project root: %v", err)
	}

	examplesDir := filepath.Join(projectRoot, "examples")
	var files []string
	files, err = filepath.Glob(filepath.Join(examplesDir, "*.crl"))
	if err != nil {
		t.Fatalf("Failed to list example files: %v", err)
	}

	if len(files) == 0 {
		t.Skip("No example files found in any expected location")
		return
	}

	// Files that are expected to fail (contain invalid syntax or test error conditions)
	expectedFailures := map[string]bool{
		"test_cannot_declare_priv.crl": true,  // Tests access control errors
		"test_custom_error.crl":        true,  // Tests error raising
		"test_throw_error.crl":         true,  // Tests error throwing
		"test_raise.crl":               true,  // Tests error raising
		"failing_with_call.crl":        true,  // Intentionally fails assertion
	}

	// Files to skip (interactive, require external resources, have syntax issues)
	skipFiles := map[string]bool{
		"http_kv_store_demo.crl":     true,  // Starts HTTP server
		"http_rest_api_demo.crl":     true,  // Starts HTTP server
		"quick_webserver_demo.crl":   true,  // Starts HTTP server
		"webserver_demo.crl":         true,  // Starts HTTP server
		"simple_check_test.crl":      true,  // Has syntax errors (check function)
		"test_return.crl":            true,  // Has syntax errors (check function)
	}
	
	passed := 0
	failed := 0
	skipped := 0
	
	for _, file := range files {
		filename := filepath.Base(file)
		
		t.Run(filename, func(t *testing.T) {
			if skipFiles[filename] {
				t.Skipf("Skipping %s - starts server or requires external resources", filename)
				skipped++
				return
			}
			
			// Run the Carrion interpreter on the file
			mainPath := filepath.Join(projectRoot, "src", "main.go")
			cmd := exec.Command("go", "run", mainPath, file)
			cmd.Dir = projectRoot
			output, err := cmd.CombinedOutput()
			
			if expectedFailures[filename] {
				if err == nil {
					t.Errorf("Expected %s to fail, but it succeeded", filename)
					failed++
				} else {
					t.Logf("Expected failure for %s: %v", filename, err)
					passed++
				}
			} else {
				if err != nil {
					t.Errorf("Failed to run %s: %v\nOutput: %s", filename, err, output)
					failed++
				} else {
					t.Logf("Successfully ran %s", filename)
					passed++
				}
			}
		})
	}
	
	// Print summary
	total := passed + failed
	if total > 0 {
		fmt.Printf("\nIntegration Test Summary:\n")
		fmt.Printf("Passed: %d\n", passed)
		fmt.Printf("Failed: %d\n", failed) 
		fmt.Printf("Skipped: %d\n", skipped)
		fmt.Printf("Total: %d\n", len(files))
		fmt.Printf("Success Rate: %.1f%%\n", float64(passed)/float64(total)*100)
	}
}

// TestBasicFunctionality tests core language features that should work
func TestBasicFunctionality(t *testing.T) {
	projectRoot, err := findProjectRoot()
	if err != nil {
		t.Fatalf("Failed to find project root: %v", err)
	}

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
			mainPath := filepath.Join(projectRoot, "src", "main.go")
			cmd := exec.Command("go", "run", mainPath, tmpFile)
			cmd.Dir = projectRoot
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