package converter

import "fmt"

type arrayConverter struct {
}

func (con arrayConverter) SchemaToExample(schema map[string]interface{}) interface{} {
	output := []interface{}{}

	items := schema["items"].(map[string]interface{})
	conv := NewConverter(items["type"].(string))
	fmt.Println(items)
	example := conv.SchemaToExample(items)

	output = append(output, example)
	return output
}

func (con arrayConverter) ExampleToSchema(example interface{}) (schema map[string]interface{}) {
	arr := example.([]interface{})
	schema = make(map[string]interface{})
	schema["type"] = "array"
	if len(arr) == 0 {
		return
	}

	elem := arr[0]
	converter := NewConverterFromInterface(elem)
	schema["items"] = converter.ExampleToSchema(elem)

	return
}
