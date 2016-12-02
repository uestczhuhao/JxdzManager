package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Employee struct {
	Id        int
	Name      string
	Job       string
	Telphone  string
	Email     string
	Script    string
	Workplace string
}

func AddEmployee(messages ...string) error {
	o := orm.NewOrm()
	var mes [7]string
	var i int = 0
	for _, message := range messages {
		mes[i] = message
		i = i + 1
	}
	employee := &Employee{
		Name:      mes[0],
		Job:       mes[1],
		Telphone:  mes[2],
		Email:     mes[3],
		Script:    mes[4],
		Workplace: mes[5]}

	_, err := o.Insert(employee)
	return err
}

//根据目标的姓名和职位删除,二者必须完全符合要求
func DelEmployee(name string, job string) error {
	o := orm.NewOrm()

	emp := Employee{Name: name, Job: job}
	var err error
	beego.Debug(emp)
	beego.Debug(o.Read(&emp, "Name") == nil)
	beego.Debug(emp)
	beego.Debug(o.Read(&emp, "Job") == nil)
	if o.Read(&emp, "Name") == nil &&
		o.Read(&emp, "Job") == nil {
		_, err = o.Delete(&emp)
		beego.Debug(err)
	}

	return err
}
