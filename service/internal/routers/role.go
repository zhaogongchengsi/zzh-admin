package routers


import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/api/v1"
	"github/gin-react-admin/internal/middlewares"
)

func Role(g *gin.Engine)  {
	app := g.Group("/api/v1/role").Use(middlewares.JWTAuth())
	{
		app.POST("/create_role", v1.CreateRole)
		app.GET("/get_roles", v1.GetAuthLIst)
	}
}