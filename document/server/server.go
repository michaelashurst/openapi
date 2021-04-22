package server

type serverVariable struct {
	Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	Default     string   `json:"default,omitempty" yaml:"default,omitempty"`
	Description string   `json:"description,omitempty" yaml:"description,omitempty"`
}

type Server struct {
	Url         string                    `json:"url,omitempty" yaml:"url,omitempty"`
	Description string                    `json:"description,omitempty" yaml:"description,omitempty"`
	Variables   map[string]serverVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}
