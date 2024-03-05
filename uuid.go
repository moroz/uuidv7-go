package uuidv7

import (
	"database/sql/driver"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type UUID [16]byte

var Nil UUID

// Value implements the `driver.Valuer` interface of `database/sql/driver`.
// Returns the UUID serialized to string format, e. g. `"018b1492-8002-7338-bae4-adb32c1b949a"`.
// This is the format expected by PostgreSQL. To get a byte slice, use the slicing operator.
// The error return value is always nil.
func (u UUID) Value() (driver.Value, error) {
	return u.String(), nil
}

// Dump returns the UUID as a simple hex string, without hyphens.
func (u UUID) Dump() string {
	return hex.EncodeToString(u[:])
}

// String returns the UUID in string format, e. g. `"018b1492-8002-7338-bae4-adb32c1b949a"`.
func (u UUID) String() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hex.EncodeToString(u[0:4]),
		hex.EncodeToString(u[4:6]),
		hex.EncodeToString(u[6:8]),
		hex.EncodeToString(u[8:10]),
		hex.EncodeToString(u[10:]))
}

// Base64 returns the UUID encoded as a Base64 string.
func (u UUID) Base64() string {
	return base64.StdEncoding.EncodeToString(u[:])
}

// IsNil checks whether an UUID is a nil UUID, i. e. if all bits are set to zero.
func (u UUID) IsNil() bool {
	return u == Nil
}

// MarshalJSON encodes the UUID as a JSON string. Implements `json.Marshaler` interface.
func (u UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

// UnmarshalJSON decodes a UUID from a JSON string.
func (u *UUID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	uuid, err := Parse(s)
	if err != nil {
		return err
	}

	*u = uuid
	return nil
}
