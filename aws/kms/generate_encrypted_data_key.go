package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

// GenerateEncryptedDataKey.
func GenerateEncryptedDataKey(ctx context.Context, keyID string) ([]byte, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	client := kms.NewFromConfig(cfg)

	out, err := client.GenerateDataKeyWithoutPlaintext(ctx, &kms.GenerateDataKeyWithoutPlaintextInput{
		KeyId:   aws.String(keyID),
		KeySpec: types.DataKeySpecAes256,
	})
	if err != nil {
		return nil, err
	}

	return out.CiphertextBlob, nil
}
