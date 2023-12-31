package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// ComparePassword compares a hashed password with a plaintext password.
func ComparePassword(hashedPassword, plaintextPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintextPassword))
	return err
}

// HashPassword hashes the provided password using bcrypt and returns the base64-encoded hash.
func HashPassword(password string) (string, error) {
	// Generate the bcrypt hash for the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Convert the hash to a string
	hashString := string(hash)

	return hashString, nil
}
