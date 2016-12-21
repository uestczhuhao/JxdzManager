package Front

import (
	"JxdzManager/models"

	"github.com/astaxie/beego"
)

type JxdzN5Controller struct {
	beego.Controller
}

func (c *JxdzN5Controller) Get() {
	id, err := c.GetInt("Id")
	if err != nil {
		beego.Error(err)
	}
	models.UpdateArticle(id)
	article, err := models.SearchArticleById(id)
	c.Data["Article"] = article
	cateId := article.BelongsId
	c.Data["CateName"], _ = models.SearchCategory(cateId)
	if err != nil {
		beego.Error(err)
	}
	c.Data["CatesForMenu"], err = models.CreateCateList()
	c.Data["link"] = models.GetAllLinker()
	hotNewsList := models.SearchArticleByBelongid(23)
	c.Data["hotNews"] = hotNewsList[0:5]

	beid, _ := c.GetInt("BeId")
	becate, _ := models.SearchCategory(beid)

	if beid == 23 {
		c.Data["isShowHotNews"] = false
	} else {
		c.Data["isShowHotNews"] = true
	}

	if beid == 35 || id == 36 {
		c.Data["IsShowRoot"] = false
	} else {
		c.Data["IsShowRoot"] = true
	}
	articles := models.SearchArticleByBelongid(beid)
	var nextid int = 0
	for i := 0; i < len(articles); i++ {
		maxlen := len(articles) - 1
		if articles[maxlen].Id == id {
			nextid = 0
			break
		}
		if articles[i].Id == id {
			nextid = articles[i+1].Id
			break
		}
	}
	c.Data["NextArticle"], _ = models.SearchArticleById(nextid)

	if becate.Depth > 0 {
		c.Data["Path"] = models.FindAllFather(beid)
		c.Data["LeftMenu"] = models.FindLeftMenu(beid)
	} else {
		c.Data["Path"] = models.FindAllFather(id)
		c.Data["LeftMenu"] = models.FindLeftMenu(id)

	}

	c.TplName = "Front/jxdzN5.html"
}
