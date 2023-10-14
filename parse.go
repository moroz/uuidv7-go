package uuidv7

import (
	"encoding/hex"
	"errors"
	"strings"
)

var ErrorInvalidFormat = errors.New("invalid UUID format")

func Parse(source string) (UUID, error) {
	var uuid UUID

	switch len(source) {
	// 018b1492-8002-7338-bae4-adb32c1b949a
	case 36:
		if source[8] != '-' || source[13] != '-' || source[18] != '-' || source[23] != '-' {
			return uuid, ErrorInvalidFormat
		}

		compact := strings.ReplaceAll(source, "-", "")
		return Parse(compact)

	// 018b149280027338bae4adb32c1b949a
	case 32:

	default:
		return uuid, ErrorInvalidFormat
	}

	bytes, err := hex.DecodeString(source)
	if err != nil {
		return uuid, ErrorInvalidFormat
	}

	return UUID(bytes), nil
}
