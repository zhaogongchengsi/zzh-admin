package service

import (
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
	"github/gin-react-admin/pkg/request"
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
	//var art []model.ArticleTags
	//err = global.DBEngine.Find(&art).Error
	//article.ArticleTags = art
	err = global.DBEngine.Create(&article).Error
	return article, err
}

// 分页获取文章
func GetArticleList(lo request.LimitOffset) (arts []model.Article, num int64, err error) {
	var art []model.Article
	var number int64
	err = global.DBEngine.Preload("ArticleTags").Offset(lo.Offset).Limit(lo.Limit).Find(&art).Offset(-1).Limit(-1).Count(&number).Error
	return art,number, err
}

func CreateTag(tag *model.ArticleTags) (ar *model.ArticleTags, err error)  {
	err = global.DBEngine.Create(&tag).Error
	return tag, err
}

func GetTagLis(lo request.LimitOffset) (arts []model.ArticleTags, num int64, err error) {
	var art []model.ArticleTags
	var number int64
	err = global.DBEngine.Offset(lo.Offset).Limit(lo.Limit).Find(&art).Offset(-1).Limit(-1).Count(&number).Error
	return art,number, err
}