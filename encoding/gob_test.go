package encoding

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Sample struct {
	Str string
	N   int
}

func TestSerializeAndDeserialize(t *testing.T) {
	tests := []struct {
		name string
		obj  Sample
	}{
		{
			name: "Serialization",
			obj: Sample{
				Str: "str",
				N:   0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			b, err := Serialize(tt.obj)
			require.NoError(err)
			obj, err := Deserialize[Sample](b)
			require.NoError(err)
			require.Equal(tt.obj, obj)
		})
	}
}
