package backend

type IndexController struct {
	BaseControoler
}

func (self *IndexController) Get() {
	self.view("首页")
}
