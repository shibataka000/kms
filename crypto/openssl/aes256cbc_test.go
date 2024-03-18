package openssl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAES256CBCEncryptAndDecrypt(t *testing.T) {
	tests := []struct {
		name      string
		plaintext string
	}{
		{
			name:      "EncryptAndDecrypt",
			plaintext: "Hello World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			pass, err := Rand(32, RandOption{
				Base64: true,
			})
			require.NoError(err)

			var iter uint64 = 10000

			ciphertext, err := AES256CBC(AES256CBCOption{
				Encrypt: true,
				In:      []byte(tt.plaintext),
				Pass:    pass,
				PBKDF2:  true,
				Iter:    iter,
			})
			require.NoError(err)

			plaintext, err := AES256CBC(AES256CBCOption{
				Decrypt: true,
				In:      []byte(ciphertext),
				Pass:    pass,
				PBKDF2:  true,
				Iter:    iter,
			})
			require.NoError(err)

			require.Equal(tt.plaintext, string(plaintext))
		})
	}
}
