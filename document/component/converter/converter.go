package converter

import (
	"fmt"
)

type converter interface {
	SchemaToExample(schema map[string]interface{}) interface{}
	ExampleToSchema(example interface{}) map[string]interface{}
}

func NewConverter(s string) converter {
	fmt.Println("Creating component converter for type", s)
	switch s {
	case "string":
		return stringConverter{}
	case "integer":
		return integerConverter{}
	case "boolean":
		return booleanConverter{}
	case "number":
		return numberConverter{}
	case "array":
		return arrayConverter{}
	case "object":
		return objectConverter{}
	}
	panic("No converter found for type " + s)
}

func NewConverterFromInterface(i interface{}) converter {
	switch i.(type) {
	case string:
		return NewConverter("string")
	case int:
		return NewConverter("integer")
	case float64:
		return NewConverter("number")
	case bool:
		return NewConverter("boolean")
	case []interface{}:
		return NewConverter("array")
	default:
		return NewConverter("object")
	}
}
