package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Index() {
	this.Data["Website"] = "oliver.me"
	this.Data["Email"] = "oliver@gmail.com"
	this.TplNames = "base.html"
}
