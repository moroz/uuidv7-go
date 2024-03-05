package uuidv7

import "fmt"

// Scan implements the `sql.Scanner` interface.
func (u *UUID) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil

	case string:
		// return a null UUID when called with empty string
		if src == "" {
			return nil
		}

		parsed, err := Parse(src)
		if err != nil {
			return fmt.Errorf("Scan: %v", err)
		}

		*u = parsed

	case []byte:
		// return a null UUID when called with an empty slice
		if len(src) == 0 {
			return nil
		}

		// If the length of the slice is 16 bytes, treat it as
		// an existing UUID
		if len(src) == 16 {
			copy((*u)[:], src)
			return nil
		}

		// For any other length, try parsing it as a string
		return u.Scan(string(src))

	default:
		return fmt.Errorf("Scan: unable to scan type %T into UUID", src)
	}
	return nil
}
