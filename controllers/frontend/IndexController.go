package frontend

type IndexController struct {
	BaseControoler
}

func (self *IndexController) Get() {
	self.view("蔡江的博客")
}
