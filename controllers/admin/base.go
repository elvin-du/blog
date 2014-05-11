package admin

import "github.com/astaxie/beego"

type baseController struct {
	controllerName string
	actionName     string
	beego.Controller
}

func (this *baseController) Prepare() {
	this.Layout = "layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "admin/nav.html"
	this.LayoutSections["CSS"] = "admin/css.html"
	this.LayoutSections["JS"] = "admin/js.html"
	this.controllerName, this.actionName = this.GetControllerAndAction()
	this.auth()
}

func (this *baseController) auth() {
	switch {
	case this.controllerName == "LoginController" && this.actionName == "Login":
	case this.controllerName == "LoginController" && this.actionName == "Index":
		if this.validSess() {
			this.Redirect("/admin/add", 302)
		}
	default:
		if !this.validSess() {
			this.Redirect("/admin/login", 302)
		}
	}
}

func (this *baseController) validSess() bool {
	//sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	//defer sess.SessionRelease(this.Ctx.ResponseWriter)
	//cookie := sess.Get(C_COOKIE_NAME)
	cookie := this.Ctx.GetCookie(C_COOKIE_NAME)
	log.Debug("%s", cookie)
	if "" == cookie {
		return false
	}
	return true
}
