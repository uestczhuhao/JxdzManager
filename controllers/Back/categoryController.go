package Back

import (
	"JxdzManager/models"
	"strconv"

	"github.com/astaxie/beego"
)

type FormController struct {
	beego.Controller
}

func (c *FormController) Get() {
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

	c.Data["CateNameDepthOne"] = models.GetAllCategoriesDepthIsOne()
	c.Data["CateName"] = models.SortCategory()
}

func (c *FormController) Post() {
	labelname := c.Input().Get("catename")
	parentidstr := c.Input().Get("cid")
	parentid, err := strconv.Atoi(parentidstr)
	typ := c.Input().Get("typ")
	cont := c.Input().Get("content")
	beego.Debug(labelname, parentid)
	beego.Debug(typ, cont)
	err = models.AddCategory(labelname, parentid, typ, cont)
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/category", 302)
}
