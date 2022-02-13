package utils

import "golang.org/x/crypto/bcrypt"

func HashBCrypt(data []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
}
