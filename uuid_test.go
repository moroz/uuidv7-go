package uuidv7_test

import (
	"testing"

	"github.com/moroz/uuidv7-go"
)

func TestUUIDIsNil(t *testing.T) {
	uuid := uuidv7.Generate()
	if uuid.IsNil() {
		t.Errorf("Generated UUIDv7 should not be nil")
	}

	if !uuidv7.Nil.IsNil() {
		t.Errorf("Nil UUID should be nil")
	}

	empty := uuidv7.UUID{}
	if !empty.IsNil() {
		t.Errorf("Empty UUID should be nil")
	}
}
