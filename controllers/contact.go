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
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "contact/nav.html"
}
