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
	ParentID int
	ChildID  string `orm:"size(100)"`
	Depth    int
}

//作为测试用的初始化函数
func InitAndClear() {
	AddCategory("学院概况", 0)
	AddCategory("学院机构", 1)
	AddCategory("行政机构", 2)
	AddCategory("院长寄语", 1)
	AddCategory("科研机构", 2)
	AddCategory("学院简介", 1)

	return
}

func AddCategory(name string, parentid int) error {
	var err error
	// var desid string
	var resultid int
	var curid int64
	o := orm.NewOrm()
	if parentid == 0 {
		cur := &Category{Name: name, ParentID: 0, ChildID: "0", Depth: 1}
		_, err = o.Insert(cur)
		return err
	}

	resultid = parentid

	parentcate := Category{Id: resultid}

	if o.Read(&parentcate, "Id") == nil {
		depth := parentcate.Depth + 1

		cur := &Category{Name: name, ParentID: parentid, ChildID: "0", Depth: depth}

		curid, err = o.Insert(cur)
		if err != nil {
			return err
		}
	} else {
		beego.Debug("输入父节点错误！")
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

func DelCategory(id int, name string, parentid int) int {
	o := orm.NewOrm()
	q := o.QueryTable("Category")
	var category Category
	// var idstr string

	err := q.Filter("Id", id).Filter("Name", name).Filter("ParentID", parentid).One(&category)

	if err != nil {
		beego.Debug(err)
		return 1

	} else {
		//err==nil  即搜索到了目标栏目category
		if category.ChildID == "0" {
			//当要删除的这个栏目没有子栏目时才可以删除,否则拒绝删除命令
			var resultidDel int
			var newchildidOfDel string
			var DelIdStr string
			resultidDel = category.ParentID
			parentcateOfDel := Category{Id: resultidDel}
			if o.Read(&parentcateOfDel, "Id") == nil {
				if strings.Contains(parentcateOfDel.ChildID, ",") == true {
					newchildidOfDel = parentcateOfDel.ChildID
					DelIdStr = strconv.Itoa(category.Id)
					newchildidOfDel = strings.Replace(newchildidOfDel, DelIdStr, "", -1)
					newchildidOfDel = strings.Replace(newchildidOfDel, ",,", ",", -1)
					newchildidOfDel = strings.Trim(newchildidOfDel, ",")
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

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cate := make([]*Category, 0)
	_, err := o.QueryTable("category").All(&cate)
	return cate, err
}

func GetSortedCategories() ([]*Category, error) {
	CateNotSort, err := GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	CateSort := make([]*Category, 0)
	CateSort = MakeSort(0, CateNotSort)
	return CateSort, err
}

func StandardOut() ([]string, []*Category) {
	StanOut := make([]string, 0)
	var i int
	var CateOutPut []*Category
	CateSort, err := GetSortedCategories()
	CateNotSort, err := GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	for i = 0; i < len(CateNotSort); i++ {
		j := CateSort[i].Depth - 1
		StanOut = append(StanOut, strings.Repeat("--", j)+CateSort[i].Name)
	}
	CateOutPut = CateSort[0:i]
	return StanOut, CateOutPut
}

var CateSort []*Category //全局变量，用以存储已经找到的栏目分类
func MakeSort(id int, CateNotSort []*Category) []*Category {
	var i int
	if id == 0 {
		CateSort = make([]*Category, 0)
	}
	for i = 0; i < len(CateNotSort); i++ {
		if id == CateNotSort[i].ParentID {
			CateSort = append(CateSort, CateNotSort[i])
			MakeSort(CateNotSort[i].Id, CateNotSort)
		}
	}
	return CateSort
}
