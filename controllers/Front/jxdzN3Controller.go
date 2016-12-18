package Front

import "github.com/astaxie/beego"

type JxdzN3Controller struct {
	beego.Controller
}

func (c *JxdzN3Controller) Get() {
	c.TplName = "Front/jxdzN3.html"
}
