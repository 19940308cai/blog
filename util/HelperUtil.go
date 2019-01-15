package util

import (
	"github.com/astaxie/beego/context"
	"strconv"
	"fmt"
)

type helper struct {
}

func NewHelper() (*helper) {
	return &helper{}
}

func (self *helper) MakeUriAndSearchMap(uriTmp string, request *context.BeegoInput) (string, []map[string]string) {
	var elements []map[string]string
	var elemt map[string]string = make(map[string]string)
	key := request.Query("key")
	if len(key) > 0 {
		elemt["key"] = key
	}
	value := request.Query("value")
	fmt.Println(value)
	if len(value) > 0 {
		elemt["value"] = value
	}
	elements = append(elements, elemt)
	if _, ok := elements[0]["key"]; ok == true {
		for _, elemt := range elements {
			uriTmp += "&key=" + elemt["key"] + "&value=" + elemt["value"]
		}
	}
	return uriTmp, elements
}

func (self *helper) MakePageAndSize(request *context.BeegoInput) (int, int) {
	const defaultPage = 1
	const defualtSize = 10
	page, err := strconv.Atoi(request.Query("page"))
	if err != nil {
		page = defaultPage
	}
	size, err := strconv.Atoi(request.Query("size"))
	if err != nil {
		size = defualtSize
	}
	return page, size
}
