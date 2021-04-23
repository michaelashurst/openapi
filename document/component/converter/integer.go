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

func (con integerConverter) ExampleToSchema(example map[string]interface{}) (schema interface{}) {
	return nil
}
