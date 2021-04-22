package component

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/michaelashurst/openapi/document/operation"
	"github.com/michaelashurst/openapi/document/parameter"
)

type Components struct {
	Schemas         map[string]*json.RawMessage      `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       map[string]operation.Response    `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      map[string]parameter.Parameter   `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        map[string]parameter.Example     `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   map[string]operation.RequestBody `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         *json.RawMessage                 `json:"headers,omitempty" yaml:"headers,omitempty"`
	SecuritySchemes *json.RawMessage                 `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Links           *json.RawMessage                 `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       *json.RawMessage                 `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
}

func getDefaultValue(properties map[string]interface{}) interface{} {
	if properties["example"] != nil {
		return properties["example"]
	}

	if properties["default"] != nil {
		return properties["default"]
	}

	switch properties["type"] {
	case "string":
		if properties["format"] == "date-time" {
			return "0000-00-00T00:00Z"
		}
		if properties["enum"] != nil {
			enums := properties["enum"]
			var values []string

			for _, enum := range enums.([]interface{}) {
				values = append(values, fmt.Sprintf("%s", enum))
				// val += enum
			}
			return strings.Join(values, "|")
		}
		return ""
	case "integer":
		return 0
	case "boolean":
		return false
	default:
		return ""
	}
}

func ConvertSchemaToExample(schema *json.RawMessage) (example *json.RawMessage) {
	// var output map[string]interface{}
	output := make(map[string]interface{})

	jstr, err := schema.MarshalJSON()
	if err != nil {
		panic(err)
	}

	var jobj map[string]interface{}
	err = json.Unmarshal(jstr, &jobj)
	if err != nil {
		panic(err)
	}

	if jobj["type"] == "object" {
		//This is an object and can keep going
		props := jobj["properties"].(map[string]interface{})
		for key, val := range props {
			if val.(map[string]interface{})["type"] != "object" {
				output[key] = getDefaultValue(val.(map[string]interface{}))
				// output[key] = val.(map[string]interface{})["example"]
			}
		}
	}

	outputData, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}

	example = &json.RawMessage{}
	err = example.UnmarshalJSON(outputData)
	if err != nil {
		panic(err)
	}
	return
}
