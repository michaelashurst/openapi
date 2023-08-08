package basicdocument

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/michaelashurst/openapi/document"
	"github.com/michaelashurst/openapi/document/info"
	"github.com/michaelashurst/openapi/document/operation"
	"github.com/michaelashurst/openapi/document/parameter"
	"github.com/michaelashurst/openapi/document/path"
	"github.com/michaelashurst/openapi/document/server"
	"github.com/michaelashurst/openapi/document/tag"
)

type basicDocument struct {
	Info         basicInfo
	Servers      []server.Server `json:"servers,omitempty" yaml:"servers,omitempty"`
	Operations   []basicOperation
	Security     *json.RawMessage `json:"security,omitempty" yaml:"security,omitempty"`
	Tags         []tag.Tag        `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs *json.RawMessage `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

type basicInfo struct {
	Openapi   string `json:"openapi,omitempty" yaml:"openapi,omitempty"`
	info.Info `yaml:",inline"`
}

type fileInfo struct {
	path   string
	format string
}

type basicDocumentFiles struct {
	infoPath       fileInfo
	tagPath        fileInfo
	serverPath     fileInfo
	operationPaths []fileInfo
}

func newBasicDocumentFiles(path string) basicDocumentFiles {
	files := basicDocumentFiles{}
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		fileName := strings.ToLower(strings.Split(info.Name(), ".")[0])
		fileFormat := strings.ToLower(strings.Split(info.Name(), ".")[1])

		if fileName == "info" {
			files.infoPath = fileInfo{path: p, format: fileFormat}
			return nil
		}
		if fileName == "tags" {
			files.tagPath = fileInfo{path: p, format: fileFormat}
			return nil
		}
		if fileName == "servers" {
			files.serverPath = fileInfo{path: p, format: fileFormat}
			return nil
		}
		subDirs := strings.Split(p, "/")
		if len(subDirs) < 2 {
			subDirs = strings.Split(p, "\\")
		}
		if subDirs[len(subDirs)-2] == "operations" {
			files.operationPaths = append(files.operationPaths, fileInfo{path: p, format: fileFormat})
			return nil
		}

		return errors.New("Unexpected file found " + info.Name() + " with path " + p)
	})
	if err != nil {
		panic(err)
	}

	return files
}

func NewBasicDocument(path string) basicDocument {
	doc := basicDocument{}

	files := newBasicDocumentFiles(path)
	NewDecoder(files.infoPath.format).decode(files.infoPath.path, &doc.Info)
	NewDecoder(files.serverPath.format).decode(files.serverPath.path, &doc.Servers)
	NewDecoder(files.tagPath.format).decode(files.tagPath.path, &doc.Tags)
	for _, path := range files.operationPaths {
		var operation basicOperation
		NewDecoder(path.format).decode(path.path, &operation)
		doc.Operations = append(doc.Operations, operation)
	}
	return doc
}

func (basic basicDocument) Document(outputPath string) (doc document.Document) {
	doc.FilePath = outputPath
	if doc.FilePath[len(doc.FilePath)-1] == '/' {
		doc.FilePath += "openapi.json"
	}

	doc.Openapi = basic.Info.Openapi
	doc.Info = basic.Info.Info
	doc.Servers = basic.Servers
	doc.Tags = basic.Tags

	doc.Paths = make(map[string]path.Path)
	for _, op := range basic.Operations {
		path := path.Path{}
		if _, ok := doc.Paths[op.Path]; ok {
			path = doc.Paths[op.Path]
		}
		operation, method := op.operation()
		setOperation(&path, operation, method)
		path.Parameters = parameter.GetParameters(op.Path)
		doc.Paths[removeQuery(op.Path)] = path
	}
	return
}

func removeQuery(path string) string {
	return strings.Split(path, "?")[0]
}

func setOperation(path *path.Path, operation operation.Operation, method string) {
	switch method {
	case "Get":
		path.Get = &operation
	case "Put":
		path.Put = &operation
	case "Post":
		path.Post = &operation
	case "Patch":
		path.Patch = &operation
	case "Delete":
		path.Delete = &operation
	case "Head":
		path.Head = &operation
	case "Trace":
		path.Trace = &operation
	}
}
