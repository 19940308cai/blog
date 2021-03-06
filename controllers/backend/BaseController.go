package backend

import (
	"github.com/astaxie/beego"
	"blog/controllers"
	"strings"
)

type returnjson struct {
	Code int
	Msg  string
	Data interface{}
}

type BaseControoler struct {
	beego.Controller
}

var needSeesion bool = false

var UserId int

func (self *BaseControoler) GetPageAndSize() (int, int) {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	if page <= 0 {
		page = 1
	}
	size, err := self.GetInt("size")
	if err != nil {
		size = 10
	}
	return page, size
}

func (self *BaseControoler) Prepare() {
	if strings.Index(self.Ctx.Request.RequestURI, "login") != -1 {
		needSeesion = false
	} else {
		needSeesion = true
		session := self.GetSession("user")
		if session == nil {
			self.Ctx.Redirect(302, "/backend/login")
		} else {
			UserId, _ = session.(int)
		}
	}
}

func (self *BaseControoler) MakeSuccessJson(code int, data interface{}) {
	self.Data["json"] = &returnjson{
		code,
		"success",
		data,
	}
	self.ServeJSON()
}

func (self *BaseControoler) MakeErrorJson(code int, data interface{}) {
	self.Data["json"] = &returnjson{
		code,
		"error",
		data,
	}
	self.ServeJSON()
}

func (self *BaseControoler) view(pageTitle string) {
	controller, _ := self.GetControllerAndAction()
	lowerController := strings.ToLower(controller)

	controllerPrefix := lowerController[:len(lowerController)-len(controllers.Constroller)]
	if controllerPrefix == "login" {
		self.Layout = "backend/login_layout.html"
	} else {
		self.Layout = "backend/layout.html"
	}
	self.Data["pageTitle"] = "博客后台--" + pageTitle

	pathElems := []string{controllers.Backend, controllers.Hierarchy, controllerPrefix, controllers.Breakpoint, controllers.Html}
	self.TplName = strings.Join(pathElems, "")
}
