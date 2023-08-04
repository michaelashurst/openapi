package converter

import (
	"fmt"
	"strings"
	"time"
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
		return "0000-00-00T00:00:00Z"
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

func (con stringConverter) ExampleToSchema(example interface{}) (schema map[string]interface{}) {
	str := example.(string)
	schema = make(map[string]interface{})

	schema["type"] = "string"

	if len(str) > 0 {
		if strings.ContainsAny(str, "|") {
			enums := strings.Split(str, "|")
			schema["example"] = enums[0]
			schema["enum"] = enums
			if enums[len(enums)-1] == "" {
				enums = enums[:len(enums)-1]
			}
			return
		}
		_, err := time.Parse("2006-01-02T15:04:05Z", str)
		if err == nil {
			schema["format"] = "date-time"
		}
		schema["example"] = example
	}
	return
}
