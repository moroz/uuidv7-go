package uuidv7_test

import (
	"encoding/json"
	"testing"

	"github.com/moroz/uuidv7-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ExampleUUID = uuidv7.UUID{0x01, 0x8b, 0x14, 0x92, 0x80, 0x02, 0x73, 0x38, 0xba, 0xe4, 0xad, 0xb3, 0x2c, 0x1b, 0x94, 0x9a}

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

func TestUUIDUnmarshal(t *testing.T) {
	str := []byte(`"018b1492-8002-7338-bae4-adb32c1b949a"`)
	var result uuidv7.UUID
	err := json.Unmarshal(str, &result)
	require.NoError(t, err)
	assert.Equal(t, result, ExampleUUID)
}
