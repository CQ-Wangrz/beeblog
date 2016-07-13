package controllers

import (
	"beeblog/models"

	"fmt"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	op := this.Input().Get("op")
	name := this.Input().Get("name")
	fmt.Println("op value:" + op + " name value:" + name)
	switch op {
	case "add":
		fmt.Println("categoryADD")
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		fmt.Println("categoryDEL")
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategories(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	}
	this.Data["IsCategory"] = true
	this.TplName = "category.html"

	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}
