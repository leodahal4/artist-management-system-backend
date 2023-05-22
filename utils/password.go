package utils

import (
	"github.com/leodahal4/artist-management-system-backend/internal/config"
	passwordValidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword function to generate the bcrypt hash of the plain text password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash function to compare a bcrypt hashed password with its possible plain text equivalent.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckPasswordStrength(password string) error {
	minEntropyBits := config.GetConfig().PassEntropy
	return passwordValidator.Validate(password, minEntropyBits)
}
