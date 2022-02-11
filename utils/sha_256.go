package utils

import "crypto/sha256"

func HashStringSHA256Bytes(stringToHash string) ([]byte, error) {
	hashedString, err := hashStringSHA256(stringToHash)
	if err != nil {
		return nil, err
	}

	return hashedString, nil
}

func hashStringSHA256(stringToHash string) ([]byte, error) {
	hasher := sha256.New()

	if _, err := hasher.Write([]byte(stringToHash)); err != nil {
		return nil, err
	}

	hash := hasher.Sum(nil)
	return hash, nil
}
