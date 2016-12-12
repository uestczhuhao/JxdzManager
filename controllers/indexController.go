package controllers

import (
	"JxdzManager/models"

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
	l := c.GetSession("Login")
	if l != true {
		c.Redirect("/", 302)
	}
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
	c.Data["Html"] = "<em><strong>1234</strong></em>"

	c.TplName = "index.html"
	// beego.Debug(models.StandardOut())
	// cates := models.StandardOut()
	// beego.Debug(len(cates))
	// beego.Debug(cates)
	c.Data["CateName"], c.Data["Categories"] = models.StandardOut()
}

func (c *IndexController) Post() {

	beego.Debug("this is IndexController Post()")
	a := c.Input().Get("content")

	beego.Debug(a)
	c.Redirect("/index", 302)
}
