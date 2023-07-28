// utils/token.go

package utils

import (
	"assign3/models"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(t time.Duration, payload models.User, privateKey string) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{
		"sub": payload.ID,
		"exp": time.Now().Add(t).Unix(),
		"iat": time.Now().Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the provided private key
	tokenString, err := token.SignedString([]byte(privateKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, publicKey string) (*jwt.Token, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(publicKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
