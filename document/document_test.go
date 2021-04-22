package document

import (
	"fmt"
	"testing"
)

func TestNewDocument(t *testing.T) {
	document := NewDocument("../swagger.json")
	fmt.Println(*document)
	if document == nil {
		t.Errorf("Failed to create document")
	}
}
