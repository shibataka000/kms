package aes

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/shibataka000/kms/encoding"
)

// Decrypt ciphertext by AES-CBC.
func Decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	cipherObj, err := encoding.Deserialize[Ciphertext](ciphertext)
	if err != nil {
		return nil, err
	}

	plaintext, err := decrypt(key, cipherObj.Blob, cipherObj.IV)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// decrypt ciphertext by AES-CBC.
func decrypt(key []byte, ciphertext []byte, iv []byte) ([]byte, error) {
	plaintext := make([]byte, len(ciphertext))

	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bm := cipher.NewCBCDecrypter(b, iv)
	bm.CryptBlocks(plaintext, ciphertext)

	return unpadding(plaintext), nil
}

// unpadding by PKCS#7.
func unpadding(b []byte) []byte {
	l := len(b)
	padLen := int(b[l-1])
	return b[:l-padLen]
}
