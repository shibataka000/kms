package crypto

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptAndDecrypt(t *testing.T) {
	tests := []struct {
		name      string
		keyID     string
		plaintext string
		iter      uint64
	}{
		{
			name:      "EncryptAndDecrypt",
			keyID:     "alias/shibataka000/kms",
			plaintext: "Hello World!",
			iter:      10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			ciphertext, err := Encrypt(ctx, tt.keyID, []byte(tt.plaintext), tt.iter)
			require.NoError(err)
			plaintext, err := Decrypt(ctx, ciphertext)
			require.NoError(err)
			require.Equal(tt.plaintext, string(plaintext))
		})
	}
}
