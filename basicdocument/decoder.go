package basicdocument

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type decoder interface {
	decode(path string, ptr interface{})
}

type yamlDecoder struct {
}

type jsonDecoder struct {
}

func NewDecoder(typ string) decoder {
	if typ == "yaml" {
		return yamlDecoder{}
	}
	if typ == "json" {
		return jsonDecoder{}
	}
	panic("Invalid decoder type")
}

func (d yamlDecoder) decode(path string, ptr interface{}) {
	bytes, err := ioutil.ReadFile(path) // For read access.
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(bytes, ptr)
}

func (d jsonDecoder) decode(path string, ptr interface{}) {
	bytes, err := ioutil.ReadFile(path) // For read access.
	if err != nil {
		panic(err)
	}
	json.Unmarshal(bytes, ptr)
}
