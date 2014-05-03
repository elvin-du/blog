package controllers

import "github.com/astaxie/beego"

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Index() {
	this.Layout = "layout.html"
	this.TplNames = "admin/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "admin/nav.html"
	this.LayoutSections["CSS"] = "admin/css.html"
	this.LayoutSections["JS"] = "admin/js.html"
}

//Method-POST; URL-/admin/login
func (this *AdminController) Login() {

}
