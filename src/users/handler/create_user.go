package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
	"hvalfangst/gin-api-with-auth/src/users/model"
	"hvalfangst/gin-api-with-auth/src/users/repository"
)

func CreateUser(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Map values from RequestBody to User struct 'input'
		var input *model.User
		if err := c.BindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		// Hash the user's password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to hash password"})
			return
		}
		input.Password = string(hashedPassword)

		// Attempt to persist a new user to the 'users' table in DB
		err = repository.Create(db, input)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to create user", "message": err.Error()})
			return
		}
		c.JSON(201, gin.H{"user": input})
	}
}
