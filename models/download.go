package models

import (
	"github.com/astaxie/beego/orm"
)

type Download struct {
	Id        int
	Title     string `orm:"size(45)"`
	LinkTo    string
	CreatTime string `orm:"size(100)"`
}

func AddDownload(tiltle string, linkto string, creatime string) error {
	o := orm.NewOrm()

	DownLoad := &Download{Title: tiltle, LinkTo: linkto, CreatTime: creatime}
	_, err := o.Insert(DownLoad)
	return err
}

func DeleteDownload(title string) error {
	o := orm.NewOrm()

	download := Download{Title: title}
	if o.Read(&download, "Title") == nil {
		_, err := o.Delete(&download)
		if err != nil {
			return err
		}
	}

	return nil
}

func SearchDownload(id int) (Download, error) {
	o := orm.NewOrm()

	down := Download{Id: id}
	err := o.Read(&down, "Id")
	if err == nil {
		return down, nil
	} else {
		return down, err
	}
}

func GetAllDownload() ([]*Download, error) {
	o := orm.NewOrm()

	down := make([]*Download, 0)

	_, err := o.QueryTable("Download").All(&down)
	return down, err
}
