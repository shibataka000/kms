package rand

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBytes(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{
			name: "32Byte",
			n:    32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			b, err := Bytes(tt.n)
			require.NoError(err)
			require.Equal(tt.n, len(b))
			require.NotEqual(make([]byte, tt.n), b)
		})
	}
}
