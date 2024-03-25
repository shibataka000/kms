package cmd

import (
	"context"
	"os"

	"github.com/shibataka000/kms/crypto"
	"github.com/spf13/cobra"
)

// NewEncryptCommand return cobra command object for `kms encrypt` command.
func NewEncryptCommand() *cobra.Command {
	var (
		keyID string
		in    string
		out   string
		iter  uint64
		rm    bool
	)

	command := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypt file by envelope encryption using AWS KMS.",
		RunE: func(_ *cobra.Command, _ []string) error {
			ctx := context.Background()
			plaintext, err := os.ReadFile(in)
			if err != nil {
				return err
			}
			ciphertext, err := crypto.Encrypt(ctx, keyID, plaintext, iter)
			if err != nil {
				return err
			}
			err = os.WriteFile(out, ciphertext, 0644)
			if err != nil {
				return err
			}
			if rm {
				return os.Remove(in)
			}
			return nil
		},
	}

	command.Flags().StringVar(&keyID, "key-id", "", "The symmetric encryption KMS key ID that encrypts the data key")
	command.Flags().StringVar(&in, "in", "", "The path to plaintext file")
	command.Flags().StringVar(&out, "out", "", "The path to ciphertext file written into")
	command.Flags().Uint64Var(&iter, "iter", 100000, "The iteration count for PBKDF2")
	command.Flags().BoolVar(&rm, "rm", false, "If true, delete plaintext file after encryption")
	command.MarkFlagRequired("key-id") // nolint:errcheck
	command.MarkFlagRequired("in")     // nolint:errcheck
	command.MarkFlagRequired("out")    // nolint:errcheck

	return command
}
