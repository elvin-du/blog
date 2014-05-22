package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type baseController struct {
	controllerName string
	actionName     string
	beego.Controller
}

func (this *baseController) Prepare() {
	this.Layout = "layout.html"
	this.auth()
}

func (this *baseController) auth() {
	if this.validSess() {
		this.Data["admin"] = "admin"
	} else {
		this.Data["admin"] = ""
	}
}

func (this *baseController) validSess() bool {
	cookie := this.Ctx.GetCookie(C_COOKIE_NAME)
	return models.AdminModel().ValidSession(cookie)
}
