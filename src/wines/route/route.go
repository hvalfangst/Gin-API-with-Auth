package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/common/middleware"
	"hvalfangst/gin-api-with-auth/src/wines/handler"
)

func ConfigureRoute(r *gin.Engine, db *pg.DB) {
	r.POST("/wines", middleware.Authorize(db, "WRITE"), handler.CreateWine(db), middleware.LogTokenActivity(db, "POST /wines"))
	r.GET("/wines/:id", middleware.Authorize(db, "READ"), handler.GetWineById(db), middleware.LogTokenActivity(db, "GET /wines/:id"))
	r.PUT("/wines/:id", middleware.Authorize(db, "EDIT"), handler.UpdateWine(db), middleware.LogTokenActivity(db, "PUT /wines/:id"))
	r.DELETE("/wines/:id", middleware.Authorize(db, "DELETE"), handler.DeleteWine(db), middleware.LogTokenActivity(db, "DELETE /wines/:id"))
}
