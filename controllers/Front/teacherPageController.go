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
	// beego.Debug(teacher)
	c.TplName = "Front/jxdzN6.html"
}
