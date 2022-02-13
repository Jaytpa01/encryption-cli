package service

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/Jaytpa01/encryption-cli/entities"
	"github.com/Jaytpa01/encryption-cli/entities/flags"
	"github.com/Jaytpa01/encryption-cli/utils"
)

func Hash(hashRequest *entities.HashRequest) error {

	dataToHash, err := determineDataToHash(hashRequest)
	if err != nil {
		return nil
	}

	hashedData, err := hashDataWithAlgorithm(dataToHash, hashRequest.Algorithm)
	if err != nil {
		return err
	}

	encodedData := encodeHashedData(hashedData, hashRequest.Encoding)

	// if no output specified, print hashed data to terminal
	if hashRequest.OutputFilePath == "" {
		fmt.Println(string(encodedData))
		return nil
	}

	// if no output file extension has been supplied, default to .txt
	if filepath.Ext(hashRequest.OutputFilePath) == "" {
		hashRequest.OutputFilePath += ".txt"
	}

	return utils.WriteBytesToFile(hashRequest.OutputFilePath, encodedData)

}

func encodeHashedData(data []byte, encoding flags.Encoding) []byte {
	encodedData := make([]byte, 128)

	var numberOfBytesWritten int
	switch encoding {
	case flags.HEX:
		numberOfBytesWritten = hex.Encode(encodedData, data)

	case flags.BASE64:
		numberOfBytesWritten = base64.StdEncoding.EncodedLen(len(data))
		base64.StdEncoding.Encode(encodedData, data)

	case flags.BINARY:
		return data
	default:
		numberOfBytesWritten = hex.Encode(encodedData, data)
	}

	// we trim the array to the amount of bytes actually written
	// to remove writing nil values to a file
	encodedData = encodedData[:numberOfBytesWritten]
	return encodedData

}

func determineDataToHash(r *entities.HashRequest) ([]byte, error) {
	var dataToHash []byte

	// if there is text to hash, we set that as the data to hash
	if r.TextToHash != "" {
		dataToHash = append(dataToHash, []byte(r.TextToHash)...)
	} else {
		// we assume a file is supplied, lets try and open it
		// program will exit if this fails
		dataToHash = utils.ReadFile(r.FileToHash)
	}

	if len(dataToHash) == 0 {
		return nil, errors.New("failed to hash data. could not determine data to hash")
	}

	return dataToHash, nil
}

func hashDataWithAlgorithm(data []byte, algorithm flags.Algorithm) ([]byte, error) {
	var hashedData []byte
	var err error

	switch algorithm {
	case flags.SHA256:
		hashedData, err = utils.HashBytesSHA256(data)
		if err != nil {
			return nil, err
		}

	case flags.BCRYPT:
		hashedData, err = utils.HashBCrypt(data)
		if err != nil {
			return nil, err
		}
	// default to use SHA256
	default:
		hashedData, err = utils.HashBytesSHA256(data)
		if err != nil {
			return nil, err
		}
	}

	return hashedData, nil

}
