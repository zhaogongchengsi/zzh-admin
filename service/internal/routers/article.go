package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github/gin-react-admin/internal/api/v1"
	"github/gin-react-admin/internal/middlewares"
)

func NewArticle (g *gin.Engine) {
	app := g.Group("/api/v1/article").Use(middlewares.JWTAuth())
	{
		app.POST("/create_article", v1.CreateArticle)
		//app.POST("/delete_menu", v1.DeleteMenu)
		//app.POST("/find_menu", v1.GetMenuByID)
		//app.POST("/up_menu", v1.UpMenu)
		//app.GET("/get_menu", v1.GetMenus)
		//app.POST("/get_submenu", v1.GetSubMenus)
	}
}
