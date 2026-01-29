package modules

import (
	"bytes"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"

	"github.com/javanhut/TheCarrionLanguage/src/object"
)

// Supported encodings map
var encodingMap = map[string]encoding.Encoding{
	"utf-8":        encoding.Nop,
	"utf8":         encoding.Nop,
	"utf-16":       unicode.UTF16(unicode.BigEndian, unicode.UseBOM),
	"utf16":        unicode.UTF16(unicode.BigEndian, unicode.UseBOM),
	"utf-16le":     unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM),
	"utf16le":      unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM),
	"utf-16be":     unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM),
	"utf16be":      unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM),
	"latin-1":      charmap.ISO8859_1,
	"latin1":       charmap.ISO8859_1,
	"iso-8859-1":   charmap.ISO8859_1,
	"iso88591":     charmap.ISO8859_1,
	"windows-1252": charmap.Windows1252,
	"windows1252":  charmap.Windows1252,
	"cp1252":       charmap.Windows1252,
	"ascii":        encoding.Nop,
}

// Helper function to extract string value from object
func extractStringEncoding(obj object.Object) (string, bool) {
	switch v := obj.(type) {
	case *object.String:
		return v.Value, true
	case *object.Instance:
		if value, exists := v.Env.Get("value"); exists {
			if strVal, ok := value.(*object.String); ok {
				return strVal.Value, true
			}
		}
		if v.Grimoire != nil && v.Grimoire.Name == "String" {
			if value, exists := v.Env.Get("value"); exists {
				if strVal, ok := value.(*object.String); ok {
					return strVal.Value, true
				}
			}
		}
		return v.Inspect(), true
	default:
		return obj.Inspect(), true
	}
}

// getEncoding retrieves the encoding by name (case-insensitive)
func getEncoding(name string) (encoding.Encoding, bool) {
	enc, ok := encodingMap[strings.ToLower(name)]
	return enc, ok
}

// detectBOM detects encoding from Byte Order Mark
func detectBOM(data []byte) string {
	// UTF-32 BE BOM
	if len(data) >= 4 && data[0] == 0x00 && data[1] == 0x00 && data[2] == 0xFE && data[3] == 0xFF {
		return "utf-32be"
	}
	// UTF-32 LE BOM
	if len(data) >= 4 && data[0] == 0xFF && data[1] == 0xFE && data[2] == 0x00 && data[3] == 0x00 {
		return "utf-32le"
	}
	// UTF-8 BOM
	if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		return "utf-8"
	}
	// UTF-16 BE BOM
	if len(data) >= 2 && data[0] == 0xFE && data[1] == 0xFF {
		return "utf-16be"
	}
	// UTF-16 LE BOM
	if len(data) >= 2 && data[0] == 0xFF && data[1] == 0xFE {
		return "utf-16le"
	}
	// Default to UTF-8 if no BOM detected
	return "utf-8"
}

// DecodeBytes decodes bytes from the specified encoding to UTF-8 string
func DecodeBytes(data []byte, encodingName string) (string, error) {
	enc, ok := getEncoding(encodingName)
	if !ok {
		return "", &EncodingError{Message: "unsupported encoding: " + encodingName}
	}

	decoder := enc.NewDecoder()
	decoded, err := decoder.Bytes(data)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// EncodeString encodes a UTF-8 string to bytes in the specified encoding
func EncodeString(text string, encodingName string) ([]byte, error) {
	enc, ok := getEncoding(encodingName)
	if !ok {
		return nil, &EncodingError{Message: "unsupported encoding: " + encodingName}
	}

	encoder := enc.NewEncoder()
	encoded, err := encoder.Bytes([]byte(text))
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

// EncodingError represents an encoding-related error
type EncodingError struct {
	Message string
}

func (e *EncodingError) Error() string {
	return e.Message
}

// EncodingBuiltins contains encoding-related builtin functions
var EncodingBuiltins = map[string]*object.Builtin{
	"encodingDecode": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return &object.Error{Message: "encodingDecode requires 1-2 arguments: data, [encoding]"}
			}

			var data []byte

			// Handle data argument - can be string (raw bytes) or array of integers
			switch v := args[0].(type) {
			case *object.String:
				data = []byte(v.Value)
			case *object.Array:
				data = make([]byte, len(v.Elements))
				for i, elem := range v.Elements {
					if intVal, ok := elem.(*object.Integer); ok {
						if intVal.Value < 0 || intVal.Value > 255 {
							return &object.Error{Message: "encodingDecode: byte values must be 0-255"}
						}
						data[i] = byte(intVal.Value)
					} else {
						return &object.Error{Message: "encodingDecode: array must contain integers"}
					}
				}
			default:
				return &object.Error{Message: "encodingDecode: data must be a string or array of bytes"}
			}

			encodingName := "utf-8"
			if len(args) == 2 {
				encStr, ok := extractStringEncoding(args[1])
				if !ok {
					return &object.Error{Message: "encodingDecode: encoding must be a string"}
				}
				encodingName = encStr
			}

			decoded, err := DecodeBytes(data, encodingName)
			if err != nil {
				return &object.Error{Message: "encodingDecode: " + err.Error()}
			}

			return &object.String{Value: decoded}
		},
	},

	"encodingEncode": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return &object.Error{Message: "encodingEncode requires 1-2 arguments: text, [encoding]"}
			}

			text, ok := extractStringEncoding(args[0])
			if !ok {
				return &object.Error{Message: "encodingEncode: text must be a string"}
			}

			encodingName := "utf-8"
			if len(args) == 2 {
				encStr, ok := extractStringEncoding(args[1])
				if !ok {
					return &object.Error{Message: "encodingEncode: encoding must be a string"}
				}
				encodingName = encStr
			}

			encoded, err := EncodeString(text, encodingName)
			if err != nil {
				return &object.Error{Message: "encodingEncode: " + err.Error()}
			}

			// Return as array of integers (bytes)
			elements := make([]object.Object, len(encoded))
			for i, b := range encoded {
				elements[i] = &object.Integer{Value: int64(b)}
			}
			return &object.Array{Elements: elements}
		},
	},

	"encodingList": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return &object.Error{Message: "encodingList takes no arguments"}
			}

			// Return unique encoding names (canonical names)
			encodings := []string{
				"utf-8",
				"utf-16",
				"utf-16le",
				"utf-16be",
				"latin-1",
				"iso-8859-1",
				"windows-1252",
				"ascii",
			}

			elements := make([]object.Object, len(encodings))
			for i, enc := range encodings {
				elements[i] = &object.String{Value: enc}
			}
			return &object.Array{Elements: elements}
		},
	},

	"encodingDetectBOM": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "encodingDetectBOM requires 1 argument: data"}
			}

			var data []byte

			// Handle data argument - can be string (raw bytes) or array of integers
			switch v := args[0].(type) {
			case *object.String:
				data = []byte(v.Value)
			case *object.Array:
				data = make([]byte, len(v.Elements))
				for i, elem := range v.Elements {
					if intVal, ok := elem.(*object.Integer); ok {
						if intVal.Value < 0 || intVal.Value > 255 {
							return &object.Error{Message: "encodingDetectBOM: byte values must be 0-255"}
						}
						data[i] = byte(intVal.Value)
					} else {
						return &object.Error{Message: "encodingDetectBOM: array must contain integers"}
					}
				}
			default:
				return &object.Error{Message: "encodingDetectBOM: data must be a string or array of bytes"}
			}

			detected := detectBOM(data)
			return &object.String{Value: detected}
		},
	},

	"encodingStripBOM": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "encodingStripBOM requires 1 argument: data"}
			}

			var data []byte

			switch v := args[0].(type) {
			case *object.String:
				data = []byte(v.Value)
			case *object.Array:
				data = make([]byte, len(v.Elements))
				for i, elem := range v.Elements {
					if intVal, ok := elem.(*object.Integer); ok {
						if intVal.Value < 0 || intVal.Value > 255 {
							return &object.Error{Message: "encodingStripBOM: byte values must be 0-255"}
						}
						data[i] = byte(intVal.Value)
					} else {
						return &object.Error{Message: "encodingStripBOM: array must contain integers"}
					}
				}
			default:
				return &object.Error{Message: "encodingStripBOM: data must be a string or array of bytes"}
			}

			// Strip known BOMs
			// UTF-8 BOM
			if bytes.HasPrefix(data, []byte{0xEF, 0xBB, 0xBF}) {
				data = data[3:]
			}
			// UTF-16 BE BOM
			if bytes.HasPrefix(data, []byte{0xFE, 0xFF}) {
				data = data[2:]
			}
			// UTF-16 LE BOM
			if bytes.HasPrefix(data, []byte{0xFF, 0xFE}) {
				data = data[2:]
			}

			// Return as array of integers (bytes)
			elements := make([]object.Object, len(data))
			for i, b := range data {
				elements[i] = &object.Integer{Value: int64(b)}
			}
			return &object.Array{Elements: elements}
		},
	},
}
