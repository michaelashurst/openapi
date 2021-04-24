package converter

import (
	"fmt"
	"strings"
)

type stringConverter struct {
}

func (con stringConverter) SchemaToExample(schema map[string]interface{}) (example interface{}) {
	if schema["example"] != nil {
		return schema["example"]
	}

	if schema["default"] != nil {
		return schema["default"]
	}

	if schema["format"] == "date-time" {
		return "0000-00-00T00:00.000Z"
	}

	if schema["enum"] != nil {
		enums := schema["enum"]
		var values []string

		for _, enum := range enums.([]interface{}) {
			values = append(values, fmt.Sprintf("%s", enum))
		}
		return strings.Join(values, "|")
	}
	return ""
}

func (con stringConverter) ExampleToSchema(example map[string]interface{}) (schema interface{}) {
	return nil
}
