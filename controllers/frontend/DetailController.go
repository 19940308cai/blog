package frontend

import (
	"blog/service/article"
	"blog/models"
	"blog/service/content"
)

type DetailIontroller struct {
	BaseControoler
}

func (self *DetailIontroller) Get() {
	articleId, err := self.GetInt("article_id")
	if err == nil {
		articleService := service.NewArticleService()
		articleWarp, err := articleService.GetArticle(
			models.NewConditionStruct("id", "=", articleId),
		)
		if err == nil {
			self.Data["category_id"] = articleWarp.Category_id
			self.Data["title"] = articleWarp.Title
			self.Data["created_at"] = articleWarp.Created_at

		}
		content, _ := content.NewContentService().GetArticleContent(
			models.NewConditionStruct("article_id", "=", articleId),
		)
		self.Data["content"] = content
		self.view(articleWarp.Title.(string))
	}
}
