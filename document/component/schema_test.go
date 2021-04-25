package component

import (
	"encoding/json"
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
			},
			"contacts": {
				"type": "array",
				"items": {
					"type": "string",
					"example": "test"
				}
			}
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

	if exampleMap["shipDate"] != "0000-00-00T00:00:00Z" {
		t.Error("shipDate expected to be 0000-00-00T00:00:00Z but got", exampleMap["shipDate"])
	}

	if exampleMap["status"] != "placed|approved|delivered" {
		t.Error("status expected to be placed|approved|delivered but got", exampleMap["status"])
	}

	if exampleMap["complete"] != false {
		t.Error("complete expected to be false but got", exampleMap["complete"])
	}

	if len(exampleMap["contacts"].([]interface{})) != 1 || exampleMap["contacts"].([]interface{})[0] != "test" {
		t.Error("contacts expected to be an array with 1 element but got", exampleMap["contacts"])
	}
}
