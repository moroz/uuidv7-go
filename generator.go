package uuidv7

import (
	"crypto/rand"
	"encoding/binary"
	"time"
)

// Generate generates a UUIDv7 with the current timestamp.
func Generate() UUID {
	var uuid UUID

	// First 48 bits: big-endian epoch time in milliseconds
	ts := time.Now().UnixMilli()
	binary.BigEndian.PutUint64(uuid[0:8], uint64(ts<<16))

	// Set the remaining 80 bits to random bytes
	rand.Read(uuid[6:])

	// Zero the first 4 bits of the 7th byte
	uuid[6] &= 0xF

	// Set UUID version marker in 7th byte
	uuid[6] |= 7 << 4

	// Zero the first 2 bits of the 9th byte
	uuid[8] &= 0x3F

	// Set the first 2 bits of the 9th byte to RFC variant 2
	uuid[8] |= 2 << 6

	return uuid
}
