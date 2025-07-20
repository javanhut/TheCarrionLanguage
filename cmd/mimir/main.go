package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/peterh/liner"
)

const MIMIR_LOGO = `
    MIMIR - The All-Seeing Helper
   ═══════════════════════════════════════
    Interactive Carrion Documentation
   ═══════════════════════════════════════
   "Knowledge is the greatest treasure"
`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: mimir [command] [options]\n\n")
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "  interactive           Start interactive help mode (default)\n")
		fmt.Fprintf(os.Stderr, "  scry <function>       Get help for specific function or topic\n")
		fmt.Fprintf(os.Stderr, "  list                  List all available functions and modules\n")
		fmt.Fprintf(os.Stderr, "  categories            Show function categories\n")
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  mimir                 # Start interactive mode\n")
		fmt.Fprintf(os.Stderr, "  mimir scry print      # Get help for print function\n")
		fmt.Fprintf(os.Stderr, "  mimir scry Array      # Get help for Array module\n")
		fmt.Fprintf(os.Stderr, "  mimir scry os         # Get help for OS module\n")
		fmt.Fprintf(os.Stderr, "  mimir list            # List all functions\n")
		fmt.Fprintf(os.Stderr, "  mimir categories      # Browse by category\n")
	}

	args := os.Args[1:]

	if len(args) == 0 {
		startInteractiveMode()
		return
	}

	command := args[0]
	switch command {
	case "interactive", "i":
		startInteractiveMode()
	case "scry", "s":
		if len(args) < 2 {
			fmt.Println("Error: scry command requires a function or topic name")
			fmt.Println("Usage: mimir scry <function>")
			fmt.Println("Example: mimir scry print")
			os.Exit(1)
		}
		scryFunction(args[1])
	case "list", "l":
		listAllFunctions()
	case "categories", "cat", "c":
		showCategories()
	case "help", "-h", "--help":
		flag.Usage()
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		flag.Usage()
		os.Exit(1)
	}
}

func startInteractiveMode() {
	line := liner.NewLiner()
	defer func() {
		ok := line.Close()
		if ok != nil {
			fmt.Printf("Error closing liner: %v\n", ok)
		}
	}()

	fmt.Println(MIMIR_LOGO)
	fmt.Println("")
	showHelpMenu()

	for {
		input, err := line.Prompt("mimir> ")
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nFarewell! May your code be bug-free!")
				return
			}
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		choice := strings.ToLower(strings.TrimSpace(input))

		switch choice {
		case "1", "builtins", "builtin":
			showBuiltinFunctions(line)
		case "2", "stdlib", "standard", "munin":
			showStandardLibrary(line)
		case "3", "syntax", "language":
			showLanguageFeatures(line)
		case "4", "examples", "demo":
			showExamples(line)
		case "5", "search", "find":
			searchFunctions(line)
		case "6", "tips", "tricks":
			showTipsAndTricks()
		case "h", "help", "menu":
			showHelpMenu()
		case "q", "quit", "exit", "bye":
			fmt.Println("\nFarewell! May your code be bug-free!")
			return
		case "clear", "cls":
			clearScreen()
		case "":
			continue
		default:
			// Try to find function by name
			if found := searchSpecificFunction(choice); !found {
				fmt.Printf("Unknown command '%s'. Type 'h' for menu or 'q' to quit.\n\n", input)
			}
		}
	}
}

func scryFunction(name string) {
	fmt.Printf("═══════════════════════════════════════════════════════════════════\n")
	fmt.Printf("SCRYING: %s\n", strings.ToUpper(name))
	fmt.Printf("═══════════════════════════════════════════════════════════════════\n\n")

	if !searchSpecificFunction(name) {
		fmt.Printf("No documentation found for '%s'\n\n", name)
		fmt.Println("Try one of these:")
		fmt.Println("   - mimir list           - See all available functions")
		fmt.Println("   - mimir categories     - Browse by category")
		fmt.Println("   - mimir scry print     - Get help for print function")
		fmt.Println("   - mimir scry Array     - Get help for Array module")
		os.Exit(1)
	}
}

func listAllFunctions() {
	fmt.Println("ALL AVAILABLE FUNCTIONS AND MODULES")
	fmt.Println("═══════════════════════════════════════════════════════════════════")

	fmt.Println("\nBUILT-IN FUNCTIONS:")
	fmt.Println("Type Conversion: int, float, str, bool, list, tuple")
	fmt.Println("Utility:        print, len, type, input, range, max, abs")
	fmt.Println("System:         help, version, modules, open, parseHash")
	fmt.Println("Mathematical:   max, abs, ord, chr")
	fmt.Println("Collections:    enumerate, pairs, is_sametype")

	fmt.Println("\nSTANDARD LIBRARY MODULES:")
	fmt.Println("Array:    Enhanced array operations and methods")
	fmt.Println("String:   String manipulation and processing")
	fmt.Println("Integer:  Integer utilities and conversions")
	fmt.Println("Float:    Floating-point operations")
	fmt.Println("Boolean:  Boolean logic operations")
	fmt.Println("File:     File I/O operations")
	fmt.Println("OS:       Operating system interface")
	fmt.Println("Math:     Mathematical functions")
	fmt.Println("Time:     Date and time utilities")
	fmt.Println("HTTP:     Web requests and API calls")

	fmt.Println("\nUse 'mimir scry <function>' for detailed help on any item above")
}

func showCategories() {
	showSearchCategories()
}

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func showHelpMenu() {
	fmt.Println("What knowledge do you seek?")
	fmt.Println("")
	fmt.Println("  1.  Built-in Functions    - Core language functions (print, len, type, etc.)")
	fmt.Println("  2.  Standard Library      - Munin modules (Array, String, File, OS, Time, HTTP, etc.)")
	fmt.Println("  3.  Language Features     - Syntax, control flow, OOP, error handling")
	fmt.Println("  4.  Examples & Demos      - Working code examples and tutorials")
	fmt.Println("  5.  Search Functions      - Find specific functions by name or purpose")
	fmt.Println("  6.  Tips & Tricks         - REPL shortcuts and advanced features")
	fmt.Println("")
	fmt.Println("Commands: Type number or name - '1'/'builtins', '2'/'stdlib', '3'/'syntax', '4'/'examples', '5'/'search', '6'/'tips', 'h' for menu, 'q' to quit")
	fmt.Println("Quick search: Type any function name directly (e.g., 'print', 'Array')")
	fmt.Println("")
}

// searchSpecificFunction searches for a specific function and displays its help
func searchSpecificFunction(name string) bool {
	name = strings.ToLower(name)

	// Built-in functions
	builtinFuncs := map[string]string{
		"print":       "print(*args) - Print values to console\n   Example: print(\"Hello\", 42, True)",
		"len":         "len(object) - Get length of strings, arrays, hashes, tuples\n   Example: len(\"hello\") → 5, len([1,2,3]) → 3",
		"type":        "type(object) - Get the type of an object as a string\n   Example: type(42) → \"INTEGER\", type(\"hello\") → \"STRING\"",
		"input":       "input(prompt=\"\") - Read user input with optional prompt\n   Example: name = input(\"Enter name: \")",
		"range":       "range(start, stop, step=1) - Generate sequence of numbers\n   Example: range(5) → [0,1,2,3,4], range(2,8,2) → [2,4,6]",
		"max":         "max(*args) - Return maximum value from arguments\n   Example: max(1,5,3) → 5, max([10,20,15]) → 20",
		"abs":         "abs(value) - Return absolute value of a number\n   Example: abs(-42) → 42, abs(3.14) → 3.14",
		"ord":         "ord(char) - Return ASCII/Unicode code of character\n   Example: ord(\"A\") → 65, ord(\"a\") → 97",
		"chr":         "chr(code) - Return character from ASCII/Unicode code\n   Example: chr(65) → \"A\", chr(97) → \"a\"",
		"int":         "int(value) - Convert value to integer\n   Example: int(\"42\") → 42, int(3.14) → 3",
		"float":       "float(value) - Convert value to float\n   Example: float(\"3.14\") → 3.14, float(42) → 42.0",
		"str":         "str(value) - Convert value to string\n   Example: str(42) → \"42\", str([1,2,3]) → \"[1, 2, 3]\"",
		"bool":        "bool(value) - Convert value to boolean\n   Example: bool(1) → True, bool(\"\") → False",
		"list":        "list(iterable) - Convert iterable to array\n   Example: list(\"hello\") → [\"h\",\"e\",\"l\",\"l\",\"o\"]",
		"tuple":       "tuple(iterable) - Convert iterable to tuple\n   Example: tuple([1,2,3]) → (1,2,3)",
		"enumerate":   "enumerate(array) - Return array of (index, value) tuples\n   Example: enumerate([\"a\",\"b\"]) → [(0,\"a\"),(1,\"b\")]",
		"pairs":       "pairs(hash, filter=\"\") - Return key-value pairs from hash\n   Example: pairs({\"a\":1, \"b\":2}) → [(\"a\",1),(\"b\",2)]",
		"is_sametype": "is_sametype(obj1, obj2) - Check if objects have same type\n   Example: is_sametype(42, 17) → True",
		"help":        "help() - Show basic help information",
		"version":     "version() - Show Carrion and Munin version information",
		"modules":     "modules() - List all available standard library modules",
		"open":        "open(path, mode='r') - Open a file for reading, writing, or appending\n   Example: file = open(\"data.txt\", \"r\")",
		"parsehash":   "parseHash(json_string) - Parse JSON string into hash object\n   Example: data = parseHash('{\"name\": \"Alice\", \"age\": 30}')",
	}

	// Standard library modules
	stdlibModules := map[string]string{
		"array":    "Array([elements]) - Enhanced array operations\n   Methods: .append(), .sort(), .reverse(), .contains(), .length()\n   Example: arr = Array([3,1,4]); arr.sort() → [1,3,4]",
		"string":   "String(value) - String manipulation and processing\n   Methods: .upper(), .lower(), .find(), .reverse(), .char_at(), .split(), .join()\n   Example: s = String(\"hello\"); s.upper() → \"HELLO\"",
		"integer":  "Integer(value=0) - Integer utilities and conversions\n   Methods: .to_bin(), .to_hex(), .is_prime(), .gcd(), .lcm()\n   Example: i = Integer(42); i.to_bin() → \"0b101010\"",
		"float":    "Float(value=0.0) - Floating-point operations\n   Methods: .round(), .sqrt(), .sin(), .cos(), .abs()\n   Example: f = Float(3.14159); f.round(2) → 3.14",
		"boolean":  "Boolean(value=False) - Boolean logic operations\n   Methods: .to_int(), .negate(), .and_with(), .or_with(), .xor_with()\n   Example: b = Boolean(True); b.to_int() → 1",
		"file":     "File() - File I/O operations\n   Methods: .file_read(), .file_write(), .seek(), .close()\n   Static: File.read(), File.write(), File.exists()\n   Example: f = File(\"data.txt\", \"r\"); content = f.file_read()",
		"os":       "os() - Operating system interface\n   Methods: .cwd(), .list_dir(), .getenv(), .run(), .sleep()\n   Example: os = os(); files = os.list_dir(\".\")",
		"http":     "ApiRequest() - HTTP client for API requests\n   Methods: .get(), .post(), .put(), .delete(), .get_json(), .post_json()\n   Example: api = ApiRequest(); response = api.get_json(\"https://api.example.com/data\")",
		"time":     "Time() - Date and time utilities\n   Methods: .now(), .now_timestamp(), .sleep(), .format(), .parse()\n   Example: t = Time(); timestamp = t.now_timestamp(); t.sleep(2)",
		"comments": "Comment syntax - Single-line (#) and multi-line (```)\n   Single: # comment\n   Multi: ``` comment block ```",
	}

	if desc, found := builtinFuncs[name]; found {
		fmt.Printf("\nBUILT-IN FUNCTION: %s\n", strings.ToUpper(name))
		fmt.Printf("═════════════════════════════\n")
		fmt.Printf("%s\n\n", desc)
		return true
	}

	if desc, found := stdlibModules[name]; found {
		fmt.Printf("\nSTANDARD LIBRARY: %s\n", strings.ToUpper(name))
		fmt.Printf("══════════════════════════════\n")
		fmt.Printf("%s\n\n", desc)
		return true
	}

	return false
}

// performFunctionSearch searches for functions matching the query
func performFunctionSearch(query string) []string {
	var results []string

	// Define searchable functions with keywords
	functions := map[string][]string{
		"print(*args) - Print values to console": {"print", "output", "display", "console"},
		"len(object) - Get length":               {"len", "length", "size", "count"},
		"type(object) - Get object type":         {"type", "typeof", "class"},
		"max(*args) - Find maximum value":        {"max", "maximum", "largest", "biggest"},
		"abs(value) - Absolute value":            {"abs", "absolute", "positive"},
		"int(value) - Convert to integer":        {"int", "integer", "convert", "number"},
		"float(value) - Convert to float":        {"float", "decimal", "convert", "number"},
		"str(value) - Convert to string":         {"str", "string", "text", "convert"},
		"Array([]) - Array operations":           {"array", "list", "collection", "sort", "append"},
		"String() - String manipulation":         {"string", "text", "upper", "lower", "find"},
		"Integer() - Integer utilities":          {"integer", "number", "binary", "hex", "prime"},
		"Float() - Float operations":             {"float", "decimal", "round", "sqrt", "math"},
		"File() - File operations":               {"file", "read", "write", "exists", "io"},
		"os() - System operations":               {"os", "system", "directory", "environment", "command"},
		"range() - Number sequences":             {"range", "sequence", "numbers", "iterate"},
		"enumerate() - Index-value pairs":        {"enumerate", "index", "iterate", "loop"},
		"open() - File opening":                  {"open", "file", "read", "write", "append"},
		"parseHash() - JSON parsing":             {"parsehash", "json", "parse", "hash", "object"},
		"ApiRequest() - HTTP client":             {"apirequest", "http", "get", "post", "request", "web", "api", "json"},
		"Time() - Time operations":               {"time", "now", "timestamp", "sleep", "date", "format"},
		"Comments - Comment syntax":              {"comments", "syntax", "comment", "single-line", "multi-line"},
	}

	for funcDesc, keywords := range functions {
		for _, keyword := range keywords {
			if strings.Contains(keyword, query) || strings.Contains(query, keyword) {
				results = append(results, funcDesc)
				break
			}
		}
	}

	return results
}

// showSearchCategories displays function categories for browsing
func showSearchCategories() {
	fmt.Println("")
	fmt.Println("FUNCTION CATEGORIES:")
	fmt.Println("")
	fmt.Println("Type Conversion:")
	fmt.Println("  int, float, str, bool, list, tuple")
	fmt.Println("")
	fmt.Println("Text Processing:")
	fmt.Println("  String, upper, lower, find, reverse")
	fmt.Println("")
	fmt.Println("Array/List Operations:")
	fmt.Println("  Array, append, sort, reverse, contains")
	fmt.Println("")
	fmt.Println("Mathematical:")
	fmt.Println("  Integer, Float, max, abs, round, sqrt")
	fmt.Println("")
	fmt.Println("File & System:")
	fmt.Println("  File, OS, read, write, directory, open")
	fmt.Println("")
	fmt.Println("Network & HTTP:")
	fmt.Println("  httpGet, httpPost, httpPut, httpDelete, httpParseJSON")
	fmt.Println("")
	fmt.Println("Time & Date:")
	fmt.Println("  timeNow, timeSleep, timeFormat, timeParse")
	fmt.Println("")
	fmt.Println("Comments & Syntax:")
	fmt.Println("  comments, syntax, single-line, multi-line")
	fmt.Println("")
	fmt.Println("Utility:")
	fmt.Println("  print, len, type, range, enumerate, parseHash")
	fmt.Println("")
}

// Interactive help functions
func showBuiltinFunctions(line *liner.State) {
	fmt.Println("")
	fmt.Println("BUILT-IN FUNCTIONS")
	fmt.Println("═══════════════════════")

	categories := map[string][]string{
		"1": {"Type Conversion", "int, float, str, bool, list, tuple"},
		"2": {"Utility Functions", "len, type, print, input, range"},
		"3": {"Mathematical", "max, abs, ord, chr"},
		"4": {"Collections", "enumerate, pairs, is_sametype"},
		"5": {"System Functions", "help, version, modules"},
	}

	for {
		fmt.Println("")
		fmt.Println("Select a category:")

		// Sort keys to ensure consistent ordering
		keys := make([]string, 0, len(categories))
		for k := range categories {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := categories[k]
			fmt.Printf("  %s. %s - %s\n", k, v[0], v[1])
		}
		fmt.Println("")
		fmt.Println("Commands:")
		fmt.Println("  Numbers: 1, 2, 3, 4, 5")
		fmt.Println("  Words: type/conversion, utility/util, math/mathematical, collections/collection, system/sys")
		fmt.Println("  Special: 'all' for everything, 'b' to go back, 'q' to quit")

		input, err := line.Prompt("builtins> ")
		if err != nil {
			return
		}

		choice := strings.ToLower(strings.TrimSpace(input))
		if choice == "b" {
			return
		}
		if choice == "q" || choice == "quit" || choice == "exit" {
			fmt.Println("Farewell! May your code be bug-free!")
			os.Exit(0)
		}

		switch choice {
		case "1", "type", "conversion":
			showTypeConversionFunctions()
		case "2", "utility", "util":
			showUtilityFunctions()
		case "3", "math", "mathematical":
			showMathFunctions()
		case "4", "collections", "collection":
			showCollectionFunctions()
		case "5", "system", "sys":
			showSystemFunctions()
		case "all":
			showAllBuiltinFunctions()
		case "":
			continue
		default:
			if !searchSpecificFunction(choice) {
				fmt.Printf("Unknown function '%s'\n", input)
			}
		}

		fmt.Println("\nPress Enter to continue...")
		line.Prompt("")
	}
}

func showStandardLibrary(line *liner.State) {
	fmt.Println("")
	fmt.Println("STANDARD LIBRARY (MUNIN)")
	fmt.Println("══════════════════════════")

	modules := map[string]string{
		"1":  "Array - Enhanced array operations",
		"2":  "String - String manipulation",
		"3":  "Integer - Integer utilities",
		"4":  "Float - Floating-point operations",
		"5":  "Boolean - Boolean logic operations",
		"6":  "File - File I/O operations",
		"7":  "OS - Operating system interface",
		"8":  "Math - Mathematical functions",
		"9":  "Time - Date and time utilities",
		"10": "HTTP - Web requests and API calls",
	}

	for {
		fmt.Println("")
		fmt.Println("Select a module:")

		// Sort keys to ensure consistent ordering
		keys := make([]string, 0, len(modules))
		for k := range modules {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := modules[k]
			fmt.Printf("  %s. %s\n", k, v)
		}
		fmt.Println("")
		fmt.Println("Commands:")
		fmt.Println("  Numbers: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10")
		fmt.Println("  Words: array, string, integer, float, boolean, file, os, math, time, http")
		fmt.Println("  Special: 'all' for everything, 'b' to go back, 'q' to quit")

		input, err := line.Prompt("stdlib> ")
		if err != nil {
			return
		}

		choice := strings.ToLower(strings.TrimSpace(input))
		if choice == "b" {
			return
		}
		if choice == "q" || choice == "quit" || choice == "exit" {
			fmt.Println("Farewell! May your code be bug-free!")
			os.Exit(0)
		}

		switch choice {
		case "1", "array":
			showArrayModule()
		case "2", "string":
			showStringModule()
		case "3", "integer":
			showIntegerModule()
		case "4", "float":
			showFloatModule()
		case "5", "boolean":
			showBooleanModule()
		case "6", "file":
			showFileModule()
		case "7", "os":
			showOSModule()
		case "8", "math":
			showMathModule()
		case "9", "time":
			showTimeModule()
		case "10", "http", "api":
			showHTTPModule()
		case "all":
			showAllStandardLibrary()
		case "":
			continue
		default:
			if !searchSpecificFunction(choice) {
				fmt.Printf("Unknown module '%s'\n", input)
			}
		}

		fmt.Println("\nPress Enter to continue...")
		line.Prompt("")
	}
}

func showLanguageFeatures(line *liner.State) {
	fmt.Println("")
	fmt.Println("CARRION LANGUAGE FEATURES")
	fmt.Println("═══════════════════════════")

	features := map[string]string{
		"1": "Variables & Assignment - Basic assignment, tuple unpacking, operators",
		"2": "Control Flow - if/otherwise/else, for/while loops, match/case",
		"3": "Functions (Spells) - Function definition, parameters, return values",
		"4": "Classes (Grimoires) - OOP, inheritance, methods, properties",
		"5": "Error Handling - attempt/ensnare/resolve, raising errors",
		"6": "Modules & Imports - Code organization, importing files",
		"7": "Data Types - Primitives, collections, type checking",
		"8": "Operators - Arithmetic, logical, comparison, bitwise",
	}

	for {
		fmt.Println("")
		fmt.Println("Select a topic:")

		// Sort keys to ensure consistent ordering
		keys := make([]string, 0, len(features))
		for k := range features {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := features[k]
			fmt.Printf("  %s. %s\n", k, v)
		}
		fmt.Println("")
		fmt.Println("Commands:")
		fmt.Println("  Numbers: 1, 2, 3, 4, 5, 6, 7, 8")
		fmt.Println("  Words: variables/assignment, control/flow, functions/spells, classes/grimoires/oop")
		fmt.Println("         errors/error, modules/imports, types/data, operators/operator")
		fmt.Println("  Special: 'all' for everything, 'b' to go back, 'q' to quit")

		input, err := line.Prompt("syntax> ")
		if err != nil {
			return
		}

		choice := strings.ToLower(strings.TrimSpace(input))
		if choice == "b" {
			return
		}
		if choice == "q" || choice == "quit" || choice == "exit" {
			fmt.Println("Farewell! May your code be bug-free!")
			os.Exit(0)
		}
		switch choice {
		case "1", "variables", "assignment":
			showVariablesAndAssignment()
		case "2", "control", "flow", "if", "for", "while":
			showControlFlow()
		case "3", "functions", "spells", "function":
			showFunctions()
		case "4", "classes", "grimoires", "oop", "class":
			showClasses()
		case "5", "errors", "error", "exceptions":
			showErrorHandling()
		case "6", "modules", "imports", "import":
			showModules()
		case "7", "types", "data":
			showDataTypes()
		case "8", "operators", "operator":
			showOperators()
		case "all":
			showAllLanguageFeatures()
		case "":
			continue
		default:
			fmt.Printf("Unknown topic '%s'\n", input)
		}

		fmt.Println("\nPress Enter to continue...")
		line.Prompt("")
	}
}

func showExamples(line *liner.State) {
	fmt.Println("")
	fmt.Println("EXAMPLES & TUTORIALS")
	fmt.Println("══════════════════════")

	examples := map[string]string{
		"1": "Hello World & Basics - Getting started with Carrion",
		"2": "Working with Arrays - Manipulation, sorting, searching",
		"3": "String Processing - Text manipulation and formatting",
		"4": "File Operations - Reading, writing, file management",
		"5": "Mathematical Calculations - Numbers, formulas, algorithms",
		"6": "Object-Oriented Programming - Classes, inheritance, methods",
		"7": "Error Handling Examples - Robust error management",
		"8": "Complete Mini Programs - Full working applications",
	}

	for {
		fmt.Println("")
		fmt.Println("Select an example category:")

		// Sort keys to ensure consistent ordering
		keys := make([]string, 0, len(examples))
		for k := range examples {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			v := examples[k]
			fmt.Printf("  %s. %s\n", k, v)
		}
		fmt.Println("")
		fmt.Println("Commands:")
		fmt.Println("  Numbers: 1, 2, 3, 4, 5, 6, 7, 8")
		fmt.Println("  Words: hello/basics, arrays/array, strings/string, files/file")
		fmt.Println("         math/mathematical, oop/classes, errors/error, programs/mini")
		fmt.Println("  Special: 'all' for everything, 'b' to go back, 'q' to quit")

		input, err := line.Prompt("examples> ")
		if err != nil {
			return
		}

		choice := strings.ToLower(strings.TrimSpace(input))
		if choice == "b" {
			return
		}
		if choice == "q" || choice == "quit" || choice == "exit" {
			fmt.Println("Farewell! May your code be bug-free!")
			os.Exit(0)
		}

		switch choice {
		case "1", "hello", "basics":
			showBasicExamples()
		case "2", "arrays", "array":
			showArrayExamples()
		case "3", "strings", "string":
			showStringExamples()
		case "4", "files", "file":
			showFileExamples()
		case "5", "math", "mathematical":
			showMathExamples()
		case "6", "oop", "classes":
			showOOPExamples()
		case "7", "errors", "error":
			showErrorExamples()
		case "8", "programs", "mini":
			showMiniPrograms()
		case "all":
			showAllExamples()
		case "":
			continue
		default:
			fmt.Printf("Unknown category '%s'\n", input)
		}

		fmt.Println("\nPress Enter to continue...")
		line.Prompt("")
	}
}

func searchFunctions(line *liner.State) {
	fmt.Println("")
	fmt.Println("FUNCTION SEARCH")
	fmt.Println("════════════════")
	fmt.Println("")
	fmt.Println("Enter keywords to search for functions, or type 'categories' to browse by category.")
	fmt.Println("Commands: function names/keywords, 'categories', 'help', 'b' to go back, 'q' to quit")

	for {
		input, err := line.Prompt("search> ")
		if err != nil {
			return
		}

		query := strings.ToLower(strings.TrimSpace(input))
		if query == "" {
			continue
		}

		if query == "b" {
			return
		}
		if query == "q" || query == "quit" || query == "exit" {
			fmt.Println("Farewell! May your code be bug-free!")
			os.Exit(0)
		}

		if query == "categories" {
			showSearchCategories()
			continue
		}

		if query == "help" {
			fmt.Println("")
			fmt.Println("Search Help:")
			fmt.Println("   • Enter function names or keywords (e.g., 'print', 'array', 'math')")
			fmt.Println("   • 'categories' - Browse function categories")
			fmt.Println("   • 'b' - Go back to main menu")
			continue
		}

		results := performFunctionSearch(query)
		if len(results) == 0 {
			fmt.Printf("No functions found matching '%s'\n", input)
		} else {
			fmt.Printf("\nFound %d function(s) matching '%s':\n\n", len(results), input)
			for _, result := range results {
				fmt.Printf("  • %s\n", result)
			}
		}
		fmt.Println("")
	}
}

func showTipsAndTricks() {
	fmt.Println("")
	fmt.Println("TIPS & TRICKS")
	fmt.Println("═══════════════")
	fmt.Println("")
	fmt.Println("REPL Usage:")
	fmt.Println("   • Tab - Auto-complete functions and keywords")
	fmt.Println("   • ↑/↓ - Navigate command history")
	fmt.Println("   • Ctrl+L - Clear screen")
	fmt.Println("   • Ctrl+C - Interrupt execution")
	fmt.Println("   • Ctrl+D - Exit REPL")
	fmt.Println("")
	fmt.Println("Language Tips:")
	fmt.Println("   • Use auto-wrapping: 42.to_bin() instead of Integer(42).to_bin()")
	fmt.Println("   • String interpolation: f\"Hello {name}!\"")
	fmt.Println("   • Tuple unpacking: x, y = (10, 20)")
	fmt.Println("   • Use 'attempt/ensnare' for error handling")
	fmt.Println("   • Functions are called 'spells', classes are 'grimoires'")
	fmt.Println("")
	fmt.Println("Debugging:")
	fmt.Println("   • Use type() to check object types")
	fmt.Println("   • Print intermediate values to debug calculations")
	fmt.Println("   • Use attempt/ensnare blocks for safe operations")
	fmt.Println("")
	fmt.Println("Documentation:")
	fmt.Println("   • 'mimir' - Open this interactive help")
	fmt.Println("   • 'mimir scry <function>' - Quick function lookup")
	fmt.Println("   • 'help()' - Basic help in REPL")
	fmt.Println("   • 'version()' - Show version information")
	fmt.Println("   • 'modules()' - List available modules")
	fmt.Println("")
}

