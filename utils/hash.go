package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14) // bcrypt.DefaultCost=10
	return string(hashedPassword[:]), err
}

func CheckPasswordHash(password, hashPassword string) bool {
	// Compare the given password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

