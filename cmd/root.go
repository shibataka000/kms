package cmd

import (
	"github.com/spf13/cobra"
)

// NewCommand return cobra command.
func NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "kms",
		Short: "Encrypt and decrypt file using AWS KMS.",
	}
	return command
}
