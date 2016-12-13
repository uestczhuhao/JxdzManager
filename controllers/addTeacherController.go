package controllers

import (
	"JxdzManager/models"
	"strconv"

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

	c.Data["CateName"] = models.SortCategory()
	c.TplName = "addTeacher.html"
}

func (c *AddTeacherController) Post() {
	cateIdstr := c.Input().Get("cid")
	name := c.Input().Get("name")
	job := c.Input().Get("job")
	telphone := c.Input().Get("telphone")
	email := c.Input().Get("email")
	script := c.Input().Get("script")
	educationBackground := c.Input().Get("workexperience")
	isDocTeacher := c.Input().Get("isdocteacher")
	depart := c.Input().Get("depart")
	cateId, _ := strconv.Atoi(cateIdstr)
	cateSearch, _ := models.SearchCategory(cateId)

	belongs := cateSearch.Name
	beego.Debug(belongs, name, job, telphone, email, script, educationBackground, isDocTeacher, depart)

	err := models.AddEmployee(belongs, name, job, telphone, email, script, educationBackground, isDocTeacher, depart)
	if err != nil {
		beego.Error(err)
	}

	c.Redirect("/addTeacher", 302)
}
