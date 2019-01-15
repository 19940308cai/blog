package content

import (
	"time"
	"github.com/astaxie/beego/orm"
	"blog/models"
)

type ArticlesContent struct {
	Id           int
	Article_id   int
	Content      string
	Html_content string
	Created_at   time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModelWithPrefix("blog_", new(ArticlesContent))
}

type contentModal struct {
	*ArticlesContent
	*models.BaseModal
}

func NewContentModal() *contentModal {
	return &contentModal{
		&ArticlesContent{},
		&models.BaseModal{"blog_articles_content"},
	}
}

func (self *contentModal) FindOneCanSetCondition(conditionStructs ...*models.ConditionStruct) (*ArticlesContent, error) {
	articleContentEntity := &ArticlesContent{}
	err := self.AppendConditionFilter(conditionStructs...).One(articleContentEntity)
	return articleContentEntity, err
}

func (self *contentModal) FindAllCanSetCondition(conditionStructs ...*models.ConditionStruct) (*[]ArticlesContent, error) {
	var articlesContents []ArticlesContent
	_, err := self.AppendConditionFilter(conditionStructs...).All(&articlesContents)
	return &articlesContents, err
}

func (self *contentModal) Add(articleContentEntity *ArticlesContent) (int64, error) {
	return self.InsertDataToTable(articleContentEntity)
}

func (self *contentModal) Update(contentEntity *ArticlesContent, updateFields []string, conditionStructs ...*models.ConditionStruct) (int64, error) {
	params, err := self.MakeOrmParamsByStruct(contentEntity, updateFields...)
	if err != nil {
		return 0, err
	}
	return self.UpdateDataInTableCanSetCondition(params, conditionStructs...)
}
