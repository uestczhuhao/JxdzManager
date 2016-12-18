package Back

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type LinkerListController struct {
	beego.Controller
}

func (c *LinkerListController) Get() {
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
	c.Data["CateNameDepthOne"] = models.GetAllCategoriesDepthIsOne()
	c.Data["CateName"] = models.SortCategory()

	c.Data["linkers"] = models.GetAllLinker()

	c.TplName = "linklist.html"
}
