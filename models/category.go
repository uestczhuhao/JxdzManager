package models

import (
	"strconv"
	"strings"

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

func AddCategory(name string, parentid string) error {
	var err error
	var desid string
	var resultid int
	var curid int64
	o := orm.NewOrm()

	if strings.Contains(parentid, ",") == true {
		temp := strings.Split(parentid, ",")
		// beego.Debug(len(temp))
		for i := 0; i < len(temp); i++ {
			desid = temp[i]
		}
	} else {
		desid = parentid
	}
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

func DelCategory(id int, name string) int {
	o := orm.NewOrm()
	q := o.QueryTable("Category")
	var category Category
	// var idstr string

	err := q.Filter("Id", id).Filter("Name", name).One(&category)
	if err != nil {
		return 1
	} else {
		if category.ChildID == "0" {
			o.Delete(&category)
			return 0
		} else {
			return 1
		}
	}

}
