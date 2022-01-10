package routers

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/api/cos"
	"github/gin-react-admin/internal/middlewares"
)

func CosRouter (g *gin.Engine)  {
	app := g.Group("/api/v1/cos").Use(middlewares.JWTAuth())
	{
		app.POST("/get_temporary_key", cos.GetTempKey)
	}
}