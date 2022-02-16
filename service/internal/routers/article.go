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
		app.GET("/get_article_list", v1.GetArticleList)
		app.GET("/tags/get_tags", v1.GetTagList)
		app.POST("/tags/create_tag", v1.CreateTag)
	}
}
