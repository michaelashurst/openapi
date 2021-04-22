package parameter

import "encoding/json"

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
	Schema        schema             `json:"schema,omitempty" yaml:"schema,omitempty"`
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

type schema struct {
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
