package uuidv7_test

import (
	"testing"

	"github.com/moroz/uuidv7-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseValidStringWithHyphens(t *testing.T) {
	str := "018b1492-8002-7338-bae4-adb32c1b949a"
	parsed, err := uuidv7.Parse(str)
	require.NoError(t, err)
	assert.Equal(t, parsed, ExampleUUID)
}

func TestParseInvalidStringWithHyphens(t *testing.T) {
	// note the misplaced hyphens
	str := "018b149280-0273-38ba-e4-adb32c1b949a"

	_, err := uuidv7.Parse(str)
	require.Error(t, err)
}

func TestParseValidStringWithoutHyphens(t *testing.T) {
	str := "018b149280027338bae4adb32c1b949a"
	parsed, err := uuidv7.Parse(str)
	require.NoError(t, err)
	assert.Equal(t, parsed, ExampleUUID)
}

func TestParseInvalidHex(t *testing.T) {
	str := "018b149280027338baz4adb32c1b949a"
	parsed, err := uuidv7.Parse(str)
	require.Error(t, err)
	assert.Equal(t, parsed, uuidv7.UUID{})
}
