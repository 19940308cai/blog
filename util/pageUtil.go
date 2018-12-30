package util

import (
	"sort"
	"fmt"
)

const _default_left = 4
const _default_right = 5

/**
	offset
	limit
	总页数
	总数量
	是否第一页
	是否最后一页
 */

type PageUtil struct {
	CurrentPage   int
	PageNo        int
	PageSize      int
	TotalPage     int
	TotalCount    int
	LeftNextPage  int
	RightNextPage int
	FirstPage     bool
	LastPage      bool
	Navs          []int
}

func NewPageUtil(count int, pageNo int, pageSize int) PageUtil {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}

	pageUtil := PageUtil{
		CurrentPage: pageNo,
		PageNo:      pageSize * (pageNo - 1),
		PageSize:    pageSize,
		TotalPage:   tp,
		TotalCount:  count,
		FirstPage:   pageNo == 1,
		LastPage:    pageNo == tp,
	}

	if tp > 0 {
		//左/右两边的导航的数量
		var _left_tmp int
		var _right_tmp int
		//数据容器
		var nav_num_warp []int
		//计算左右两边的下一次值
		var _next_left int
		var _next_right int
		/**
			计算原则:
				以当前页面为中心，向左右两边扩展，超出的部分跳过
				左: 极值是1
				右: 极值是PageSize
		 */
		for i := 1; i <= _default_right; i++ {
			_right_tmp = pageNo + i
			if _right_tmp > tp {
				continue
			} else {
				nav_num_warp = append(nav_num_warp, _right_tmp)
			}

			if i <= _default_left {
				_left_tmp = pageNo - i
				if _left_tmp <= 0 {
					continue
				}
				nav_num_warp = append(nav_num_warp, _left_tmp)
			}
		}
		//追加当前页
		nav_num_warp = append(nav_num_warp, pageNo)
		sort.Ints(nav_num_warp)
		fmt.Println(nav_num_warp)
		if len(nav_num_warp) < pageSize {
			//需要补充的数据
			var diff_page int = pageSize - len(nav_num_warp)
			//根据左右两边最大最小值来区分方向
			if nav_num_warp[0] == 1 {
				for i := 1; i <= diff_page; i++ {
					right := nav_num_warp[len(nav_num_warp)-1] + 1
					if right >= pageUtil.TotalPage {
						break
					}
					nav_num_warp = append(nav_num_warp, nav_num_warp[len(nav_num_warp)-1]+1)
				}
			} else {
				for i := 1; i <= diff_page; i++ {
					nav_num_warp = append(nav_num_warp, nav_num_warp[0]-i)
				}
			}
		}
		sort.Ints(nav_num_warp)
		pageUtil.Navs = nav_num_warp
		_next_left = pageUtil.CurrentPage - 1
		_next_right = pageUtil.CurrentPage + 1
		if _next_left <= 0 {
			_next_left = pageUtil.CurrentPage
		}
		if _next_right > tp {
			_next_right = pageUtil.CurrentPage
		}
		pageUtil.LeftNextPage = _next_left
		pageUtil.RightNextPage = _next_right
	}
	return pageUtil
}
