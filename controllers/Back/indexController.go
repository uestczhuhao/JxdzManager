package Back

import (
	"JxdzManager/models"
	"strconv"
	"time"

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
	c.Data["CateNameDepthOne"] = models.GetAllCategoriesDepthIsOne()

	onedepthidstr := c.Input().Get("OneDepthId")
	onedepthid, _ := strconv.Atoi(onedepthidstr)
	c.Data["CateNameDepthOther"], _ = models.FindAllChild(onedepthid)

	c.Data["CateName"] = models.SortCategory()
}

func (c *IndexController) Post() {
	cateIdstr := c.Input().Get("cid")
	cateId, _ := strconv.Atoi(cateIdstr)

	title := c.Input().Get("title")
	from := c.Input().Get("from")
	content := c.Input().Get("content")
	timenow := time.Now().Format("2006-01-02")
	err := models.AddContent(cateId, title, content, from, timenow)
	if err != nil {
		beego.Debug(err)
	}
	c.Redirect("/index", 302)
}
