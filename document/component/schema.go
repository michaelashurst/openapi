package component

import (
	"encoding/json"

	"github.com/michaelashurst/openapi/document/component/converter"
)

type Schema struct {
	*json.RawMessage
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (s Schema) GenerateExample() (example Example) {
	output := make(map[string]interface{})

	jstr, err := s.MarshalJSON()
	checkErr(err)

	var jobj map[string]interface{}
	err = json.Unmarshal(jstr, &jobj)
	checkErr(err)

	converter := converter.NewConverter(jobj["type"].(string))
	output = converter.SchemaToExample(jobj).(map[string]interface{})

	outputData, err := json.Marshal(output)
	checkErr(err)

	example = Example{&json.RawMessage{}}
	err = example.UnmarshalJSON(outputData)
	checkErr(err)

	return
}
