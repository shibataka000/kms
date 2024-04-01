package crypto

import (
	"context"

	"github.com/shibataka000/kms/crypto/aes"
	"github.com/shibataka000/kms/crypto/kms"
	"github.com/shibataka000/kms/encoding"
)

// Decrypt ciphertext by envelope encryption.
func Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error) {
	cipherObj, err := encoding.Deserialize[Ciphertext](ciphertext)
	if err != nil {
		return nil, err
	}

	dataKey, err := kms.Decrypt(ctx, cipherObj.EncryptedDataKey)
	if err != nil {
		return nil, err
	}

	plaintext, err := aes.Decrypt(dataKey, cipherObj.Blob)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
