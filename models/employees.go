package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Employee struct {
	Id                  int
	Belongs             string
	Name                string
	Job                 string
	Telphone            string
	Email               string
	Script              string `orm:"size(800)"`
	EducationBackground string `orm:"size(5000)"`
	IsDocTeacher        string `orm:"size(5)"`
	Department          string
}

//其中workplace一栏被用作是学历

func AddEmployee(messages ...string) error {
	o := orm.NewOrm()
	var mes [20]string
	var i int = 0
	for _, message := range messages {
		mes[i] = message
		i = i + 1
	}
	employee := &Employee{
		Belongs:             mes[0],
		Name:                mes[1],
		Job:                 mes[2],
		Telphone:            mes[3],
		Email:               mes[4],
		Script:              mes[5],
		EducationBackground: mes[6],
		IsDocTeacher:        mes[7],
		Department:          mes[8]}

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

func SelectEmployeeByJob(job string) ([]*Employee, error) {
	o := orm.NewOrm()
	//func make(Type, size IntegerType) Type
	//返回的是type的引用
	emps := make([]*Employee, 0)

	_, err := o.QueryTable("Employee").Filter("Job", job).All(&emps)

	//返回的emps是slice值,用len(emps)来获取其长度
	return emps, err
}

func SelectEmployeeByDocTeacher(isornot string) ([]*Employee, error) {
	o := orm.NewOrm()
	emps := make([]*Employee, 0)

	_, err := o.QueryTable("Employee").Filter("IsDocTeacher", isornot).All(&emps)

	//返回的emps是slice值,用len(emps)来获取其长度
	return emps, err
}

func SelectEmployeeByDepartment(depart string) ([]*Employee, error) {
	o := orm.NewOrm()
	emps := make([]*Employee, 0)

	_, err := o.QueryTable("Employee").Filter("Department", depart).All(&emps)

	//返回的emps是slice值,用len(emps)来获取其长度
	return emps, err
}

func GetAllEmployees() []*Employee {
	o := orm.NewOrm()
	emp := make([]*Employee, 0)

	_, err := o.QueryTable("Employee").All(&emp)
	if err != nil {
		beego.Error(err)
	}
	return emp
}
