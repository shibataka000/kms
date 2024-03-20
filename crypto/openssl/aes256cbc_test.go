package openssl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAES256CBCEncryptAndDecrypt(t *testing.T) {
	tests := []struct {
		name          string
		plaintext     string
		genKeyOption  RandOption
		encryptOption AES256CBCOption
		decryptOption AES256CBCOption
	}{
		{
			name:      "EncryptAndDecrypt",
			plaintext: "Hello World!",
			genKeyOption: RandOption{
				Base64: true,
			},
			encryptOption: AES256CBCOption{
				Encrypt: true,
				Decrypt: false,
				In:      []byte{}, // Overwritten in test.
				Pass:    []byte{}, // Overwritten in test.
				Salt:    true,
				PBKDF2:  true,
				Iter:    10000,
			},
			decryptOption: AES256CBCOption{
				Encrypt: false,
				Decrypt: true,
				In:      []byte{}, // Overwritten in test.
				Pass:    []byte{}, // Overwritten in test.
				Salt:    true,
				PBKDF2:  true,
				Iter:    10000,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			pass, err := Rand(32, tt.genKeyOption)
			require.NoError(err)

			tt.encryptOption.In = []byte(tt.plaintext)
			tt.encryptOption.Pass = pass
			ciphertext, err := AES256CBC(tt.encryptOption)
			require.NoError(err)

			tt.decryptOption.In = ciphertext
			tt.decryptOption.Pass = pass
			plaintext, err := AES256CBC(tt.decryptOption)
			require.NoError(err)

			require.Equal(tt.plaintext, string(plaintext))
		})
	}
}
