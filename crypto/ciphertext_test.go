package crypto

import (
	"testing"

	"github.com/shibataka000/kms/encoding"
	"github.com/stretchr/testify/require"
)

func TestSerializeAndDeserialize(t *testing.T) {
	tests := []struct {
		name       string
		ciphertext Ciphertext
	}{
		{
			name: "SerializeAndDeserialize",
			ciphertext: Ciphertext{
				Blob:             []byte("abc"),
				EncryptedDataKey: []byte("def"),
				Iter:             10000,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			b, err := encoding.Serialize(tt.ciphertext)
			require.NoError(err)
			ciphertext, err := encoding.Deserialize[Ciphertext](b)
			require.NoError(err)
			require.Equal(tt.ciphertext, ciphertext)
		})
	}
}
