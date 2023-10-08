package uuidv7_test

import (
	"testing"

	"github.com/moroz/uuidv7-go"
)

func TestGenerate(t *testing.T) {
	uuid := uuidv7.Generate()
	for i, byteVal := range uuid {
		if byteVal == 0 {
			t.Errorf("Unexpected zero byte at index %d\n", i)
		}
	}

	versionMarker := uuid[6]
	if versionMarker&0xF0 != 7<<4 {
		t.Errorf("Invalid UUID version marker: %#02x\n", versionMarker)
	}

	rfcMarker := uuid[8]
	if rfcMarker&0xC0 != 2<<6 {
		t.Errorf("Invalid RFC variant marker: %#02x\n", rfcMarker)
	}
}
