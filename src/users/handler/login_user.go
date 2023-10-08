package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/common/security/jwt"
	"hvalfangst/gin-api-with-auth/src/users/repository"
)

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
