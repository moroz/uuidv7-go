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

func TestUUIDBase64(t *testing.T) {
	str := "018b1492-8002-7338-bae4-adb32c1b949a"
	parsed, _ := uuidv7.Parse(str)
	actual := parsed.Base64()
	expected := "AYsUkoACczi65K2zLBuUmg=="
	if actual != expected {
		t.Errorf("Expected UUID %s to be encoded as %s, got %s", str, expected, actual)
	}
}
