package auth

import (
	"log"
	"github.com/dgrijalva/jwt-go"
)

// Authenticates a auth by username and password. Returns JWT, true on success, "", false otherwise
func Authenticate(username string, password string, userService UserService) (string, bool) {
	if !userService.Authenticate(username, password) {
		return "", false
	}

	// Build and sign JWS token
	secret := []byte("kalle4ever")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Println(err)
		return "", false
	}

	return tokenString, true
}
