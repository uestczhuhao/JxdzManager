package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Content struct {
	Id      int
	Title   string `orm:"size(30)"`
	Content string `orm:"size(500)"`
}

func AddContent(title string, cont string) error {

	o := orm.NewOrm()

	content := &Content{Title: title, Content: cont}

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
