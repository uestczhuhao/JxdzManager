package main

import (
	"JxdzManager/models"
	_ "JxdzManager/routers"
	"fmt"
	// "time"

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
	// models.AddCategory("寄语", 25, "文章列表", "")

	// i := models.DelCategory(24, "寄语1222", 5)
	// beego.Debug(i)
	// beego.Debug(strings.Replace("a,,d,1,2", ",", "", -1))

	// cate, _ := models.CreateCateList()
	//
	// for i := 0; i < len(cate); i++ {
	// 	fmt.Println(cate[i].CateOne)
	// 	fmt.Println(cate[i].DepthNumber)
	// 	fmt.Println(cate[i].Id)
	// 	// for j := 0; j < len(cate[i].CateTwo); j++ {
	// 	// 	fmt.Println(cate[i].CateTwo[j])
	// 	// }
	// 	for j := 0; j < len(cate[i].CateOnesChild); j++ {
	// 		fmt.Println(cate[i].CateOnesChild[j])
	// 	}
	// }

	// for i := 0; i < len(cates); i++ {
	// 	fmt.Println(cates[i])
	//
	// }

	// tNow := time.Now()
	// timeNow := tNow.Format("2006-01-02")
	// beego.Debug(timeNow)

	users, _ := models.GetAllUsers()
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
	str, Cates := models.StandardOut()
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
		fmt.Println(Cates[i])
	}

	beego.AddFuncMap("Repeat", repeat)
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

func repeat(depth int) (out string) {
	for i := 1; i < depth; i++ {
		out += "--"
	}
	return out
}
