package Back

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type AddTeacherController struct {
	beego.Controller
}

func (c *AddTeacherController) Get() {
	v := c.GetSession("authority")
	if v == 1 {
		c.Data["IsSuper"] = true
		c.Data["admin"] = "超级"
	} else if v == 2 {
		c.Data["IsSuper"] = false
		c.Data["admin"] = "普通"
	} else {
		c.Redirect("/", 302)
	}
	c.Data["CateNameDepthOne"] = models.GetAllCategoriesDepthIsOne()

	c.Data["CateName"] = models.SortCategory()
	c.TplName = "addTeacher.html"
}

func (c *AddTeacherController) Post() {
	name := c.Input().Get("name")
	job := c.Input().Get("job")
	telphone := c.Input().Get("telphone")
	email := c.Input().Get("email")
	script := c.Input().Get("script")
	educationBackground := c.Input().Get("workexperience")
	isDocTeacher := c.Input().Get("isdocteacher")
	depart := c.Input().Get("depart")
	research := c.Input().Get("research")
	teach := c.Input().Get("teach")
	prize := c.Input().Get("prize")

	beego.Debug(name, job, telphone, email, script, educationBackground, isDocTeacher, depart)

	err := models.AddEmployee(name, job, telphone, email, script, educationBackground, isDocTeacher, depart, "", research, teach, prize)
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/addTeacher", 302)
}
