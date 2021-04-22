package parameter

import "encoding/json"

type encoding struct {
	ContentType   string                      `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers       map[string]*json.RawMessage `json:"headers,omitempty" yaml:"headers,omitempty"`
	Style         string                      `json:"style,omitempty" yaml:"style,omitempty"`
	Explode       bool                        `json:"explode,omitempty" yaml:"explode,omitempty"`
	AllowReserved bool                        `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}

type MediaType struct {
	Schema   schema             `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example  *json.RawMessage   `json:"example,omitempty" yaml:"example,omitempty"`
	Examples map[string]Example `json:"examples,omitempty" yaml:"examples,omitempty"`
}
