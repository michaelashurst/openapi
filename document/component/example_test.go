package component

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGenerateSchema(t *testing.T) {
	exampleJson := `{
		"id": 0,
		"shipDate": "0000-00-00T00:00:00Z",
		"status": "placed|approved|delivered",
		"complete": false,
		"contacts": [
			"test"
		]
	}`

	exampleBytes := []byte(exampleJson)
	example := Example{&json.RawMessage{}}
	err := json.Unmarshal(exampleBytes, &example)
	if err != nil {
		panic(err)
	}
	schema := example.GenerateSchema()

	schemaBytes, err := schema.MarshalJSON()
	if err != nil {
		panic(err)
	}
	var schemaMap map[string]interface{}
	json.Unmarshal(schemaBytes, &schemaMap)

	if schemaMap["id"].(map[string]interface{})["type"] != "integer" {
		t.Error("id expected to be integer but got", schemaMap["id"])
	}

	if schemaMap["shipDate"].(map[string]interface{})["type"] != "string" || schemaMap["shipDate"].(map[string]interface{})["format"] != "date-time" {
		t.Error("shipDate expected to be type string in date-time format but got", schemaMap["shipDate"])
	}

	var interfaceSlice []interface{}
	if schemaMap["status"].(map[string]interface{})["type"] != "string" ||
		reflect.TypeOf(schemaMap["status"].(map[string]interface{})["enum"]) != reflect.TypeOf(interfaceSlice) ||
		len(schemaMap["status"].(map[string]interface{})["enum"].([]interface{})) != 3 {
		t.Error("status expected to be an array of strings with 3 values but got", schemaMap["status"])
	}

	if schemaMap["complete"].(map[string]interface{})["type"] != "boolean" {
		t.Error("complete expected to be boolean type but got", schemaMap["complete"])
	}

	if schemaMap["contacts"].(map[string]interface{})["type"] != "array" || schemaMap["contacts"].(map[string]interface{})["items"].(map[string]interface{})["type"] != "string" {
		t.Error("contacts expected to be array type with string items but got", schemaMap["contacts"])
	}
}
