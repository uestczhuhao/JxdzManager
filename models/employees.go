package models

import "github.com/astaxie/beego/orm"

type Employee struct {
	Id           int
	Name         string
	Job          string
	Telphone     string
	Email        string
	Script       string `orm:"size(800)"`
	Workplace    string `orm:"size(500)"`
	IsDocTeacher string `orm:"size(1)"`
	Department   string
}

func AddEmployee(messages ...string) error {
	o := orm.NewOrm()
	var mes [10]string
	var i int = 0
	for _, message := range messages {
		mes[i] = message
		i = i + 1
	}
	employee := &Employee{
		Name:         mes[0],
		Job:          mes[1],
		Telphone:     mes[2],
		Email:        mes[3],
		Script:       mes[4],
		Workplace:    mes[5],
		IsDocTeacher: mes[6],
		Department:   mes[7]}

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
