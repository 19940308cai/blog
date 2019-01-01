package frontend

type DetailIontroller struct {
	BaseControoler
}

func (self *DetailIontroller) Get() {
	self.view("详情页面")
}
