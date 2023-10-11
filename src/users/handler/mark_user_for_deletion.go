package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/users/repository"
	"net/http"
	"strconv"
	"time"
)

func MarkUserForDeletion(db *pg.DB) gin.HandlerFunc {
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

		// Check if the user is already marked for deletion
		if user.ForDeletion {
			// Calculate and return the time difference
			deletionTime := user.DeletionTime
			currentTime := time.Now()
			timeDifference := deletionTime.Sub(currentTime).Seconds()

			c.JSON(http.StatusOK, gin.H{"message": "User is already marked for deletion", "timeUntilDeletion": timeDifference})
			return
		}

		// Parse the time in seconds from the query parameter "deletionTimeInSeconds"
		deletionTimeInSecondsStr := c.DefaultQuery("deletionTimeInSeconds", "0")
		deletionTimeInSeconds, err := strconv.Atoi(deletionTimeInSecondsStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid deletion time"})
			return
		}

		// Calculate the deletion time by adding the desired seconds to the current time
		deletionTime := time.Now().Add(time.Second * time.Duration(deletionTimeInSeconds))

		// Mark the user for deletion by updating the "ForDeletion" and "DeletionTime" fields
		user.ForDeletion = true
		user.DeletionTime = deletionTime

		// Update the user in the database
		if err := repository.Update(db, user.ID, user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark user for deletion"})
			return
		}

		// Logical delete succeeded
		c.JSON(http.StatusOK, gin.H{"message": "User marked for deletion"})
	}
}
