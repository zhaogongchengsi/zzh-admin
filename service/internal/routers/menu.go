package routers


import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/api/v1"
	"github/gin-react-admin/internal/middlewares"
)

func Menu(g *gin.Engine)  {
	app := g.Group("/api/v1").Use(middlewares.JWTAuth())
	{
		app.POST("/menu/create_menu", v1.CreateMenu)
		app.POST("/menu/delete_menu", v1.DeleteMenu)
		app.POST("/menu/find_menu", v1.GetMenuByID)
		app.POST("/menu/up_menu", v1.UpMenu)
		app.GET("/menu/get_menu", v1.GetMenus)
		app.POST("/menu/get_submenu", v1.GetSubMenus)
	}
}