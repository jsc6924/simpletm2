package utils

import (
	"backend/config"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
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

// genToken generates a secure random token of the specified length.
func GenToken(k int) (string, error) {
	// Create a byte slice of the specified length
	b := make([]byte, k)
	// Fill the byte slice with random bytes
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// Encode the byte slice to a URL-safe base64 string
	return base64.URLEncoding.EncodeToString(b), nil
}

func MakeSharedURL(user, apitoken, game string) string {
	return fmt.Sprintf("simpletm://%s/%s/%s/%s/%s", config.GenUrlProtocal, config.GenUrlHost, user, apitoken, game)
}
