package main

import (
	"JxdzManager/models"
	_ "JxdzManager/routers"
	"fmt"

	"github.com/astaxie/beego"
)

func main() {

	// models.AddUserSup("zhu", "123456")
	// models.AddUserCom("hao", "123456")

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
	// models.InitAndClear()
	// models.AddCategory("寄语1222", 6)

	// i := models.DelCategory(9, "55555", 2)
	// beego.Debug(i)
	// beego.Debug(strings.Replace("a,,d,1,2", ",", "", -1))

	// cate, err := models.GetAllCategories()
	// if err != nil {
	// 	beego.Error(err)
	// }
	// for i := 0; i < len(cate); i++ {
	// 	beego.Debug(cate[i])
	// }
	users, _ := models.GetAllUsers()
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
	str, Cates := models.StandardOut()
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
		fmt.Println(Cates[i])
	}

	// beego.Debug(strings.Trim(",12,3,", ","))
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
