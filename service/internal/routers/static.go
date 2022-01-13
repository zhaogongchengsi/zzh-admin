package routers

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/api/fileService"
	"github/gin-react-admin/internal/middlewares"
)

func Static (g *gin.Engine) {
	var static string = global.FileService.StaticPath
	g.Static(static, static)
	NewFileService(g)
	//g.StaticFS("/more_static", http.Dir("my_file_system"))
	//g.StaticFile("/favicon.ico", "./resources/favicon.ico")
}


func NewFileService (g *gin.Engine) {

	app := g.Group("/api/v1/upload_download").Use(middlewares.JWTAuth())
	{
		app.POST("/upload", fileService.UpLoad)
	}
}