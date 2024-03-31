package crypto

// Ciphertext.
type Ciphertext struct {
	Blob             []byte
	EncryptedDataKey []byte
	Iter             uint64
}
