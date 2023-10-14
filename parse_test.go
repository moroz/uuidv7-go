package uuidv7_test

import (
	"testing"

	"github.com/moroz/uuidv7-go"
)

func TestParseValidStringWithHyphens(t *testing.T) {
	str := "018b1492-8002-7338-bae4-adb32c1b949a"
	parsed, err := uuidv7.Parse(str)
	if err != nil {
		t.Errorf("Expected Parse to succeed, encountered error: %v\n", err)
	}
	assertValidUUIDv7(t, parsed)
}

func TestParseInvalidStringWithHyphens(t *testing.T) {
	// note the misplaced hyphens
	str := "018b149280-0273-38ba-e4-adb32c1b949a"

	_, err := uuidv7.Parse(str)
	if err == nil {
		t.Errorf("Expected Parse to reject string %s", str)
	}
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
