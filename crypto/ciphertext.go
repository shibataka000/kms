package crypto

import (
	"bytes"
	"encoding/gob"
)

// Ciphertext.
type Ciphertext struct {
	Blob             []byte
	EncryptedDataKey []byte
	Iter             uint64
}

// Serialize ciphertext object to bytes.
func serialize(c Ciphertext) ([]byte, error) {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(c)
	return b.Bytes(), err
}

// Deserialize bytes to ciphertext object.
func deserialize(b []byte) (Ciphertext, error) {
	var c Ciphertext
	r := bytes.NewReader(b)
	err := gob.NewDecoder(r).Decode(&c)
	return c, err
}
