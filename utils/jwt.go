// Package utils provides utility functions for JWT handling.
package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(userID uint, role string) (string, error) {
	// Create JWT with claims: userID, role, exp
	claims := jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 72).Unix(), // Token expiration time
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with a secret key

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		// Handle error
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method is valid using withvalidmethod.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method: %v", token.Method)
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		// Handle error
		return nil, fmt.Errorf("error parsing token: %w", err)
	}
	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	// Return the token if valid
	return token, nil

}
