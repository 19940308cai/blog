package open

import "blog/service/category"

type CategoryController struct {
	BaseControoler
}


func (self *CategoryController) Get() {
	categoryService := category.CategoryService{}
	categorys := categoryService.GetAllCategory()
	if self.IsAjax() {
		self.MakeSuccessJson(200, categorys)
	} else{
		self.MakeErrorJson(500,nil)
	}
}