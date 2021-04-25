package converter

type booleanConverter struct{}

func (con booleanConverter) SchemaToExample(schema map[string]interface{}) (example interface{}) {
	if schema["example"] != nil {
		return schema["example"]
	}

	if schema["default"] != nil {
		return schema["default"]
	}

	return false
}

func (con booleanConverter) ExampleToSchema(example interface{}) (schema map[string]interface{}) {
	b := example.(bool)
	schema = make(map[string]interface{})
	schema["type"] = "boolean"
	schema["example"] = b
	return
}
