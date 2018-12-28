package content

import (
	"blog/models/content"
	"blog/models"
)

type ContentService struct {
}

func NewContentService() *ContentService {
	return &ContentService{}
}

/**
	serice里头尽量操作指针
	返回的时候可以将指针实体包装成warp
 */
func (self *ContentService) GetArticleContent(conditionStruct ...*models.ConditionStruct) (string, error) {
	var articleContent string
	contentModel := content.NewContentModal()
	articleContentEntity, err := contentModel.FindOneCanSetCondition(conditionStruct...)
	if err == nil {
		return articleContentEntity.Content, nil
	}
	return articleContent, err
}

func (self *ContentService) ChangeContent(articleId int, userInput string) bool {
	contentModel := content.NewContentModal()
	condition := models.NewConditionStruct("article_id", "=", articleId)
	articleContent, err := contentModel.FindOneCanSetCondition(condition)
	if err == nil {
		articleContent.Content = userInput
		var updateFields = []string{"Content"}
		_, err = contentModel.Update(articleContent, updateFields, condition)
	} else {
		articleContent = &content.ArticlesContent{
			Content:    userInput,
			Article_id: articleId,
		}
		_, err = contentModel.Add(articleContent)
	}
	if err != nil {
		return false
	} else {
		return true
	}
}
