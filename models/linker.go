package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Linker struct {
	Id      int
	Belongs string
	Title   string
	Url     string
}

func AddLinker(belong string, title string, url string) error {
	o := orm.NewOrm()
	newlink := &Linker{Belongs: belong, Title: title, Url: url}

	_, err := o.Insert(newlink)

	return err

}

func DelLinker(title string) error {
	o := orm.NewOrm()
	link := Linker{Title: title}

	var err error
	if o.Read(&link, "Title") == nil {
		_, err = o.Delete(&link)
	}
	return err

}

func GetAllLinker() []*Linker {
	o := orm.NewOrm()
	link := make([]*Linker, 0)

	_, err := o.QueryTable("Linker").All(&link)
	if err != nil {
		beego.Error(err)
	}
	return link
}
