package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ArticleName string `gorm:"comment:文章名字" json:"article_name"`
	ArticleTitle string `gorm:"comment:文章标题" json:"article_title"`
	ArticleUrl string `gorm:"comment:文章的存储路径" json:"article_url"` // 文章存储在cos 或者本地时的文件路径
	// 文章的存储类型 可能存储在服务器本地（local） 也可能存储在cos等对象存储服务器中（cos） 也可能以字符串的形式存储在数据库中(db)
	ArticleStorageType string `gorm:"comment:文章的存储类型; default:cos" json:"article_storage_type"`
	ArticleTags []ArticleTags `gorm:"comment:文章标签;many2many:article_tag;" json:"article_tags"`
	ArticleAuthor string `gorm:"comment:文章作者" json:"article_author"`
	ArticleContext string `gorm:"type:text;comment:文章内容(作为字符串存储在数据库时)" json:"article_context"`
	ArticleFileName string `gorm:"comment: 文章的文件名字" json:"article_file_name"`
	ArticleType string `gorm:"comment:文章类型" json:"article_type"`
	Likes int `gorm:"comment:点赞数" json:"likes"`
	NumberViews int `gorm:"comment:观看数" json:"number_views"`
	UUID  uuid.UUID `json:"uuid" gorm:"comment:创建文章用户的唯一Id" json:"uuid"` // 用户UUID
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