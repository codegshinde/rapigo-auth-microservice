package utils

import (
	"rapigo/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTSecretKey is the secret key used for JWT token generation.
var JWTSecretKey = []byte(config.GetEnvVariable("JWT_SECRET"))

// Claims represents the claims in a JWT token.
type Claims struct {
	AdminID string `json:"adminId"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token for the provided AdminID.
func GenerateToken(adminID string) (string, error) {
	claims := &Claims{
		AdminID: adminID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "probys-auth",
			Subject:   "probys-auth-token",
			ID:        "1",
			Audience:  []string{"probys-small-capital"},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWTSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
