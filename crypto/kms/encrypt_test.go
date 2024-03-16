package kms

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name      string
		keyID     string
		plaintext string
	}{
		{
			name:      "EncryptAndDecrypt",
			keyID:     "alias/shibataka000/kms",
			plaintext: "Hello World!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			ciphertext, err := Encrypt(ctx, tt.keyID, []byte(tt.plaintext))
			require.NoError(err)
			plaintext, err := Decrypt(ctx, ciphertext)
			require.NoError(err)
			require.Equal(tt.plaintext, string(plaintext))
		})
	}
}
