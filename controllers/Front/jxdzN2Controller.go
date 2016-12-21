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
	c.Data["CateName"], _ = models.SearchCategory(id)
	c.Data["CatesForMenu"], err = models.CreateCateList()
	c.Data["link"] = models.GetAllLinker()

	if id == 23 {
		c.Data["isShowHotNews"] = false
	} else {
		c.Data["isShowHotNews"] = true
	}
	if id == 35 || id == 36 {
		c.Data["IsShowRoot"] = false
	} else {
		c.Data["IsShowRoot"] = true
	}
	hotNewsList := models.SearchArticleByBelongid(23)
	c.Data["hotNews"] = hotNewsList[0:5]

	c.Data["Path"] = models.FindAllFather(id)
	c.Data["LeftMenu"] = models.FindLeftMenu(id)
	c.TplName = "Front/jxdzN2.html"
}

func (c *JxdzController) Post() {
	c.Data["CatesForMenu"], _ = models.CreateCateList()
	c.Data["Path"] = models.FindAllFather(36)
	c.Data["link"] = models.GetAllLinker()
	c.Data["isShowHotNews"] = true
	hotNewsList := models.SearchArticleByBelongid(23)
	c.Data["hotNews"] = hotNewsList[0:5]
	search := c.Input().Get("search")
	Articles := models.SearchFunction(search)
	c.Data["articles"] = Articles
	c.TplName = "Front/jxdzN2.html"

}
