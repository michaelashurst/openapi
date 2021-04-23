package component

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/michaelashurst/openapi/document/component/converter"
)

type Schema struct {
	*json.RawMessage
}

func (s Schema) GenerateExample() (example Example) {
	output := make(map[string]interface{})

	jstr, err := s.MarshalJSON()
	if err != nil {
		panic(err)
	}

	var jobj map[string]interface{}
	err = json.Unmarshal(jstr, &jobj)
	if err != nil {
		panic(err)
	}

	converter := converter.NewConverter(jobj["type"].(string))
	output = converter.SchemaToExample(jobj).(map[string]interface{})
	// if jobj["type"] == "object" {
	// 	props := jobj["properties"].(map[string]interface{})
	// 	for key, val := range props {
	// 		if val.(map[string]interface{})["type"] != "object" {
	// 			output[key] = getDefaultValue(val.(map[string]interface{}))
	// 		}
	// 	}
	// }

	outputData, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}

	example = Example{&json.RawMessage{}}
	err = example.UnmarshalJSON(outputData)
	if err != nil {
		panic(err)
	}
	return
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
