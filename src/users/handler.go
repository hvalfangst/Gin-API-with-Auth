package users

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"strconv"
)

func CreateUserHandler(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input *User
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		err := create(db, input)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create user", "message": err.Error()})
			return
		}
		c.JSON(201, gin.H{"user": input})
	}
}

func GetUserByIDHandler(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDParam := c.Param("id")
		userID, err := strconv.ParseInt(userIDParam, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}

		user, err := getByID(db, userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve user by ID"})
			return
		}
		c.JSON(200, gin.H{"user": user})
	}
}

func GetUserByEmailHandler(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")

		user, err := getByEmail(db, email)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to retrieve user by email"})
			return
		}
		c.JSON(200, gin.H{"user": user})
	}
}
