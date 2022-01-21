package service

import (
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
)

// CreateArticle 创建文章
func CreateArticle(article *model.Article) (ar *model.Article, err error)  {
	if article.ArticleStorageType == "services" {
		file, fer := SaveString(article.ArticleContext, article.ArticleFileName, false)
		if fer != nil {
			return article, fer
		}
		article.ArticleContext = " "
		article.ArticleUrl = file.SavaPath
		article.ArticleFileName = file.FileName
	}
	err = global.DBEngine.Create(&article).Error
	return article, err
}
