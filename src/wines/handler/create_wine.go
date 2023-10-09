package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/wines/model"
	"hvalfangst/gin-api-with-auth/src/wines/repository"
)

func CreateWine(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Map values from RequestBody to Wine struct 'input'
		var input *model.Wine
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		// Attempt to persist a new wine to the 'wines' table in DB
		err := repository.Create(db, input)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create wine", "message": err.Error()})
			return
		}

		c.Set("snowman", "ERNST")
		c.JSON(201, gin.H{"wine": input})
	}
}
