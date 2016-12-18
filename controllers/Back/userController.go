package Back

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {

	v := u.GetSession("authority")
	if v == 1 {
		u.Data["IsSuper"] = true
		u.Data["admin"] = "超级"
	} else if v == 2 {
		u.Data["IsSuper"] = false
		u.Data["admin"] = "普通"
	} else {
		u.Redirect("/", 302)
	}

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
	u.Data["CateNameDepthOne"] = models.GetAllCategoriesDepthIsOne()
	u.Data["CateName"] = models.SortCategory()
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
