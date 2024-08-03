package hashing

import (

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	// log.Printf("password that we get %v", password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// log.Printf("hash that we get %v",hashedPassword)
	return string(hashedPassword), nil
}
