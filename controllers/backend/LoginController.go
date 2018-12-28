package backend

import (
	"blog/service/login"
)

type LoginController struct {
	BaseControoler
}

func (self *LoginController) Get() {
	self.view("登录")
}

func (self *LoginController) Post() {
	username := self.GetString("username")
	password := self.GetString("password")
	if username == "" {
		self.MakeErrorJson(406, "请输入用户名")
	}
	if password == "" {
		self.MakeErrorJson(406, "请输入密码")
	}
	loginService := new(login.LoginService)
	status := loginService.LoginProcess(username, password)
	if status == false {
		self.MakeErrorJson(404, "账号密码错误")
	} else {
		self.SetSession("user", 1)
		self.MakeSuccessJson(200, "欢迎回来,即将进入后台!")
	}
}
