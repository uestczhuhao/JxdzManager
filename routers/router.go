package routers

import (
	"JxdzManager/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/",&controllers.LoginController{})
	beego.Router("/", &controllers.MainController{}, "get:LoginIndex;post:Post")
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/category", &controllers.FormController{})
	beego.Router("/add", &controllers.AddController{})
	beego.Router("/test", &controllers.ImageController{})
}
