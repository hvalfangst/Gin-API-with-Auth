package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
	"hvalfangst/gin-api-with-auth/src/users/repository"
)

func Authenticate(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Extract username and password from Basic Auth headers
		username, password, hasAuth := c.Request.BasicAuth()

		if !hasAuth {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Fetch user based on email
		user, err := repository.GetByEmail(db, username)

		// Hash and compare password derived from header with that of hashed password fetched from database
		if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Store the username in the Gin context so that it may be used in the 'Login' handler
		c.Set("username", username)

		// User associated with request have been successfully authenticated
		c.Next()
	}
}
