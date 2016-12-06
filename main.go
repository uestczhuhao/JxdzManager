package main

import (
	"JxdzManager/models"
	_ "JxdzManager/routers"

	"github.com/astaxie/beego"
)

func main() {

	models.AddUserSup("zhu", "123456")
	models.AddUserCom("hao", "123456")

	// var err error
	// err = models.AddEmployee("aaa", "某某职位qw", "22333322", "843728@123.com", "", "", "1")
	// err = models.AddEmployee("adddaa", "某某qwq职位", "22333322", "843728@123.com", "", "", "1", "jixie")
	// err = models.AddEmployee("a", "某某职eqe位", "22333322", "843728@123.com", "", "", "1", "jixie")
	// err = models.AddEmployee("sa", "某rw某职位", "22333322", "843728@123.com", "", "", "1", "jixie")

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

	// IsError := models.DelCategory(12, "寄语11")
	// beego.Debug(strings.Replace("a,,d,1,2", ",", "", -1))
	// emps, err := models.SelectEmployeeByDepartment("jixie")
	// // beego.Debug(len(emps))
	//
	// for i := 0; i < len(emps); i = i + 1 {
	// 	beego.Debug(emps[i].Name)
	// }
	// // beego.Debug(emps)
	// if err != nil {
	// 	beego.Error(err)
	// }
	beego.Run()
}
