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
	c.Data["link"] = models.GetAllLinker()

	c.Data["HotNews"] = models.SearchArticleByBelongid(23)

	News1 := models.SearchArticleByBelongid(27)
	c.Data["News1"] = News1
	c.Data["News2"] = models.SearchArticleByBelongid(28)
	c.Data["News3"] = models.SearchArticleByBelongid(29)
	c.Data["News4"] = models.SearchArticleByBelongid(30)
	c.Data["News5"] = models.SearchArticleByBelongid(31)
	c.Data["News6"] = models.SearchArticleByBelongid(32)
	c.Data["News7"] = models.SearchArticleByBelongid(33)
	c.Data["News8"] = models.SearchArticleByBelongid(34)

	// for i := 0; i < len(links); i++ {
	// 	beego.Debug(links[i].Title)
	// }
	c.TplName = "Front/jxdz.html"
}
