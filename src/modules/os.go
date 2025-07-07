package modules

import (
	"os"
	"os/exec"
	"time"

	"github.com/javanhut/TheCarrionLanguage/src/object"
)

// Helper function to extract string value from object (handles both direct strings and instance-wrapped strings)
func extractStringOS(obj object.Object) (string, bool) {
	switch v := obj.(type) {
	case *object.String:
		return v.Value, true
	case *object.Instance:
		// Check if this is a String instance with a value property
		if value, exists := v.Env.Get("value"); exists {
			if strVal, ok := value.(*object.String); ok {
				return strVal.Value, true
			}
		}
		// Check if this is a String instance with a direct string representation
		if v.Grimoire != nil && v.Grimoire.Name == "String" {
			if value, exists := v.Env.Get("value"); exists {
				if strVal, ok := value.(*object.String); ok {
					return strVal.Value, true
				}
			}
		}
		return "", false
	case *object.Builtin:
		// Handle builtin functions that might represent strings
		return "", false
	default:
		return "", false
	}
}

// Helper function to extract boolean value from object
func extractBoolOS(obj object.Object) (bool, bool) {
	switch v := obj.(type) {
	case *object.Boolean:
		return v.Value, true
	case *object.Instance:
		// Check if this is a Boolean instance with a value property
		if value, exists := v.Env.Get("value"); exists {
			if boolVal, ok := value.(*object.Boolean); ok {
				return boolVal.Value, true
			}
		}
		return false, false
	default:
		return false, false
	}
}

// Helper function to extract integer value from object
func extractIntOS(obj object.Object) (int64, bool) {
	switch v := obj.(type) {
	case *object.Integer:
		return v.Value, true
	case *object.Instance:
		// Check if this is an Integer instance with a value property
		if value, exists := v.Env.Get("value"); exists {
			if intVal, ok := value.(*object.Integer); ok {
				return intVal.Value, true
			}
		}
		return 0, false
	default:
		return 0, false
	}
}

// Helper function to extract float value from object
func extractFloatOS(obj object.Object) (float64, bool) {
	switch v := obj.(type) {
	case *object.Float:
		return v.Value, true
	case *object.Instance:
		// Check if this is a Float instance with a value property
		if value, exists := v.Env.Get("value"); exists {
			if floatVal, ok := value.(*object.Float); ok {
				return floatVal.Value, true
			}
		}
		return 0, false
	default:
		return 0, false
	}
}

var OSBuiltins = map[string]*object.Builtin{
	"osRunCommand": {
		Fn: func(args ...object.Object) object.Object {
			var command string
			var cmdArgs []string
			var capture bool

			if len(args) < 1 {
				return &object.Error{Message: "osRunCommand requires at least 1 argument (command)"}
			}

			command, ok := extractStringOS(args[0])
			if !ok {
				errMsg := "osRunCommand command must be a STRING, got=" + string(args[0].Type())
				if inst, isInst := args[0].(*object.Instance); isInst {
					if inst.Grimoire != nil {
						errMsg += " (grimoire: " + inst.Grimoire.Name + ")"
					}
					errMsg += " with inspect: " + args[0].Inspect()
				}
				return &object.Error{Message: errMsg}
			}

			if len(args) > 1 {
				arrArg, isArr := args[1].(*object.Array)
				if isArr {
					for _, elem := range arrArg.Elements {
						elemStr, ok := extractStringOS(elem)
						if !ok {
							return &object.Error{Message: "osRunCommand arg array must contain only STRINGs"}
						}
						cmdArgs = append(cmdArgs, elemStr)
					}
				}
			}

			if len(args) > 2 {
				capture, ok = extractBoolOS(args[2])
				if !ok {
					return &object.Error{Message: "osRunCommand third arg must be BOOLEAN for captureOutput"}
				}
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
				return &object.Error{Message: "error running command '" + command + "': " + err.Error()}
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
				return &object.Error{Message: "osGetEnv requires 1 argument: key"}
			}
			keyStr, ok := extractStringOS(args[0])
			if !ok {
				return &object.Error{Message: "osGetEnv argument must be STRING, got=" + string(args[0].Type())}
			}
			val := os.Getenv(keyStr)
			return &object.String{Value: val}
		},
	},

	"osSetEnv": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "osSetEnv requires 2 arguments: key, value"}
			}
			key, okKey := extractStringOS(args[0])
			val, okVal := extractStringOS(args[1])
			if !okKey || !okVal {
				return &object.Error{Message: "osSetEnv arguments must be STRINGS, got types: " + string(args[0].Type()) + " and " + string(args[1].Type())}
			}
			err := os.Setenv(key, val)
			if err != nil {
				return &object.Error{Message: "failed to set env var: " + err.Error()}
			}
			return &object.None{}
		},
	},

	"osGetCwd": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return &object.Error{Message: "osGetCwd takes no arguments"}
			}
			dir, err := os.Getwd()
			if err != nil {
				return &object.Error{Message: "failed to get current directory: " + err.Error()}
			}
			return &object.String{Value: dir}
		},
	},

	"osChdir": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "osChdir requires 1 argument: directory path"}
			}
			dirPath, ok := extractStringOS(args[0])
			if !ok {
				return &object.Error{Message: "osChdir argument must be STRING"}
			}
			err := os.Chdir(dirPath)
			if err != nil {
				return &object.Error{Message: "failed to chdir to '" + dirPath + "': " + err.Error()}
			}
			return &object.None{}
		},
	},

	"osSleep": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "osSleep requires 1 argument: seconds (INT or FLOAT)"}
			}

			// Try integer first
			if intVal, ok := extractIntOS(args[0]); ok {
				time.Sleep(time.Duration(intVal) * time.Second)
			} else if floatVal, ok := extractFloatOS(args[0]); ok {
				nanos := int64(floatVal * 1_000_000_000)
				time.Sleep(time.Duration(nanos))
			} else {
				return &object.Error{Message: "osSleep argument must be INTEGER or FLOAT, got " + string(args[0].Type())}
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
				dirStr, ok := extractStringOS(args[0])
				if !ok {
					return &object.Error{Message: "osListDir argument must be STRING, got=" + string(args[0].Type()) + " with value: " + args[0].Inspect()}
				}
				dir = dirStr
			} else {
				return &object.Error{Message: "osListDir requires 0 or 1 arguments"}
			}

			entries, err := os.ReadDir(dir)
			if err != nil {
				return &object.Error{Message: "failed to read directory '" + dir + "': " + err.Error()}
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
				return &object.Error{Message: "osRemove requires 1 argument: path"}
			}
			pathStr, ok := extractStringOS(args[0])
			if !ok {
				return &object.Error{Message: "osRemove argument must be STRING"}
			}
			err := os.Remove(pathStr)
			if err != nil {
				return &object.Error{Message: "failed to remove '" + pathStr + "': " + err.Error()}
			}
			return &object.None{}
		},
	},

	"osMkdir": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return &object.Error{Message: "osMkdir requires 1 or 2 arguments: path, [perm int]"}
			}
			pathStr, ok := extractStringOS(args[0])
			if !ok {
				return &object.Error{Message: "osMkdir path must be STRING, got=" + string(args[0].Type())}
			}
			perm := os.FileMode(0755)
			if len(args) == 2 {
				permVal, ok := extractIntOS(args[1])
				if !ok {
					return &object.Error{Message: "osMkdir second arg must be an INTEGER for permissions"}
				}
				perm = os.FileMode(permVal)
			}
			err := os.Mkdir(pathStr, perm)
			if err != nil {
				return &object.Error{Message: "failed to create directory '" + pathStr + "': " + err.Error()}
			}
			return &object.None{}
		},
	},

	"osExpandEnv": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "osExpandEnv requires 1 argument: string"}
			}
			strArg, ok := extractStringOS(args[0])
			if !ok {
				return &object.Error{Message: "osExpandEnv argument must be STRING, got " + string(args[0].Type())}
			}
			expanded := os.ExpandEnv(strArg)
			return &object.String{Value: expanded}
		},
	},
}