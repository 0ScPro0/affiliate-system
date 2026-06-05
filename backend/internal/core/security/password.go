package core_security

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a plain text password and returns a bcrypt hash of the password.
// It uses bcrypt's DefaultCost which provides a good balance between security and performance.
// Returns the hashed password as a string or an error if hashing fails.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// VerifyPassword compares a plain text password with a bcrypt hash.
// Returns true if the password matches the hash, false otherwise.
func VerifyPassword(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}