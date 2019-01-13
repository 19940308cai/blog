package routers

import (
	"blog/controllers/backend"
	"github.com/astaxie/beego"
	"blog/controllers/frontend"
	"blog/controllers/open"
)

func init() {


	/**
		应用层-前台
	 */
	beego.Router("/backend/login", &backend.LoginController{})
	beego.Router("/backend/dologin", &backend.LoginController{})

	//首页
	beego.Router("/backend/index", &backend.IndexController{})

	//类别
	beego.Router("/backend/category", &backend.CategoryController{})
	beego.Router("/backend/docategory", &backend.CategoryController{})

	//博客(复数)
	beego.Router("/backend/articles", &backend.ArticlesController{})
	beego.Router("/backend/doarticles", &backend.ArticlesController{})

	//博客(非复数)
	beego.Router("/backend/article", &backend.ArticleController{})
	beego.Router("/backend/doarticle", &backend.ArticleController{})

	//博客正文
	beego.Router("/backend/content", &backend.ContentController{})


	/**
		应用层-后台
	 */
	beego.Router("/", &frontend.IndexController{}, "get:Get")
	beego.Router("/frontend/index", &frontend.IndexController{})
	beego.Router("/frontend/detail", &frontend.DetailIontroller{})


	/**
		服务层
	 */
	beego.Router("/open/categorys", &open.CategoryController{})
	beego.Router("/open/articles", &open.ArticlesController{})

}
