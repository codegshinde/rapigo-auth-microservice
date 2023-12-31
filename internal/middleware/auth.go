package middleware

import (
	"net/http"
	"rapigo/internal/models"
	"rapigo/internal/service"
	"rapigo/pkg/utils"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// GetAuthenticatedAdmin extracts and validates the JWT token from the Authorization header.
func GetAuthenticatedAdmin(r *http.Request) (*models.AdminResponse, error) {
	// Extract the JWT token from the Authorization header
	tokenString := extractTokenFromHeader(r)

	if tokenString == "" {
		return nil, jwt.ErrSignatureInvalid
	}

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return utils.JWTSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	// Access admin information from the token claims
	claims, ok := token.Claims.(*utils.Claims)
	if !ok {
		return nil, jwt.ErrSignatureInvalid
	}

	// Assuming you have a function to retrieve admin information based on AdminID
	admin, err := service.GetAdminByID(claims.AdminID)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

// extractTokenFromHeader extracts the JWT token from the Authorization header.
func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	// The Authorization header should be in the format "Bearer <token>"
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}
