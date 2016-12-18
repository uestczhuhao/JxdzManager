package Back

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type OneAticleController struct {
	beego.Controller
}

func (c *OneAticleController) Get() {
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
	articletitle := c.Input().Get("articletitle")
	c.Data["CateNameDepthOne"] = models.GetAllCategoriesDepthIsOne()
	c.Data["CateName"] = models.SortCategory()

	c.Data["Ariticle"], _ = models.SearchContent(articletitle)
	beego.Debug(c.Input().Get("articletitle"))
	c.TplName = "onearticle.html"
}
