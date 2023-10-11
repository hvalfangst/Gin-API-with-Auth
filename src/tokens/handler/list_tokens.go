package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/tokens/repository"
)

func ListTokens(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokens, err := repository.ListTokens(db)
		if err != nil {
			c.JSON(404, gin.H{"error": "Could not list tokens"})
			return
		}

		if len(tokens) == 0 {
			c.JSON(200, gin.H{"message": "No tokens were found"})
			return
		}

		c.JSON(200, gin.H{"tokens": tokens})
	}
}
