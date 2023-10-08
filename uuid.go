package uuidv7

import (
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

type UUID [16]byte

func (u UUID) Value() (driver.Value, error) {
	return driver.Value(u[:]), nil
}

func (u *UUID) Scan(src interface{}) error {
	var source []byte
	switch src.(type) {
	case []byte:
		fmt.Println(string(src.([]byte)))
		source = src.([]byte)
		if len(source) == 32 || len(source) == 36 {
			uuid, err := Parse(string(source))
			if err != nil {
				return err
			}
			*u = uuid
			return nil
		}
		if len(source) == 16 {
			*u = UUID(source)
			return nil
		}
	default:
		return errors.New("Incompatible type for UUID")
	}
	return nil
}

func Parse(source string) (UUID, error) {
	var uuid UUID
	if len(source) != 36 && len(source) != 32 {
		return UUID{}, errors.New("Invalid length")
	}
	if len(source) == 36 {
		source = strings.ReplaceAll(source, "-", "")
	}
	bytes, err := hex.DecodeString(source)
	if err != nil {
		return uuid, errors.New("Invalid UUID hex")
	}
	return UUID(bytes), nil
}

func (u UUID) String() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hex.EncodeToString(u[0:4]),
		hex.EncodeToString(u[4:6]),
		hex.EncodeToString(u[6:8]),
		hex.EncodeToString(u[8:10]),
		hex.EncodeToString(u[10:]))
}
