package cmd

import (
	"errors"
	"fmt"

	"github.com/Jaytpa01/encryption-cli/entities"
	"github.com/Jaytpa01/encryption-cli/entities/flags"
	"github.com/Jaytpa01/encryption-cli/service"
	"github.com/spf13/cobra"
)

func hashCommand() *cobra.Command {

	var algorithm flags.Algorithm
	var encoding flags.Encoding
	var textToHash string
	var fileToHash string
	var outputFilepath string

	hashCmd := &cobra.Command{
		Use:     "hash",
		Short:   "",
		Long:    "",
		Example: `hash "this is text to be hashed" [-a algorithm] [-e encoding] [-o output-filename.txt]`,
		Args: func(cmd *cobra.Command, args []string) error {
			// if there are no arguments && no filepath, throw an error
			if len(args) < 1 && fileToHash == "" {
				return fmt.Errorf("needs either a supplied string, or a file to hash")
			}

			// if there are arguments AND a filepath, throw an error
			if len(args) > 0 && fileToHash != "" {
				return errors.New("provide either one argument, or a file to hash using [-f file.txt] NOT both")
			}
			return nil
		},

		PreRun: func(cmd *cobra.Command, args []string) {
			// set the textToHash if a command line argument has been supplied
			if len(args) > 0 {
				textToHash = args[0]
			}
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			// fmt.Println("hash", args, algorithm, fileToHash, outputFilepath, encoding)

			hashReq := &entities.HashRequest{
				Algorithm:      algorithm,
				Encoding:       encoding,
				TextToHash:     textToHash,
				FileToHash:     fileToHash,
				OutputFilePath: outputFilepath,
			}

			if err := hashReq.ValidateAndSetDefaults(); err != nil {
				return err
			}

			return service.Hash(hashReq)
		},
	}

	hashCmd.Flags().VarP(&algorithm, "algorithm", "a", fmt.Sprintf("selects the hashing algorithm to use.\nallowed values: \"%s\", \"%s\"\ndefault: \"%s\"", flags.SHA256, flags.BCRYPT, flags.SHA256))
	hashCmd.Flags().VarP(&encoding, "encoding", "e", fmt.Sprintf("selects the encoding to output the hash to.\nallowed values: \"%s\", \"%s\", \"%s\"\ndefault: \"%s\"", flags.HEX, flags.BASE64, flags.BINARY, flags.HEX))

	hashCmd.Flags().StringVarP(&fileToHash, "file", "f", "", "optionally select a file to hash")
	hashCmd.Flags().StringVarP(&outputFilepath, "output", "o", "", "optionally save the hash to a file instead of writing to the console")

	return hashCmd
}
