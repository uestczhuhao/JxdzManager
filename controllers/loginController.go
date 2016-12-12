package controllers

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) LoginIndex() {
	c.TplName = "login.html"

}

func (c *MainController) Post() {

	account := c.Input().Get("name")
	password := c.Input().Get("password")
	// passcode:=c.Input().Get("code")

	user, err := models.SearchUser(account)
	// beego.Debug(user.Account)
	// beego.Debug(user.Password)
	if err != nil {
		beego.Error(err)
	}
	if user.Account != "" {
		if password == user.Password {
			if user.Authority == "super" {
				c.SetSession("authority", int(1))
			} else if user.Authority == "common" {
				c.SetSession("authority", int(2))
			}
			c.SetSession("Login", true)
			c.Redirect("/index", 302)
		} else {
			beego.Debug("密码输入错误!")
			c.Redirect("/", 302)
			///想办法把验证错误的信息传到前端
		}
	} else {
		beego.Debug("此账号不存在,请查正后重新输入!")
		c.Redirect("/", 302)
		///想办法把验证错误的信息传到前端
	}

	RegUsername := c.Input().Get("username1")
	RegPaaword := c.Input().Get("password1")
	Authority := c.Input().Get("right")
	// beego.Debug("Comming Register!")
	if Authority == "super" {
		models.AddUserSup(RegUsername, RegPaaword)
	} else if Authority == "normal" {
		models.AddUserCom(RegUsername, RegPaaword)
	}
}
