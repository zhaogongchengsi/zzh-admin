package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github/gin-react-admin/docs"
	v1 "github/gin-react-admin/internal/api/v1"
	"github/gin-react-admin/internal/middlewares"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors()) // 配置请求跨域
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/AutoMigrate", v1.AutoMigrated) // ！！！！数据库迁移接口 开发环境记得删除
	r.GET("/api/v1/verify_code", v1.CreateVerificationCode)
	Static(r)
	Menu(r)
	Role(r)
	User(r)
	CosRouter(r)
	NewArticle(r)
	return r
}
