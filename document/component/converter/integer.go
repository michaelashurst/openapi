package converter

type integerConverter struct{}

func (con integerConverter) SchemaToExample(schema map[string]interface{}) (example interface{}) {
	if schema["example"] != nil {
		return schema["example"]
	}

	if schema["default"] != nil {
		return schema["default"]
	}

	return 0
}

func (con integerConverter) ExampleToSchema(example interface{}) (schema map[string]interface{}) {
	i := example.(int64)
	schema = make(map[string]interface{})
	schema["type"] = "integer"
	if i > 0 {
		schema["example"] = i
	}
	return
}
