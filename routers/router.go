package routers

import (
	"JxdzManager/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/",&controllers.LoginController{})
	beego.Router("/", &controllers.MainController{}, "get:LoginIndex;post:Post")
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/form", &controllers.FormController{})
	beego.Router("/add", &controllers.AddController{})

}
