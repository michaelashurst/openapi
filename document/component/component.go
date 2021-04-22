package component

import (
	"encoding/json"

	"github.com/michaelashurst/openapi/document/operation"
	"github.com/michaelashurst/openapi/document/parameter"
)

type Components struct {
	Schemas         map[string]Schema                `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses       map[string]operation.Response    `json:"responses,omitempty" yaml:"responses,omitempty"`
	Parameters      map[string]parameter.Parameter   `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Examples        map[string]parameter.Example     `json:"examples,omitempty" yaml:"examples,omitempty"`
	RequestBodies   map[string]operation.RequestBody `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
	Headers         *json.RawMessage                 `json:"headers,omitempty" yaml:"headers,omitempty"`
	SecuritySchemes *json.RawMessage                 `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
	Links           *json.RawMessage                 `json:"links,omitempty" yaml:"links,omitempty"`
	Callbacks       *json.RawMessage                 `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
}
