package Front

import "github.com/astaxie/beego"

type JxdzN1Controller struct {
	beego.Controller
}

func (c *JxdzN1Controller) Get() {
	c.TplName = "Front/jxdzN1.html"
}
