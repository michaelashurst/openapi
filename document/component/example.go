package component

import "encoding/json"

type Example struct {
	*json.RawMessage
}

func (e Example) GenerateSchema() (schema Schema) {
	return
}
