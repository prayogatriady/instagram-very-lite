package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(hashPassword, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(userPassword))

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
