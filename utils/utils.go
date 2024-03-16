package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"sanbercode-final-project-Sahrul-Yoyo/config"
	"sanbercode-final-project-Sahrul-Yoyo/models"
)

// HashPassword hashes a plain text password
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(user *models.User, cfg *config.Config) (string, error) {
	// ... implement JWT token generation logic using cfg.JWTSecret
	token := "JWT token"
	return token, nil
}
