package document

import (
	"encoding/json"
	"testing"

	"github.com/michaelashurst/openapi/document/component"
)

// func TestNewDocument(t *testing.T) {
// 	document := NewDocument("../swagger.json")
// 	fmt.Println(*document)
// 	if document == nil {
// 		t.Errorf("Failed to create document")
// 	}
// }

func TestGetComponentSchema(t *testing.T) {
	schemas := make(map[string]*json.RawMessage)
	successSchemaJson := `{
		"type": "object",
		"properties": {
			"name": {
				"type": "string",
				"example": "TestSuccess"
			}
		}
	}`
	failSchemaJson := `{
		"type": "object",
		"properties": {
			"name": {
				"type": "string",
				"example": "TestFail"
			}
		}
	}`

	schemaBytes := []byte(successSchemaJson)
	successSchema := &json.RawMessage{}
	err := json.Unmarshal(schemaBytes, &successSchema)
	if err != nil {
		panic(err)
	}

	failSchemaBytes := []byte(failSchemaJson)
	failSchema := &json.RawMessage{}
	err = json.Unmarshal(failSchemaBytes, &failSchema)
	if err != nil {
		panic(err)
	}

	schemas["SUCCESS"] = successSchema
	schemas["FAIL"] = failSchema

	document := Document{
		Components: component.Components{
			Schemas: schemas,
		},
	}

	example, err := document.GetComponentExample("SUCCESS")
	if err != nil {
		t.Error("error when getting component example", err.Error())
	}
	exampleBytes, err := example.MarshalJSON()
	if err != nil {
		panic(err)
	}
	var exampleMap map[string]interface{}
	json.Unmarshal(exampleBytes, &exampleMap)
	if exampleMap["name"] != "TestSuccess" {
		t.Error("Expected component name of TestSuccess but got ", exampleMap["name"])
	}
}
