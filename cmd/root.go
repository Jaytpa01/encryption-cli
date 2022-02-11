package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ecli",
	Short: "ECLI Encryption is a file encryption/decryption CLI",
	Long:  `ECLI is a small CLI used to encrypt/decrypt files locally given a password`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var (
	password       string
	inputFilePath  string
	outputFileName string
)

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "REQUIRED - password used to encrypt or decrypt your file")
	rootCmd.MarkPersistentFlagRequired("password")

	rootCmd.PersistentFlags().StringVarP(&inputFilePath, "file", "f", "", "REQUIRED - the name of the file you would like to encrypt")
	rootCmd.MarkPersistentFlagRequired("file")

	rootCmd.PersistentFlags().StringVarP(&outputFileName, "output", "o", "", "the name of the encrypted/decrypted output file")

	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
}
