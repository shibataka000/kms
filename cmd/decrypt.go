package cmd

import (
	"context"
	"os"

	"github.com/shibataka000/kms/crypto"
	"github.com/spf13/cobra"
)

// NewDecryptCommand return cobra command object for `kms decrypt` command.
func NewDecryptCommand() *cobra.Command {
	var (
		in  string
		out string
		rm  bool
	)

	command := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt file by envelope encryption using AWS KMS.",
		RunE: func(_ *cobra.Command, _ []string) error {
			ctx := context.Background()
			ciphertext, err := os.ReadFile(in)
			if err != nil {
				return err
			}
			plaintext, err := crypto.Decrypt(ctx, ciphertext)
			if err != nil {
				return err
			}
			err = os.WriteFile(out, plaintext, 0644)
			if err != nil {
				return err
			}
			if rm {
				return os.Remove(in)
			}
			return nil
		},
	}

	command.Flags().StringVar(&in, "in", "", "The path to ciphertext file")
	command.Flags().StringVar(&out, "out", "", "The path to plaintext file written into")
	command.Flags().BoolVar(&rm, "rm", false, "If true, delete ciphertext file after decryption")
	command.MarkFlagRequired("in")  // nolint:errcheck
	command.MarkFlagRequired("out") // nolint:errcheck

	return command
}
