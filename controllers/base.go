package controllers

import "github.com/astaxie/beego"

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
	//sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	//defer sess.SessionRelease(this.Ctx.ResponseWriter)
	//cookie := sess.Get(C_COOKIE_NAME)
	cookie := this.Ctx.GetCookie(C_COOKIE_NAME)
	if "macs" != cookie {
		return false
	}
	return true
}
