package open

import (
	"blog/service/article"
	"blog/util"
	"fmt"
)

type ArticlesController struct {
	BaseControoler
}

func (self *ArticlesController) Get() {
	articleService := service.NewArticleService()
	var uriFormatModal string = "/frontend/index?page=%d&size=10"
	helper := util.NewHelper()
	uriFormatModal, searchMap := helper.MakeUriAndSearchMap(uriFormatModal, self.Ctx.Input)
	page, size := helper.MakePageAndSize(self.Ctx.Input)
	articles, pageUtil, err := articleService.GetAllArticles(page, size, searchMap...)
	var data map[string]interface{} = make(map[string]interface{})
	if self.IsAjax() {
		if err == nil {
			data["uriFormatModal"] = fmt.Sprintf(uriFormatModal, pageUtil.RightNextPage)
			data["articles"] = articles
			self.MakeSuccessJson(200, data)
			return
		} else {
			self.MakeErrorJson(500, nil)
		}
	}else{
		self.MakeErrorJson(500, nil)
	}

}
