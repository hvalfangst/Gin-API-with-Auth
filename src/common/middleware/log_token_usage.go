package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	tokensModel "hvalfangst/gin-api-with-auth/src/common/security/jwt/tokens/model"
	tokensRepo "hvalfangst/gin-api-with-auth/src/common/security/jwt/tokens/repository"
	"time"
)

func PersistTokenUsage(db *pg.DB, endpoint string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Retrieve the snowman name from the Gin context
		tokenID := c.MustGet("tokenID").(string)

		// Parse the 'uuid' string to an uuid.UUID type
		parsedUUID, err := uuid.Parse(tokenID)
		if err != nil {
			fmt.Println("Error parsing UUID:", err)
			c.JSON(400, gin.H{"error": "Invalid UUID"})
			c.Abort()
			return
		}

		tokenUsageStruct := tokensModel.TokenUsage{
			TokenID:  parsedUUID,
			Token:    nil,
			Endpoint: endpoint,
			UsedAt:   time.Now(),
		}

		err = tokensRepo.CreateTokenUsage(db, &tokenUsageStruct)

		if err != nil {
			fmt.Printf("Failed to persist TokenUsage associated with Token ID: %v\n", tokenID)
			c.Abort()
			return
		}

		// Chain of handlers has succeeded
		c.Next()
	}
}
