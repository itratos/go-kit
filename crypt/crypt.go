package crypt

import "crypto/sha256"

func Encrypt(text string) string {
	sum := sha256.Sum256([]byte(text))
	return string(sum[:])
}
