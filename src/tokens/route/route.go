package route

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"hvalfangst/gin-api-with-auth/src/tokens/handler"
)

func ConfigureRoute(r *gin.Engine, db *pg.DB) {
	r.GET("/tokens/:id", handler.GetToken(db))
	r.GET("/token-activities/:id", handler.GetTokenActivity(db))
}
