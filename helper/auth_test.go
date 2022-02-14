package helper

import (
	"log"
	"testing"
)

func TestVerifyPassword(t *testing.T) {
	pass1 := "password1"
	hash := HashPassword(pass1)

	isTrue := VerifyPassword(hash, pass1)
	if isTrue {
		log.Println("true")
	} else {
		log.Println("false")
	}
}

func TestHashPassword(t *testing.T) {
	pass1 := "password1"

	hashpass := HashPassword(pass1)
	log.Println(hashpass)
}
