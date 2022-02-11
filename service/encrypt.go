package service

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Jaytpa01/encryption-cli/entities"
	"github.com/Jaytpa01/encryption-cli/utils"
)

func EncryptFile(encryptionRequest entities.EncryptionDecryptionRequest) {

	dataToEncrypt := utils.ReadFile(encryptionRequest.InputFilePath)

	// we get the original file extension and append it to the end of the encrypted file
	// this is so we can retrieve the correct file extension and set it when we decrypt the file
	fileExtensionInBytes := []byte(filepath.Ext(encryptionRequest.InputFilePath))
	dataToEncrypt = append(dataToEncrypt, fileExtensionInBytes...)

	key, err := utils.HashStringSHA256Bytes(encryptionRequest.Password)
	if err != nil {
		log.Fatalf("error encrypting file, failed to generate a hash from provided password: %s", err.Error())
	}

	encryptedData, err := utils.EncryptDataAES256(key, dataToEncrypt)
	if err != nil {
		log.Fatalf("error encrypting file: %s", err.Error())
	}

	encryptionRequest.PrepareOutputFilename(".aes")

	err = utils.WriteBytesToFile(encryptionRequest.OutputFileName, encryptedData)
	if err != nil {
		log.Fatalf("error writing encrypted data to file: %s", err.Error())
	}

	fmt.Printf("successfully encrypted file: %s\n", encryptionRequest.OutputFileName)
}
