package admin

import (
	"blog/controllers"
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
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "admin/nav.html"
	this.LayoutSections["CSS"] = "admin/css.html"
	this.LayoutSections["JS"] = "admin/js.html"
	this.controllerName, this.actionName = this.GetControllerAndAction()
	this.auth()
}

func (this *baseController) auth() {
	beego.Debug("autho")
	switch {
	case this.controllerName == "AdminController" && this.actionName == "Login":
	case this.controllerName == "AdminController" && this.actionName == "Index":
		if this.validSess() {
			this.Data["admin"] = true
			this.Redirect("/admin/article/add", 302)
			return
		}
		this.Data["admin"] = false
	default:
		if !this.validSess() {
			this.Data["admin"] = false
			this.Redirect("/admin/login", 302)
			return
		}
		this.Data["admin"] = true
	}
}

func (this *baseController) validSess() bool {
	cookie := this.Ctx.GetCookie(controllers.C_COOKIE_NAME)
	return models.AdminModel().ValidSession(cookie)
}
