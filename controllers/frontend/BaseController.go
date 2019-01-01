package frontend

import (
	"github.com/astaxie/beego"
	"strings"
	"blog/controllers"
)

type BaseControoler struct {
	beego.Controller
}

func (self *BaseControoler) view(pageTitle string) {
	controller, _ := self.GetControllerAndAction()
	lowerController := strings.ToLower(controller)
	controllerPrefix := lowerController[:len(lowerController)-len(controllers.Constroller)]
	self.Layout = "frontend/layout.html"
	self.Data["pageTitle"] = pageTitle
	if controllerPrefix == "index" {
		self.Data["bodyClass"] = ""
	}else{
		self.Data["bodyClass"] = "single"
	}
	pathElems := []string{controllers.Frontend, controllers.Hierarchy, controllerPrefix, controllers.Breakpoint, controllers.Html}
	self.TplName = strings.Join(pathElems, "")
}
