package path

import (
	"github.com/michaelashurst/openapi/document/operation"
	"github.com/michaelashurst/openapi/document/parameter"
	"github.com/michaelashurst/openapi/document/server"
)

type Path struct {
	Ref         string                `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Summary     string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	Get         *operation.Operation  `json:"get,omitempty" yaml:"get,omitempty"`
	Post        *operation.Operation  `json:"post,omitempty" yaml:"post,omitempty"`
	Put         *operation.Operation  `json:"put,omitempty" yaml:"put,omitempty"`
	Delete      *operation.Operation  `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options     *operation.Operation  `json:"options,omitempty" yaml:"options,omitempty"`
	Head        *operation.Operation  `json:"head,omitempty" yaml:"head,omitempty"`
	Patch       *operation.Operation  `json:"patch,omitempty" yaml:"patch,omitempty"`
	Trace       *operation.Operation  `json:"trace,omitempty" yaml:"trace,omitempty"`
	Servers     []server.Server       `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  []parameter.Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
