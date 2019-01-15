package backend

import (
	"blog/service/content"
)

type ContentController struct {
	BaseControoler
}

func (self *ContentController) Post() {
	if self.IsAjax() == false {
		self.MakeErrorJson(500, "不是合法的请求...")
		return
	} else {
		articleId, err := self.GetInt("article_id")
		if err != nil {
			self.MakeErrorJson(500, "缺少文章ID...")
			return
		}
		if articleId <= 0 {
			self.MakeErrorJson(500, "不是合法的文章ID...")
			return
		}
		markdownContent := self.GetString("content")
		htmlContent := self.GetString("html_content")
		contentService := content.NewContentService()
		status := contentService.ChangeContent(articleId, markdownContent, htmlContent)
		if status == false {
			self.MakeErrorJson(500, "更新文章内容失败...")
			return
		}else{
			self.MakeSuccessJson(200, "更新内容成功...")
			return
		}
	}
}
