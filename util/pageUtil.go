package util

import (
	"fmt"
	"sort"
)

/**
	offset
	limit
	总页数
	总数量
	是否第一页
	是否最后一页
 */

type PageUtil struct {
	CurrentPage int
	PageNo      int
	PageSize    int
	TotalPage   int
	TotalCount  int
	FirstPage   bool
	LastPage    bool
	Navs        []int
}

func NewPageUtil(count int, pageNo int, pageSize int) PageUtil {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}

	pageUtil := PageUtil{
		CurrentPage: pageNo,
		PageNo:     pageSize * (pageNo - 1),
		PageSize:   pageSize,
		TotalPage:  tp,
		TotalCount: count,
		FirstPage:  pageNo == 1,
		LastPage:   pageNo == tp,
	}

	if tp > 0 {
		//预计左边的页数
		_left := 4
		//预计右边的页数
		_right := 5
		//还剩多少页
		diffPage := tp - pageNo
		var nav_num_warp []int

		//不满足左边补充
		if diffPage < _right {
			_left += _right - diffPage
		}
		for i := 0; i <= _right; i++ {
			appendInt := pageNo + i + 1
			if appendInt > tp {
				break
			}
			nav_num_warp = append(nav_num_warp, appendInt)
		}
		fmt.Println(nav_num_warp)

		nav_num_warp = append(nav_num_warp, pageNo)
		fmt.Println(nav_num_warp)

		for i := 0; i <= _left; i++ {
			appendInt := pageNo - i - 1
			if appendInt <= 0 {
				break
			}
			nav_num_warp = append(nav_num_warp, appendInt)
		}
		sort.Ints(nav_num_warp)
		pageUtil.Navs = nav_num_warp
	}
	return pageUtil
}
