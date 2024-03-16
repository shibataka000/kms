package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

// Decrypt ciphertext.
func Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := kms.NewFromConfig(cfg)

	out, err := client.Decrypt(ctx, &kms.DecryptInput{
		CiphertextBlob: ciphertext,
	})
	if err != nil {
		return nil, err
	}

	return out.Plaintext, nil
}
