package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(s string) string {
	// Create a new SHA-256 hash object
	h := sha256.New()

	// Write the string to the hash object
	h.Write([]byte(s))

	// Get the resulting hash as a byte slice
	bs := h.Sum(nil)

	// Encode the byte slice to a hexadecimal string
	return hex.EncodeToString(bs)
}
