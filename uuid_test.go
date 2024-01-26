package uuidv7_test

import (
	"encoding/json"
	"testing"

	"github.com/moroz/uuidv7-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUUIDIsNil(t *testing.T) {
	uuid := uuidv7.Generate()
	assert.NotEqual(t, uuid.IsNil(), true)
	assert.Equal(t, uuidv7.Nil.IsNil(), true)

	empty := uuidv7.UUID{}
	assert.Equal(t, empty.IsNil(), true)
}

func TestUUIDBase64(t *testing.T) {
	str := "018b1492-8002-7338-bae4-adb32c1b949a"
	parsed, _ := uuidv7.Parse(str)
	actual := parsed.Base64()
	expected := "AYsUkoACczi65K2zLBuUmg=="
	assert.Equal(t, actual, expected)
}

func TestUUIDMarshalToJSON(t *testing.T) {
	str := "018b1492-8002-7338-bae4-adb32c1b949a"
	parsed, _ := uuidv7.Parse(str)
	expected := `"` + str + `"`
	actual, err := json.Marshal(parsed)
	require.NoError(t, err)
	assert.Equal(t, actual, []byte(expected))
}
