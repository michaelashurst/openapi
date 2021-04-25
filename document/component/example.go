package component

import (
	"encoding/json"

	"github.com/michaelashurst/openapi/document/component/converter"
)

type Example struct {
	*json.RawMessage
}

func (e Example) GenerateSchema() (schema Schema) {
	output := make(map[string]interface{})

	jstr, err := e.MarshalJSON()
	checkErr(err)

	var jobj interface{}
	err = json.Unmarshal(jstr, &jobj)
	checkErr(err)

	converter := converter.NewConverterFromInterface(jobj)
	output = converter.ExampleToSchema(jobj)

	outputData, err := json.Marshal(output)
	checkErr(err)

	schema = Schema{&json.RawMessage{}}
	err = schema.UnmarshalJSON(outputData)
	checkErr(err)

	return
}
