package main

import (
	"JxdzManager/models"
	_ "JxdzManager/routers"

	"github.com/astaxie/beego"
)

func main() {

	models.AddUserSup("zhu", "123456")
	models.AddUserCom("hao", "123456")

	var err error
	err = models.AddEmployee("zhu", "某某职位", "22333322", "843728@123.com")
	// err = models.AddEmployee("zhu", "某某职位", "22333322", "843728@123.com")
	// err = models.AddEmployee("zhu", "某某职位", "22333322", "843728@123.com")
	// err = models.AddEmployee("zhu", "某某职位", "22333322", "843728@123.com")
	// err = models.AddEmployee("zhu", "某某职位", "22333322", "843728@123.com")
	// err = models.AddEmployee("zhu", "某某职位", "22333322", "843728@123.com")
	// err = models.AddEmployee("zhu", "某某职位", "22333322", "843728@123.com")
	// err = models.AddEmployee("zhu", "某某职位", "22333322", "843728@123.com")
	// err = models.AddEmployee("zhu", "某某职位", "22333322", "843728@123.com")
	// err = models.AddEmployee("zhu", "某某职位", "22333322", "843728@123.com")
	// err = models.AddEmployee("zhu", "某某职位", "22333322", "843728@123.com")
	// // err := models.DelContent("adddaa")
	// if err != nil {
	// 	beego.Error(err)
	// }
	err = models.DelEmployee("zhu", "某某")
	if err != nil {
		beego.Error(err)
	}
	beego.Run()
}
