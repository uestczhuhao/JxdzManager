package models

import (
	"github.com/astaxie/beego/orm"
)

type Institution struct {
	Id         int
	Title      string
	Duty       string `orm:"size(500)"`
	MemberName string
	MemberJob  string
}

func AddInstitution(title string, duty string) error {
	o := orm.NewOrm()
	newinst := &Institution{Title: title, Duty: duty}

	_, err := o.Insert(newinst)

	return err

}

func DelInstitution(title string) error {
	o := orm.NewOrm()
	institution := Institution{Title: title}

	var err error
	if o.Read(&institution, "Title") == nil {
		_, err = o.Delete(&institution)
	}
	return err

}
