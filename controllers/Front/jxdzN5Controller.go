package Front

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type JxdzN5Controller struct {
	beego.Controller
}

func (c *JxdzN5Controller) Get() {
	id, err := c.GetInt("Id")
	if err != nil {
		beego.Error(err)
	}
	models.UpdateArticle(id)
	c.Data["Article"], err = models.SearchArticleById(id)
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "Front/jxdzN5.html"
}
