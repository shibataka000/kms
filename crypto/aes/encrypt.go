package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"

	"github.com/shibataka000/kms/crypto/rand"
	"github.com/shibataka000/kms/encoding"
)

// Encrypt plaintext by AES-CBC.
func Encrypt(key []byte, plaintext []byte) ([]byte, error) {
	iv, err := generateIV()
	if err != nil {
		return nil, err
	}

	ciphertext, err := encrypt(key, plaintext, iv)
	if err != nil {
		return nil, err
	}

	return encoding.Serialize(Ciphertext{
		Blob: ciphertext,
		IV:   iv,
	})
}

// encrypt plaintext by AES-CBC.
func encrypt(key []byte, plaintext []byte, iv []byte) ([]byte, error) {
	padded := padding(plaintext)
	ciphertext := make([]byte, len(padded))

	b, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bm := cipher.NewCBCEncrypter(b, iv)
	bm.CryptBlocks(ciphertext, padded)

	return ciphertext, nil
}

// padding by PKCS#7.
func padding(b []byte) []byte {
	padLen := aes.BlockSize - (len(b) % aes.BlockSize)
	return append(b, bytes.Repeat([]byte{byte(padLen)}, padLen)...)
}

// generate IV.
func generateIV() ([]byte, error) {
	return rand.Bytes(aes.BlockSize)
}
