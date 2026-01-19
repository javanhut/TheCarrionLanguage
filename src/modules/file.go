package modules

import (
	"io"
	"os"
	"sync"

	"github.com/javanhut/TheCarrionLanguage/src/object"
)

// Global file handle registry
var (
	fileHandles = make(map[int64]*os.File)
	nextHandle  int64 = 1
	handleMutex sync.RWMutex
)

// Helper function to extract string value from object
func extractStringFile(obj object.Object) (string, bool) {
	switch v := obj.(type) {
	case *object.String:
		return v.Value, true
	case *object.Instance:
		// Handle primitive wrapper instances (like String grimoire instances)
		if value, exists := v.Env.Get("value"); exists {
			if strVal, ok := value.(*object.String); ok {
				return strVal.Value, true
			}
		}
		// Check if it's a String grimoire instance by checking the grimoire name
		if v.Grimoire != nil && v.Grimoire.Name == "String" {
			if value, exists := v.Env.Get("value"); exists {
				if strVal, ok := value.(*object.String); ok {
					return strVal.Value, true
				}
			}
		}
		// Fall back to using the instance's Inspect method for string representation
		return v.Inspect(), true
	default:
		// For any other object type, use its string representation
		return obj.Inspect(), true
	}
}

// Helper function to extract integer value from object
func extractIntFile(obj object.Object) (int64, bool) {
	switch v := obj.(type) {
	case *object.Integer:
		return v.Value, true
	default:
		return 0, false
	}
}

// Get file handle from registry
func getFileHandle(handleID int64) (*os.File, bool) {
	handleMutex.RLock()
	defer handleMutex.RUnlock()
	file, exists := fileHandles[handleID]
	return file, exists
}

// Store file handle in registry
func storeFileHandle(file *os.File) int64 {
	handleMutex.Lock()
	defer handleMutex.Unlock()
	handleID := nextHandle
	nextHandle++
	fileHandles[handleID] = file
	return handleID
}

// Remove file handle from registry
func removeFileHandle(handleID int64) {
	handleMutex.Lock()
	defer handleMutex.Unlock()
	delete(fileHandles, handleID)
}

var FileBuiltins = map[string]*object.Builtin{
	"fileOpen": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return &object.Error{Message: "fileOpen requires 1-2 arguments: path, [mode]"}
			}
			
			pathStr, ok := extractStringFile(args[0])
			if !ok {
				return &object.Error{Message: "fileOpen: path must be a string"}
			}
			
			mode := "r"
			if len(args) == 2 {
				modeStr, ok := extractStringFile(args[1])
				if !ok {
					return &object.Error{Message: "fileOpen: mode must be a string"}
				}
				mode = modeStr
			}
			
			var file *os.File
			var err error
			
			// Handle different modes with conditionals
			switch mode {
			case "r":
				file, err = os.Open(pathStr)
			case "w":
				file, err = os.Create(pathStr)
			case "a":
				file, err = os.OpenFile(pathStr, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			case "r+":
				file, err = os.OpenFile(pathStr, os.O_RDWR, 0644)
			case "w+":
				file, err = os.OpenFile(pathStr, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			case "a+":
				file, err = os.OpenFile(pathStr, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
			default:
				return &object.Error{Message: "fileOpen: unsupported mode '" + mode + "'. Use r, w, a, r+, w+, a+"}
			}
			
			if err != nil {
				return &object.Error{Message: "failed to open file '" + pathStr + "': " + err.Error()}
			}
			
			handleID := storeFileHandle(file)
			return &object.Integer{Value: handleID}
		},
	},

	"fileReadHandle": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return &object.Error{Message: "fileReadHandle requires 1-2 arguments: handleID, [size]"}
			}
			
			handleID, ok := extractIntFile(args[0])
			if !ok {
				return &object.Error{Message: "fileReadHandle: handleID must be an integer"}
			}
			
			file, exists := getFileHandle(handleID)
			if !exists {
				return &object.Error{Message: "fileReadHandle: invalid file handle"}
			}
			
			// Conditional size parameter
			var data []byte
			var err error
			
			if len(args) == 2 {
				size, ok := extractIntFile(args[1])
				if !ok {
					return &object.Error{Message: "fileReadHandle: size must be an integer"}
				}
				if size <= 0 {
					// Read all
					data, err = io.ReadAll(file)
				} else {
					// Read specific size
					data = make([]byte, size)
					n, readErr := file.Read(data)
					if readErr != nil && readErr != io.EOF {
						err = readErr
					} else {
						data = data[:n]
					}
				}
			} else {
				// Read all by default
				data, err = io.ReadAll(file)
			}
			
			if err != nil {
				return &object.Error{Message: "failed to read from file: " + err.Error()}
			}
			
			return &object.String{Value: string(data)}
		},
	},

	"fileWriteHandle": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "fileWriteHandle requires 2 arguments: handleID, content"}
			}
			
			handleID, ok := extractIntFile(args[0])
			if !ok {
				return &object.Error{Message: "fileWriteHandle: handleID must be an integer"}
			}
			
			contentStr, ok := extractStringFile(args[1])
			if !ok {
				return &object.Error{Message: "fileWriteHandle: content must be a string"}
			}
			
			file, exists := getFileHandle(handleID)
			if !exists {
				return &object.Error{Message: "fileWriteHandle: invalid file handle"}
			}
			
			n, err := file.WriteString(contentStr)
			if err != nil {
				return &object.Error{Message: "failed to write to file: " + err.Error()}
			}
			
			return &object.Integer{Value: int64(n)}
		},
	},

	"fileClose": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "fileClose requires 1 argument: handleID"}
			}
			
			handleID, ok := extractIntFile(args[0])
			if !ok {
				return &object.Error{Message: "fileClose: handleID must be an integer"}
			}
			
			file, exists := getFileHandle(handleID)
			if !exists {
				return &object.Error{Message: "fileClose: invalid file handle"}
			}
			
			err := file.Close()
			removeFileHandle(handleID)
			
			if err != nil {
				return &object.Error{Message: "failed to close file: " + err.Error()}
			}
			
			return &object.None{}
		},
	},

	"fileSeek": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return &object.Error{Message: "fileSeek requires 2-3 arguments: handleID, offset, [whence]"}
			}
			
			handleID, ok := extractIntFile(args[0])
			if !ok {
				return &object.Error{Message: "fileSeek: handleID must be an integer"}
			}
			
			offset, ok := extractIntFile(args[1])
			if !ok {
				return &object.Error{Message: "fileSeek: offset must be an integer"}
			}
			
			whence := int(io.SeekStart) // Default to start
			if len(args) == 3 {
				whenceVal, ok := extractIntFile(args[2])
				if !ok {
					return &object.Error{Message: "fileSeek: whence must be an integer"}
				}
				whence = int(whenceVal)
			}
			
			file, exists := getFileHandle(handleID)
			if !exists {
				return &object.Error{Message: "fileSeek: invalid file handle"}
			}
			
			newPos, err := file.Seek(offset, whence)
			if err != nil {
				return &object.Error{Message: "failed to seek in file: " + err.Error()}
			}
			
			return &object.Integer{Value: newPos}
		},
	},

	"fileTell": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "fileTell requires 1 argument: handleID"}
			}
			
			handleID, ok := extractIntFile(args[0])
			if !ok {
				return &object.Error{Message: "fileTell: handleID must be an integer"}
			}
			
			file, exists := getFileHandle(handleID)
			if !exists {
				return &object.Error{Message: "fileTell: invalid file handle"}
			}
			
			pos, err := file.Seek(0, io.SeekCurrent)
			if err != nil {
				return &object.Error{Message: "failed to get file position: " + err.Error()}
			}
			
			return &object.Integer{Value: pos}
		},
	},

	"fileFlush": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "fileFlush requires 1 argument: handleID"}
			}
			
			handleID, ok := extractIntFile(args[0])
			if !ok {
				return &object.Error{Message: "fileFlush: handleID must be an integer"}
			}
			
			file, exists := getFileHandle(handleID)
			if !exists {
				return &object.Error{Message: "fileFlush: invalid file handle"}
			}
			
			err := file.Sync()
			if err != nil {
				return &object.Error{Message: "failed to flush file: " + err.Error()}
			}
			
			return &object.None{}
		},
	},

	// Keep convenience functions for backward compatibility
	"fileReadPath": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "fileReadPath requires 1 argument: path"}
			}
			pathStr, ok := extractStringFile(args[0])
			if !ok {
				return &object.Error{Message: "fileReadPath: path must be a string"}
			}
			data, err := os.ReadFile(pathStr)
			if err != nil {
				return &object.Error{Message: "failed to read file '" + pathStr + "': " + err.Error()}
			}
			return &object.String{Value: string(data)}
		},
	},

	"fileWritePath": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "fileWritePath requires 2 arguments: path, content"}
			}
			pathStr, ok1 := extractStringFile(args[0])
			contentStr, ok2 := extractStringFile(args[1])
			if !ok1 || !ok2 {
				return &object.Error{Message: "fileWritePath: path/content must be strings"}
			}
			err := os.WriteFile(pathStr, []byte(contentStr), 0644)
			if err != nil {
				return &object.Error{Message: "failed to write file '" + pathStr + "': " + err.Error()}
			}
			return &object.None{}
		},
	},

	"fileAppendPath": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "fileAppendPath requires 2 arguments: path, content"}
			}
			pathStr, ok1 := extractStringFile(args[0])
			contentStr, ok2 := extractStringFile(args[1])
			if !ok1 || !ok2 {
				return &object.Error{Message: "fileAppendPath: path/content must be strings"}
			}
			f, err := os.OpenFile(pathStr, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				return &object.Error{Message: "failed to open file '" + pathStr + "' for append: " + err.Error()}
			}
			defer f.Close()
			_, err = f.WriteString(contentStr)
			if err != nil {
				return &object.Error{Message: "failed to append to file '" + pathStr + "': " + err.Error()}
			}
			return &object.None{}
		},
	},

	"fileExists": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "fileExists requires 1 argument: path"}
			}
			pathStr, ok := extractStringFile(args[0])
			if !ok {
				return &object.Error{Message: "fileExists: path must be a string, got " + string(args[0].Type()) + " with value: " + args[0].Inspect()}
			}
			stat, err := os.Stat(pathStr)
			if err != nil {
				if os.IsNotExist(err) {
					return &object.Boolean{Value: false}
				}
				return &object.Error{Message: "error checking fileExists for '" + pathStr + "': " + err.Error()}
			}
			// Debug: log what we found
			fileType := "file"
			if stat.IsDir() {
				fileType = "directory"
			}
			_ = fileType // avoid unused variable warning
			return &object.Boolean{Value: true}
		},
	},

	// Backward compatibility aliases - redirect to path-based functions
	"fileRead": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "fileRead requires 1 argument: path"}
			}
			pathStr, ok := extractStringFile(args[0])
			if !ok {
				return &object.Error{Message: "fileRead: path must be a string"}
			}
			data, err := os.ReadFile(pathStr)
			if err != nil {
				return &object.Error{Message: "failed to read file '" + pathStr + "': " + err.Error()}
			}
			return &object.String{Value: string(data)}
		},
	},

	"fileWrite": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "fileWrite requires 2 arguments: path, content"}
			}
			pathStr, ok1 := extractStringFile(args[0])
			contentStr, ok2 := extractStringFile(args[1])
			if !ok1 || !ok2 {
				return &object.Error{Message: "fileWrite: path/content must be strings"}
			}
			err := os.WriteFile(pathStr, []byte(contentStr), 0644)
			if err != nil {
				return &object.Error{Message: "failed to write file '" + pathStr + "': " + err.Error()}
			}
			return &object.None{}
		},
	},

	"fileAppend": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return &object.Error{Message: "fileAppend requires 2 arguments: path, content"}
			}
			pathStr, ok1 := extractStringFile(args[0])
			contentStr, ok2 := extractStringFile(args[1])
			if !ok1 || !ok2 {
				return &object.Error{Message: "fileAppend: path/content must be strings"}
			}
			f, err := os.OpenFile(pathStr, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				return &object.Error{Message: "failed to open file '" + pathStr + "' for append: " + err.Error()}
			}
			defer f.Close()
			_, err = f.WriteString(contentStr)
			if err != nil {
				return &object.Error{Message: "failed to append to file '" + pathStr + "': " + err.Error()}
			}
			return &object.None{}
		},
	},
}