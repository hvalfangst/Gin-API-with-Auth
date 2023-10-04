package users

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func ConfigureRoute(r *gin.Engine, db *pg.DB) {
	r.POST("/users", CreateUserHandler(db))
	r.GET("/users/:id", GetUserByIDHandler(db))
	r.DELETE("/users/:id", DeleteUserByIdHandler(db))
	r.GET("/users/email/:email", GetUserByEmailHandler(db))
	r.DELETE("/users/email/:email", DeleteUserByEmailHandler(db))
}
