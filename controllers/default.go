package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	fmt.Println(this.Data["IsLogin"])
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopic(true)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}

	fmt.Println(this.Data["IsLogin"])
}
