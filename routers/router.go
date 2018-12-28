package routers

import (
	"blog/controllers/backend"
	"github.com/astaxie/beego"
)

func init() {
	//登录
	beego.Router("/login", &backend.LoginController{})
	beego.Router("/dologin", &backend.LoginController{})

	//首页
	beego.Router("/index", &backend.IndexController{})
	beego.Router("/", &backend.IndexController{}, "get:Get")

	//类别
	beego.Router("/category", &backend.CategoryController{})
	beego.Router("/docategory", &backend.CategoryController{})

	//博客(复数)
	beego.Router("/articles", &backend.ArticlesController{})
	beego.Router("/doarticles", &backend.ArticlesController{})

	//博客(非复数)
	beego.Router("/article", &backend.ArticleController{})
	beego.Router("/doarticle", &backend.ArticleController{})

	//博客正文
	beego.Router("/content", &backend.ContentController{})

}
