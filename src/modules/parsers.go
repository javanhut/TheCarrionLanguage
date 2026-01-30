package modules

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/javanhut/TheCarrionLanguage/src/object"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
)

// Helper function to extract string value from object
func extractStringParser(obj object.Object) (string, bool) {
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

// convertToCarrionObject converts Go values to Carrion objects
func convertToCarrionObject(data interface{}) object.Object {
	switch v := data.(type) {
	case nil:
		return &object.None{}
	case bool:
		return &object.Boolean{Value: v}
	case int:
		return &object.Integer{Value: int64(v)}
	case int64:
		return &object.Integer{Value: v}
	case float64:
		return &object.Float{Value: v}
	case float32:
		return &object.Float{Value: float64(v)}
	case string:
		return &object.String{Value: v}
	case []interface{}:
		elements := make([]object.Object, len(v))
		for i, elem := range v {
			elements[i] = convertToCarrionObject(elem)
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
				Value: convertToCarrionObject(value),
			}
		}
		return result
	case map[interface{}]interface{}:
		// YAML sometimes returns this type
		result := &object.Hash{
			Pairs: make(map[object.HashKey]object.HashPair),
		}
		for key, value := range v {
			var keyStr string
			switch k := key.(type) {
			case string:
				keyStr = k
			default:
				// Convert any other type to string
				keyStr = fmt.Sprintf("%v", k)
			}
			keyObj := &object.String{Value: keyStr}
			result.Pairs[keyObj.HashKey()] = object.HashPair{
				Key:   keyObj,
				Value: convertToCarrionObject(value),
			}
		}
		return result
	default:
		// Try to handle as string
		return &object.String{Value: ""}
	}
}

var ParserBuiltins = map[string]*object.Builtin{
	// ==================== JSON ====================
	"jsonReadFile": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "jsonReadFile requires 1 argument: path"}
			}

			pathStr, ok := extractStringParser(args[0])
			if !ok {
				return &object.Error{Message: "jsonReadFile: path must be a string"}
			}

			data, err := os.ReadFile(pathStr)
			if err != nil {
				return &object.Error{Message: "jsonReadFile: failed to read file '" + pathStr + "': " + err.Error()}
			}

			var result interface{}
			if err := json.Unmarshal(data, &result); err != nil {
				return &object.Error{Message: "jsonReadFile: failed to parse JSON: " + err.Error()}
			}

			return convertToCarrionObject(result)
		},
	},

	"jsonParse": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "jsonParse requires 1 argument: jsonString"}
			}

			jsonStr, ok := extractStringParser(args[0])
			if !ok {
				return &object.Error{Message: "jsonParse: argument must be a string"}
			}

			var result interface{}
			if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
				return &object.Error{Message: "jsonParse: failed to parse JSON: " + err.Error()}
			}

			return convertToCarrionObject(result)
		},
	},

	// ==================== YAML ====================
	"yamlReadFile": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "yamlReadFile requires 1 argument: path"}
			}

			pathStr, ok := extractStringParser(args[0])
			if !ok {
				return &object.Error{Message: "yamlReadFile: path must be a string"}
			}

			data, err := os.ReadFile(pathStr)
			if err != nil {
				return &object.Error{Message: "yamlReadFile: failed to read file '" + pathStr + "': " + err.Error()}
			}

			var result interface{}
			if err := yaml.Unmarshal(data, &result); err != nil {
				return &object.Error{Message: "yamlReadFile: failed to parse YAML: " + err.Error()}
			}

			return convertToCarrionObject(result)
		},
	},

	"yamlParse": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "yamlParse requires 1 argument: yamlString"}
			}

			yamlStr, ok := extractStringParser(args[0])
			if !ok {
				return &object.Error{Message: "yamlParse: argument must be a string"}
			}

			var result interface{}
			if err := yaml.Unmarshal([]byte(yamlStr), &result); err != nil {
				return &object.Error{Message: "yamlParse: failed to parse YAML: " + err.Error()}
			}

			return convertToCarrionObject(result)
		},
	},

	// ==================== TOML ====================
	"tomlReadFile": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "tomlReadFile requires 1 argument: path"}
			}

			pathStr, ok := extractStringParser(args[0])
			if !ok {
				return &object.Error{Message: "tomlReadFile: path must be a string"}
			}

			var result map[string]interface{}
			if _, err := toml.DecodeFile(pathStr, &result); err != nil {
				return &object.Error{Message: "tomlReadFile: failed to parse TOML: " + err.Error()}
			}

			return convertToCarrionObject(result)
		},
	},

	"tomlParse": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "tomlParse requires 1 argument: tomlString"}
			}

			tomlStr, ok := extractStringParser(args[0])
			if !ok {
				return &object.Error{Message: "tomlParse: argument must be a string"}
			}

			var result map[string]interface{}
			if _, err := toml.Decode(tomlStr, &result); err != nil {
				return &object.Error{Message: "tomlParse: failed to parse TOML: " + err.Error()}
			}

			return convertToCarrionObject(result)
		},
	},

	// ==================== XML ====================
	"xmlReadFile": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "xmlReadFile requires 1 argument: path"}
			}

			pathStr, ok := extractStringParser(args[0])
			if !ok {
				return &object.Error{Message: "xmlReadFile: path must be a string"}
			}

			data, err := os.ReadFile(pathStr)
			if err != nil {
				return &object.Error{Message: "xmlReadFile: failed to read file '" + pathStr + "': " + err.Error()}
			}

			return parseXMLToCarrion(data)
		},
	},

	"xmlParse": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "xmlParse requires 1 argument: xmlString"}
			}

			xmlStr, ok := extractStringParser(args[0])
			if !ok {
				return &object.Error{Message: "xmlParse: argument must be a string"}
			}

			return parseXMLToCarrion([]byte(xmlStr))
		},
	},

	// ==================== INI ====================
	"iniReadFile": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "iniReadFile requires 1 argument: path"}
			}

			pathStr, ok := extractStringParser(args[0])
			if !ok {
				return &object.Error{Message: "iniReadFile: path must be a string"}
			}

			cfg, err := ini.Load(pathStr)
			if err != nil {
				return &object.Error{Message: "iniReadFile: failed to parse INI: " + err.Error()}
			}

			return iniToCarrion(cfg)
		},
	},

	"iniParse": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "iniParse requires 1 argument: iniString"}
			}

			iniStr, ok := extractStringParser(args[0])
			if !ok {
				return &object.Error{Message: "iniParse: argument must be a string"}
			}

			cfg, err := ini.Load([]byte(iniStr))
			if err != nil {
				return &object.Error{Message: "iniParse: failed to parse INI: " + err.Error()}
			}

			return iniToCarrion(cfg)
		},
	},

	// ==================== Properties (Java-style) ====================
	"propertiesReadFile": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "propertiesReadFile requires 1 argument: path"}
			}

			pathStr, ok := extractStringParser(args[0])
			if !ok {
				return &object.Error{Message: "propertiesReadFile: path must be a string"}
			}

			data, err := os.ReadFile(pathStr)
			if err != nil {
				return &object.Error{Message: "propertiesReadFile: failed to read file '" + pathStr + "': " + err.Error()}
			}

			return parseProperties(string(data))
		},
	},

	"propertiesParse": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "propertiesParse requires 1 argument: propertiesString"}
			}

			propStr, ok := extractStringParser(args[0])
			if !ok {
				return &object.Error{Message: "propertiesParse: argument must be a string"}
			}

			return parseProperties(propStr)
		},
	},
}

// parseXMLToCarrion converts XML to a Carrion Hash structure
func parseXMLToCarrion(data []byte) object.Object {
	decoder := xml.NewDecoder(strings.NewReader(string(data)))

	// Build a nested structure from XML
	root := make(map[string]interface{})
	var stack []map[string]interface{}
	var currentKey string

	stack = append(stack, root)

	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}

		switch t := token.(type) {
		case xml.StartElement:
			newElement := make(map[string]interface{})

			// Add attributes
			if len(t.Attr) > 0 {
				attrs := make(map[string]interface{})
				for _, attr := range t.Attr {
					attrs[attr.Name.Local] = attr.Value
				}
				newElement["@attributes"] = attrs
			}

			current := stack[len(stack)-1]
			currentKey = t.Name.Local

			// Handle multiple elements with same name
			if existing, exists := current[currentKey]; exists {
				switch v := existing.(type) {
				case []interface{}:
					current[currentKey] = append(v, newElement)
				default:
					current[currentKey] = []interface{}{v, newElement}
				}
			} else {
				current[currentKey] = newElement
			}

			stack = append(stack, newElement)

		case xml.EndElement:
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}

		case xml.CharData:
			text := strings.TrimSpace(string(t))
			if text != "" && len(stack) > 0 {
				current := stack[len(stack)-1]
				current["#text"] = text
			}
		}
	}

	return convertToCarrionObject(root)
}

// iniToCarrion converts an INI file to a Carrion Hash
func iniToCarrion(cfg *ini.File) object.Object {
	result := &object.Hash{
		Pairs: make(map[object.HashKey]object.HashPair),
	}

	for _, section := range cfg.Sections() {
		sectionName := section.Name()
		sectionHash := &object.Hash{
			Pairs: make(map[object.HashKey]object.HashPair),
		}

		for _, key := range section.Keys() {
			keyObj := &object.String{Value: key.Name()}
			valueObj := &object.String{Value: key.Value()}
			sectionHash.Pairs[keyObj.HashKey()] = object.HashPair{
				Key:   keyObj,
				Value: valueObj,
			}
		}

		sectionKeyObj := &object.String{Value: sectionName}
		result.Pairs[sectionKeyObj.HashKey()] = object.HashPair{
			Key:   sectionKeyObj,
			Value: sectionHash,
		}
	}

	return result
}

// parseProperties parses Java-style properties files
func parseProperties(data string) object.Object {
	result := &object.Hash{
		Pairs: make(map[object.HashKey]object.HashPair),
	}

	lines := strings.Split(data, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "!") {
			continue
		}

		// Find separator (= or :)
		sepIdx := -1
		for i, c := range line {
			if c == '=' || c == ':' {
				sepIdx = i
				break
			}
		}

		if sepIdx > 0 {
			key := strings.TrimSpace(line[:sepIdx])
			value := strings.TrimSpace(line[sepIdx+1:])

			keyObj := &object.String{Value: key}
			valueObj := &object.String{Value: value}
			result.Pairs[keyObj.HashKey()] = object.HashPair{
				Key:   keyObj,
				Value: valueObj,
			}
		}
	}

	return result
}
