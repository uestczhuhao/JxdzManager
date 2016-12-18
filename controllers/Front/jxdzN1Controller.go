package Front

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type JxdzN1Controller struct {
	beego.Controller
}

func (c *JxdzN1Controller) Get() {
	id, err := c.GetInt("Id")
	if err != nil {
		beego.Error(err)
	}

	cateOnePage, err := models.SearchCategory(id)
	if err != nil {
		beego.Error(err)
	}
	c.Data["OnePage"] = cateOnePage
	c.Data["pageContent"] = cateOnePage.Content
	c.Data["link"] = models.GetAllLinker()

	c.Data["CatesForMenu"], err = models.CreateCateList()

	catetest, _ := models.SearchCategory(id)
	catetestfather, _ := models.SearchCategory(catetest.ParentID)
	// beego.Debug(flag)
	if catetestfather.Type == "目录" && catetestfather.Depth != 1 {
		id = catetest.ParentID
	}
	c.Data["Path"] = models.FindAllFather(id)
	c.Data["LeftMenu"] = models.FindLeftMenu(id)
	c.TplName = "Front/jxdzN1.html"

}
