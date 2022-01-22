package routers

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/api/v1"
	"github/gin-react-admin/internal/middlewares"
)

func User(g *gin.Engine)  {
	app := g.Group("/api/v1/user").Use(middlewares.JWTAuth())
	{
		app.POST("/login", v1.Login)
		app.POST("/register", v1.Register)
		app.POST("/change_password", v1.ChangePassword)
		app.POST("/users", v1.GetUserList)
		app.GET("/userinfo", v1.GetUserInfo)
	}
}