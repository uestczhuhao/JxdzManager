package main

import (
	"JxdzManager/models"
	_ "JxdzManager/routers"
	"time"

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

	// var cates []*models.Category
	// cate1, _ := models.SearchCategory(1)
	// cates = append(cates, &cate1)
	// cate2, _ := models.SearchCategory(2)
	// cates = append(cates, &cate2)
	// cates := models.FindLeftMenu(3)
	//
	// for i := 0; i < len(cates); i++ {
	// 	fmt.Println(cates[i])
	//
	// }
	// search := models.SearchFunction("学子")
	//
	// for i := 0; i < len(search); i++ {
	// 	beego.Debug(search[i].Title)
	// }
	// tNow := time.Now()
	// timeNow := tNow.Format("2006-01-02")
	// beego.Debug(timeNow)

	// users, _ := models.GetAllUsers()
	// for i := 0; i < len(users); i++ {
	// 	fmt.Println(users[i])
	// }
	// str, Cates := models.StandardOut()
	// for i := 0; i < len(str); i++ {
	// 	fmt.Println(str[i])
	// 	fmt.Println(Cates[i])
	// }

	beego.AddFuncMap("Repeat", repeat)
	beego.AddFuncMap("Isnew", IsNewArticle)
	beego.AddFuncMap("Counttop", CountTopOfImport)
	beego.AddFuncMap("Idconvert", IdConvert)
	beego.AddFuncMap("Istoolong", TilIsToolong)
	beego.AddFuncMap("Idconvrtforimg", IdConverForImg)
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

// func Truncation(title string)(trun string){
//   lenth:=len(title)
// 	if lenth<15{
// 		trun=title
// 	}else {
//
// 	}
// }
func IdConverForImg(id int) (idout int) {
	HotNews := models.SearchArticleByBelongid(23)
	NewsShow := HotNews[0:5]
	for i := 0; i < 5; i++ {
		if NewsShow[i].Id == id {
			idout = i + 1
			return idout
		}
	}
	return idout
}

func TilIsToolong(title string) (IsorNot bool) {
	if len(title) > 37*3 {
		IsorNot = true
	} else {
		IsorNot = false
	}
	return IsorNot
}

func IdConvert(idin int) (idout int) {
	idout = 33 - idin
	return idout
}
func CountTopOfImport(idin int) (top int) {
	id := 33 - idin
	if id < 3 {
		top = 75 * (id + 5)
	} else {
		top = 75 * (id - 3)
	}
	return top
}

func IsNewArticle(creattime string) (isnew bool) {
	const shortForm = "2006-01-02"
	d, _ := time.Parse(shortForm, creattime)
	if time.Since(d).Hours() < 72 {
		isnew = true
	} else {
		isnew = false
	}
	return isnew
}

func repeat(depth int) (out string) {
	if depth > 20 {
		depth = depth % 50
	}
	if depth == 1 || depth == 2 {
		out += ""
	} else {
		for i := 1; i < depth-1; i++ {
			out += "--"
		}
	}
	return out
}
