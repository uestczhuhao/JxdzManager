package Back

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
	articleTitle := c.Input().Get("delarticletitle")

	err := models.DelContent(articleTitle)
	if err != nil {
		beego.Debug(err)
	}
	c.Data["CateNameDepthOne"] = models.GetAllCategoriesDepthIsOne()
	c.Data["CateName"] = models.SortCategory()

	c.Data["articles"] = models.GetAllContent()
	c.TplName = "articlelist.html"
}
