package cmd

import (
	"github.com/spf13/cobra"
)

var (
	password       string
	inputFilePath  string
	outputFileName string
)

func RootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ecli",
		Short: "ECLI Encryption is a file encryption/decryption CLI",
		Long:  `ECLI is a small CLI used to encrypt/decrypt files locally given a password`,
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "REQUIRED - password used to encrypt or decrypt your file")

	// rootCmd.PersistentFlags().StringVarP(&inputFilePath, "file", "f", "", "REQUIRED - the name of the file you would like to encrypt")

	// rootCmd.PersistentFlags().StringVarP(&outputFileName, "output", "o", "", "the name of the encrypted/decrypted output file")

	rootCmd.AddCommand(encryptCommand())
	rootCmd.AddCommand(decryptCommand())
	rootCmd.AddCommand(hashCommand())

	return rootCmd
}
