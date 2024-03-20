package crypto

import (
	"context"

	"github.com/shibataka000/kms/crypto/kms"
	"github.com/shibataka000/kms/crypto/openssl"
)

// Encrypt plaintext by envelope encryption.
func Encrypt(ctx context.Context, kmsKeyID string, plaintext []byte, iter uint64) ([]byte, error) {
	dataKey, err := openssl.Rand(32, openssl.RandOption{
		Base64: true,
	})
	if err != nil {
		return nil, err
	}

	ciphertext, err := openssl.AES256CBC(openssl.AES256CBCOption{
		Encrypt: true,
		In:      plaintext,
		Pass:    dataKey,
		Salt:    true,
		PBKDF2:  true,
		Iter:    iter,
	})
	if err != nil {
		return nil, err
	}

	encryptedDataKey, err := kms.Encrypt(ctx, kmsKeyID, dataKey)
	if err != nil {
		return nil, err
	}

	ciphertextObj := Ciphertext{
		Blob:             ciphertext,
		EncryptedDataKey: encryptedDataKey,
		Iter:             iter,
	}

	return serialize(ciphertextObj)
}
