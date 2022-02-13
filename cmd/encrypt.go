package cmd

import (
	"github.com/Jaytpa01/encryption-cli/entities"
	"github.com/Jaytpa01/encryption-cli/service"
	"github.com/spf13/cobra"
)

func encryptCommand() *cobra.Command {
	encryptCmd := &cobra.Command{
		Use:   "encrypt",
		Short: "Encrypts the specified file",
		Long:  `Encrypts the specified file using the Advanced Encryption Standard (AES-256). Requires a password and input filename. You may also supply an optional output file name`,
		Run: func(cmd *cobra.Command, args []string) {

			encryptionRequest := &entities.EncryptionDecryptionRequest{Password: password, InputFilePath: inputFilePath, OutputFileName: outputFileName}
			encryptionRequest.ValidateRequest()
			service.EncryptFile(*encryptionRequest)
		},
	}

	encryptCmd.MarkPersistentFlagRequired("file")
	encryptCmd.MarkPersistentFlagRequired("password")

	return encryptCmd
}
