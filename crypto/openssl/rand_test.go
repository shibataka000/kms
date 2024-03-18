package openssl

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRand(t *testing.T) {
	tests := []struct {
		name string
		opts RandOption
		num  int
	}{
		{
			name: "Generate32BitKey",
			opts: RandOption{
				Base64: true,
			},
			num: 32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)

			rand, err := Rand(tt.num, tt.opts)
			require.NoError(err)

			if tt.opts.Base64 {
				raw, err := base64.StdEncoding.DecodeString(string(rand))
				require.NoError(err)
				require.Equal(tt.num, len(raw))
			}
		})
	}
}
