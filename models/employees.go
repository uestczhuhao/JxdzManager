package models

import "github.com/astaxie/beego/orm"

type Employee struct {
	Id        int
	Name      string
	Job       string
	Telphone  string
	Email     string
	Script    string `orm:"size(800)"`
	Workplace string `orm:"size(500)"`
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
	q := o.QueryTable("Employee")
	var emp Employee

	err := q.Filter("Name", name).Filter("Job", job).One(&emp)
	// beego.Debug(emp)
	_, err = o.Delete(&emp)
	return err
}
