package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Content struct {
	Id        int
	BelongsId int
	Title     string `orm:"size(30)"`
	Content   string `orm:"size(15000)"`
	From      string
	CreatTime string
	ReadTime  int
}

func AddContent(belong int, title string, cont string, from string, ctime string) error {

	o := orm.NewOrm()

	content := &Content{BelongsId: belong, Title: title, Content: cont, From: from, CreatTime: ctime}

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

func SearchArticleById(id int) (cont Content, err error) {
	o := orm.NewOrm()
	cont = Content{Id: id}
	err = o.Read(&cont, "Id")
	if err == nil {
		return cont, err
	} else {
		cont = Content{Title: ""}
	}
	return cont, err
}

func SearchArticleByBelongid(id int) []*Content {
	allarticle := GetAllContent()
	var articles []*Content
	for i := 0; i < len(allarticle); i++ {
		if id == allarticle[i].BelongsId {
			articles = append(articles, allarticle[i])
		}
	}
	var lenth int = len(articles)
	var articlesout []*Content
	for j := 0; j < lenth; j++ {
		var temp *Content
		temp = articles[lenth-1-j]
		articlesout = append(articlesout, temp)
	}
	return articlesout
}

func UpdateArticle(id int) error {
	var err error
	o := orm.NewOrm()
	article, err := SearchArticleById(id)
	if err == nil {
		readtime := article.ReadTime

		readtime = readtime + 1
		article.ReadTime = readtime
		o.Update(&article)
	}
	return err
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
