package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	tokensRepo "hvalfangst/gin-api-with-auth/src/common/security/jwt/tokens/repository"
	"hvalfangst/gin-api-with-auth/src/common/utils/configuration"
	usersRepo "hvalfangst/gin-api-with-auth/src/users/repository"
	"strings"
	"time"
)

func Authorize(db *pg.DB, requiredAccess string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Extract the bearer token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			fmt.Println("Authorization header missing")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the header has the "Bearer " prefix
		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			fmt.Println("Bearer token missing")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract the token without the prefix
		tokenString := authHeader[len(bearerPrefix):]

		// Extract JWT configuration
		jwtConfig, _ := configuration.Get("jwt")

		// Extract encryption key from JWT configuration
		encryptionKey := jwtConfig.(configuration.Jwt).EncryptionKey

		// Decode the token with our encryption key
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				fmt.Printf("invalid signing method: %v\n", token.Header["alg"])
				c.JSON(401, gin.H{"error": "Unauthorized"})
				c.Abort()
				return nil, nil
			}
			return []byte(encryptionKey), nil
		})

		// Return 401 on decode errors
		if err != nil {
			fmt.Println("Error parsing token:", err)
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check for invalid token
		if !token.Valid {
			fmt.Println("Invalid token")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("Failed to extract claims")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract fields associated with claims
		email := claims["sub"].(string)
		access := claims["access"].(string)
		expiration := time.Unix(int64(claims["exp"].(float64)), 0)
		tokenID := claims["id"].(string)

		// Check if the token is expired
		if time.Now().After(expiration) {
			fmt.Println("Token has expired")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check whether user associated with email extracted from claims exists in DB
		user, err := usersRepo.GetByEmail(db, email)
		if err != nil {
			fmt.Println("User associated with claims not present in DB")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Parse the 'uuid' string to an uuid.UUID type
		parsedUUID, err := uuid.Parse(tokenID)
		if err != nil {
			fmt.Println("Error parsing UUID:", err)
			c.JSON(400, gin.H{"error": "Invalid UUID"})
			c.Abort()
			return
		}

		// Check whether a token associated with uuid derived for claims exists in 'tokens' table
		_, err = tokensRepo.GetTokenByID(db, parsedUUID)

		if err != nil {
			fmt.Println("Token associated with uuid derived from claims not present in DB")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Ensure that access rights derived from claims matched that stored in DB
		if user.Access != access {
			fmt.Println("Access rights mismatch between claims and DB")
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Check if the user's role grants access to the requiredAccess or its higher-level counterparts
		if !hasAccess(user.Access, requiredAccess) {
			fmt.Printf("User has [%v] access, but the minimum required access for this endpoint is [%v]\n", user.Access, requiredAccess)
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set tokenID on context as it is likely to be utilized in the final middleware function 'PersistTokenUsage'
		c.Set("tokenID", tokenID)

		// Request has been successfully authorized
		c.Next()
	}
}

func hasAccess(userAccess, requiredAccess string) bool {

	if userAccess == "DELETE" {
		return true // DELETE role grants all access
	}

	// Define hierarchy with access inheritance for higher-level rights
	accessRightsHierarchy := map[string][]string{
		"EDIT":  {"EDIT", "WRITE", "READ"},
		"WRITE": {"WRITE", "READ"},
		"READ":  {"READ"},
	}

	// Check if the required access right is contained in the access hierarchy
	for _, role := range accessRightsHierarchy[userAccess] {
		if role == requiredAccess {
			return true
		}
	}

	return false
}
