package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/wines/repository"
	"strconv"
)

func DeleteWine(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Fetch request parameter string value associated with key 'id' and convert it to Integer
		wineIdParameter := c.Param("id")
		wineId, err := strconv.ParseInt(wineIdParameter, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid wine ID"})
			return
		}

		// Delete wine by ID
		err = repository.Delete(db, wineId)
		if err != nil {
			c.JSON(404, gin.H{"error": "Wine doesn't exist"})
			return
		}
		c.JSON(200, gin.H{"message": "Wine has been deleted"})
	}
}
