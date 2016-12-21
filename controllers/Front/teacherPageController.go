package Front

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type JxdzN6Controller struct {
	beego.Controller
}

func (c *JxdzN6Controller) Get() {

	teacherId, err := c.GetInt("ID")
	if err != nil {
		beego.Error(err)
	}
	c.Data["teacher"] = models.SearchTeacherById(teacherId)
	c.Data["CatesForMenu"], err = models.CreateCateList()
	c.Data["link"] = models.GetAllLinker()

	hotNewsList := models.SearchArticleByBelongid(23)
	c.Data["hotNews"] = hotNewsList[0:5]

	c.Data["Path"] = models.FindAllFather(7)
	c.Data["LeftMenu"] = models.FindLeftMenu(7)
	// beego.Debug(teacher)
	c.TplName = "Front/jxdzN6.html"
}
