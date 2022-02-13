package cmd

import (
	"github.com/Jaytpa01/encryption-cli/entities"
	"github.com/Jaytpa01/encryption-cli/service"
	"github.com/spf13/cobra"
)

func decryptCommand() *cobra.Command {

	decryptCmd := &cobra.Command{
		Use:   "decrypt",
		Short: "Decrypt the specified file.",
		Long: `Decrypts the specified file using the Advanced Encryption Standard (AES-256). Requires a password and input filename.
	You may also supply an optional output file name. When specifying an output file, you do NOT need to supply a file extension.`,
		Run: func(cmd *cobra.Command, args []string) {

			decryptionRequest := &entities.EncryptionDecryptionRequest{Password: password, InputFilePath: inputFilePath, OutputFileName: outputFileName}
			decryptionRequest.ValidateRequest()
			service.DecryptFile(*decryptionRequest)
		},
	}

	decryptCmd.MarkPersistentFlagRequired("file")
	decryptCmd.MarkPersistentFlagRequired("password")

	return decryptCmd
}
