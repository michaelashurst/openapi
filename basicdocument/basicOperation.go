package basicdocument

import (
	"encoding/json"
	"strings"

	"github.com/michaelashurst/openapi/document/component"
	"github.com/michaelashurst/openapi/document/operation"
	"github.com/michaelashurst/openapi/document/parameter"
)

type basicOperation struct {
	Path        string                 `json:"path,omitempty" yaml:"path,omitempty"`
	Description string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Summary     string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Method      string                 `json:"method,omitempty" yaml:"method,omitempty"`
	Tags        []string               `json:"tags,omitempty" yaml:"tags,omitempty"`
	Request     *json.RawMessage       `json:"request,omitempty" yaml:"request,omitempty"`
	Responses   map[string]interface{} `json:"responses,omitempty" yaml:"responses,omitempty"`
}

func (basic basicOperation) operation() (op operation.Operation, method string) {
	method = strings.Title(strings.ToLower(basic.Method))
	op.Description = basic.Description
	op.Summary = basic.Summary
	op.Tags = basic.Tags

	if basic.Request != nil {
		bytes, err := json.Marshal(basic.Request)
		if err != nil {
			panic(err)
		}

		ex := component.Example{RawMessage: &json.RawMessage{}}
		ex.UnmarshalJSON(bytes)
		reqExample := make(map[string]interface{})
		json.Unmarshal(bytes, &reqExample)
		req := make(map[string]parameter.MediaType)
		req["application/json"] = parameter.MediaType{
			Example: reqExample,
			Schema:  ex.GenerateSchema().RawMessage,
		}

		op.RequestBody = &operation.RequestBody{
			Required:    true,
			Content:     req,
			Description: "",
		}
	}

	if basic.Responses != nil {
		op.Responses = make(map[string]operation.Response)
		for key, response := range basic.Responses {
			bytes, err := json.Marshal(response)
			if err != nil {
				panic(err)
			}

			ex := component.Example{RawMessage: &json.RawMessage{}}
			ex.UnmarshalJSON(bytes)
			resExample := make(map[string]interface{})
			json.Unmarshal(bytes, &resExample)
			res := make(map[string]parameter.MediaType)
			res["application/json"] = parameter.MediaType{
				Example: resExample,
				Schema:  ex.GenerateSchema().RawMessage,
			}

			op.Responses[key] = operation.Response{
				Content:     res,
				Description: "",
			}
		}
	}

	return
}
