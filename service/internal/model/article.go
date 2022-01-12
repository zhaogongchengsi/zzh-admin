package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ArticleName string `gorm:"comment:文章名字"`
	ArticleTitle string `gorm:"comment:文章标题"`
	ArticleUrl string `gorm:"comment:文章的存储路径"` // 文章存储在cos 或者本地时的文件路径

	// 文章的存储类型 可能存储在服务器本地（local） 也可能存储在cos等对象存储服务器中（cos） 也可能以字符串的形式存储在数据库中(db)
	ArticleStorageType string `gorm:"comment:文章的存储类型; default:cos"`

	ArticleTags []ArticleTags `gorm:"comment:文章标签;many2many:article_tag;"`
	ArticleAuthor string `gorm:"comment:文章作者"`
	ArticleContext string `gorm:"type:text;comment:文章内容(作为字符串存储在数据库时)"`
	ArticleFileName string `gorm:"comment: 文章的文件名字"`
	ArticleType string `gorm:"comment:文章类型"`
}

type ArticleTags struct {
	gorm.Model
	Tag string `gorm:"tag_name; comment: 标签名字"`
	Details string `gorm:"details; comment: 标签详情"`
}
//
//type ArticleStorageType struct {
//	ArticleStorage string `gorm:"article_storage; comment:文章存储类型; default: cos"`
//	ArticleType string `gorm:"article_type; comment:文章的类型"`
//}