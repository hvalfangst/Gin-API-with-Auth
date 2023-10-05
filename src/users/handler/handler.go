package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
	"hvalfangst/gin-api-with-auth/src/common/security/jwt"
	"hvalfangst/gin-api-with-auth/src/users/model"
	"hvalfangst/gin-api-with-auth/src/users/repository"
	"strconv"
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

func LoginUser(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Retrieve the username from the Gin context
		username := c.MustGet("username").(string)

		user, err := repository.GetByEmail(db, username)

		if err != nil {
			c.JSON(401, gin.H{"error": err})
			return
		}

		token, err := jwt.GenerateToken(user)

		if err != nil {
			c.JSON(401, gin.H{"error": err})
			return
		}

		c.JSON(200, gin.H{"token": token})
	}
}

func GetUserById(db *pg.DB) gin.HandlerFunc {
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
		c.JSON(200, gin.H{"user": user})
	}
}

func GetUserByEmail(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Fetch request parameter string associated with key 'email'
		email := c.Param("email")

		// Query user by email
		user, err := repository.GetByEmail(db, email)
		if err != nil {
			c.JSON(404, gin.H{"error": "User doesn't exist"})
			return
		}
		c.JSON(200, gin.H{"user": user})
	}
}

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
