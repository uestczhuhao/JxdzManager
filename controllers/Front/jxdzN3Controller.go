package Front

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type JxdzN3Controller struct {
	beego.Controller
}

func (c *JxdzN3Controller) Get() {
	id, err := c.GetInt("Id")
	if err != nil {
		beego.Error(err)
	}
	c.Data["downlist"] = models.SearchDownByBelongid(id)
	c.Data["CateName"], _ = models.SearchCategory(id)
	c.Data["CatesForMenu"], err = models.CreateCateList()
	c.Data["link"] = models.GetAllLinker()

	c.Data["Path"] = models.FindAllFather(id)
	c.Data["LeftMenu"] = models.FindLeftMenu(id)
	c.TplName = "Front/jxdzN3.html"
}
