package article

import (
	"time"
	"github.com/astaxie/beego/orm"
	"blog/models"
	"blog/util"
)

type Article struct {
	Id          int
	Title       string
	Category_id int
	Author_id   int
	Remove      int
	Created_at  time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModelWithPrefix("blog_", new(Article))
}

type articleModal struct {
	*Article
	*models.BaseModal
}

func NewArticleModel() *articleModal {
	return &articleModal{
		&Article{},
		&models.BaseModal{"blog_article"},
	}
}

func (self *articleModal) FindOneCanSetCondition(conditionStructs ...*models.ConditionStruct) (*Article, error) {
	articleEntity := &Article{}
	err := self.AppendConditionFilter(conditionStructs...).One(articleEntity)
	return articleEntity, err
}

func (self *articleModal) FindAllCanSetCondition(conditionStructs ...*models.ConditionStruct) (*[]Article, error) {
	var articles []Article
	_, err := self.AppendConditionFilter(conditionStructs...).All(&articles)
	return &articles, err
}

func (self *articleModal) FindListCanSetConditionWithPage(pageUtil util.PageUtil, conditionStructs ...*models.ConditionStruct) (*[]Article, error) {
	var articles []Article
	_, err := self.GetDataCanSetConditionAndLimit(pageUtil, conditionStructs...).All(&articles)
	return &articles, err
}

func (self *articleModal) Add(article *Article) (error) {
	status, err := self.InsertDataToTable(article)
	if status == 1 {
		return nil
	}
	return err
}

func (self *articleModal) Update(articleEntity *Article, updateFields []string, conditionStructs ...*models.ConditionStruct) (int64, error) {
	params, err := self.MakeOrmParamsByStruct(articleEntity, updateFields...)
	if err != nil {
		return 0, err
	}
	return self.UpdateDataInTableCanSetCondition(params, conditionStructs...)
}
