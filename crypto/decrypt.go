package crypto

import (
	"context"

	"github.com/shibataka000/kms/crypto/kms"
	"github.com/shibataka000/kms/crypto/openssl"
)

// Decrypt ciphertext by envelope encryption.
func Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error) {
	ciphertextObj, err := deserialize(ciphertext)
	if err != nil {
		return nil, err
	}

	dataKey, err := kms.Decrypt(ctx, ciphertextObj.EncryptedDataKey)
	if err != nil {
		return nil, err
	}

	plaintext, err := openssl.AES256CBC(openssl.AES256CBCOption{
		Decrypt: true,
		In:      ciphertextObj.Blob,
		Pass:    dataKey,
		PBKDF2:  true,
		Iter:    ciphertextObj.Iter,
	})
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
