package routers

import (
	"JxdzManager/controllers/Back"
	"JxdzManager/controllers/Front"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/",&controllers.LoginController{})
	beego.Router("/", &Back.MainController{}, "get:LoginIndex;post:Post")
	beego.Router("/index", &Back.IndexController{})
	beego.Router("/category", &Back.FormController{})
	beego.Router("/addTeacher", &Back.AddTeacherController{})
	beego.Router("/addLink", &Back.AddLinkController{})
	beego.Router("/addDownload", &Back.AddDownloadController{})
	beego.Router("/user", &Back.UserController{})

	beego.Router("/teacherlist", &Back.TeacherListController{})
	beego.Router("/linkerlist", &Back.LinkerListController{})
	beego.Router("/articlelist", &Back.AticleListController{})
	beego.Router("/onearticle", &Back.OneAticleController{})

	//前端路由设置
	beego.Router("/Judge", &Front.JudgeController{}, "get:Judge")
	beego.Router("/jxdz", &Front.JxdzController{})
	beego.Router("/jxdzN1", &Front.JxdzN1Controller{})
	beego.Router("/jxdzN2", &Front.JxdzN2Controller{})
	beego.Router("/jxdzN3", &Front.JxdzN3Controller{})
	beego.Router("/jxdzN4", &Front.JxdzN4Controller{})
	beego.Router("/jxdzN5", &Front.JxdzN5Controller{})
	beego.Router("/jxdzN6", &Front.JxdzN6Controller{})

}
