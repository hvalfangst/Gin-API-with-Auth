package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/wines/model"
	"hvalfangst/gin-api-with-auth/src/wines/repository"
	"strconv"
)

func UpdateWine(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract wine ID from the URL parameter
		wineIDParam := c.Param("id")
		wineID, err := strconv.Atoi(wineIDParam)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid wine ID"})
			return
		}

		// Map values from RequestBody to Wine struct 'input'
		var input *model.Wine
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		// Update the wine with the given ID
		err = repository.Update(db, int64(wineID), input)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to update wine", "message": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Wine updated successfully"})
	}
}
