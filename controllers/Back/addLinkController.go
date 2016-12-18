package Back

import (
	"JxdzManager/models"
	"strconv"

	"github.com/astaxie/beego"
)

type AddLinkController struct {
	beego.Controller
}

func (c *AddLinkController) Get() {
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
	c.TplName = "addLink.html"
}

func (c *AddLinkController) Post() {
	cateIdstr := c.Input().Get("cid")
	cateId, _ := strconv.Atoi(cateIdstr)
	cateSearch, _ := models.SearchCategory(cateId)

	belongs := cateSearch.Name
	beego.Debug(belongs)
	title := c.Input().Get("title")
	linker := c.Input().Get("linker")
	err := models.AddLinker(belongs, title, linker)
	if err != nil {
		beego.Debug(err)
	}
	c.Redirect("/addLink", 302)
}
