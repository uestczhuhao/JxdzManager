package controllers

import (
	"github.com/astaxie/beego"
)

type FormController struct {
	beego.Controller
}

func (c *FormController) Get() {

	c.TplName = "form.html"
}
