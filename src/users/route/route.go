package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/common/middleware"
	"hvalfangst/gin-api-with-auth/src/users/handler"
)

func ConfigureRoute(r *gin.Engine, db *pg.DB) {
	r.POST("/users", handler.CreateUser(db))
	r.POST("/users/login", middleware.Authenticate(db), handler.LoginUser(db))
	r.GET("/users/:id", handler.GetUserById(db))
	r.DELETE("/users/:id", handler.DeleteUserById(db))
	r.GET("/users/email/:email", handler.GetUserByEmail(db))
	r.DELETE("/users/email/:email", handler.DeleteUserByEmail(db))
	r.PATCH("/users/deactivate/:id", handler.DeactivateUser(db))
	r.PATCH("/users/mark-for-deletion/:id", handler.MarkUserForDeletion(db))
}
