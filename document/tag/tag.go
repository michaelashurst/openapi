package tag

import "encoding/json"

type Tag struct {
	Name         string           `json:"name,omitempty" yaml:"name,omitempty"`
	ExternalDocs *json.RawMessage `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	Description  string           `json:"description,omitempty" yaml:"description,omitempty"`
}
