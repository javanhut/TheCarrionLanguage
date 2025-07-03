package modules

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/javanhut/TheCarrionLanguage/src/object"
)

var HttpModule = map[string]*object.Builtin{
	"httpGet": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return &object.Error{Message: "httpGet expects 1 or 2 arguments: httpGet(url, [headers])"}
			}

			urlObj, ok := args[0].(*object.String)
			if !ok {
				if instance, ok := args[0].(*object.Instance); ok {
					if value, ok := instance.Env.Get("value"); ok {
						if str, ok := value.(*object.String); ok {
							urlObj = str
						} else {
							return &object.Error{Message: "httpGet expects URL as string"}
						}
					} else {
						return &object.Error{Message: "httpGet expects URL as string"}
					}
				} else {
					return &object.Error{Message: "httpGet expects URL as string"}
				}
			}

			client := &http.Client{
				Timeout: 30 * time.Second,
			}

			req, err := http.NewRequest("GET", urlObj.Value, nil)
			if err != nil {
				return &object.Error{Message: fmt.Sprintf("Failed to create request: %v", err)}
			}

			if len(args) == 2 {
				if err := setHeaders(req, args[1]); err != nil {
					return err
				}
			}

			resp, err := client.Do(req)
			if err != nil {
				return &object.Error{Message: fmt.Sprintf("Request failed: %v", err)}
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return &object.Error{Message: fmt.Sprintf("Failed to read response: %v", err)}
			}

			result := &object.Hash{
				Pairs: make(map[object.HashKey]object.HashPair),
			}

			statusKey := &object.String{Value: "status"}
			result.Pairs[statusKey.HashKey()] = object.HashPair{
				Key:   statusKey,
				Value: &object.Integer{Value: int64(resp.StatusCode)},
			}

			bodyKey := &object.String{Value: "body"}
			result.Pairs[bodyKey.HashKey()] = object.HashPair{
				Key:   bodyKey,
				Value: &object.String{Value: string(body)},
			}

			headersKey := &object.String{Value: "headers"}
			result.Pairs[headersKey.HashKey()] = object.HashPair{
				Key:   headersKey,
				Value: headersToHash(resp.Header),
			}

			return result
		},
	},
	"httpPost": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return &object.Error{Message: "httpPost expects 2 or 3 arguments: httpPost(url, body, [headers])"}
			}

			urlObj, err := extractString(args[0], "URL")
			if err != nil {
				return err
			}

			bodyStr, err := extractString(args[1], "body")
			if err != nil {
				return err
			}

			client := &http.Client{
				Timeout: 30 * time.Second,
			}

			req, reqErr := http.NewRequest("POST", urlObj, strings.NewReader(bodyStr))
			if reqErr != nil {
				return &object.Error{Message: fmt.Sprintf("Failed to create request: %v", reqErr)}
			}

			req.Header.Set("Content-Type", "application/json")

			if len(args) == 3 {
				if err := setHeaders(req, args[2]); err != nil {
					return err
				}
			}

			resp, respErr := client.Do(req)
			if respErr != nil {
				return &object.Error{Message: fmt.Sprintf("Request failed: %v", respErr)}
			}
			defer resp.Body.Close()

			return buildResponse(resp)
		},
	},
	"httpPut": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return &object.Error{Message: "httpPut expects 2 or 3 arguments: httpPut(url, body, [headers])"}
			}

			urlObj, err := extractString(args[0], "URL")
			if err != nil {
				return err
			}

			bodyStr, err := extractString(args[1], "body")
			if err != nil {
				return err
			}

			client := &http.Client{
				Timeout: 30 * time.Second,
			}

			req, reqErr := http.NewRequest("PUT", urlObj, strings.NewReader(bodyStr))
			if reqErr != nil {
				return &object.Error{Message: fmt.Sprintf("Failed to create request: %v", reqErr)}
			}

			req.Header.Set("Content-Type", "application/json")

			if len(args) == 3 {
				if err := setHeaders(req, args[2]); err != nil {
					return err
				}
			}

			resp, respErr := client.Do(req)
			if respErr != nil {
				return &object.Error{Message: fmt.Sprintf("Request failed: %v", respErr)}
			}
			defer resp.Body.Close()

			return buildResponse(resp)
		},
	},
	"httpDelete": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return &object.Error{Message: "httpDelete expects 1 or 2 arguments: httpDelete(url, [headers])"}
			}

			urlObj, err := extractString(args[0], "URL")
			if err != nil {
				return err
			}

			client := &http.Client{
				Timeout: 30 * time.Second,
			}

			req, reqErr := http.NewRequest("DELETE", urlObj, nil)
			if reqErr != nil {
				return &object.Error{Message: fmt.Sprintf("Failed to create request: %v", reqErr)}
			}

			if len(args) == 2 {
				if err := setHeaders(req, args[1]); err != nil {
					return err
				}
			}

			resp, respErr := client.Do(req)
			if respErr != nil {
				return &object.Error{Message: fmt.Sprintf("Request failed: %v", respErr)}
			}
			defer resp.Body.Close()

			return buildResponse(resp)
		},
	},
	"httpHead": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return &object.Error{Message: "httpHead expects 1 or 2 arguments: httpHead(url, [headers])"}
			}

			urlObj, err := extractString(args[0], "URL")
			if err != nil {
				return err
			}

			client := &http.Client{
				Timeout: 30 * time.Second,
			}

			req, reqErr := http.NewRequest("HEAD", urlObj, nil)
			if reqErr != nil {
				return &object.Error{Message: fmt.Sprintf("Failed to create request: %v", reqErr)}
			}

			if len(args) == 2 {
				if err := setHeaders(req, args[1]); err != nil {
					return err
				}
			}

			resp, respErr := client.Do(req)
			if respErr != nil {
				return &object.Error{Message: fmt.Sprintf("Request failed: %v", respErr)}
			}
			defer resp.Body.Close()

			result := &object.Hash{
				Pairs: make(map[object.HashKey]object.HashPair),
			}

			statusKey := &object.String{Value: "status"}
			result.Pairs[statusKey.HashKey()] = object.HashPair{
				Key:   statusKey,
				Value: &object.Integer{Value: int64(resp.StatusCode)},
			}

			headersKey := &object.String{Value: "headers"}
			result.Pairs[headersKey.HashKey()] = object.HashPair{
				Key:   headersKey,
				Value: headersToHash(resp.Header),
			}

			return result
		},
	},
	"httpRequest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "httpRequest expects 1 argument: httpRequest(options)"}
			}

			options, ok := args[0].(*object.Hash)
			if !ok {
				return &object.Error{Message: "httpRequest expects options as hash"}
			}

			method, err := getHashString(options, "method")
			if err != nil {
				method = "GET"
			}

			url, err := getHashString(options, "url")
			if err != nil {
				return &object.Error{Message: "httpRequest requires 'url' in options"}
			}

			var bodyReader io.Reader
			if body, err := getHashString(options, "body"); err == nil {
				bodyReader = strings.NewReader(body)
			}

			timeout := 30
			if timeoutVal, err := getHashInt(options, "timeout"); err == nil {
				timeout = int(timeoutVal)
			}

			client := &http.Client{
				Timeout: time.Duration(timeout) * time.Second,
			}

			req, reqErr := http.NewRequest(method, url, bodyReader)
			if reqErr != nil {
				return &object.Error{Message: fmt.Sprintf("Failed to create request: %v", reqErr)}
			}

			if headers, err := getHashValue(options, "headers"); err == nil {
				if err := setHeaders(req, headers); err != nil {
					return err
				}
			}

			resp, respErr := client.Do(req)
			if respErr != nil {
				return &object.Error{Message: fmt.Sprintf("Request failed: %v", respErr)}
			}
			defer resp.Body.Close()

			return buildResponse(resp)
		},
	},
	"httpParseJSON": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "httpParseJSON expects 1 argument: httpParseJSON(jsonString)"}
			}

			jsonStr, err := extractString(args[0], "JSON string")
			if err != nil {
				return err
			}

			var result interface{}
			if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
				return &object.Error{Message: fmt.Sprintf("Failed to parse JSON: %v", err)}
			}

			return jsonToObject(result)
		},
	},
	"httpStringifyJSON": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "httpStringifyJSON expects 1 argument: httpStringifyJSON(object)"}
			}

			data := objectToInterface(args[0])
			jsonBytes, err := json.Marshal(data)
			if err != nil {
				return &object.Error{Message: fmt.Sprintf("Failed to stringify JSON: %v", err)}
			}

			return &object.String{Value: string(jsonBytes)}
		},
	},
	"httpBuildQuery": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "httpBuildQuery expects 1 argument: httpBuildQuery(params)"}
			}

			params, ok := args[0].(*object.Hash)
			if !ok {
				return &object.Error{Message: "httpBuildQuery expects params as hash"}
			}

			var queryParts []string
			for _, pair := range params.Pairs {
				key := pair.Key.Inspect()
				value := pair.Value.Inspect()
				queryParts = append(queryParts, fmt.Sprintf("%s=%s", key, value))
			}

			return &object.String{Value: strings.Join(queryParts, "&")}
		},
	},
}

func extractString(obj object.Object, name string) (string, object.Object) {
	if strObj, ok := obj.(*object.String); ok {
		return strObj.Value, nil
	}
	if instance, ok := obj.(*object.Instance); ok {
		if value, ok := instance.Env.Get("value"); ok {
			if str, ok := value.(*object.String); ok {
				return str.Value, nil
			}
		}
	}
	return "", &object.Error{Message: fmt.Sprintf("%s must be a string", name)}
}

func setHeaders(req *http.Request, headersObj object.Object) object.Object {
	headers, ok := headersObj.(*object.Hash)
	if !ok {
		return &object.Error{Message: "Headers must be a hash"}
	}

	for _, pair := range headers.Pairs {
		key := pair.Key.Inspect()
		value := pair.Value.Inspect()
		req.Header.Set(key, value)
	}

	return nil
}

func headersToHash(headers http.Header) *object.Hash {
	result := &object.Hash{
		Pairs: make(map[object.HashKey]object.HashPair),
	}

	for key, values := range headers {
		keyObj := &object.String{Value: key}
		valueObj := &object.String{Value: strings.Join(values, ", ")}
		result.Pairs[keyObj.HashKey()] = object.HashPair{
			Key:   keyObj,
			Value: valueObj,
		}
	}

	return result
}

func buildResponse(resp *http.Response) object.Object {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to read response: %v", err)}
	}

	result := &object.Hash{
		Pairs: make(map[object.HashKey]object.HashPair),
	}

	statusKey := &object.String{Value: "status"}
	result.Pairs[statusKey.HashKey()] = object.HashPair{
		Key:   statusKey,
		Value: &object.Integer{Value: int64(resp.StatusCode)},
	}

	bodyKey := &object.String{Value: "body"}
	result.Pairs[bodyKey.HashKey()] = object.HashPair{
		Key:   bodyKey,
		Value: &object.String{Value: string(body)},
	}

	headersKey := &object.String{Value: "headers"}
	result.Pairs[headersKey.HashKey()] = object.HashPair{
		Key:   headersKey,
		Value: headersToHash(resp.Header),
	}

	return result
}

func getHashString(hash *object.Hash, key string) (string, error) {
	keyObj := &object.String{Value: key}
	if pair, ok := hash.Pairs[keyObj.HashKey()]; ok {
		if str, ok := pair.Value.(*object.String); ok {
			return str.Value, nil
		}
		return "", fmt.Errorf("value for key %s is not a string", key)
	}
	return "", fmt.Errorf("key %s not found", key)
}

func getHashInt(hash *object.Hash, key string) (int64, error) {
	keyObj := &object.String{Value: key}
	if pair, ok := hash.Pairs[keyObj.HashKey()]; ok {
		if intObj, ok := pair.Value.(*object.Integer); ok {
			return intObj.Value, nil
		}
		return 0, fmt.Errorf("value for key %s is not an integer", key)
	}
	return 0, fmt.Errorf("key %s not found", key)
}

func getHashValue(hash *object.Hash, key string) (object.Object, error) {
	keyObj := &object.String{Value: key}
	if pair, ok := hash.Pairs[keyObj.HashKey()]; ok {
		return pair.Value, nil
	}
	return nil, fmt.Errorf("key %s not found", key)
}

func jsonToObject(data interface{}) object.Object {
	switch v := data.(type) {
	case nil:
		return &object.None{}
	case bool:
		if v {
			return &object.Boolean{Value: true}
		}
		return &object.Boolean{Value: false}
	case float64:
		if v == float64(int64(v)) {
			return &object.Integer{Value: int64(v)}
		}
		return &object.Float{Value: v}
	case string:
		return &object.String{Value: v}
	case []interface{}:
		elements := make([]object.Object, len(v))
		for i, elem := range v {
			elements[i] = jsonToObject(elem)
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
				Value: jsonToObject(value),
			}
		}
		return result
	default:
		return &object.Error{Message: fmt.Sprintf("Unsupported JSON type: %T", v)}
	}
}

func objectToInterface(obj object.Object) interface{} {
	switch o := obj.(type) {
	case *object.None:
		return nil
	case *object.Boolean:
		return o.Value
	case *object.Integer:
		return o.Value
	case *object.Float:
		return o.Value
	case *object.String:
		return o.Value
	case *object.Array:
		result := make([]interface{}, len(o.Elements))
		for i, elem := range o.Elements {
			result[i] = objectToInterface(elem)
		}
		return result
	case *object.Hash:
		result := make(map[string]interface{})
		for _, pair := range o.Pairs {
			key := pair.Key.Inspect()
			result[key] = objectToInterface(pair.Value)
		}
		return result
	default:
		return obj.Inspect()
	}
}