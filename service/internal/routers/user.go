package routers

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/api/v1"
	"github/gin-react-admin/internal/middlewares"
)

func User(g *gin.Engine)  {
	app := g.Group("/api/v1/user")
	{
		app.POST("/login", v1.Login)
		app.POST("/register", v1.Register)
		app.POST("/change_password", v1.ChangePassword).Use(middlewares.JWTAuth())
		app.POST("/users", v1.GetUserList).Use(middlewares.JWTAuth())
		app.GET("/userinfo", v1.GetUserInfo).Use(middlewares.JWTAuth())
	}
}