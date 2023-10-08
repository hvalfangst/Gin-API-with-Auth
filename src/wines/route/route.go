package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/common/middleware"
	"hvalfangst/gin-api-with-auth/src/wines/handler"
)

func ConfigureRoute(r *gin.Engine, db *pg.DB) {
	r.POST("/wines", middleware.Authorize(db, "WRITE"), handler.CreateWine(db))
	r.GET("/wines/:id", middleware.Authorize(db, "READ"), handler.GetWineById(db))
	r.PUT("/wines/:id", middleware.Authorize(db, "EDIT"), handler.UpdateWine(db))
	r.DELETE("/wines/:id", middleware.Authorize(db, "DELETE"), handler.DeleteWine(db))
}
