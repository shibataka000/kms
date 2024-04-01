package crypto

import (
	"context"

	"github.com/shibataka000/kms/crypto/aes"
	"github.com/shibataka000/kms/crypto/kms"
	"github.com/shibataka000/kms/encoding"
)

// Encrypt plaintext by envelope encryption.
func Encrypt(ctx context.Context, kmsKeyID string, plaintext []byte) ([]byte, error) {
	dataKey, err := aes.GenerateKey(32)
	if err != nil {
		return nil, err
	}

	ciphertext, err := aes.Encrypt(dataKey, plaintext)
	if err != nil {
		return nil, err
	}

	encryptedDataKey, err := kms.Encrypt(ctx, kmsKeyID, dataKey)
	if err != nil {
		return nil, err
	}

	return encoding.Serialize(Ciphertext{
		Blob:             ciphertext,
		EncryptedDataKey: encryptedDataKey,
	})
}
