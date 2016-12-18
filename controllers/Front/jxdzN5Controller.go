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
	article, err := models.SearchArticleById(id)
	c.Data["Article"] = article
	cateId := article.BelongsId
	c.Data["CateName"], _ = models.SearchCategory(cateId)
	if err != nil {
		beego.Error(err)
	}
	c.Data["CatesForMenu"], err = models.CreateCateList()
	c.Data["link"] = models.GetAllLinker()

	beid, _ := c.GetInt("BeId")
	becate, _ := models.SearchCategory(beid)
	if becate.Depth > 10 {
		c.Data["Path"] = models.FindAllFather(beid)
		c.Data["LeftMenu"] = models.FindLeftMenu(beid)
	} else {
		c.Data["Path"] = models.FindAllFather(id)
		c.Data["LeftMenu"] = models.FindLeftMenu(id)

	}

	c.TplName = "Front/jxdzN5.html"
}
