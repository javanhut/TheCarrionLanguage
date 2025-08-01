package evaluator

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/peterh/liner"

	"github.com/javanhut/TheCarrionLanguage/src/modules"
	"github.com/javanhut/TheCarrionLanguage/src/object"
)

var LineReader *liner.State

// Helper function to extract string value from object (handles both direct strings and instance-wrapped strings)
func extractStringBuiltin(obj object.Object) (string, bool) {
	switch v := obj.(type) {
	case *object.String:
		return v.Value, true
	case *object.Instance:
		if value, exists := v.Env.Get("value"); exists {
			if strVal, ok := value.(*object.String); ok {
				return strVal.Value, true
			}
		}
		return "", false
	default:
		return "", false
	}
}

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.Tuple:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.Hash:
				return &object.Integer{Value: int64(len(arg.Pairs))}
			case *object.Instance:
				// Handle instances based on their grimoire type
				switch arg.Grimoire.Name {
				case "Array":
					if elements, exists := arg.Env.Get("elements"); exists {
						// Check if elements is a direct Array
						if arr, isArray := elements.(*object.Array); isArray {
							return &object.Integer{Value: int64(len(arr.Elements))}
						}
						// Check if elements is an Instance wrapping an Array
						if elemInstance, isInstance := elements.(*object.Instance); isInstance {
							// Check if it's an Array instance containing value
							if value, valueExists := elemInstance.Env.Get("value"); valueExists {
								if arr, isArray := value.(*object.Array); isArray {
									return &object.Integer{Value: int64(len(arr.Elements))}
								}
							}
							// Try to see if it's a direct wrapped array
							if elemInstance.Grimoire.Name == "Array" {
								if innerElements, innerExists := elemInstance.Env.Get("elements"); innerExists {
									if arr, isArray := innerElements.(*object.Array); isArray {
										return &object.Integer{Value: int64(len(arr.Elements))}
									}
								}
							}
						}
					}
					return newError("invalid Array instance: missing or invalid elements")
				case "String":
					if value, exists := arg.Env.Get("value"); exists {
						if str, isString := value.(*object.String); isString {
							return &object.Integer{Value: int64(len(str.Value))}
						}
					}
					return newError("invalid String instance: missing value")
				default:
					return newError("len() not supported for %s instances", arg.Grimoire.Name)
				}
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"print": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				fmt.Println()
				return &object.None{}
			}
			
			// Build the output string from all arguments
			var parts []string
			for _, arg := range args {
				parts = append(parts, arg.Inspect())
			}
			
			// Join with spaces and print with newline (default behavior)
			output := strings.Join(parts, " ")
			fmt.Println(output)
			return &object.None{}
		},
	},
	"printn": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return &object.None{}
			}
			
			// Build the output string from all arguments
			var parts []string
			for _, arg := range args {
				parts = append(parts, arg.Inspect())
			}
			
			// Join with spaces and print without newline
			output := strings.Join(parts, " ")
			fmt.Print(output)
			return &object.None{}
		},
	},
	"printend": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				fmt.Println()
				return &object.None{}
			}
			
			// Default end character is newline
			endChar := "\n"
			printArgs := args
			
			// Check if the first argument is a hash/map containing end specification
			if len(args) > 0 {
				if hash, ok := args[0].(*object.Hash); ok {
					// Look for "end" key in the hash
					endKey := &object.String{Value: "end"}
					if endPair, exists := hash.Pairs[endKey.HashKey()]; exists {
						if endStr, ok := extractStringBuiltin(endPair.Value); ok {
							endChar = endStr
						}
					}
					
					// Look for "values" key in the hash for the actual values to print
					valuesKey := &object.String{Value: "values"}
					if valuesPair, exists := hash.Pairs[valuesKey.HashKey()]; exists {
						if array, ok := valuesPair.Value.(*object.Array); ok {
							printArgs = array.Elements
						} else {
							// If values is not an array, treat it as a single value
							printArgs = []object.Object{valuesPair.Value}
						}
					} else {
						// If no values key, print nothing (only apply end character)
						printArgs = []object.Object{}
					}
				}
			}
			
			// Build the output string from print arguments
			var parts []string
			for _, arg := range printArgs {
				parts = append(parts, arg.Inspect())
			}
			
			// Join with spaces and print with specified end character
			output := strings.Join(parts, " ")
			fmt.Print(output + endChar)
			return &object.None{}
		},
	},

	"input": {
		Fn: func(args ...object.Object) object.Object {
			prompt := ""
			if len(args) > 0 {
				if str, ok := args[0].(*object.String); ok {
					prompt = str.Value
				}
			}

			if LineReader != nil {
				userInput, err := LineReader.Prompt(prompt)
				if err != nil {
					return &object.Error{Message: "error reading input: " + err.Error()}
				}

				if userInput != "" {
					LineReader.AppendHistory(userInput)
				}
				return &object.String{Value: userInput}
			}

			fmt.Print(prompt)
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				return &object.Error{Message: "error reading input: " + err.Error()}
			}
			// Remove the trailing newline
			input = strings.TrimSuffix(input, "\n")
			input = strings.TrimSuffix(input, "\r") // Handle Windows line endings
			return &object.String{Value: input}
		},
	},
	"type": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			obj := args[0]
			
			// Use getObjectTypeString to get proper type names for instances
			switch o := obj.(type) {
			case *object.Instance:
				if o.Grimoire != nil {
					return &object.String{Value: o.Grimoire.Name}
				}
				return &object.String{Value: "instance"}
			case *object.Integer:
				return &object.String{Value: "Integer"}
			case *object.Float:
				return &object.String{Value: "Float"}
			case *object.String:
				return &object.String{Value: "String"}
			case *object.Boolean:
				return &object.String{Value: "Boolean"}
			case *object.Array:
				return &object.String{Value: "Array"}
			case *object.Hash:
				return &object.String{Value: "Map"}
			case *object.Tuple:
				return &object.String{Value: "Tuple"}
			case *object.Function:
				return &object.String{Value: "Function"}
			case *object.None:
				return &object.String{Value: "None"}
			case *object.Time:
				return &object.String{Value: "Time"}
			case *object.Duration:
				return &object.String{Value: "Duration"}
			default:
				return &object.String{Value: string(obj.Type())}
			}
		},
	},
	"int": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			
			// Unwrap Instance objects to get the underlying primitive
			arg := args[0]
			if instance, ok := arg.(*object.Instance); ok {
				if value, exists := instance.Env.Get("value"); exists {
					arg = value
				}
			}
			
			switch typedArg := arg.(type) {
			case *object.String:
				value, err := strconv.Atoi(typedArg.Value)
				if err != nil {
					return newError("cannot convert string to int: %s", err)
				}
				return &object.Integer{Value: int64(value)}
			case *object.Float:
				return &object.Integer{Value: int64(typedArg.Value)}
			case *object.Integer:
				return typedArg
			default:
				return newError("cannot convert %s to int", arg.Type())
			}
		},
	},
	"to_int": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			
			// Unwrap Instance objects to get the underlying primitive
			arg := args[0]
			if instance, ok := arg.(*object.Instance); ok {
				if value, exists := instance.Env.Get("value"); exists {
					arg = value
				}
			}
			
			switch typedArg := arg.(type) {
			case *object.String:
				value, err := strconv.Atoi(typedArg.Value)
				if err != nil {
					return newError("cannot convert string to int: %s", err)
				}
				return &object.Integer{Value: int64(value)}
			case *object.Float:
				return &object.Integer{Value: int64(typedArg.Value)}
			case *object.Integer:
				return typedArg
			default:
				return newError("cannot convert %s to int", arg.Type())
			}
		},
	},

	"float": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			
			// Unwrap Instance objects to get the underlying primitive
			arg := args[0]
			if instance, ok := arg.(*object.Instance); ok {
				if value, exists := instance.Env.Get("value"); exists {
					arg = value
				}
			}
			
			switch typedArg := arg.(type) {
			case *object.String:
				value, err := strconv.ParseFloat(typedArg.Value, 64)
				if err != nil {
					return newError("cannot convert string to float: %s", err)
				}
				return &object.Float{Value: value}
			case *object.Integer:
				return &object.Float{Value: float64(typedArg.Value)}
			case *object.Float:
				return typedArg
			default:
				return newError("cannot convert %s to float", arg.Type())
			}
		},
	},
	"str": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			
			// For Instance objects, try to get the underlying primitive value first
			arg := args[0]
			if instance, ok := arg.(*object.Instance); ok {
				if value, exists := instance.Env.Get("value"); exists {
					// Use the underlying primitive's string representation
					primitive := &object.String{Value: value.Inspect()}
					return wrapPrimitiveForBuiltin(primitive)
				}
			}
			
			// Fallback to using Inspect() on the object directly
			primitive := &object.String{Value: arg.Inspect()}
			// Create a String instance that supports method calls like .lower()
			return wrapPrimitiveForBuiltin(primitive)
		},
	},
	"String": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			// Convert input to string value
			var strValue string
			switch arg := args[0].(type) {
			case *object.String:
				strValue = arg.Value
			default:
				strValue = arg.Inspect()
			}

			// Create a String grimoire instance using the grimoire system
			primitive := &object.String{Value: strValue}
			return wrapPrimitiveForBuiltin(primitive)
		},
	},
	"bool": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Boolean:
				return arg
			case *object.Integer:
				if arg.Value == 0 {
					return FALSE
				}
				return TRUE
			case *object.Float:
				if arg.Value == 0.0 {
					return FALSE
				}
				return TRUE
			case *object.String:
				if arg.Value == "" {
					return FALSE
				}
				return TRUE
			case *object.Array:
				if len(arg.Elements) == 0 {
					return FALSE
				}
				return TRUE
			case *object.Hash:
				if len(arg.Pairs) == 0 {
					return FALSE
				}
				return TRUE
			case *object.None:
				return FALSE
			default:
				return TRUE
			}
		},
	},
	"list": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				elements := make([]object.Object, len(arg.Value))
				for i, char := range arg.Value {
					elements[i] = &object.String{Value: string(char)}
				}
				return &object.Array{Elements: elements}
			case *object.Tuple:
				return &object.Array{Elements: arg.Elements}
			default:
				return newError("cannot convert %s to list", arg.Type())
			}
		},
	},
	"tuple": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Tuple{Elements: arg.Elements}
			case *object.Tuple:
				return arg
			default:
				return newError("cannot convert %s to tuple", arg.Type())
			}
		},
	},
	"range": {
		Fn: func(args ...object.Object) object.Object {
			var start, stop, step int64

			switch len(args) {
			case 1:
				stopValue, err := extractIntegerValue(args[0])
				if err != nil {
					return newError("argument to `range` must be INTEGER, got=%s", args[0].Type())
				}
				start, stop, step = 0, stopValue, 1
			case 2:
				startValue, err1 := extractIntegerValue(args[0])
				stopValue, err2 := extractIntegerValue(args[1])
				if err1 != nil || err2 != nil {
					return newError(
						"arguments to `range` must be INTEGER, got=%s and %s",
						args[0].Type(),
						args[1].Type(),
					)
				}
				start, stop, step = startValue, stopValue, 1
			case 3:
				startValue, err1 := extractIntegerValue(args[0])
				stopValue, err2 := extractIntegerValue(args[1])
				stepValue, err3 := extractIntegerValue(args[2])
				if err1 != nil || err2 != nil || err3 != nil {
					return newError(
						"arguments to `range` must be INTEGER, got=%s, %s, %s",
						args[0].Type(),
						args[1].Type(),
						args[2].Type(),
					)
				}
				start, stop, step = startValue, stopValue, stepValue
			default:
				return newError("wrong number of arguments. got=%d, want=1..3", len(args))
			}

			if step == 0 {
				return newError("step argument to `range` cannot be zero")
			}

			var elements []object.Object
			if step > 0 {
				for i := start; i < stop; i += step {
					elements = append(elements, &object.Integer{Value: i})
				}
			} else {
				for i := start; i > stop; i += step {
					elements = append(elements, &object.Integer{Value: i})
				}
			}

			return &object.Array{Elements: elements}
		},
	},

	"max": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return newError("max requires at least one argument")
			}

			var nums []float64
			for _, arg := range args {
				switch v := arg.(type) {
				case *object.Integer:
					nums = append(nums, float64(v.Value))
				case *object.Float:
					nums = append(nums, v.Value)
				default:
					return newError("max: unsupported type %s", arg.Type())
				}
			}

			maxVal := nums[0]
			for _, n := range nums[1:] {
				if n > maxVal {
					maxVal = n
				}
			}

			if maxVal == float64(int64(maxVal)) {
				return &object.Integer{Value: int64(maxVal)}
			}
			return &object.Float{Value: maxVal}
		},
	},

	"abs": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("abs requires exactly one argument, got %d", len(args))
			}
			switch v := args[0].(type) {
			case *object.Integer:
				if v.Value < 0 {
					return &object.Integer{Value: -v.Value}
				}
				return v
			case *object.Float:
				if v.Value < 0 {
					return &object.Float{Value: -v.Value}
				}
				return v
			default:
				return newError("abs not supported for type %s", args[0].Type())
			}
		},
	},

	"Error": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return newError("Error requires 1 or 2 arguments: name, [message]")
			}
			name, ok := args[0].(*object.String)
			if !ok {
				return newError("Error name must be a string")
			}

			var message string
			if len(args) == 2 {
				msg, ok := args[1].(*object.String)
				if ok {
					message = msg.Value
				}
			}

			return &object.CustomError{
				Name:    name.Value,
				Message: message,
				Details: make(map[string]object.Object),
			}
		},
	},
	"enumerate": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("enumerate expects 1 argument, got %d", len(args))
			}

			var elements []object.Object
			
			switch arg := args[0].(type) {
			case *object.Array:
				elements = arg.Elements
			case *object.Instance:
				// Handle Array instances
				if arg.Grimoire.Name == "Array" {
					if elementsObj, exists := arg.Env.Get("elements"); exists {
						// Check if elements is a direct Array
						if arr, isArray := elementsObj.(*object.Array); isArray {
							elements = arr.Elements
						} else if elemInstance, isInstance := elementsObj.(*object.Instance); isInstance {
							// Check if it's an Array instance containing value
							if value, valueExists := elemInstance.Env.Get("value"); valueExists {
								if arr, isArray := value.(*object.Array); isArray {
									elements = arr.Elements
								}
							} else if elemInstance.Grimoire.Name == "Array" {
								// Try to see if it's a direct wrapped array
								if innerElements, innerExists := elemInstance.Env.Get("elements"); innerExists {
									if arr, isArray := innerElements.(*object.Array); isArray {
										elements = arr.Elements
									}
								}
							}
						}
					}
					if elements == nil {
						return newError("invalid Array instance: missing or invalid elements")
					}
				} else {
					return newError("enumerate expects an array, got instance of %s", arg.Grimoire.Name)
				}
			default:
				return newError("enumerate expects an array, got %s", args[0].Type())
			}
			
			var enumerated []object.Object
			for i, elem := range elements {
				tuple := &object.Tuple{
					Elements: []object.Object{
						&object.Integer{Value: int64(i)},
						elem,
					},
				}
				enumerated = append(enumerated, tuple)
			}

			return &object.Array{Elements: enumerated}
		},
	},

	"pairs": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return newError("pairs expects 1 or 2 arguments, got=%d", len(args))
			}
			// The first argument must be a hash.
			hashObj, ok := args[0].(*object.Hash)
			if !ok {
				return newError(
					"pairs expects a HASH as the first argument, got %s",
					args[0].Type(),
				)
			}

			// Determine the filter string if provided.
			filter := ""
			if len(args) == 2 {
				// Handle both raw strings and String instances
				switch arg := args[1].(type) {
				case *object.String:
					filter = arg.Value
				case *object.Instance:
					// Check if it's a String instance
					if arg.Grimoire.Name == "String" {
						if value, exists := arg.Env.Get("value"); exists {
							if str, isString := value.(*object.String); isString {
								filter = str.Value
							}
						}
					} else {
						return newError(
							"pairs second argument must be a STRING filter, got %s instance",
							arg.Grimoire.Name,
						)
					}
				default:
					return newError(
						"pairs second argument must be a STRING filter, got %s",
						args[1].Type(),
					)
				}
			}

			// Iterate over the hash's pairs.
			var result []object.Object
			for _, pair := range hashObj.Pairs {
				switch filter {
				case "":
					// Default: return both key and value in a tuple.
					result = append(result, &object.Tuple{
						Elements: []object.Object{pair.Key, pair.Value},
					})
				case "key", "k":
					result = append(result, pair.Key)
				case "value", "v":
					result = append(result, pair.Value)
				default:
					return newError(
						"pairs: invalid filter %q; expected 'key', 'value', 'k', or 'v'",
						filter,
					)
				}
			}
			// Return as Array instance so it has access to keys() and values() methods
			arrayResult := &object.Array{Elements: result}

			// Wrap the array as an Array instance if the stdlib is available
			if stdlibEnv != nil {
				if grimObj, ok := stdlibEnv.Get("Array"); ok {
					if grimoire, isGrim := grimObj.(*object.Grimoire); isGrim {
						// Create instance exactly like the normal grimoire constructor
						instance := &object.Instance{
							Grimoire: grimoire,
							Env:      object.NewEnclosedEnvironment(grimoire.Env),
						}

						// Set self reference and elements
						instance.Env.Set("self", instance)
						instance.Env.Set("elements", arrayResult)

						return instance
					}
				}
			}

			// Fallback to raw array if wrapping fails
			return arrayResult
		},
	},
	"is_sametype": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("is_sametype requires 2 arguments, got=%d", len(args))
			}

			return &object.Boolean{Value: args[0].Type() == args[1].Type()}
		},
	},
	"ord": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("ord requires exactly 1 argument, got=%d", len(args))
			}
			str, ok := args[0].(*object.String)
			if !ok {
				return newError("ord argument must be STRING, got=%s", args[0].Type())
			}
			if len(str.Value) != 1 {
				return newError("ord expects a single character string, got length %d", len(str.Value))
			}
			return &object.Integer{Value: int64(str.Value[0])}
		},
	},
	"chr": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("chr requires exactly 1 argument, got=%d", len(args))
			}
			num, ok := args[0].(*object.Integer)
			if !ok {
				return newError("chr argument must be INTEGER, got=%s", args[0].Type())
			}
			if num.Value < 0 || num.Value > 255 {
				return newError("chr argument must be in range 0-255, got=%d", num.Value)
			}
			return &object.String{Value: string(rune(num.Value))}
		},
	},
	"open": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return newError("open requires 1 or 2 arguments: path, [mode]")
			}

			// Get path argument
			pathStr, ok := extractStringBuiltin(args[0])
			if !ok {
				return newError("open path must be STRING, got=%s", args[0].Type())
			}

			// Get mode argument (default to "r")
			mode := "r"
			if len(args) == 2 {
				modeStr, ok := extractStringBuiltin(args[1])
				if !ok {
					return newError("open mode must be STRING, got=%s", args[1].Type())
				}
				mode = modeStr
			}

			// Validate mode
			if mode != "r" && mode != "w" && mode != "a" {
				return newError("invalid file mode: %s (must be 'r', 'w', or 'a')", mode)
			}

			// Get the File grimoire from stdlib environment
			if stdlibEnv == nil {
				return newError("stdlib not loaded")
			}

			fileGrimObj, exists := stdlibEnv.Get("File")
			if !exists {
				return newError("File grimoire not found in stdlib")
			}

			fileGrim, ok := fileGrimObj.(*object.Grimoire)
			if !ok {
				return newError("File is not a grimoire")
			}

			// Create a new File instance
			instance := &object.Instance{
				Grimoire: fileGrim,
				Env:      object.NewEnclosedEnvironment(fileGrim.Env),
			}

			// Set self reference
			instance.Env.Set("self", instance)

			// Initialize the File instance state according to the File grimoire's init spell
			instance.Env.Set("path", &object.String{Value: pathStr})
			instance.Env.Set("mode", &object.String{Value: mode})
			instance.Env.Set("_handle", &object.None{})
			instance.Env.Set("_closed", &object.Boolean{Value: false})

			// Handle file operations based on mode
			if mode == "r" {
				// Read mode - read the file content
				content, err := os.ReadFile(pathStr)
				if err != nil {
					return newError("failed to open file '%s' for reading: %s", pathStr, err)
				}
				instance.Env.Set("_content", &object.String{Value: string(content)})
				instance.Env.Set("_position", &object.Integer{Value: 0})
			} else if mode == "w" {
				// Write mode - clear the file
				err := os.WriteFile(pathStr, []byte(""), 0644)
				if err != nil {
					return newError("failed to open file '%s' for writing: %s", pathStr, err)
				}
			} else if mode == "a" {
				// Append mode - check if file exists and get content
				if _, err := os.Stat(pathStr); err == nil {
					content, err := os.ReadFile(pathStr)
					if err != nil {
						return newError("failed to read file '%s' for append: %s", pathStr, err)
					}
					instance.Env.Set("_content", &object.String{Value: string(content)})
				} else {
					instance.Env.Set("_content", &object.String{Value: ""})
				}
			}

			return instance
		},
	},
	"parseHash": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("parseHash requires exactly 1 argument, got=%d", len(args))
			}

			// Extract string from argument
			hashStr, ok := extractStringBuiltin(args[0])
			if !ok {
				return newError("parseHash argument must be STRING, got=%s", args[0].Type())
			}

			// Parse as JSON (similar to httpParseJSON)
			var data interface{}
			if err := json.Unmarshal([]byte(hashStr), &data); err != nil {
				// Return a more helpful error message
				return newError("parseHash: failed to parse string as JSON object: %s", err)
			}

			// Convert JSON data to Carrion object
			result := jsonToCarrionObject(data)
			
			// Ensure result is a hash
			if hash, ok := result.(*object.Hash); ok {
				return hash
			}
			
			return newError("parseHash: input must be a JSON object (dictionary/hashmap), got %s", result.Type())
		},
	},
}

// Add OS module functions to builtins when module is loaded
func init() {
	// Merge OS module functions into builtins
	for name, builtin := range modules.OSBuiltins {
		builtins[name] = builtin
	}
	// Merge File module functions into builtins
	for name, builtin := range modules.FileBuiltins {
		builtins[name] = builtin
	}
	// Merge Sockets module functions into builtins
	for name, builtin := range modules.SocketsModule {
		builtins[name] = builtin
	}
	// Merge HTTP module functions into builtins
	for name, builtin := range modules.HttpModule {
		builtins[name] = builtin
	}
}

// Global reference to the stdlib environment
var stdlibEnv *object.Environment

// SetStdlibEnv sets the global reference to the standard library environment
func SetStdlibEnv(env *object.Environment) {
	stdlibEnv = env
}

// wrapPrimitiveForBuiltin wraps a primitive object in a grimoire instance for use in builtin functions
// This creates String instances that support method calls like .lower()
func wrapPrimitiveForBuiltin(obj object.Object) object.Object {
	if stdlibEnv == nil {
		// Fallback: return primitive if no stdlib environment available
		return obj
	}

	var grimName string

	switch obj.Type() {
	case object.STRING_OBJ:
		grimName = "String"
	case object.INTEGER_OBJ:
		grimName = "Integer"
	case object.FLOAT_OBJ:
		grimName = "Float"
	case object.BOOLEAN_OBJ:
		grimName = "Boolean"
	default:
		return obj // Not a wrappable primitive, return as is
	}

	// Try to find the grimoire in the stdlib environment
	if grimObj, ok := stdlibEnv.Get(grimName); ok {
		if grimoire, isGrim := grimObj.(*object.Grimoire); isGrim {
			// Create instance exactly like the normal grimoire constructor
			instance := &object.Instance{
				Grimoire: grimoire,
				Env:      object.NewEnclosedEnvironment(grimoire.Env),
			}

			// Set up the instance environment
			instance.Env.Set("self", instance)
			instance.Env.Set("value", obj)

			return instance
		}
	}

	// If grimoire not found, return the original object
	return obj
}

// extractIntegerValue extracts an integer value from various object types
func extractIntegerValue(obj object.Object) (int64, error) {
	switch v := obj.(type) {
	case *object.Integer:
		return v.Value, nil
	case *object.Instance:
		// Check if it's an Integer instance with a value
		if v.Grimoire.Name == "Integer" {
			if valueObj, exists := v.Env.Get("value"); exists {
				if intVal, ok := valueObj.(*object.Integer); ok {
					return intVal.Value, nil
				}
			}
		}
		return 0, fmt.Errorf("instance is not an Integer")
	case *object.Float:
		// Allow converting float to int
		return int64(v.Value), nil
	default:
		return 0, fmt.Errorf("cannot convert %s to integer", obj.Type())
	}
}

// jsonToCarrionObject converts JSON data to Carrion objects
func jsonToCarrionObject(data interface{}) object.Object {
	switch v := data.(type) {
	case nil:
		return &object.None{}
	case bool:
		if v {
			return &object.Boolean{Value: true}
		}
		return &object.Boolean{Value: false}
	case float64:
		// Check if it's an integer
		if v == float64(int64(v)) {
			return &object.Integer{Value: int64(v)}
		}
		return &object.Float{Value: v}
	case string:
		return &object.String{Value: v}
	case []interface{}:
		elements := make([]object.Object, len(v))
		for i, elem := range v {
			elements[i] = jsonToCarrionObject(elem)
		}
		return &object.Array{Elements: elements}
	case map[string]interface{}:
		result := &object.Hash{
			Pairs: make(map[object.HashKey]object.HashPair),
		}
		for key, value := range v {
			keyObj := &object.String{Value: key}
			result.Pairs[keyObj.HashKey()] = object.HashPair{
				Key:   keyObj,
				Value: jsonToCarrionObject(value),
			}
		}
		return result
	default:
		return &object.Error{Message: fmt.Sprintf("Unsupported JSON type: %T", v)}
	}
}

// GetBuiltins returns a copy of the built-ins map for external access
func GetBuiltins() map[string]*object.Builtin {
	result := make(map[string]*object.Builtin)
	for name, builtin := range builtins {
		result[name] = builtin
	}
	return result
}
