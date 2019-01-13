package warp

import (
	"blog/models/article"
)

type ArticleWarp struct {
	Id            interface{}
	Title         interface{}
	Category_id   interface{}
	Author_id     interface{}
	Content       interface{}
	Remove        interface{}
	Created_at    interface{}
	Category_name interface{}
}

func (self ArticleWarp) MakeArticle(article article.Article) ArticleWarp {
	self.Id = article.Id
	self.Title = article.Title
	self.Category_id = article.Category_id
	self.Remove = article.Remove
	self.Created_at = article.Created_at
	return self
}

func (self ArticleWarp) MakeArticles(articles *[]article.Article) []ArticleWarp {
	var articleWarps []ArticleWarp
	for _, item := range *articles {
		articleWarps = append(articleWarps, self.MakeArticle(item))
	}
	return articleWarps
}
