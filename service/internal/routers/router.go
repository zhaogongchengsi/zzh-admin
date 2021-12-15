package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github/gin-react-admin/docs"
	"github/gin-react-admin/internal/middlewares"
	v1 "github/gin-react-admin/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors()) // 配置请求跨域
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/AutoMigrate", v1.AutoMigrated) // ！！！！数据库迁移接口 开发环境记得删除

	app1 := r.Group("/api/v1").Use(middlewares.JWTAuth())
	{
		app1.POST("/menu/create_menu", v1.CreateMenu)
		app1.POST("/menu/delete_menu", v1.DeleteMenu)
		app1.GET("/menu/get_menu", v1.GetMenus)
		app1.GET("/verify_code", v1.CreateVerificationCode)
		app1.POST("/user/login", v1.Login)
		app1.POST("/user/register", v1.Register)
		app1.POST("/user/change_password", v1.ChangePassword)
		app1.POST("/user/users", v1.GetUserList)
		app1.POST("/role/create_role", v1.CreateRole)
		app1.GET("/role/get_roles", v1.GetAuthLIst)
	}

	return r
}
