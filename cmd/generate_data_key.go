package cmd

import (
	"context"
	"os"

	"github.com/shibataka000/kms/aws/kms"
	"github.com/spf13/cobra"
)

// NewRootCommand return cobra command for generate-data-key sub command.
func NewGenerateDataKeyCommand() *cobra.Command {
	var (
		keyID string
		out   string
	)

	command := &cobra.Command{
		Use:   "generate-data-key",
		Short: "Generate data key.",
		RunE: func(_ *cobra.Command, _ []string) error {
			ctx := context.Background()
			dataKey, err := kms.GenerateDataKey(ctx, keyID)
			if err != nil {
				return err
			}
			return os.WriteFile(out, dataKey, 0644)
		},
	}

	command.Flags().StringVar(&keyID, "key-id", "", "The symmetric encryption KMS key that encrypts the data key")
	command.Flags().StringVar(&out, "out", "", "The path written data key into")
	command.MarkFlagRequired("key-id")
	command.MarkFlagRequired("out")

	return command
}
