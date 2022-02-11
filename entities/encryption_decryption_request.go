package entities

import (
	"log"
	"path/filepath"

	"github.com/Jaytpa01/encryption-cli/utils"
)

type EncryptionDecryptionRequest struct {
	Password       string
	InputFilePath  string
	OutputFileName string
}

func (r *EncryptionDecryptionRequest) ValidateRequest() {
	if ext := filepath.Ext(r.InputFilePath); len(ext) == 0 {
		log.Fatal("error encrypting file. invalid file specified")
	}

	// all other input fields are validated by cobra when trying to use the command line
}

func (r *EncryptionDecryptionRequest) PrepareOutputFilename(fileExtension string) {
	if r.OutputFileName == "" {
		r.OutputFileName = utils.GetFileNameWithoutExtension(r.InputFilePath)
	}
	r.OutputFileName += fileExtension
}
