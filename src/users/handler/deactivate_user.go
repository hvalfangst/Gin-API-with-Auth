package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/users/repository"
	"net/http"
	"strconv"
)

func DeactivateUser(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Fetch request parameter string value associated with key 'id' and convert it to Integer
		userIDParam := c.Param("id")
		userID, err := strconv.ParseInt(userIDParam, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}

		// Query user by ID
		user, err := repository.GetById(db, userID)
		if err != nil {
			c.JSON(404, gin.H{"error": "User doesn't exist"})
			return
		}

		// Check if the user is already deactivated
		if user.Deactivated {
			c.JSON(http.StatusOK, gin.H{"message": "User is already deactivated"})
			return
		}

		user.Deactivated = true

		// Update the user in the database
		if err := repository.Update(db, user.ID, user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deactivate user"})
			return
		}

		// Deactivation succeeded
		c.JSON(http.StatusOK, gin.H{"message": "User deactivated"})
	}
}
