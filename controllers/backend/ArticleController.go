package backend

import (
	"blog/service/article"
	"blog/models"
	"blog/service/content"
	"blog/util"
)

//博客列表
type ArticlesController struct {
	BaseControoler
}

func (self *ArticlesController) Get() {
	articleService := service.NewArticleService()
	var uriFormatModal string = "/backend/articles?page=%d&size=10"
	helper := util.NewHelper()
	uriFormatModal, searchMap := helper.MakeUriAndSearchMap(uriFormatModal, self.Ctx.Input)
	page, size := helper.MakePageAndSize(self.Ctx.Input)
	articles, pageUtil, err := articleService.GetAllArticles(page, size, searchMap...)
	if err == nil {
		self.Data["uriFormatModal"] = uriFormatModal
		self.Data["page"] = pageUtil
		self.Data["articles"] = articles
	}
	self.view("博客列表")
}

func (self *ArticlesController) Post() {
	article_id, err := self.GetInt("article_id")
	if err != nil {
		self.MakeErrorJson(500, "article_id不是合法的参数...")
		return
	}
	remove, err := self.GetInt("remove")
	if err != nil {
		self.MakeErrorJson(500, "remove不是合法的参数...")
		return
	}
	articleService := service.NewArticleService()
	var params map[string]interface{} = make(map[string]interface{})
	params["remove"] = remove
	articleService.Update(article_id, params)
	self.MakeSuccessJson(200, "移除成功...")
	return
}

//博客
type ArticleController struct {
	BaseControoler
}

func (self *ArticleController) Get() {
	var pageTitle string
	articleService := service.NewArticleService()
	if self.GetString("action") != "update" {
		articleId, content := articleService.InitBlog()
		self.Data["articleId"] = articleId
		self.Data["content"] = content
		pageTitle = "新建博客"
	} else {
		articleId, err := self.GetInt("article_id")
		if err == nil {
			articleWarp, err := articleService.GetArticle(
				models.NewConditionStruct("id", "=", articleId),
			)
			if err == nil {
				self.Data["category_id"] = articleWarp.Category_id
				self.Data["title"] = articleWarp.Title
			}
			content, _ := content.NewContentService().GetArticleContent(
				models.NewConditionStruct("article_id", "=", articleId),
			)
			self.Data["content"] = content
		}
		self.Data["articleId"] = articleId
		pageTitle = "修改博客"
	}
	self.view(pageTitle)
}

func (self *ArticleController) Post() {
	if self.IsAjax() == false {
		self.MakeErrorJson(500, "请求不是合法的")
	} else {
		articleId, _ := self.GetInt("article_id")
		if articleId <= 0 {
			self.MakeErrorJson(500, "参数不是合法...")
			return
		}
		title := self.GetString("title")
		if title == "" {
			self.MakeErrorJson(500, "缺少博客名称...")
			return
		}
		describe := self.GetString("describe")
		if describe == "" {
			self.MakeErrorJson(500,"缺少博客描述...")
			return
		}
		categoryId, _ := self.GetInt("category_id");
		if categoryId <= 0 {
			self.MakeErrorJson(500, "参数不合法....")
			return
		}
		remove, _ := self.GetInt("remove")
		articleService := service.NewArticleService()
		var params map[string]interface{} = make(map[string]interface{})
		params["title"] = title
		params["categoryId"] = categoryId
		params["remove"] = remove
		params["describe"] = describe
		_, err := articleService.Update(articleId, params)
		if err != nil {
			self.MakeErrorJson(500, "发布博客失败...")
			return
		}
		self.MakeSuccessJson(200, "发布成功...")
	}
}
