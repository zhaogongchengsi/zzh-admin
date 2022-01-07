package routers

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/api/v1"
	"github/gin-react-admin/internal/middlewares"
)

func User(g *gin.Engine)  {
	app := g.Group("/api/v1").Use(middlewares.JWTAuth())
	{
		app.POST("/user/login", v1.Login)
		app.POST("/user/register", v1.Register)
		app.POST("/user/change_password", v1.ChangePassword)
		app.POST("/user/users", v1.GetUserList)
	}
}