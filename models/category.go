package models

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id       int
	Name     string `orm:"size(45)"`
	ParentID string `orm:"size(45)"`
	ChildID  string `orm:"size(100)"`
	Depth    int
	Path     string
}

//作为测试用的初始化函数
func InitAndClear() {
	AddCategory("学校概况", "0")
	AddCategory("学院简介", "1")
	AddCategory("学院机构", "1")
	AddCategory("行政机构", "3")
	AddCategory("科研机构", "3")
	AddCategory("院长寄语", "1")
	return
}

func AddCategory(name string, parentid string) error {
	var err error
	var desid string
	var resultid int
	var curid int64
	o := orm.NewOrm()
	if parentid == "0" {
		cur := &Category{Name: name, ParentID: "0", ChildID: "0", Depth: 1, Path: name}
		_, err = o.Insert(cur)
		return err
	}

	desid = parentid
	resultid, _ = strconv.Atoi(desid)

	parentcate := Category{Id: resultid}

	if o.Read(&parentcate, "Id") == nil {
		depth := parentcate.Depth + 1

		path := parentcate.Path + "/" + name

		cur := &Category{Name: name, ParentID: parentid, ChildID: "0", Depth: depth, Path: path}

		curid, err = o.Insert(cur)
		if err != nil {
			return err
		}
	}

	if o.Read(&parentcate, "Id") == nil {
		var newchildid string
		var curidstr string
		curidstr = strconv.FormatInt(curid, 10)

		if parentcate.ChildID != "0" {
			newchildid = parentcate.ChildID + "," + curidstr
			parentcate.ChildID = newchildid
			o.Update(&parentcate, "ChildID")
		} else {
			parentcate.ChildID = curidstr
			o.Update(&parentcate, "ChildID")
		}
	}
	return err
}

func DelCategory(id int, name string, parentid string) int {
	o := orm.NewOrm()
	q := o.QueryTable("Category")
	var category Category
	// var idstr string

	err := q.Filter("Id", id).Filter("Name", name).Filter("ParentID", parentid).One(&category)
	if err != nil {
		return 1
	} else {
		//err==nil  即搜索到了目标栏目category
		if category.ChildID == "0" {
			//当要删除的这个栏目没有子栏目时才可以删除,否则拒绝删除命令
			var resultidDel int
			var newchildidOfDel string
			var DelIdStr string
			resultidDel, _ = strconv.Atoi(category.ParentID)
			parentcateOfDel := Category{Id: resultidDel}
			if o.Read(&parentcateOfDel, "Id") == nil {
				beego.Debug(strings.Contains(parentcateOfDel.ChildID, ",") == true)
				if strings.Contains(parentcateOfDel.ChildID, ",") == true {
					newchildidOfDel = parentcateOfDel.ChildID
					beego.Debug(newchildidOfDel)
					DelIdStr = strconv.Itoa(category.Id)
					beego.Debug(DelIdStr)
					newchildidOfDel = strings.Replace(newchildidOfDel, DelIdStr, "", -1)
					beego.Debug(newchildidOfDel)
					newchildidOfDel = strings.Replace(newchildidOfDel, ",,", ",", -1)
					beego.Debug(newchildidOfDel)
					newchildidOfDel = strings.Trim(newchildidOfDel, ",")
					beego.Debug(newchildidOfDel)
					parentcateOfDel.ChildID = newchildidOfDel
					o.Update(&parentcateOfDel, "ChildID")
				} else {
					parentcateOfDel.ChildID = "0"
					o.Update(&parentcateOfDel, "ChildID")
				}
			}

			o.Delete(&category)
			return 0
		} else {
			return 1
		}
	}

}
