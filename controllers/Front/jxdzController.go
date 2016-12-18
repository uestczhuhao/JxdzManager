package Front

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type JxdzController struct {
	beego.Controller
}

func (c *JxdzController) Get() {
	var err error
	c.Data["CatesForMenu"], err = models.CreateCateList()

	if err != nil {
		beego.Error(err)
	}
	// c.Data["number"]=
	// c.Data["CateDepthTwo1"], _ = models.FindAllSon(1)
	// c.Data["CateDepthTwo2"], _ = models.FindAllSon(6)
	// c.Data["CateDepthTwo3"], _ = models.FindAllSon(11)
	c.TplName = "Front/jxdz.html"
}
