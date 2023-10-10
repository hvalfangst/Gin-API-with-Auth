package jwt

import (
	"github.com/go-pg/pg/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"hvalfangst/gin-api-with-auth/src/common/utils/configuration"
	tokenModel "hvalfangst/gin-api-with-auth/src/tokens/model"
	tokenRepo "hvalfangst/gin-api-with-auth/src/tokens/repository"
	"hvalfangst/gin-api-with-auth/src/users/model"
	"time"
)

func GenerateToken(db *pg.DB, user *model.User) (string, error) {

	// Generate a new v4 UUID as identifier for token. This value will be used as primary key
	tokenID := uuid.New()

	// Define the JWT claims
	claims := jwt.MapClaims{
		"sub":    user.Email,
		"access": user.Access,
		"id":     tokenID,                          // PK utilized to query table 'tokens'
		"exp":    time.Now().Add(time.Hour).Unix(), // Token expires in 1 hour
	}

	// CreateToken the JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Extract JWT configuration
	jwtConfig, _ := configuration.Get("jwt")

	// Sign the token with our encryption key derived from JWT configuration
	tokenString, err := token.SignedString([]byte(jwtConfig.(configuration.Jwt).EncryptionKey))
	if err != nil {
		return "", err
	}

	// Insert struct containing necessary metadata in order to query tokens based from claims
	err = persistToken(db, user, tokenID, err)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func persistToken(db *pg.DB, user *model.User, tokenID uuid.UUID, err error) error {

	tokenStruct := tokenModel.Token{
		ID:           tokenID,
		CreationDate: time.Now(),
		UserID:       user.ID,
	}

	err = tokenRepo.CreateToken(db, &tokenStruct)
	return err
}
