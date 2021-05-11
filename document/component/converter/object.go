package converter

import "strings"

type objectConverter struct {
}

func (con objectConverter) SchemaToExample(schema map[string]interface{}) interface{} {
	output := make(map[string]interface{})

	props := schema["properties"].(map[string]interface{})
	for key, val := range props {
		var formattedKey string
		if val.(map[string]interface{})["description"] != nil {
			formattedKey = key + " | " + val.(map[string]interface{})["description"].(string)
		} else {
			formattedKey = key
		}
		conv := NewConverter(val.(map[string]interface{})["type"].(string))
		output[formattedKey] = conv.SchemaToExample(val.(map[string]interface{}))
	}

	return output
}

func (con objectConverter) ExampleToSchema(example interface{}) (schema map[string]interface{}) {
	obj := example.(map[string]interface{})
	schema = make(map[string]interface{})
	schema["type"] = "object"
	schema["properties"] = make(map[string]interface{})
	for key, value := range obj {

		converter := NewConverterFromInterface(value)
		sub := converter.ExampleToSchema(value)

		var cleansedKey string
		if strings.Contains(key, "|") {
			parts := strings.Split(key, "|")
			cleansedKey = strings.Trim(parts[0], " ")
			sub["description"] = strings.Trim(parts[1], " ")
		} else {
			cleansedKey = key
		}

		schema["properties"].(map[string]interface{})[cleansedKey] = sub
	}
	return
}
