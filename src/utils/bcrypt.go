package utils

import (
    "golang.org/x/crypto/bcrypt"
)

// HashPassword Password encryption, I'll leave the cost at 10 for testing purposes and not take too long to login
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
