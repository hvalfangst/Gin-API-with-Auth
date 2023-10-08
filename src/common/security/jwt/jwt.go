package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"hvalfangst/gin-api-with-auth/src/common/utils/configuration"
	"hvalfangst/gin-api-with-auth/src/users/model"
	"time"
)

func GenerateToken(user *model.User) (string, error) {

	// Define the JWT claims
	claims := jwt.MapClaims{
		"sub":    user.Email,
		"access": user.Access,
		"exp":    time.Now().Add(time.Hour).Unix(), // Token expires in 1 hour
	}

	// Create the JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Extract JWT configuration
	jwtConfig, _ := configuration.Get("jwt")

	// Sign the token with our encryption key derived from JWT configuration
	tokenString, err := token.SignedString([]byte(jwtConfig.(configuration.Jwt).EncryptionKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
