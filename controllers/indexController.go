package controllers

import (
	"JxdzManager/models"
	"strconv"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {

	// l := c.GetSession("Login")
	// if l != true {
	// 	c.Redirect("/", 302)
	// }
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
	// c.Data["Html"] = "<em><strong>1234</strong></em>"

	c.TplName = "index.html"
	c.Data["CateName"] = models.SortCategory()
}

func (c *IndexController) Post() {
	cateIdstr := c.Input().Get("cid")
	cateId, _ := strconv.Atoi(cateIdstr)
	cateSearch, _ := models.SearchCategory(cateId)

	belongs := cateSearch.Name
	beego.Debug(belongs)

	title := c.Input().Get("title")
	content := c.Input().Get("content")
	err := models.AddContent(belongs, title, content)
	if err != nil {
		beego.Debug(err)
	}
	c.Redirect("/index", 302)
}
