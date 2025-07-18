package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sort"

	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/ast"
)

type TestResult struct {
	FunctionName string
	Passed       bool
	ErrorMessage string
}

type FileTestResult struct {
	FilePath     string
	RelativePath string
	Tests        []TestResult
	Passed       int
	Failed       int
}

type TestRunner struct {
	env      *object.Environment
	detailed bool
	results  map[string][]TestResult
}

func main() {
	var detailed bool
	flag.BoolVar(&detailed, "d", false, "Show detailed test output")
	flag.BoolVar(&detailed, "detailed", false, "Show detailed test output")
	
	// Custom usage function
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: sindri appraise [options] [path]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  sindri appraise                    # Run all test files in current directory\n")
		fmt.Fprintf(os.Stderr, "  sindri appraise test.crl           # Run specific test file\n")
		fmt.Fprintf(os.Stderr, "  sindri appraise ./appraise         # Run all test files in directory\n")
		fmt.Fprintf(os.Stderr, "  sindri appraise -d test.crl        # Run with detailed output\n")
	}
	
	// Parse command line
	args := os.Args[1:]
	if len(args) < 1 || args[0] != "appraise" {
		flag.Usage()
		os.Exit(1)
	}
	
	// Parse flags after "appraise"
	flag.CommandLine.Parse(args[1:])
	
	// Get the path argument (if any)
	pathArg := ""
	if flag.NArg() > 0 {
		pathArg = flag.Arg(0)
	}
	
	runner := NewTestRunner(detailed)
	
	// Find test files
	testFiles, err := findTestFiles(pathArg)
	if err != nil {
		fmt.Printf("Error finding test files: %v\n", err)
		os.Exit(1)
	}
	
	if len(testFiles) == 0 {
		fmt.Println("No test files found")
		os.Exit(0)
	}
	
	// Run tests on all files
	results := runner.RunMultipleFiles(testFiles)
	
	// Display results
	if detailed || len(testFiles) == 1 {
		displayDetailedResults(results)
	} else {
		displaySummaryResults(results)
	}
	
	// Exit with error if any tests failed
	for _, result := range results {
		if result.Failed > 0 {
			os.Exit(1)
		}
	}
}

func NewTestRunner(detailed bool) *TestRunner {
	return &TestRunner{
		env:      object.NewEnvironment(),
		detailed: detailed,
		results:  make(map[string][]TestResult),
	}
}

func findTestFiles(pathArg string) ([]string, error) {
	var testFiles []string
	
	if pathArg == "" {
		// No path specified, search current directory
		pathArg = "."
	}
	
	info, err := os.Stat(pathArg)
	if err != nil {
		return nil, err
	}
	
	if info.IsDir() {
		// Search directory for test files
		err := filepath.Walk(pathArg, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			
			if !info.IsDir() && strings.HasSuffix(path, ".crl") {
				// Check if filename contains "appraise"
				basename := filepath.Base(path)
				if strings.Contains(basename, "appraise") {
					testFiles = append(testFiles, path)
				}
			}
			
			return nil
		})
		
		if err != nil {
			return nil, err
		}
	} else {
		// Single file specified
		if !strings.HasSuffix(pathArg, ".crl") {
			return nil, fmt.Errorf("test file must have .crl extension")
		}
		testFiles = append(testFiles, pathArg)
	}
	
	// Sort files for consistent output
	sort.Strings(testFiles)
	
	return testFiles, nil
}

func (tr *TestRunner) RunMultipleFiles(files []string) []FileTestResult {
	var results []FileTestResult
	
	for _, file := range files {
		result := tr.RunSingleFile(file)
		results = append(results, result)
	}
	
	return results
}

func (tr *TestRunner) RunSingleFile(filename string) FileTestResult {
	result := FileTestResult{
		FilePath:     filename,
		RelativePath: getRelativePath(filename),
		Tests:        []TestResult{},
		Passed:       0,
		Failed:       0,
	}
	
	content, err := os.ReadFile(filename)
	if err != nil {
		result.Tests = append(result.Tests, TestResult{
			FunctionName: "FILE_READ_ERROR",
			Passed:       false,
			ErrorMessage: fmt.Sprintf("Failed to read file: %v", err),
		})
		result.Failed++
		return result
	}
	
	l := lexer.New(string(content))
	p := parser.New(l)
	program := p.ParseProgram()
	
	if len(p.Errors()) > 0 {
		result.Tests = append(result.Tests, TestResult{
			FunctionName: "PARSE_ERROR",
			Passed:       false,
			ErrorMessage: strings.Join(p.Errors(), "; "),
		})
		result.Failed++
		return result
	}
	
	// Find all appraise functions
	appraiseFunctions := tr.findAppraiseFunctions(program)
	
	if len(appraiseFunctions) == 0 {
		return result
	}
	
	// Execute the program to define all functions
	ctx := &evaluator.CallContext{
		FunctionName: "main",
		IsDirectExecution: true,
	}
	
	// Create a fresh environment for this file
	fileEnv := object.NewEnvironment()
	evaluator.Eval(program, fileEnv, ctx)
	
	// Run each appraise function
	for _, funcName := range appraiseFunctions {
		testResult := tr.runSingleTest(funcName, fileEnv)
		result.Tests = append(result.Tests, testResult)
		
		if testResult.Passed {
			result.Passed++
		} else {
			result.Failed++
		}
	}
	
	return result
}

func (tr *TestRunner) findAppraiseFunctions(program *ast.Program) []string {
	var appraiseFunctions []string
	
	for _, stmt := range program.Statements {
		// Check for function definitions
		if funcDef, ok := stmt.(*ast.FunctionDefinition); ok {
			name := funcDef.Name.Value
			if strings.Contains(name, "appraise") {
				appraiseFunctions = append(appraiseFunctions, name)
			}
		}
		
		// Check for grimoire methods that contain appraise
		if grimoireDef, ok := stmt.(*ast.GrimoireDefinition); ok {
			for _, method := range grimoireDef.Methods {
				name := method.Name.Value
				if strings.Contains(name, "appraise") {
					// For grimoire methods, we need the full path
					methodName := fmt.Sprintf("%s.%s", grimoireDef.Name.Value, name)
					appraiseFunctions = append(appraiseFunctions, methodName)
				}
			}
		}
	}
	
	return appraiseFunctions
}

func (tr *TestRunner) runSingleTest(funcName string, env *object.Environment) TestResult {
	result := TestResult{
		FunctionName: funcName,
		Passed:       false,
		ErrorMessage: "",
	}
	
	// Handle grimoire methods vs standalone functions
	var evalResult object.Object
	ctx := &evaluator.CallContext{
		FunctionName: funcName,
	}
	
	if strings.Contains(funcName, ".") {
		// For now, grimoire method testing is not supported
		evalResult = &object.Error{Message: "Grimoire method testing not yet supported"}
	} else {
		// Standalone function - create a call expression and evaluate it
		callExpr := &ast.CallExpression{
			Function: &ast.Identifier{Value: funcName},
			Arguments: []ast.Expression{},
		}
		ctx.Node = callExpr
		evalResult = evaluator.Eval(callExpr, env, ctx)
	}
	
	// Check if the test passed
	switch errorObj := evalResult.(type) {
	case *object.Error:
		result.ErrorMessage = errorObj.Message
	case *object.ErrorWithTrace:
		result.ErrorMessage = errorObj.Message
	default:
		result.Passed = true
	}
	
	return result
}

func getRelativePath(filePath string) string {
	cwd, err := os.Getwd()
	if err != nil {
		return filePath
	}
	
	relPath, err := filepath.Rel(cwd, filePath)
	if err != nil {
		return filePath
	}
	
	// If the relative path goes up directories, use absolute path
	if strings.HasPrefix(relPath, "..") {
		absPath, err := filepath.Abs(filePath)
		if err != nil {
			return filePath
		}
		return absPath
	}
	
	return relPath
}

func displayDetailedResults(results []FileTestResult) {
	for _, fileResult := range results {
		fmt.Printf("\nRunning tests in %s...\n", fileResult.RelativePath)
		
		if len(fileResult.Tests) == 0 {
			fmt.Println("No test functions found")
			continue
		}
		
		fmt.Printf("Found %d test function(s)\n\n", len(fileResult.Tests))
		
		for _, test := range fileResult.Tests {
			if test.Passed {
				fmt.Printf("Running %s... \033[32mPASSED\033[0m\n", test.FunctionName)
			} else {
				fmt.Printf("Running %s... \033[31mFAILED\033[0m", test.FunctionName)
				if test.ErrorMessage != "" {
					fmt.Printf(" - %s", test.ErrorMessage)
				}
				fmt.Println()
			}
		}
		
		fmt.Printf("\n=== Test Summary for %s ===\n", fileResult.RelativePath)
		fmt.Printf("\033[32mPassed: %d\033[0m\n", fileResult.Passed)
		fmt.Printf("\033[31mFailed: %d\033[0m\n", fileResult.Failed)
		fmt.Printf("Total:  %d\n", fileResult.Passed + fileResult.Failed)
		
		if fileResult.Failed > 0 {
			fmt.Printf("\n\033[31mSome tests failed!\033[0m\n")
		} else if fileResult.Passed > 0 {
			fmt.Printf("\n\033[32mAll tests passed!\033[0m\n")
		}
	}
}

func displaySummaryResults(results []FileTestResult) {
	totalPassed := 0
	totalFailed := 0
	
	fmt.Println("\n=== Test Summary ===")
	
	for _, fileResult := range results {
		if len(fileResult.Tests) == 0 {
			continue
		}
		
		status := "\033[32mPASSED\033[0m"
		if fileResult.Failed > 0 {
			status = "\033[31mFAILED\033[0m"
		}
		
		fmt.Printf("%-50s - %s\n", fileResult.RelativePath, status)
		
		totalPassed += fileResult.Passed
		totalFailed += fileResult.Failed
	}
	
	fmt.Printf("\n\033[32mPassed: %d/%d\033[0m\n", totalPassed, totalPassed+totalFailed)
	fmt.Printf("\033[31mFailed: %d/%d\033[0m\n", totalFailed, totalPassed+totalFailed)
	
	if totalFailed > 0 {
		fmt.Printf("\n\033[31mSome tests failed! Use -d or --detailed for more information.\033[0m\n")
	} else if totalPassed > 0 {
		fmt.Printf("\n\033[32mAll tests passed!\033[0m\n")
	}
}