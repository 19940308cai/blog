package warp

import (
	"blog/models/category"
)

type CategoryWarp struct {
	Id     interface{}
	Name   interface{}
	Remove interface{}
}

func (self CategoryWarp) MakeCategory(category category.Category) CategoryWarp {
	self.Id = category.Id
	self.Name = category.Name
	self.Remove = category.Remove
	return self
}

func (self CategoryWarp) MakeCategorys(categorys *[]category.Category) []CategoryWarp {
	var categoryWarps []CategoryWarp
	for _, item := range *categorys {
		categoryWarps = append(categoryWarps, self.MakeCategory(item))
	}
	return categoryWarps
}
