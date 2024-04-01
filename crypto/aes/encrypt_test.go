package aes

import (
	"bytes"
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptAndDecrypt(t *testing.T) {
	tests := []struct {
		name      string
		keyLength int
		plaintext []byte
	}{
		{
			name:      "EncryptAndDecrypt",
			keyLength: 32,
			plaintext: []byte("this is test message"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			key, err := GenerateKey(tt.keyLength)
			require.NoError(err)
			ciphertext, err := Encrypt(key, tt.plaintext)
			require.NoError(err)
			plaintext, err := Decrypt(key, ciphertext)
			require.NoError(err)
			require.Equal(tt.plaintext, plaintext)
		})
	}
}

func TestEncryptAndDecrypt2(t *testing.T) {
	tests := []struct {
		name       string
		key        []byte
		plaintext  []byte
		iv         []byte
		ciphertext []byte
	}{
		{
			// echo -n 'this is test message' | openssl aes-256-cbc -e -base64 -nosalt -K 645E739A7F9F162725C1533DC2C5E827645E739A7F9F162725C1533DC2C5E827 -iv 71fbf00383b6e214dc08b8b94183cf30 | base64 -d | hexdump -C
			name:       "this is test message",
			key:        []byte{0x64, 0x5E, 0x73, 0x9A, 0x7F, 0x9F, 0x16, 0x27, 0x25, 0xC1, 0x53, 0x3D, 0xC2, 0xC5, 0xE8, 0x27, 0x64, 0x5E, 0x73, 0x9A, 0x7F, 0x9F, 0x16, 0x27, 0x25, 0xC1, 0x53, 0x3D, 0xC2, 0xC5, 0xE8, 0x27},
			plaintext:  []byte("this is test message"),
			iv:         []byte{0x71, 0xfb, 0xf0, 0x03, 0x83, 0xb6, 0xe2, 0x14, 0xdc, 0x08, 0xb8, 0xb9, 0x41, 0x83, 0xcf, 0x30},
			ciphertext: []byte{0x69, 0xbe, 0x9c, 0xfc, 0x06, 0x2e, 0x4b, 0x9e, 0x86, 0x46, 0x1c, 0xf7, 0xfc, 0xe5, 0x98, 0x4c, 0x79, 0xd6, 0xe6, 0xa7, 0x2c, 0x5d, 0x86, 0x7a, 0x95, 0x8b, 0x67, 0x8f, 0xfe, 0x38, 0x9c, 0x05},
		},
		{
			// echo -n 'this is test message 2' | openssl aes-256-cbc -e -base64 -nosalt -K 645E739A7F9F162725C1533DC2C5E827645E739A7F9F162725C1533DC2C5E827 -iv 71fbf00383b6e214dc08b8b94183cf30 | base64 -d | hexdump -C
			name:       "this is test message 2",
			key:        []byte{0x64, 0x5E, 0x73, 0x9A, 0x7F, 0x9F, 0x16, 0x27, 0x25, 0xC1, 0x53, 0x3D, 0xC2, 0xC5, 0xE8, 0x27, 0x64, 0x5E, 0x73, 0x9A, 0x7F, 0x9F, 0x16, 0x27, 0x25, 0xC1, 0x53, 0x3D, 0xC2, 0xC5, 0xE8, 0x27},
			plaintext:  []byte("this is test message 2"),
			iv:         []byte{0x71, 0xfb, 0xf0, 0x03, 0x83, 0xb6, 0xe2, 0x14, 0xdc, 0x08, 0xb8, 0xb9, 0x41, 0x83, 0xcf, 0x30},
			ciphertext: []byte{0x69, 0xbe, 0x9c, 0xfc, 0x06, 0x2e, 0x4b, 0x9e, 0x86, 0x46, 0x1c, 0xf7, 0xfc, 0xe5, 0x98, 0x4c, 0xe3, 0xe7, 0x37, 0x9a, 0xa4, 0x28, 0xc0, 0xfd, 0xb6, 0x4a, 0x4a, 0x2b, 0x85, 0x15, 0x74, 0xbe},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			// Test encrypting
			ciphertext, err := encrypt(tt.key, padding(tt.plaintext), tt.iv)
			require.NoError(err)
			require.Equal(tt.ciphertext, ciphertext)

			// Test decrypting
			plaintext, err := decrypt(tt.key, tt.ciphertext, tt.iv)
			require.NoError(err)
			require.Equal(tt.plaintext, unpadding(plaintext))
		})
	}
}

func TestPaddingAndUnpadding(t *testing.T) {
	tests := []struct {
		name   string
		b      []byte
		padded []byte
	}{
		{
			name:   "0Byte",
			b:      []byte{},
			padded: []byte{0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10},
		},
		{
			name:   "1Byte",
			b:      []byte{0x00},
			padded: []byte{0x00, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f},
		},
		{
			name:   "15Byte",
			b:      []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			padded: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01},
		},
		{
			name:   "16Byte",
			b:      []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			padded: append(bytes.Repeat([]byte{0x00}, 16), bytes.Repeat([]byte{0x10}, 16)...),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			require.Equal(tt.padded, padding(tt.b))
			require.Equal(tt.b, unpadding(tt.padded))
		})
	}
}

func TestGenerateIV(t *testing.T) {
	require := require.New(t)
	iv, err := generateIV()
	require.NoError(err)
	require.Equal(aes.BlockSize, len(iv))
}