package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rapigo/config"
	"time"

	//third parties import
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Token   string `json:"token"`
	AdminId string `json:"adminId"`
}

var jwtSecretKey = []byte(config.GetEnvVariables("JWT_SECRET"))

type Claims struct {
	AdminId string `json:"adminId"`
	jwt.RegisteredClaims
}

func GenerateToken(AdminId string) (string, error) {
	claims := &Claims{
		AdminId: AdminId,
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

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func WriteErrorResponse(w http.ResponseWriter, status int, errMsg string) {
	w.WriteHeader(status)
	w.Write([]byte(errMsg))
}

func LogError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// createJSONResponse creates a JSON response from the provided data
func CreateJSONResponse(data interface{}) ([]byte, error) {
	responseJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return responseJSON, nil
}

func ComparePassword(hashedPassword []byte, plaintextPassword string) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(plaintextPassword))
	return err
}
