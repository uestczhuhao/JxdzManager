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
	f, h, err := this.GetFile("imgFile")
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		this.SaveToFile("imgFile", "static/upload/"+h.Filename) // 保存位置在 static/upload,没有文件夹要先创建
	}
	this.TplName = "index.html"
}
