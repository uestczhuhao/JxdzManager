package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	// login := c.GetSession("Login")
	// if login == false {
	// 	c.Redirect("/", 302)
	// }

	v := c.GetSession("authority")
	// beego.Debug(v)
	if v == 1 {
		c.Data["IsSuper"] = true
		c.Data["admin"] = "超级"
	} else if v == 2 {
		c.Data["IsSuper"] = false
		c.Data["admin"] = "普通"
	} else {
		c.Redirect("/", 302)
	}

	c.TplName = "index.html"
}

func (c *IndexController) Post() {
	c.TplName = "index.html"
}
