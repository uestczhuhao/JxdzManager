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
	beego.Router("/addTeacher", &controllers.AddTeacherController{})
	beego.Router("/addLink", &controllers.AddLinkController{})
	beego.Router("/user", &controllers.UserController{})

	beego.Router("/teacherlist", &controllers.TeacherListController{})
	beego.Router("/linkerlist", &controllers.LinkerListController{})
	beego.Router("/articlelist", &controllers.AticleListController{})
	beego.Router("/onearticle", &controllers.OneAticleController{})

}
