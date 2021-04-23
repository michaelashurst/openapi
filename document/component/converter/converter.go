package converter

type converter interface {
	SchemaToExample(schema map[string]interface{}) interface{}
	ExampleToSchema(example map[string]interface{}) interface{}
}

func NewConverter(s string) converter {
	switch s {
	case "string":
		return stringConverter{}
	case "integer":
		return integerConverter{}
	case "boolean":
		return booleanConverter{}
	case "object":
		return objectConverter{}
	}
	panic("No converter found")
}
