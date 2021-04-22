package document

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/michaelashurst/openapi/document/component"
	"github.com/michaelashurst/openapi/document/info"
	"github.com/michaelashurst/openapi/document/operation"
	"github.com/michaelashurst/openapi/document/path"
	"github.com/michaelashurst/openapi/document/server"
	"github.com/michaelashurst/openapi/document/tag"
)

type Document struct {
	filePath     string
	Openapi      string               `json:"openapi,omitempty" yaml:"openapi,omitempty"`
	Info         info.Info            `json:"info,omitempty" yaml:"info,omitempty"`
	Servers      []server.Server      `json:"servers,omitempty" yaml:"servers,omitempty"`
	Paths        map[string]path.Path `json:"paths,omitempty" yaml:"paths,omitempty"`
	Components   component.Components `json:"components,omitempty" yaml:"components,omitempty"`
	Security     *json.RawMessage     `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         []tag.Tag            `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs *json.RawMessage     `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

func NewDocument(filepath string) *Document {
	file, _ := ioutil.ReadFile(filepath)
	doc := Document{filePath: filepath}
	_ = json.Unmarshal([]byte(file), &doc)

	return &doc
}

func contains(arr []string, val string) bool {
	for _, s := range arr {
		if s == val {
			return true
		}
	}
	return false
}

func (doc *Document) GetOperationsByTag(tag string) (ops []operation.PathOperation) {
	for key, path := range doc.Paths {
		for _, op := range []*operation.Operation{path.Get, path.Delete, path.Post, path.Put, path.Patch, path.Head, path.Options, path.Trace} {
			if op != nil && contains(op.Tags, tag) {
				ops = append(ops, operation.PathOperation{Operation: *op, Path: key})
			}
		}
	}
	return
}

func (doc *Document) Save() {
	fmt.Println("Saving file to ", doc.filePath)
	bytes, err := json.MarshalIndent(doc, "", "    ")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(doc.filePath, bytes, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func (doc *Document) UpdateInfo(info info.Info) {
	doc.Info = info
	doc.Save()
}