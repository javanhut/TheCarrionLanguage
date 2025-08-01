package main

import (
	"flag"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/javanhut/TheCarrionLanguage/src/ast"
	"github.com/javanhut/TheCarrionLanguage/src/evaluator"
	"github.com/javanhut/TheCarrionLanguage/src/lexer"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"github.com/javanhut/TheCarrionLanguage/src/parser"
)

type TestResult struct {
	FunctionName string
	Passed       bool
	ErrorMessage string
	Duration     time.Duration
	StartTime    time.Time
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
	timeout  time.Duration
}

const hammer = `
	@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#*%@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@#*%@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%*%@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%*%@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%*%@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%+%@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@+%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%#####%@%@@@@@@@@@@@@@@@@@@@@@%+%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%#*#%%%@@%%%#%%%%@@@@@@@@@@@@@@@@@+%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%*#@@%%@@@@@@@@@@@@@@@@@@@@@@@@@@@@*%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%**%@@@@@@%@%@@@@@@@@@%%@@@@@@@@@@@*%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@*#%@@@@@@@@@@@%%%%@@@@@@@@@@@@@@@@@*%@%%**#%@@@@@@@@@@@@@%%@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@##%%@@@@@@@%@%###%%@@@%@@@@@@@@@@%**#%%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@*+%%%%%@@@%%#@@@#%%*****%%%%%@%%@@@@@@@*%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@*+=*##%%#@@#%%@@%#%%%@@@@%##****#%#*###@@@@@@#%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%%@
@@@@@@@@@@@@@@@@@@@@@@@@@#*+*#%@@@%@@@@%@%%%@@@@%%##@@@%######++-+#*#%%@%#@@%*@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@%++=+#%@@@@@@@@@@@@@@@@##%@@#*#@@@#*#*****+=--:-::=%@#@@*@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@##%@@@@@@@@@@@@@@@@%#%@@@%##*%@@@@@###*+*%@%-=-=**=-#@@@@%#@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@*+#@@@@@@@@@@@@@@@@@%%#%@@%#*%@@@@@@%****#**#%*==+-:=%@@@@*%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@####@@@@@@@@@@@@@@@@@@###@@@@@@@@@%%@%*=:.-+*%@@%*#%@@@@@@@*@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@#*%@@@@@@@@@@@@@@@@@@@@#**#@@@@@@@@@@@@%%#+::=###*-**@@@@@@@*%@@@@@@@@@@%#######**###%@@@%%@@@@@@@@@@@
@@@@@@@@@@@@@@*#@@@@@@@@@@@@@@@@@@%@@@%***@@@@@@@@%%@@@%%#%%%+**++*@@@@@@@@@%%%%%%%%%%%%%%%%%%%%%%%%%%####+=#@@@@@@@@
@@@@@@@@@@@@#=+#@@@@@@@@@@@@%%%@@@@%%%@**+%@@@@@@%@@%@@@@@%#%%%--%%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%@%#%@@%%%#*+%@@@@
@@@@@@@@@@@=%@@@@@@%%@@@@@@@@*#@@@@**#@#*=%@@@@@@@@@@@@%*###+*@##@@@@@@%%@@@@@@@@@@@@@@@@@@@@@@@@@@%@@%@%@%%#%%@#+#@@
@@@@@@@@@@+@@@@@@@@@@@@@@@%@@%+%@%@**%@#*=#@@@@%%%@@@%%#%=+#*-*@##%@@%%@#@@@@@@@@@@@@@@@@@@@@@@@@@@@%@@@@*@@%#%#%@#=%
@@@@@@@@@+#@@@@@@@@@@%%@@@%##%+*@*%*%@@%@%@@@@@@@@@%%%#@@@@@@@@@%%%@@%#@%@@@@@@@@@@@@@@@@@@@@%@@@@@@%@@@**@@%%%*#%@@+
@@@@@@@@*@@%@@@@@@%%@@@@#@%**##+###%@%##%%=%#%@@@@@%#%#@@@%%#%%%@%%@#*%@%#@@@@@@@@@@@@@@@@@@##@@@@@@@@@#+%%#%##*+#@@@
@@@@@@@#+#@@@@@@@@@@@##@@*##+#%*@%@%@##%+**#++*%@@@##*@%%#*#***+%%#%+#@@*@%@@@@@@@@@@@@@@@@@%%@@@@%%%%@#+##+#***#%@@@
@@@@@@*#@@@@@@@%#*#%#*@%*@##*#%@@@%@@**####+++*%@@@#@##%*+***+*#%#@#%@%*@@#@@@@@@@@@@@@@@@@@@@@@@*@@%##@==+==-+*#%#@@
@@@@@%%@@@@@@%@@@@%*+++*%%*##@@@#%@@#+++++++++*@@@@%%*#%%%@@%#%%%@%@@#%@*+##@@@@@@@@@@@@@@@@@@%@@%@%***#+=+--=*###%%@
@@@@@##@@@@@@@%%@@@@%%@@@#*%@@@@@@@%+**++++++#%%%@@@@#%@%%%%@%@@@@@@%@##%@@@#@@@@@@@@@@@@@%%@@@#@@@+=**+=-=+++#*##@@@
@@@@%*@@@@@@@@%@@@@%@@@@#*--#@@@@@#+#**++++*##*+#@@@@%%%##%#%*%#@@@@#+*%%#++#@@@@@@@@@@@@@@@@@@@%@%====+---++*%*++%%@
@@@@%%@@@@@@@@@@%%@@@#***+-..=@@@*%%**+***+++++++%@@%@@%###%**@*@%#%@@#****#%%@@@@@@@@@@@##@@@@##%%+-:---===-+*#=-%#@
@@@@@%%@@@@@@@%@@@%##%*+##*++=:*%####%#*+++++++++*@@@%@@%%%###*%@@@@@@@##*##@**@@@@@@@@@%*@@%@#*%##*=..:-==-::**=:+##
@@@@@@%#@@%@@@@@@%##*##*****+++-:###*#@@@@#**++*#%%@@@%%%%%#%%@@@@@@@#*****#**%@@@@@@#%@@@@@@@++##**=. .-=-..:++=-=*#
@@@@@@@@#@#@@@@@####%****#*******=:=%%*++++++**#%@%%@@@@@%@@%@@@@@###%%##+----*%@@@@@@@@@@@@#%+=*=-:.  :-...:=--==+*+
@@@@@@@@@@#%@@@%#%%%#*###**####*+**+==**++==-::::-=+%@@%%%**%@#@%@%####**+-:-=-*@@@@@@@@#%@%#@*:-:-.   :. .:=-:---==+
@@@@@@@@@@@*#@@%%@%%@%%%###%%#*##+=#+++--**+++*+*#%@%@@%#**%@@%%@%%%%##**+-:::.+@@@@@@##@##+%%+..=:        -:.=--=.=+
@@@@@@@@@@@@@@@@#*%%@@@@@@%%##%**#****+*=:.=*++**+**@@@%%%###@@#@%@%%%%%#**+#==%@@@@@@#%%+#+++- .-:.        ..==-::=#
@@@@@@@@@@@@@@@#+#@@@**##%@@%%#%*********+:.:-*++#*=@@@####*#@@@#@@@%%##*+=:=-:*@@@@#@@@@*=--+:.  -:         ..-:.-==
@@@@@@@@@@@@@@@%##%%@@@@@@@@%%%@*#****#*##=.:*+*#*++*@@@@@@@@@@@@%@@%#%%###*=*=:-%@@@#@@@%=-:-=.   .          ...-=:=
@@@@@@@@@@@@@@*#%@%@@@%@#@%%@@#@###*###-:*#=-***#+#**+####@@@@@@@@@%@@@%@@#*+#**-+@@@@@##*--:.=-               .--:==
@@@@@@@@@@@@@*%%%@%@@@@@###%#%%#@###+=#*=+#*=-@%*+##***%@@@@@@@@@@@#**+**=%#=-=#+=%@@@#++=::...=                ---=+
@@@@@@@@@@@@*%@%%@#@@@@@@@%%#%@%##+#*+*#*-**+=*#+=*-..+++=+++*##***##%%%#--#*--=+%%+=+##+--:   .                ..:-*
@@@@@@@@@@@**@@%@@%@@@%@@%%@@%##%%%##*+*#+:%***+*#++=-:.::::       .--+*+++++**++++*#####*-=: ..                .::-*
@@@@@@@@@@@+@@@@@@@%%@@%%@%##%@@@%%@@%*#%+%=###+**+%%%+*#*@@@@@@@%%#==-::::...    ..:::=====-    .  --..         ..:=
@@@@@@@@@@**@@@@@@@%%%@%%#@%#%####%%@@@%%%@@@@%*#%#*+**%**%@@@@@@%%#%%#*##*+++***###*+.    ..        ...:-*+      .:+
@@@@@@@@@@=%@@@@@@@%@#@@%#%@#*##***#**+++++++*#++****%@##%%#****+**+++#@%++*%%%@@@%*#*=-.         :..:.           :+#
@@@@@@@@@#+*@@@@@@@%%#%@%##%%*+*%#****###**+++***%#**+++**#*****+=-+#%@@@@@@@@%#*#%@@#+-:.                       :=+#
@@@@@@@@@++%@@@@@@@%%#%@%#**%%*++%@#***#%%%@@@@@@@@%**###%@@@%%####%##*##*+++++====+%#==-::                    :-+#%@
@@@@@@@@@++@@@@@@@@%%##@@#***%#+++*@%##***###%%%%#%####++*****##%@@@@@@@@@@%%#*++*#%@%%@###+=*: .  .   :.  :+++###%#%
@@@@@@@@##*@@@@@@@@%%##%@##*+#%#*++**%@%#**#%*=*@@@%@@@%###***++++*==+**++*#@@@@@@@@@@*+=+#%@@#+@#*#*++%*##%%#%%@@@@%
@@@@@@@@*#%@@@@@@@@%%###@%#**+##*##**##%@@@%@@@@@@@*@@@@@@@#++**+***#%%#*+=----=++**=+*%%##*+**==++=+*#%@*++*+===----
@@@@@@@#%#@@@@@@@@@%#%###%##**+###%%###%%%%%@@%%%@@#%#@#****%@@@%@%%##**++=====+++#%#*=--:-::::.-:::---:-++-:----===+
@@@@@@@*%#@@@@@@@@@%%%###%%####*#%%%%%%#######%%%@@%%#%@@@@@@@@@%@@@@@@@%%%*+*%@@@@%%%****************####*%##%%@@@@@
@@@@@@#%%@@@@%@@@@@##%###%%#######%%%%%%%%######%@@*@@@@@@@@@@@@@@@@@@@%@*%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@%@@@@@@@@@
@@@@@@%%#@@@@%@@@@@#####%%@%#######%%%@@@@@%%%%%@@@*%@@@@@@@@@@@@@@@@@@@%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@%#%@@#@@@@@@#####%%%%%########%@@@@@@@@@@@@@*%@@@@@@@@@@@@%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
	`

func main() {
	var detailed bool
	var report bool
	flag.BoolVar(&detailed, "d", false, "Show detailed test output")
	flag.BoolVar(&detailed, "detailed", false, "Show detailed test output")
	flag.BoolVar(&report, "r", false, "Generate HTML report")
	flag.BoolVar(&report, "report", false, "Generate HTML report")

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
		fmt.Fprintf(os.Stderr, "  sindri appraise -r test.crl        # Generate HTML report\n")
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
	if detailed {
		displayDetailedResults(results)
	} else {
		displaySummaryResults(results)
	}

	// Generate HTML report if requested
	if report {
		reportFile, err := generateHTMLReport(results)
		if err != nil {
			fmt.Printf("Error generating report: %v\n", err)
		} else {
			fmt.Printf("\nHTML report generated: %s\n", reportFile)
		}
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
		timeout:  30 * time.Second, // Default timeout of 30 seconds per test
	}
}

func findTestFiles(pathArg string) ([]string, error) {
	var testFiles []string
	testFileSet := make(map[string]bool) // To avoid duplicates

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
				shouldInclude := false

				// Check if filename contains "appraise"
				basename := filepath.Base(path)
				if strings.Contains(basename, "appraise") {
					shouldInclude = true
				}

				// Also check if we're in an "appraise" directory
				dirName := filepath.Base(filepath.Dir(path))
				if strings.Contains(dirName, "appraise") {
					shouldInclude = true
				}

				if shouldInclude && !testFileSet[path] {
					testFiles = append(testFiles, path)
					testFileSet[path] = true
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
		FunctionName:      "main",
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
	startTime := time.Now()
	result := TestResult{
		FunctionName: funcName,
		Passed:       false,
		ErrorMessage: "",
		StartTime:    startTime,
	}

	// Channel to receive the evaluation result
	type evalResultWithError struct {
		result object.Object
		err    error
	}
	resultChan := make(chan evalResultWithError, 1)

	// Run the test evaluation in a goroutine
	go func() {
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
				Function:  &ast.Identifier{Value: funcName},
				Arguments: []ast.Expression{},
			}
			ctx.Node = callExpr
			evalResult = evaluator.Eval(callExpr, env, ctx)
		}

		resultChan <- evalResultWithError{result: evalResult, err: nil}
	}()

	// Wait for either the result or timeout
	select {
	case evalRes := <-resultChan:
		// Test completed within timeout
		result.Duration = time.Since(startTime)

		// Check if the test passed
		switch errorObj := evalRes.result.(type) {
		case *object.Error:
			result.ErrorMessage = errorObj.Message
		case *object.ErrorWithTrace:
			result.ErrorMessage = errorObj.Message
		default:
			result.Passed = true
		}

	case <-time.After(tr.timeout):
		// Test timed out
		result.Duration = time.Since(startTime)
		result.ErrorMessage = fmt.Sprintf("Test execution timed out after %v", tr.timeout)
		result.Passed = false
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
		fmt.Printf("Total:  %d\n", fileResult.Passed+fileResult.Failed)

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

	// If there are multiple files, show file-based summary
	if len(results) > 1 {
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
	} else {
		// Single file, show individual test results
		for _, fileResult := range results {
			if len(fileResult.Tests) == 0 {
				continue
			}

			// Display each test function result
			for _, test := range fileResult.Tests {
				if test.Passed {
					fmt.Printf("Running %s... \033[32mPASSED\033[0m\n", test.FunctionName)
				} else {
					fmt.Printf("Running %s... \033[31mFAILED\033[0m\n", test.FunctionName)
				}
			}

			totalPassed += fileResult.Passed
			totalFailed += fileResult.Failed
		}
	}

	fmt.Printf("\n\033[32mPassed: %d/%d\033[0m\n", totalPassed, totalPassed+totalFailed)
	fmt.Printf("\033[31mFailed: %d/%d\033[0m\n", totalFailed, totalPassed+totalFailed)

	if totalFailed > 0 {
		fmt.Printf("\n\033[31mSome tests failed! Use -d or --detailed for more information.\033[0m\n")
	} else if totalPassed > 0 {
		fmt.Printf("\n\033[32mAll tests passed!\033[0m\n")
	}
}

func generateHTMLReport(results []FileTestResult) (string, error) {
	// Generate timestamp-based filename
	timestamp := time.Now().Format("20060102_150405")
	reportFile := fmt.Sprintf("sindri_report_%s.html", timestamp)

	// Calculate overall statistics
	totalTests := 0
	totalPassed := 0
	totalFailed := 0
	totalDuration := time.Duration(0)

	for _, fileResult := range results {
		totalTests += len(fileResult.Tests)
		totalPassed += fileResult.Passed
		totalFailed += fileResult.Failed
		for _, test := range fileResult.Tests {
			totalDuration += test.Duration
		}
	}

	// Create HTML report
	htmlContent := generateHTMLContent(results, totalTests, totalPassed, totalFailed, totalDuration)

	// Write to file
	err := os.WriteFile(reportFile, []byte(htmlContent), 0644)
	if err != nil {
		return "", err
	}

	return reportFile, nil
}

func generateHTMLContent(results []FileTestResult, totalTests, totalPassed, totalFailed int, totalDuration time.Duration) string {
	passRate := float64(0)
	if totalTests > 0 {
		passRate = float64(totalPassed) / float64(totalTests) * 100
	}

	// Build test results HTML
	var testResultsHTML strings.Builder
	for _, fileResult := range results {
		if len(fileResult.Tests) == 0 {
			continue
		}

		testResultsHTML.WriteString(fmt.Sprintf(`
		<div class="file-section">
			<h2 class="file-header">%s</h2>
			<div class="test-list">`, html.EscapeString(fileResult.RelativePath)))

		for _, test := range fileResult.Tests {
			statusClass := "passed"
			statusText := "PASSED"
			errorSection := ""

			if !test.Passed {
				statusClass = "failed"
				statusText = "FAILED"
				if test.ErrorMessage != "" {
					errorSection = fmt.Sprintf(`
					<div class="error-message">
						<strong>Error:</strong> %s
					</div>`, html.EscapeString(test.ErrorMessage))
				}
			}

			testResultsHTML.WriteString(fmt.Sprintf(`
				<div class="test-case %s">
					<div class="test-header">
						<span class="test-name">%s</span>
						<span class="test-status %s">%s</span>
						<span class="test-duration">%s</span>
					</div>
					%s
				</div>`, statusClass, html.EscapeString(test.FunctionName), statusClass, statusText, test.Duration, errorSection))
		}

		testResultsHTML.WriteString(`
			</div>
		</div>`)
	}

	// Generate the complete HTML
	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Sindri Test Report - %s</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background-color: #0a0a0a;
            color: #e5e5e5;
            line-height: 1.6;
        }
        
        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
        }
        
        header {
            background: linear-gradient(135deg, #1a1a1a 0%%, #2a1810 50%%, #3a1f10 100%%);
            background-image: linear-gradient(135deg, rgba(26,26,26,0.95) 0%%, rgba(42,24,16,0.95) 50%%, rgba(58,31,16,0.95) 100%%);
            color: #ffd700;
            padding: 40px 0;
            text-align: center;
            box-shadow: 0 4px 8px rgba(255,69,0,0.3);
            border-bottom: 3px solid #ff4500;
            position: relative;
        }
        
        .header-content {
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 30px;
            flex-wrap: wrap;
        }
        
        .logo {
            width: 120px;
            height: auto;
            filter: drop-shadow(0 4px 8px rgba(255,69,0,0.5));
        }
        
        h1 {
            font-size: 2.5em;
            margin-bottom: 10px;
            text-shadow: 2px 2px 4px rgba(255,69,0,0.5);
        }
        
        .subtitle {
            font-size: 1.1em;
            opacity: 0.9;
            color: #ffb347;
        }
        
        .summary {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
            gap: 20px;
            margin: 30px 0;
        }
        
        .summary-card {
            background: linear-gradient(135deg, #1a1a1a 0%%, #2a2a2a 100%%);
            border: 1px solid #3a3a3a;
            border-radius: 8px;
            padding: 25px;
            box-shadow: 0 4px 12px rgba(255,69,0,0.2);
            text-align: center;
            transition: all 0.3s;
        }
        
        .summary-card:hover {
            transform: translateY(-2px);
            box-shadow: 0 6px 20px rgba(255,69,0,0.3);
            border-color: #ff4500;
        }
        
        .summary-card h3 {
            color: #ffb347;
            font-size: 0.9em;
            text-transform: uppercase;
            letter-spacing: 1px;
            margin-bottom: 10px;
        }
        
        .summary-card .value {
            font-size: 2.5em;
            font-weight: bold;
            margin-bottom: 5px;
            text-shadow: 1px 1px 2px rgba(0,0,0,0.5);
        }
        
        .summary-card.total .value { color: #ffd700; }
        .summary-card.passed .value { color: #32cd32; }
        .summary-card.failed .value { color: #ff4500; }
        .summary-card.duration .value { color: #ff8c00; font-size: 2em; }
        .summary-card.rate .value { color: #ffa500; }
        
        .progress-bar {
            width: 100%%;
            height: 30px;
            background-color: #1a1a1a;
            border: 1px solid #3a3a3a;
            border-radius: 15px;
            overflow: hidden;
            margin: 30px 0;
            box-shadow: inset 0 2px 4px rgba(0,0,0,0.5), 0 2px 8px rgba(255,69,0,0.2);
        }
        
        .progress-fill {
            height: 100%%;
            background: linear-gradient(90deg, #228b22 0%%, #32cd32 50%%, #3cb371 100%%);
            transition: width 0.5s ease;
            display: flex;
            align-items: center;
            justify-content: center;
            color: white;
            font-weight: bold;
            text-shadow: 1px 1px 2px rgba(0,0,0,0.7);
        }
        
        .file-section {
            background: #1a1a1a;
            border: 1px solid #3a3a3a;
            border-radius: 8px;
            margin-bottom: 20px;
            box-shadow: 0 4px 12px rgba(255,69,0,0.15);
            overflow: hidden;
        }
        
        .file-header {
            background: linear-gradient(135deg, #2a2a2a 0%%, #1a1a1a 100%%);
            padding: 20px;
            border-bottom: 2px solid #ff4500;
            font-size: 1.3em;
            color: #ffd700;
        }
        
        .test-list {
            padding: 20px;
        }
        
        .test-case {
            border-left: 4px solid #3a3a3a;
            padding: 15px;
            margin-bottom: 15px;
            background-color: #2a2a2a;
            border-radius: 4px;
            transition: all 0.2s;
        }
        
        .test-case.passed {
            border-left-color: #32cd32;
            background: linear-gradient(135deg, #1a2a1a 0%%, #2a2a2a 100%%);
        }
        
        .test-case.failed {
            border-left-color: #ff4500;
            background: linear-gradient(135deg, #2a1a1a 0%%, #2a2a2a 100%%);
        }
        
        .test-case:hover {
            box-shadow: 0 4px 8px rgba(255,69,0,0.2);
            transform: translateX(2px);
        }
        
        .test-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            flex-wrap: wrap;
            gap: 10px;
        }
        
        .test-name {
            font-weight: 600;
            font-size: 1.1em;
            color: #ffd700;
        }
        
        .test-status {
            padding: 4px 12px;
            border-radius: 20px;
            font-size: 0.85em;
            font-weight: 600;
        }
        
        .test-status.passed {
            background-color: rgba(50,205,50,0.2);
            color: #32cd32;
            border: 1px solid #32cd32;
        }
        
        .test-status.failed {
            background-color: rgba(255,69,0,0.2);
            color: #ff4500;
            border: 1px solid #ff4500;
        }
        
        .test-duration {
            color: #ffb347;
            font-size: 0.9em;
        }
        
        .error-message {
            margin-top: 10px;
            padding: 15px;
            background-color: rgba(255,69,0,0.1);
            border: 1px solid #ff4500;
            border-radius: 4px;
            color: #ff8c00;
            font-family: 'Courier New', monospace;
            font-size: 0.9em;
            line-height: 1.5;
            overflow-x: auto;
        }
        
        footer {
            text-align: center;
            padding: 30px 0;
            color: #ffb347;
            font-size: 0.9em;
            border-top: 1px solid #3a3a3a;
            margin-top: 50px;
        }
        
        .charts-section {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
            gap: 30px;
            margin: 30px 0;
        }
        
        .chart-container {
            background: #1a1a1a;
            border: 1px solid #3a3a3a;
            border-radius: 8px;
            padding: 25px;
            box-shadow: 0 4px 12px rgba(255,69,0,0.15);
        }
        
        .chart-container h3 {
            margin-bottom: 20px;
            color: #ffd700;
        }
        
        h2 {
            color: #ffd700;
            text-shadow: 1px 1px 2px rgba(255,69,0,0.3);
        }
        
        @media (max-width: 768px) {
            .container {
                padding: 10px;
            }
            
            h1 {
                font-size: 2em;
            }
            
            .summary {
                grid-template-columns: 1fr;
                gap: 15px;
            }
            
            .test-header {
                flex-direction: column;
                align-items: flex-start;
            }
        }
    </style>
</head>
<body>
    <header>
        <div class="container">
            <div class="header-content">
                <img src="../assets/dwarf_in_forge.png" alt="Sindri Forge" class="logo">
                <div>
                    <h1>Sindri Test Report</h1>
                    <p class="subtitle">Generated on %s</p>
                </div>
                <img src="../assets/dwarf_in_forge.png" alt="Sindri Forge" class="logo">
            </div>
        </div>
    </header>
    
    <div class="container">
        <div class="summary">
            <div class="summary-card total">
                <h3>Total Tests</h3>
                <div class="value">%d</div>
            </div>
            <div class="summary-card passed">
                <h3>Passed</h3>
                <div class="value">%d</div>
            </div>
            <div class="summary-card failed">
                <h3>Failed</h3>
                <div class="value">%d</div>
            </div>
            <div class="summary-card duration">
                <h3>Total Duration</h3>
                <div class="value">%s</div>
            </div>
            <div class="summary-card rate">
                <h3>Pass Rate</h3>
                <div class="value">%.1f%%</div>
            </div>
        </div>
        
        <div class="progress-bar">
            <div class="progress-fill" style="width: %.1f%%" data-width="%.1f">
                %.1f%% Passed
            </div>
        </div>
        
        <h2 style="margin: 30px 0 20px; color: #ffd700;">Test Results</h2>
        
        %s
        
        <footer>
            <p>Generated by Sindri Testing Framework</p>
        </footer>
    </div>
    
    <script>
        // Animate progress bar on load
        window.addEventListener('load', function() {
            const progressFill = document.querySelector('.progress-fill');
            const targetWidth = progressFill.getAttribute('data-width');
            progressFill.style.width = '0%%';
            setTimeout(() => {
                progressFill.style.width = targetWidth + '%%';
            }, 100);
        });
        
        // Add click to copy for error messages
        document.querySelectorAll('.error-message').forEach(el => {
            el.style.cursor = 'pointer';
            el.title = 'Click to copy';
            el.addEventListener('click', function() {
                const text = this.textContent.replace('Error: ', '').trim();
                navigator.clipboard.writeText(text).then(() => {
                    const original = this.innerHTML;
                    this.innerHTML = '<strong>Copied!</strong>';
                    setTimeout(() => {
                        this.innerHTML = original;
                    }, 1000);
                });
            });
        });
    </script>
</body>
</html>`, time.Now().Format("January 2, 2006 at 3:04 PM"), time.Now().Format("January 2, 2006 at 3:04 PM"), totalTests, totalPassed, totalFailed, totalDuration, passRate, passRate, passRate, passRate, testResultsHTML.String())
}
