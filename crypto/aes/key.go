package aes

import "github.com/shibataka000/kms/crypto/rand"

// GenerateKey generate n byte key.
func GenerateKey(n int) ([]byte, error) {
	return rand.Bytes(n)
}
