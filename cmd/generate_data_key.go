package cmd

import (
	"context"
	"os"

	"github.com/shibataka000/kms/aws/kms"
	"github.com/spf13/cobra"
)

// NewGenerateDataKeyCommand return cobra command for generate-data-key sub command.
func NewGenerateDataKeyCommand() *cobra.Command {
	var (
		key string
		out string
	)

	command := &cobra.Command{
		Use:   "generate-data-key",
		Short: "Generate data key.",
		RunE: func(_ *cobra.Command, _ []string) error {
			ctx := context.Background()
			dataKey, err := kms.GenerateDataKey(ctx, key)
			if err != nil {
				return err
			}
			return os.WriteFile(out, dataKey, 0644)
		},
	}

	command.Flags().StringVar(&key, "kms-key", "", "The symmetric encryption KMS key that encrypts the data key")
	command.Flags().StringVar(&out, "out", "", "The path written encrypted data key into")
	command.MarkFlagRequired("kms-key") // nolint:errcheck
	command.MarkFlagRequired("out")     // nolint:errcheck

	return command
}
