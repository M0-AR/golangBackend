package middleware

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// hashes the new incomming password with bcrypt
func GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
