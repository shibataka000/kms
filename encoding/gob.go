package encoding

import (
	"bytes"
	"encoding/gob"
)

// Serialize object and return bytes.
func Serialize[T any](obj T) ([]byte, error) {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(obj)
	return b.Bytes(), err
}

// Deserialize bytes and return object.
func Deserialize[T any](b []byte) (T, error) {
	var obj T
	err := gob.NewDecoder(bytes.NewReader(b)).Decode(&obj)
	return obj, err
}
