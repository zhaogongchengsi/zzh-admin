package service

import (
	"github/gin-react-admin/global"
	"github/gin-react-admin/internal/model"
)

func CreateArticle(article *model.Article) (ar *model.Article, err error)  {
	err = global.DBEngine.Create(&article).Error
	return ar, err
}
