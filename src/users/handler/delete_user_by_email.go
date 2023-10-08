package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/users/repository"
)

func DeleteUserByEmail(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Fetch request parameter string associated with key 'email'
		email := c.Param("email")

		// Delete user by Email
		err := repository.DeleteByEmail(db, email)
		if err != nil {
			c.JSON(404, gin.H{"error": "User doesn't exist"})
			return
		}
		c.JSON(200, gin.H{"message": "User has been deleted"})
	}
}
