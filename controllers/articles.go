package controllers

import (
	"github.com/astaxie/beego"
)

type ArticlesController struct {
	beego.Controller
}

func (this *ArticlesController) Index() {
	this.Data["Website"] = "oliver.me"
	this.Data["Email"] = "oliver@gmail.com"
	this.TplNames = "base.html"
}
