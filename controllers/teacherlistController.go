package controllers

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type TeacherListController struct {
	beego.Controller
}

func (c *TeacherListController) Get() {
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
	teachername := c.Input().Get("teachername")
	teacherjob := c.Input().Get("teacherjob")
	beego.Debug(teachername, teacherjob)
	err := models.DelEmployee(teachername, teacherjob)
	if err != nil {
		beego.Debug(err)
	}
	c.Data["AllTeachers"] = models.GetAllEmployees()
	c.TplName = "teacherlist.html"
}

func (c *TeacherListController) Post() {

	c.Redirect("/teacherlist", 302)
}
