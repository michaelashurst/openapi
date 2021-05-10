package parameter

import (
	"encoding/json"
	"strings"
)

type Parameter struct {
	Name            string `json:"name,omitempty" yaml:"name,omitempty"`
	In              string `json:"in,omitempty" yaml:"in,omitempty"`
	Description     string `json:"description,omitempty" yaml:"description,omitempty"`
	Required        bool   `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated      bool   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	AllowEmptyValue bool   `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`

	Style         string             `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool               `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool               `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
	Schema        *json.RawMessage   `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example       *json.RawMessage   `json:"example,omitempty" yaml:"example,omitempty"`
	Examples      map[string]Example `json:"examples,omitempty" yaml:"examples,omitempty"`

	Content map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`
}

type Example struct {
	Summary       string           `json:"summary" yaml:"summary,omitempty"`
	Description   string           `json:"description" yaml:"description,omitempty"`
	Value         *json.RawMessage `json:"value,omitempty" yaml:"value,omitempty"`
	ExternalValue string           `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
}

func GetParameters(p string) []Parameter {
	params := []Parameter{}
	urlParts := strings.Split(p, "?")

	path := urlParts[0]
	pathParts := strings.Split(path, "/")

	for _, part := range pathParts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			schema := make(map[string]interface{})
			schema["type"] = "string"

			bytes, _ := json.Marshal(schema)
			rawSchema := json.RawMessage{}
			rawSchema.UnmarshalJSON(bytes)

			params = append(params, Parameter{
				Name:            part[1 : len(part)-1],
				In:              "path",
				Required:        true,
				AllowEmptyValue: false,
				Schema:          &rawSchema,
			})
		}
	}

	if len(urlParts) > 1 {
		query := urlParts[1]
		queryParts := strings.Split(query, "&")
		for _, part := range queryParts {
			schema := make(map[string]interface{})
			schema["type"] = "string"

			bytes, _ := json.Marshal(schema)
			rawSchema := json.RawMessage{}
			rawSchema.UnmarshalJSON(bytes)

			partSplit := strings.Split(part, "=")
			params = append(params, Parameter{
				Name:            partSplit[0],
				In:              "query",
				Required:        false,
				AllowEmptyValue: true,
				Schema:          &rawSchema,
			})
		}
	}

	return params
}
