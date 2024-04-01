package rand

import "crypto/rand"

// Bytes return random bytes.
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}
