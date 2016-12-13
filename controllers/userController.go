package controllers

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	u.TplName = "user.html"
	var err error

	account := u.Input().Get("account")
	if account != "" {
		err = models.DelUser(account)
		if err != nil {
			beego.Error(err)
		}
	}

	u.Data["AllUsers"], err = models.GetAllUsers()
	if err != nil {
		beego.Error(err)
	}
	u.Data["CateName"], u.Data["Categories"] = models.StandardOut()
}

func (u *UserController) Post() {
	newaccount := u.Input().Get("newaccount")
	newpassword := u.Input().Get("newpassword")
	authority := u.Input().Get("right")
	beego.Debug(newaccount, newpassword, authority)
	if newaccount != "" {
		if authority == "super" {
			err := models.AddUserSup(newaccount, newpassword)
			if err != nil {
				beego.Error(err)
			}
		} else if authority == "normal" {
			err := models.AddUserCom(newaccount, newpassword)
			if err != nil {
				beego.Error(err)
			}
		}
	}
	u.Redirect("/user", 302)
}
