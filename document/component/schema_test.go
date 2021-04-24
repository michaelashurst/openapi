package component

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGenerateExample(t *testing.T) {
	schemaJson := `{
		"type": "object",
		"properties": {
			"id": {
				"type": "integer",
				"format": "int64"
			},
			"quantity": {
				"type": "integer",
				"format": "int32"
			},
			"shipDate": {
				"type": "string",
				"format": "date-time"
			},
			"status": {
				"type": "string",
				"description": "Order Status",
				"enum": [
					"placed",
					"approved",
					"delivered"
				]
			},
			"complete": {
				"type": "boolean",
				"default": false
			}
		},
		"xml": {
			"name": "Order"
		}
	}`

	schemaBytes := []byte(schemaJson)
	schema := Schema{&json.RawMessage{}}
	err := json.Unmarshal(schemaBytes, &schema)
	if err != nil {
		panic(err)
	}
	example := schema.GenerateExample()

	exampleBytes, err := example.MarshalJSON()
	if err != nil {
		panic(err)
	}
	var exampleMap map[string]interface{}
	json.Unmarshal(exampleBytes, &exampleMap)

	if exampleMap["id"].(float64) != 0 {
		t.Error("id expected to be 0 but got", exampleMap["id"])
	}

	fmt.Println("Test5")
	if exampleMap["shipDate"] != "0000-00-00T00:00.000Z" {
		t.Error("shipDate expected to be 0000-00-00T00:00.000Z but got", exampleMap["shipDate"])
	}

	if exampleMap["status"] != "placed|approved|delivered" {
		t.Error("status expected to be placed|approved|delivered but got", exampleMap["status"])
	}

	if exampleMap["complete"] != false {
		t.Error("complete expected to be false but got", exampleMap["complete"])
	}
}
