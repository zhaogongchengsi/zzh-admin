package v1

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/pkg/request"
	"github/gin-react-admin/pkg/response"
)


// @Tags Article
// @Summary 创建文章
// @Produce  application/json
// @Success 200 {string} string "{state:{code:0, msg:"成功"},"data":{ }}"
// @Router /api/v1/article/create_article [post]
func CreateArticle(c *gin.Context)  {
	var newArticle request.Article
	if err := c.ShouldBindJSON(&newArticle); err != nil {
		res := response.Response{Err: err}
		res.SendInputParameterError(c)
		return
	}
	ar := model.Article {
			ArticleFileName: newArticle.FileName,
			ArticleContext: newArticle.ArticleContext,
			ArticleAuthor: newArticle.ArticleAuthor,
			ArticleName: newArticle.ArticleName,
			ArticleStorageType: newArticle.ArticleStorageType,
			ArticleType: newArticle.ArticleType,
			ArticleTags: newArticle.ArticleTags,
	}


}