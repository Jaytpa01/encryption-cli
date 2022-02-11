package service

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"

	"github.com/Jaytpa01/encryption-cli/entities"
	"github.com/Jaytpa01/encryption-cli/utils"
)

func DecryptFile(decryptionRequest entities.EncryptionDecryptionRequest) {

	dataToDecrypt := utils.ReadFile(decryptionRequest.InputFilePath)

	key, err := utils.HashStringSHA256Bytes(decryptionRequest.Password)
	if err != nil {
		log.Fatalf("error decrypting file, failed to generate a hash from provided password: %s", err.Error())
	}

	decryptedData, err := utils.DecryptDataAES256(key, dataToDecrypt)
	if err != nil {
		log.Fatalf("failed to decrypt data: %s", err.Error())
	}

	fileExt := filepath.Ext(string(decryptedData))

	// we trim the file extension from the rest of the data
	decryptedData = bytes.TrimSuffix(decryptedData, []byte(fileExt))

	decryptionRequest.PrepareOutputFilename(fileExt)

	err = utils.WriteBytesToFile(decryptionRequest.OutputFileName, decryptedData)
	if err != nil {
		log.Fatalf("error writing decrypted data to file: %s", err.Error())
	}

	fmt.Printf("successfully decrypted file: %s\n", decryptionRequest.OutputFileName)
}
