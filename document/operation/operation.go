package operation

import (
	"encoding/json"

	"github.com/michaelashurst/openapi/document/parameter"
)

type PathOperation struct {
	Operation
	Path   string `json:"path,omitempty" yaml:"path,omitempty"`
	Method string `json:"method,omitempty" yaml:"method,omitempty"`
}

type Operation struct {
	Tags         []string              `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *json.RawMessage      `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationId  string                `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	RequestBody  *RequestBody          `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    map[string]Response   `json:"responses" yaml:"responses"`
	Callbacks    *json.RawMessage      `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   bool                  `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     *json.RawMessage      `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      *json.RawMessage      `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters   []parameter.Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

type Response struct {
	Description string                         `json:"description,omitempty" yaml:"description,omitempty"`
	Headers     map[string]*json.RawMessage    `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]parameter.MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*json.RawMessage    `json:"links,omitempty" yaml:"links,omitempty"`
}

type RequestBody struct {
	Description string                         `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]parameter.MediaType `json:"content" yaml:"content"`
	Required    bool                           `json:"required,omitempty" yaml:"required,omitempty"`
}
