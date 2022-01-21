package request

type Article struct {
	FileName string `json:"fileName"` // 文件名字
	ArticleName string `json:"articleName"` // 文章名字
	ArticleTitle string `json:"articleTitle"` // 文章标题
	ArticleUrl string `json:"articleUrl"` // 文章存储的Url 若文章是存储在 数据库 则为空
	ArticleStorageType string `json:"articleStorageType"` // 文章的存储类型
	ArticleAuthor string `json:"articleAuthor"` // 文章作者
	ArticleType string `json:"articleType"` // 文章的类型
	ArticleContext string `json:"articleContext"` // 文章的内容主题 若文章不是存储在数据库 则本字段为空
	//ArticleTags []model.ArticleTags `json:"articleTags"`
}
