package routers


import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/api/v1"
	"github/gin-react-admin/internal/middlewares"
)

func Role(g *gin.Engine)  {
	app := g.Group("/api/v1").Use(middlewares.JWTAuth())
	{
		app.POST("/role/create_role", v1.CreateRole)
		app.GET("/role/get_roles", v1.GetAuthLIst)
	}
}