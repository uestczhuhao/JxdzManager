package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type ImageController struct {
	beego.Controller
}

func (this *ImageController) Post() {
	beego.Debug("this is ImageController Post")
	f, h, err := c.GetFile("imgFile")
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		c.SaveToFile("imgFile", "static/upload/"+h.Filename) // 保存位置在 static/upload,没有文件夹要先创建
	}
	beego.Redirect("/index")
}
