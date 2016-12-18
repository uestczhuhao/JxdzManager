package Front

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type JxdzN2Controller struct {
	beego.Controller
}

func (c *JxdzN2Controller) Get() {
	id, err := c.GetInt("Id")
	if err != nil {
		beego.Error(err)
	}
	c.Data["articles"] = models.SearchArticleByBelongid(id)
	c.TplName = "Front/jxdzN2.html"
}
