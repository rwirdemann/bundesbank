package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// ValidateToken returns the username on successful token validation
func ValidateToken(tokenString string) (string, error) {
	secret := []byte("kalle4ever")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["username"].(string), nil
	}
	return "", err
}
