package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func MakeSHA256(password string) string {
	encryptedPassword := sha256.New()
	encryptedPassword.Write([]byte(password))
	hashedPassword := hex.EncodeToString(encryptedPassword.Sum(nil))

	return hashedPassword
}
