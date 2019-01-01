package routers

import (
	"blog/controllers/backend"
	"github.com/astaxie/beego"
	"blog/controllers/frontend"
	"blog/controllers/open"
)

func init() {


	/**
		应用层
	 */
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


	//前端
	beego.Router("/frontend/index", &frontend.IndexController{})
	beego.Router("/frontend/detail", &frontend.DetailIontroller{})


	/**
		服务层
	 */
	beego.Router("/open/categorys", &open.CategoryController{})

}
