package controllers

import (
	"github.com/astaxie/beego"
)

type AddTeacherController struct {
	beego.Controller
}

func (c *AddTeacherController) Get() {
	v := c.GetSession("authority")
	if v == 1 {
		c.Data["IsSuper"] = true
		c.Data["admin"] = "超级"
	} else if v == 2 {
		c.Data["IsSuper"] = false
		c.Data["admin"] = "普通"
	} else {
		c.Redirect("/", 302)
	}
	c.TplName = "addTeacher.html"
}
