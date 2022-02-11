package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func EncryptDataAES256(key, dataToEncrypt []byte) ([]byte, error) {
	keyLength := len(key)
	if keyLength < 32 {
		return nil, fmt.Errorf("encryption key must be 32 bytes, received: %d byte(s)", keyLength)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encryptedData := aesGcm.Seal(nonce, nonce, dataToEncrypt, nil)
	return encryptedData, nil
}

func DecryptDataAES256(key, dataToDecrypt []byte) ([]byte, error) {
	keyLength := len(key)
	if keyLength < 32 {
		return nil, fmt.Errorf("encryption key must be 32 bytes, received: %d byte(s)", keyLength)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGcm.NonceSize()
	// the nonce is prepended to the ciphertext, so we extract the nonce and the cipertext as their own seperate byte array
	nonce, ciphertext := dataToDecrypt[:nonceSize], dataToDecrypt[nonceSize:]

	decryptedData, err := aesGcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return decryptedData, nil
}
