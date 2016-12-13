package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Content struct {
	Id      int
	Belongs string
	Title   string `orm:"size(30)"`
	Content string `orm:"size(5000)"`
}

func AddContent(belong string, title string, cont string) error {

	o := orm.NewOrm()

	content := &Content{Belongs: belong, Title: title, Content: cont}

	_, err := o.Insert(content)
	return err

}

func DelContent(til string) error {
	o := orm.NewOrm()

	cont := Content{Title: til}
	var err error
	if o.Read(&cont, "Title") == nil {
		_, err = o.Delete(&cont)
	}
	return err
}

func SearchContent(til string) (cont Content, err error) {
	o := orm.NewOrm()
	cont = Content{Title: til}
	err = o.Read(&cont, "Title")
	if err == nil {
		return cont, err
	} else {
		cont = Content{Title: ""}

	}
	return cont, err
}

func GetAllContent() []*Content {
	o := orm.NewOrm()
	cont := make([]*Content, 0)

	_, err := o.QueryTable("Content").All(&cont)
	if err != nil {
		beego.Error(err)
	}
	return cont
}
