package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"hvalfangst/gin-api-with-auth/src/tokens/repository"
)

func DeleteToken(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Fetch request parameter string value associated with key 'id' and convert it to Integer
		tokenIDParameter := c.Param("id")

		// Parse the string to an uuid.UUID type
		tokenID, err := uuid.Parse(tokenIDParameter)
		if err != nil {
			fmt.Println("Error parsing UUID:", err)
			c.JSON(400, gin.H{"error": "Invalid UUID"})
			c.Abort()
			return
		}

		// First we need to delete any activity rows associated with the given token ID
		err = repository.DeleteTokenActivity(db, tokenID)
		if err != nil {
			c.JSON(404, gin.H{"error": "Token activity doesn't exist"})
			return
		}

		// Delete token by PK
		err = repository.DeleteToken(db, tokenID)
		if err != nil {
			c.JSON(404, gin.H{"error": "Token doesn't exist"})
			return
		}
		c.JSON(200, gin.H{"result": "Token has been deleted"})
	}
}
