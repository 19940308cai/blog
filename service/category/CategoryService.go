package category

import (
	"blog/models/category"
	"blog/warp"
	"blog/models"
)

type CategoryService struct {
}

func (self *CategoryService) GetAllCategory() []warp.CategoryWarp {
	categoryModal := category.NewCategoryModel()
	categorys, err := categoryModal.FindAllCanSetCondition(
		models.NewConditionStruct("remove", "=", 0),
	)
	if err != nil {
		return nil
	}
	categoryWarp := warp.CategoryWarp{}
	categoryWarps := categoryWarp.MakeCategorys(categorys)
	return categoryWarps
}

func (self *CategoryService) AddCategory(categoryName string) (int64, error) {
	categoryModal := category.NewCategoryModel()
	categoryEntity := &category.Category{
		Name: categoryName,
	}
	return categoryModal.Add(categoryEntity)
}

func (self *CategoryService) UpdateCategory(id int, params map[string]interface{}) bool {
	//查找修改的是否存在
	categoryModal := category.NewCategoryModel()
	category, err := categoryModal.FindOneCanSetCondition(
		models.NewConditionStruct("id", "=", id),
	)
	if err == nil {
		return false
	}
	value, ok := params["remove"]
	if ok {
		remove, ok := value.(int)
		if ok {
			category.Remove = remove
		}
	}
	/**
		生成修改字段以 字符串数组存储
	 */
	var updateFields []string
	for key, _ := range params {
		updateFields = append(updateFields, key)
	}
	_, err = categoryModal.Update(category, updateFields,
		models.NewConditionStruct("Id", "=", id),
	)
	if err == nil {
		return true
	} else {
		return false
	}
}
