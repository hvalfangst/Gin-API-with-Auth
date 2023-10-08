package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/users/repository"
	"strconv"
)

func DeleteUserById(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Fetch request parameter string value associated with key 'id' and convert it to Integer
		userIDParam := c.Param("id")
		userID, err := strconv.ParseInt(userIDParam, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}

		// Delete user by ID
		err = repository.DeleteByID(db, userID)
		if err != nil {
			c.JSON(404, gin.H{"error": "User doesn't exist"})
			return
		}
		c.JSON(200, gin.H{"message": "User has been deleted"})
	}
}
