package controllers

import (
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}

func (this *AboutController) Index() {
	this.Data["Content"] = "This is just a private blog. I am a programmer, love coding."
	this.Layout = "layout.html"
	this.TplNames = "about/index.html"
}
