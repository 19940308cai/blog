package service

import (
	"blog/models/article"
	"blog/models"
	"blog/models/content"
	"github.com/astaxie/beego/logs"
	"blog/warp"
	"blog/models/category"
	"fmt"
	"blog/util"
)

type ArticleService struct {
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (self *ArticleService) InitBlog() (int, string) {
	var articleContentText string
	articleModel := article.NewArticleModel()
	articleEntity, err := articleModel.FindOneCanSetCondition(models.NewConditionStruct("remove", "=", 2), )
	if err == nil {
		contentModel := content.NewContentModal()
		articleContent, err := contentModel.FindOneCanSetCondition(models.NewConditionStruct("article_id", "=", articleEntity.Id))
		if err == nil {
			articleContentText = articleContent.Content
		}
	} else {
		articleEntity = &article.Article{
			Remove: 2,
		}
		err := articleModel.Add(articleEntity)
		if err != nil {
			logs.Debug("初始化博客失败")
			return 0, ""
		}
	}
	return articleEntity.Id, articleContentText
}

func (self *ArticleService) GetArticle(conditionStruct ...*models.ConditionStruct) (warp.ArticleWarp, error) {
	articleModel := article.NewArticleModel()
	article, err := articleModel.FindOneCanSetCondition(conditionStruct...)
	articleWarp := warp.ArticleWarp{}
	if err == nil {
		articleWarp = articleWarp.MakeArticle(*article)
		return articleWarp, nil
	} else {
		return articleWarp, err
	}

}

func (self *ArticleService) GetAllArticles(page, size int) ([]warp.ArticleWarp, util.PageUtil, error) {
	articleModel := article.NewArticleModel()
	articleCount, err := articleModel.GetCountTableCanSetCondition(
		models.NewConditionStruct("Remove", "=", 0),
	)
	pageUtil := util.NewPageUtil(int(articleCount), page, size)
	if err != nil {
		return nil, pageUtil, err
	}

	articlesEntity, err := articleModel.FindListCanSetConditionWithPage(pageUtil,
		models.NewConditionStruct("Remove", "=", 0),
	)
	if err != nil {
		return nil, pageUtil, err
	}

	categoryModel := category.NewCategoryModel()
	categorysEntitys, err := categoryModel.FindAllCanSetCondition(
		models.NewConditionStruct("Remove", "=", 0),
	)
	if err != nil {
		logs.Debug("查询文章,扫描到没有文章分类....")
		return nil, pageUtil, err
	}
	var categoryIdNameMap map[int]string = make(map[int]string)
	for _, category := range *categorysEntitys {
		categoryIdNameMap[category.Id] = category.Name
	}
	articlesWarp := warp.ArticleWarp{}.MakeArticles(articlesEntity)
	for step, article := range articlesWarp {
		categoryName, ok := categoryIdNameMap[article.Category_id.(int)]
		if ok {
			articlesWarp[step].Category_name = categoryName
		} else {
			logs.Debug(fmt.Sprintf("博客ID: %d 没有映射到对应的CategoryName", article.Id.(int)))
			articlesWarp[step].Category_name = "无"
		}
	}
	return articlesWarp, pageUtil, nil
}

func (self *ArticleService) Update(articleId int, params map[string]interface{}) (int64, error) {
	articleModel := article.NewArticleModel()
	articleEm := &article.Article{}
	var updateFields []string
	title, ok := params["title"]
	if ok {
		articleEm.Title = title.(string)
		updateFields = append(updateFields, "Title")
	}
	categoryId, ok := params["categoryId"]
	if ok {
		articleEm.Category_id = categoryId.(int)
		updateFields = append(updateFields, "Category_id")
	}
	remove, ok := params["remove"]
	if ok {
		articleEm.Remove = remove.(int)
		updateFields = append(updateFields, "Remove")
	}
	return articleModel.Update(articleEm, updateFields, models.NewConditionStruct("Id", "=", articleId))
}
