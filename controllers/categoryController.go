package controllers

import (
	"JxdzManager/models"
	"strconv"

	"github.com/astaxie/beego"
)

type FormController struct {
	beego.Controller
}

func (c *FormController) Get() {
	c.TplName = "category.html"
	idstr := c.Input().Get("id")
	name := c.Input().Get("name")
	parentidstr := c.Input().Get("parentid")

	if idstr != "" {
		id, err := strconv.Atoi(idstr)
		if err != nil {
			beego.Error(err)
		}
		parentid, err := strconv.Atoi(parentidstr)
		if err != nil {
			beego.Error(err)
		}
		i := models.DelCategory(id, name, parentid)
		if i == 1 {
			beego.Debug("该栏目有子栏目，拒绝删除")
		} else if i == 0 {
			beego.Debug("删除成功")
		}
	}
	c.Data["CateName"], c.Data["Categories"] = models.StandardOut()
	c.Data["CateName"] = models.SortCategory()
}

func (c *FormController) Post() {
	labelname := c.Input().Get("catename")
	parentidstr := c.Input().Get("cid")
	parentid, err := strconv.Atoi(parentidstr)
	beego.Debug(labelname)
	beego.Debug(parentid)
	err = models.AddCategory(labelname, parentid)
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/category", 302)
}
