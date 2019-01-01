package open

import (
	"github.com/astaxie/beego"
)

type returnjson struct {
	Code int
	Msg  string
	Data interface{}
}

type BaseControoler struct {
	beego.Controller
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