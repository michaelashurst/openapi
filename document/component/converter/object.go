package converter

type objectConverter struct {
}

func (con objectConverter) SchemaToExample(schema map[string]interface{}) interface{} {
	output := make(map[string]interface{})

	props := schema["properties"].(map[string]interface{})
	for key, val := range props {
		conv := NewConverter(val.(map[string]interface{})["type"].(string))
		output[key] = conv.SchemaToExample(val.(map[string]interface{}))
	}

	return output
}

func (con objectConverter) ExampleToSchema(example interface{}) (schema map[string]interface{}) {
	obj := example.(map[string]interface{})
	schema = make(map[string]interface{})
	for key, value := range obj {
		converter := NewConverterFromInterface(value)
		schema[key] = converter.ExampleToSchema(value)
	}
	return
}
