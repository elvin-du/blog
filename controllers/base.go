package controllers

import "github.com/astaxie/beego"

type baseController struct {
	beego.Controller
}

func (this *baseController) Prepare() {
	this.Layout = "layout.html"

}
