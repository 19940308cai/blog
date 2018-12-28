package category

import (
	"github.com/astaxie/beego/orm"
	"blog/models"
)

type Category struct {
	Id     int
	Name   string
	Remove int
}

func init() {
	orm.RegisterModelWithPrefix("blog_", new(Category))
}

type categoryModal struct {
	*Category
	*models.BaseModal
}

func NewCategoryModel() *categoryModal {
	return &categoryModal{
		&Category{},
		&models.BaseModal{"blog_category"},
	}
}

func (self *categoryModal) FindAllCanSetCondition(conditionStructs ...*models.ConditionStruct) (*[]Category, error) {
	var categorys []Category
	_, err := self.AppendConditionFilter(conditionStructs...).All(&categorys)
	return &categorys, err
}

func (self *categoryModal) FindOneCanSetCondition(conditionStructs ...*models.ConditionStruct) (*Category, error) {
	var category Category
	err := self.AppendConditionFilter(conditionStructs...).One(&category)
	return &category, err
}

func (self *categoryModal) Add(category *Category) (int64, error) {
	return self.InsertDataToTable(category)
}

func (self *categoryModal) Update(categoryEntity *Category, updateFields []string, conditionStructs ...*models.ConditionStruct) (int64, error) {
	params, err := self.MakeOrmParamsByStruct(categoryEntity, updateFields...)
	if err != nil {
		return 0, err
	}
	return self.UpdateDataInTableCanSetCondition(params, conditionStructs...)
}
