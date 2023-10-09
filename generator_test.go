package uuidv7_test

import (
	"testing"

	"github.com/moroz/uuidv7-go"
)

func assertValidUUIDv7(t *testing.T, uuid uuidv7.UUID) {
	versionMarker := uuid[6]
	if versionMarker&0xF0 != 7<<4 {
		t.Errorf("Invalid UUID version marker: %#02x\n", versionMarker)
	}

	rfcMarker := uuid[8]
	if rfcMarker&0xC0 != 2<<6 {
		t.Errorf("Invalid RFC variant marker: %#02x\n", rfcMarker)
	}
}

func TestGenerate(t *testing.T) {
	uuid := uuidv7.Generate()
	for i, byteVal := range uuid {
		if byteVal == 0 {
			t.Errorf("Unexpected zero byte at index %d\n", i)
		}
	}

	assertValidUUIDv7(t, uuid)
}

func TestParseValidStringWithHyphens(t *testing.T) {
	str := "018b1492-8002-7338-bae4-adb32c1b949a"
	parsed, err := uuidv7.Parse(str)
	if err != nil {
		t.Errorf("Expected Parse to succeed, encountered error: %v\n", err)
	}
	assertValidUUIDv7(t, parsed)
}

func TestParseValidStringWithoutHyphens(t *testing.T) {
	str := "018b149280027338bae4adb32c1b949a"
	parsed, err := uuidv7.Parse(str)
	if err != nil {
		t.Errorf("Expected Parse to succeed, encountered error: %v\n", err)
	}
	assertValidUUIDv7(t, parsed)
}

func TestParseInvalidHex(t *testing.T) {
	str := "018b149280027338baz4adb32c1b949a"
	parsed, err := uuidv7.Parse(str)
	if err == nil {
		t.Error("Expected Parse to fail", err)
	}

	var nullUUID uuidv7.UUID
	if parsed != nullUUID {
		t.Errorf("Expected Parse to return null UUID, got: %v\n", parsed)
	}
}
