package converter

type numberConverter struct{}

func (con numberConverter) SchemaToExample(schema map[string]interface{}) (example interface{}) {
	if schema["example"] != nil {
		return schema["example"]
	}

	if schema["default"] != nil {
		return schema["default"]
	}

	return 0
}

func (con numberConverter) ExampleToSchema(example interface{}) (schema map[string]interface{}) {
	f := example.(float64)
	schema = make(map[string]interface{})
	if f == float64(int64(f)) {
		return integerConverter{}.ExampleToSchema(int64(f))
	}

	schema["type"] = "number"
	schema["format"] = "float"
	if f > 0 {
		schema["example"] = f
	}
	return
}
