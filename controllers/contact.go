package controllers

import (
	"github.com/astaxie/beego"
)

type ContactController struct {
	beego.Controller
}

func (this *ContactController) Index() {
	this.Layout = "layout.html"
	this.TplNames = "contact/index.html"
}
