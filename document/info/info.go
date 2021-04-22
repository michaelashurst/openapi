package info

type Info struct {
	Title          string  `json:"title,omitempty" yaml:"title,omitempty"`
	Description    string  `json:"description,omitempty" yaml:"description,omitempty"`
	TermsOfService string  `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
	Version        string  `json:"version,omitempty" yaml:"version,omitempty"`
	Licenae        license `json:"license,omitempty" yaml:"license,omitempty"`
	Contact        contact `json:"contact,omitempty" yaml:"contact,omitempty"`
}

type contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	Url   string `json:"url,omitempty" yaml:"url,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}

type license struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Url  string `json:"url,omitempty" yaml:"url,omitempty"`
}
