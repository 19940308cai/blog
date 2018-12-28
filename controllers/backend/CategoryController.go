package backend

import (
	"blog/service/category"
)

type CategoryController struct {
	BaseControoler
}

func (self *CategoryController) Get() {
	categoryService := category.CategoryService{}
	categorys := categoryService.GetAllCategory()
	if self.IsAjax() {
		self.MakeSuccessJson(200, categorys)
	} else {
		self.Data["Categorys"] = categorys
		self.view("类别管理")
	}
}

func (self *CategoryController) Post() {
	id, _ := self.GetInt("id")
	if id > 0 {
		params := make(map[string]interface{})
		remove, _ := self.GetInt("remove")
		if remove >= 0 {
			params["remove"] = remove
		}
		name := self.GetString("name")
		if name != "" {
			params["name"] = name
		}
		self._update(id, params)
	} else {
		categoryName := self.GetString("name")
		if categoryName == "" {
			self.MakeErrorJson(406, "不能输入空字符")
		}
		self._add(categoryName)
	}
}

func (self *CategoryController) _update(id int, params map[string]interface{}) {
	categoryService := category.CategoryService{}
	status := categoryService.UpdateCategory(id, params)
	if status == true {
		self.MakeSuccessJson(200, "删除类别成功...")
	} else {
		self.MakeErrorJson(500, "删除类别失败...")
	}

}

func (self *CategoryController) _add(categoryName string) {
	categoryService := category.CategoryService{}
	_, err := categoryService.AddCategory(categoryName)
	if err != nil {
		self.MakeErrorJson(500, "添加文章类别失败")
	} else {
		self.MakeSuccessJson(200, "添加文章类别成功")
	}
}
