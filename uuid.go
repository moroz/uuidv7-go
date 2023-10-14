package uuidv7

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
)

type UUID [16]byte

func (u UUID) Value() (driver.Value, error) {
	return driver.Value(u[:]), nil
}
func (u UUID) Dump() string {
	return hex.EncodeToString(u[:])
}

func (u UUID) String() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hex.EncodeToString(u[0:4]),
		hex.EncodeToString(u[4:6]),
		hex.EncodeToString(u[6:8]),
		hex.EncodeToString(u[8:10]),
		hex.EncodeToString(u[10:]))
}
