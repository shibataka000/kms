package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCommand return cobra command for root command.
func NewRootCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "kms",
		Short: "Encrypt and decrypt file using AWS KMS.",
	}
	command.AddCommand(NewGenerateDataKeyCommand())
	return command
}
