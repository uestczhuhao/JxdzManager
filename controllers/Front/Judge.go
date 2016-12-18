package Front

import (
	"JxdzManager/models"
	"strconv"

	"github.com/astaxie/beego"
)

type JudgeController struct {
	beego.Controller
}

func (c *JudgeController) Judge() {

	id, err := c.GetInt("ID")

	if err != nil {
		beego.Error(err)
	}
	cate, err := models.SearchCategory(id)
	if err != nil {
		beego.Error(err)
	}
	cateType := cate.Type
	if cateType == "目录" {
		cateSon, _ := models.FindAllSon(id)
		if cateSon[0].Type != "目录" {
			cateType = cateSon[0].Type
		} else {
			cateGrandson, _ := models.FindAllSon(cateSon[0].Id)
			cateType = cateGrandson[0].Type
		}
	}

	if cateType == "文章列表" {
		idstr := strconv.Itoa(id)
		c.Redirect("/jxdzN2?Id="+idstr, 302)
	} else if cateType == "单页面" {
		c.Redirect("/index", 302)
	} else if cateType == "教师列表" {
		c.Redirect("/jxdzN4?Job=职称", 302)
	} else if cateType == "链接列表" {
		c.Redirect("/jxdzN3", 302)
	} else if cateType == "下载列表" {
		c.Redirect("/jxdzN4", 302)
	} else {
		beego.Error("类型错误")
	}
	// c.Redirect("/index", 302)
}

// func BindMenu(id int) error{
//   cate, err := models.SearchCategory(id)
//   if err != nil {
// 		beego.Error(err)
// 	}
//   cateSon,_:=models.FindAllSon(id)
//
// }
// } else if cateType == "目录"{
//   BindMenu(id)
// }
