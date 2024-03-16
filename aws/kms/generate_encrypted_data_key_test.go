package kms

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateEncryptedDataKey(t *testing.T) {
	tests := []struct {
		name  string
		keyID string
	}{
		{
			name:  "GenerateDataKeyUsingAliasShibataka000KmsKey",
			keyID: "alias/shibataka000/kms",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require := require.New(t)
			ctx := context.Background()
			_, err := GenerateEncryptedDataKey(ctx, tt.keyID)
			require.NoError(err)
		})
	}
}