package controllers

import (
	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}

func (this *AboutController) Index() {
	this.Layout = "layout.html"
	this.TplNames = "about/index.html"
	this.Data["Content"] = "This is just a private blog. I am a programmer, love coding."
	this.LayoutSections = make(map[string]string)
	//this.LayoutSections["CSS"] = "about/css.html"
	this.LayoutSections["CSS"] = ""
	this.LayoutSections["JS"] = ""
	this.LayoutSections["Nav"] = "about/nav.html"
}
