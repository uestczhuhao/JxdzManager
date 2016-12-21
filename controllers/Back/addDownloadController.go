package Back

import (
	"JxdzManager/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type AddDownloadController struct {
	beego.Controller
}

func (c *AddDownloadController) Get() {
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
	c.TplName = "addDownload.html"
}

func (c *AddDownloadController) Post() {
	cateIdstr := c.Input().Get("cid")
	cateId, _ := strconv.Atoi(cateIdstr)
	cateSearch, _ := models.SearchCategory(cateId)

	belongs := cateSearch.Id
	title := c.Input().Get("title")
	linker := c.Input().Get("linker")
	timenow := time.Now().Format("2006-01-02")
	err := models.AddDownload(belongs, title, linker, timenow)
	if err != nil {
		beego.Debug(err)
	}
	c.Redirect("/addDownload", 302)
}
