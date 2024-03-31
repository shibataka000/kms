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
	r := bytes.NewReader(b)
	err := gob.NewDecoder(r).Decode(&obj)
	return obj, err
}
