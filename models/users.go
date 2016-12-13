package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id        int
	Account   string `orm:"unique"`
	Password  string `orm:"size(30)"`
	Authority string
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Users), new(Category), new(Content), new(Employee), new(Linker))

	// 设置基本的数据库
	orm.RegisterDataBase("default", "mysql", "root:golang@/user?charset=utf8", 30)

	// create table
	//第二个参数为true时每次运行重新建表
	orm.RunSyncdb("default", false, true)
}

func AddUserCom(account string, password string) error {
	o := orm.NewOrm()

	commonuser := &Users{
		Account:   account,
		Password:  password,
		Authority: "common"}

	_, err := o.Insert(commonuser)
	return err
}

func AddUserSup(account string, password string) error {
	o := orm.NewOrm()

	superuser := &Users{
		Account:   account,
		Password:  password,
		Authority: "super"}

	_, err := o.Insert(superuser)

	return err
}

func DelUser(account string) error {
	o := orm.NewOrm()

	user := Users{Account: account}
	var err error
	if o.Read(&user, "Account") == nil {
		_, err = o.Delete(&user)
	}
	return err
}

func SearchUser(account string) (users Users, err error) {
	o := orm.NewOrm()

	users = Users{Account: account}
	if o.Read(&users, "Account") == nil {
		return users, err
	} else {
		users = Users{Account: ""}
	}
	return users, nil

}

func GetAllUsers() ([]*Users, error) {
	o := orm.NewOrm()

	users := make([]*Users, 0)

	qs := o.QueryTable("users")
	_, err := qs.All(&users)
	return users, err
}
