package Front

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type JxdzN4Controller struct {
	beego.Controller
}

func (c *JxdzN4Controller) Get() {
	job := c.Input().Get("Job")
	doc := c.Input().Get("Doc")
	dep := c.Input().Get("Dep")
	if job != "" {
		c.Data["Show1"] = true
		c.Data["TeachersPro"], _ = models.SelectEmployeeByJob("教授")
		c.Data["TeachersDepuPro"], _ = models.SelectEmployeeByJob("副教授")
		c.Data["Teachers"], _ = models.SelectEmployeeByJob("讲师")

	} else if doc != "" {
		c.Data["Show2"] = true
		c.Data["Teachers"], _ = models.SelectEmployeeByDocTeacher("是")
	} else if dep != "" {
		c.Data["Show3"] = true
		c.Data["Teachers1"], _ = models.SelectEmployeeByDepartment("机械工程系")
		c.Data["Teachers2"], _ = models.SelectEmployeeByDepartment("工业工程系")
		c.Data["Teachers3"], _ = models.SelectEmployeeByDepartment("电力电子系")
		c.Data["Teachers4"], _ = models.SelectEmployeeByDepartment("工程训练中心")

	} else {
		c.Data["Show1"] = false
		c.Data["Show2"] = false
		c.Data["Show3"] = false
	}

	c.TplName = "Front/jxdzN4.html"
}
