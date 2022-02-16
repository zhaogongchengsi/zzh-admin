package v1

import (
	"github.com/gin-gonic/gin"
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/internal/service"
	"github/gin-react-admin/pkg/errcode"
	"github/gin-react-admin/pkg/request"
	"github/gin-react-admin/pkg/response"
	"github/gin-react-admin/utils"
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
	value, _ := utils.GetHeaderUser(c)
 	ar := model.Article {
			ArticleFileName: newArticle.FileName,
			ArticleContext: newArticle.ArticleContext,
			ArticleAuthor: value.NickName,
			ArticleName: newArticle.FileName,
			ArticleStorageType: newArticle.ArticleStorageType,
			ArticleType: newArticle.ArticleType,
			ArticleTitle: newArticle.ArticleTitle,
			ArticleUrl: newArticle.ArticleUrl,
			UUID: value.UUID,
			ArticleDesc: newArticle.ArticleDesc,
	}
	art, er := service.CreateArticle(&ar, newArticle.ArticleTags)
	if er != nil {
		res := response.Response{
			Err: er,
			State: response.State{Message: "文章创建失败", Code: errcode.CreateError},
		}
		res.SendError(c)
		return
	}
	resOk := response.Response{Data: *art}
	resOk.Send(c)
}


// @Tags Article
// @Summary 获取文章列表
// @Produce  application/json
// @Success 200 {string} string "{state:{code:0, msg:"成功"},"data":{ }}"
// @Router /api/v1/article/get_article_list [get]
func GetArticleList(c *gin.Context) {
	lo := request.GetLimitOffset(c)
	artist, con,  err := service.GetArticleList(*lo)
	if err != nil {
		res := response.Response{
			Err: err,
			State: response.State{Message: "文章获取失败", Code: errcode.FindError},
		}
		res.SendError(c)
		return
	}
	resOk := response.Response{Data:map[string]interface{}{
		"article_list":artist,
		"limit_offset": lo,
		"count":con,
	}}
	resOk.Send(c)
}


// @Tags Article
// @Summary 创建标签
// @Produce  application/json
// @Success 200 {string} string "{state:{code:0, msg:"成功"},"data":{ }}"
// @Router /api/v1/article/tags/create_tag [post]
func CreateTag (c *gin.Context)  {
	var newtag request.Tag
	if err := c.ShouldBindJSON(&newtag); err != nil {
		res := response.Response{Err: err}
		res.SendInputParameterError(c)
		return
	}

	ar := model.ArticleTags {
		Tag: newtag.Tag,
		TagColor: newtag.TagColor,
		TagDesc: newtag.TagDesc,
	}

	tag, er := service.CreateTag(&ar)
	if er != nil {
		res := response.Response{
			Err: er,
			State: response.State{Message: "标签创建失败", Code: errcode.CreateError},
		}
		res.SendError(c)
		return
	}
	resOk := response.Response{Data: *tag}
	resOk.Send(c)
}



func GetTagList (c *gin.Context) {
	lo := request.GetLimitOffset(c)
	tags, con , err := service.GetTagLis(*lo)
	if err != nil {
		res := response.Response{
			Err: err,
			State: response.State{Message: "标签获取失败", Code: errcode.FindError},
		}
		res.SendError(c)
		return
	}
	resOk := response.Response{Data:map[string]interface{}{
		"tag_list":tags,
		"limit_offset": lo,
		"count":con,
	}}
	resOk.Send(c)
}