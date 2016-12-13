package controllers

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type AticleListController struct {
	beego.Controller
}

func (c *AticleListController) Get() {
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

	c.Data["articles"] = models.GetAllContent()
	beego.Debug(c.Data["articles"])
	c.TplName = "articlelist.html"
}
