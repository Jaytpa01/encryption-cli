package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func HashStringSHA256Bytes(stringToHash string) ([]byte, error) {
	hashedString, err := hashStringSHA256(stringToHash)
	if err != nil {
		return nil, err
	}

	return hashedString, nil
}

func HashStringSHA256Hex(stringToHash string) (string, error) {
	hashedString, err := hashStringSHA256(stringToHash)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hashedString), nil
}

func HashStringSHA256Base64(stringToHash string) (string, error) {
	hashedString, err := hashStringSHA256(stringToHash)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(hashedString), nil
}

func HashBytesSHA256(data []byte) ([]byte, error) {
	return hashDataSHA256(data)
}

func hashStringSHA256(stringToHash string) ([]byte, error) {
	return hashDataSHA256([]byte(stringToHash))
}

func hashDataSHA256(data []byte) ([]byte, error) {
	hasher := sha256.New()

	if _, err := hasher.Write(data); err != nil {
		return nil, err
	}

	hash := hasher.Sum(nil)
	return hash, nil
}
