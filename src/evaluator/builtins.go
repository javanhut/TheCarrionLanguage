package evaluator

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/peterh/liner"

	"github.com/javanhut/TheCarrionLanguage/src/object"
)

func flushInputBuffer() {
	var discard string
	fmt.Scanln(&discard)
}

var LineReader *liner.State

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
			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"print": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect(), " ")
			}
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
			var input string
			fmt.Scanln(&input)
			flushInputBuffer()
			return &object.String{Value: input}
		},
	},
	"type": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			obj := args[0]
			return &object.String{Value: string(obj.Type())}
		},
	},
	"int": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				value, err := strconv.Atoi(arg.Value)
				if err != nil {
					return newError("cannot convert string to int: %s", err)
				}
				return &object.Integer{Value: int64(value)}
			case *object.Float:
				return &object.Integer{Value: int64(arg.Value)}
			case *object.Integer:
				return arg
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
			switch arg := args[0].(type) {
			case *object.String:
				value, err := strconv.Atoi(arg.Value)
				if err != nil {
					return newError("cannot convert string to int: %s", err)
				}
				return &object.Integer{Value: int64(value)}
			case *object.Float:
				return &object.Integer{Value: int64(arg.Value)}
			case *object.Integer:
				return arg
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
			switch arg := args[0].(type) {
			case *object.String:
				value, err := strconv.ParseFloat(arg.Value, 64)
				if err != nil {
					return newError("cannot convert string to float: %s", err)
				}
				return &object.Float{Value: value}
			case *object.Integer:
				return &object.Float{Value: float64(arg.Value)}
			case *object.Float:
				return arg
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
			return &object.String{Value: args[0].Inspect()}
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

				stopObj, ok := args[0].(*object.Integer)
				if !ok {
					return newError("argument to `range` must be INTEGER, got=%s", args[0].Type())
				}
				start, stop, step = 0, stopObj.Value, 1
			case 2:

				startObj, ok1 := args[0].(*object.Integer)
				stopObj, ok2 := args[1].(*object.Integer)
				if !ok1 || !ok2 {
					return newError(
						"arguments to `range` must be INTEGER, got=%s and %s",
						args[0].Type(),
						args[1].Type(),
					)
				}
				start, stop, step = startObj.Value, stopObj.Value, 1
			case 3:

				startObj, ok1 := args[0].(*object.Integer)
				stopObj, ok2 := args[1].(*object.Integer)
				stepObj, ok3 := args[2].(*object.Integer)
				if !ok1 || !ok2 || !ok3 {
					return newError(
						"arguments to `range` must be INTEGER, got=%s, %s, %s",
						args[0].Type(),
						args[1].Type(),
						args[2].Type(),
					)
				}
				start, stop, step = startObj.Value, stopObj.Value, stepObj.Value
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

	"osRunCommand": {
		Fn: func(args ...object.Object) object.Object {
			var command string
			var cmdArgs []string
			var capture bool

			if len(args) < 1 {
				return newError("osRunCommand requires at least 1 argument (command)")
			}

			strArg, ok := args[0].(*object.String)
			if !ok {
				return newError("osRunCommand command must be a STRING, got=%s", args[0].Type())
			}
			command = strArg.Value

			if len(args) > 1 {
				arrArg, isArr := args[1].(*object.Array)
				if isArr {
					for _, elem := range arrArg.Elements {
						strElem, ok := elem.(*object.String)
						if !ok {
							return newError("osRunCommand arg array must contain only STRINGs")
						}
						cmdArgs = append(cmdArgs, strElem.Value)
					}
				}
			}

			if len(args) > 2 {
				boolArg, isBool := args[2].(*object.Boolean)
				if !isBool {
					return newError("osRunCommand third arg must be BOOLEAN for captureOutput")
				}
				capture = boolArg.Value
			}

			cmd := exec.Command(command, cmdArgs...)
			var outputBytes []byte
			var err error
			if capture {
				outputBytes, err = cmd.CombinedOutput()
			} else {
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
			}

			if err != nil {
				return newError("error running command '%s': %s", command, err)
			}

			if capture {
				return &object.String{Value: string(outputBytes)}
			}
			return &object.None{}
		},
	},

	"osGetEnv": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("osGetEnv requires 1 argument: key")
			}
			keyArg, ok := args[0].(*object.String)
			if !ok {
				return newError("osGetEnv argument must be STRING, got=%s", args[0].Type())
			}
			val := os.Getenv(keyArg.Value)
			return &object.String{Value: val}
		},
	},

	"osSetEnv": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("osSetEnv requires 2 arguments: key, value")
			}
			key, okKey := args[0].(*object.String)
			val, okVal := args[1].(*object.String)
			if !okKey || !okVal {
				return newError("osSetEnv arguments must be STRINGS")
			}
			err := os.Setenv(key.Value, val.Value)
			if err != nil {
				return newError("failed to set env var: %s", err)
			}
			return &object.None{}
		},
	},

	"osGetCwd": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return newError("osGetCwd takes no arguments")
			}
			dir, err := os.Getwd()
			if err != nil {
				return newError("failed to get current directory: %s", err)
			}
			return &object.String{Value: dir}
		},
	},

	"osChdir": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("osChdir requires 1 argument: directory path")
			}
			dirArg, ok := args[0].(*object.String)
			if !ok {
				return newError("osChdir argument must be STRING")
			}
			err := os.Chdir(dirArg.Value)
			if err != nil {
				return newError("failed to chdir to '%s': %s", dirArg.Value, err)
			}
			return &object.None{}
		},
	},

	"osSleep": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("osSleep requires 1 argument: seconds (INT or FLOAT)")
			}

			switch val := args[0].(type) {
			case *object.Integer:
				time.Sleep(time.Duration(val.Value) * time.Second)
			case *object.Float:

				nanos := int64(val.Value * 1_000_000_000)
				time.Sleep(time.Duration(nanos))
			default:
				return newError("osSleep argument must be INTEGER or FLOAT, got %s", args[0].Type())
			}

			return &object.None{}
		},
	},

	"osListDir": {
		Fn: func(args ...object.Object) object.Object {
			var dir string
			if len(args) == 0 {
				dir = "."
			} else if len(args) == 1 {
				strArg, ok := args[0].(*object.String)
				if !ok {
					return newError("osListDir argument must be STRING, got=%s", args[0].Type())
				}
				dir = strArg.Value
			} else {
				return newError("osListDir requires 0 or 1 arguments, got=%d", len(args))
			}

			entries, err := os.ReadDir(dir)
			if err != nil {
				return newError("failed to read directory '%s': %s", dir, err)
			}

			var results []object.Object
			for _, e := range entries {
				results = append(results, &object.String{Value: e.Name()})
			}
			return &object.Array{Elements: results}
		},
	},

	"osRemove": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("osRemove requires 1 argument: path")
			}
			pathArg, ok := args[0].(*object.String)
			if !ok {
				return newError("osRemove argument must be STRING")
			}
			err := os.Remove(pathArg.Value)
			if err != nil {
				return newError("failed to remove '%s': %s", pathArg.Value, err)
			}
			return &object.None{}
		},
	},

	"osMkdir": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return newError("osMkdir requires 1 or 2 arguments: path, [perm int]")
			}
			pathArg, ok := args[0].(*object.String)
			if !ok {
				return newError("osMkdir path must be STRING, got=%s", args[0].Type())
			}
			perm := os.FileMode(0755)
			if len(args) == 2 {
				intArg, ok := args[1].(*object.Integer)
				if !ok {
					return newError("osMkdir second arg must be an INTEGER for permissions")
				}
				perm = os.FileMode(intArg.Value)
			}
			err := os.Mkdir(pathArg.Value, perm)
			if err != nil {
				return newError("failed to create directory '%s': %s", pathArg.Value, err)
			}
			return &object.None{}
		},
	},

	"osExpandEnv": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("osExpandEnv requires 1 argument: string")
			}
			strArg, ok := args[0].(*object.String)
			if !ok {
				return newError("osExpandEnv argument must be STRING")
			}
			expanded := os.ExpandEnv(strArg.Value)
			return &object.String{Value: expanded}
		},
	},

	"fileRead": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("fileRead requires 1 argument: path")
			}
			pathArg, ok := args[0].(*object.String)
			if !ok {
				return newError("fileRead: path must be a string")
			}
			data, err := os.ReadFile(pathArg.Value)
			if err != nil {
				return newError("failed to read file '%s': %s", pathArg.Value, err)
			}
			return &object.String{Value: string(data)}
		},
	},

	"fileWrite": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("fileWrite requires 2 arguments: path, content")
			}
			pathArg, ok1 := args[0].(*object.String)
			contentArg, ok2 := args[1].(*object.String)
			if !ok1 || !ok2 {
				return newError("fileWrite: path/content must be STRINGs")
			}

			err := os.WriteFile(pathArg.Value, []byte(contentArg.Value), 0644)
			if err != nil {
				return newError("failed to write file '%s': %s", pathArg.Value, err)
			}
			return &object.None{}
		},
	},

	"fileAppend": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("fileAppend requires 2 arguments: path, content")
			}
			pathArg, ok1 := args[0].(*object.String)
			contentArg, ok2 := args[1].(*object.String)
			if !ok1 || !ok2 {
				return newError("fileAppend: path/content must be STRINGs")
			}

			f, err := os.OpenFile(pathArg.Value, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				return newError("failed to open file '%s' for append: %s", pathArg.Value, err)
			}
			defer f.Close()

			_, err = f.WriteString(contentArg.Value)
			if err != nil {
				return newError("failed to append to file '%s': %s", pathArg.Value, err)
			}
			return &object.None{}
		},
	},

	"fileExists": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("fileExists requires 1 argument: path")
			}
			pathArg, ok := args[0].(*object.String)
			if !ok {
				return newError("fileExists: path must be a string")
			}

			_, err := os.Stat(pathArg.Value)
			if err != nil {
				if os.IsNotExist(err) {
					return &object.Boolean{Value: false}
				}

				return newError("error checking fileExists for '%s': %s", pathArg.Value, err)
			}
			return &object.Boolean{Value: true}
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

			arr, ok := args[0].(*object.Array)
			if !ok {
				return newError("enumerate expects an array, got %s", args[0].Type())
			}
			var enumerated []object.Object
			for i, elem := range arr.Elements {

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
}
