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
	this.Layout = "admin/layout.html"
	this.LayoutSections = make(map[string]string)
	this.controllerName, this.actionName = this.GetControllerAndAction()
	this.auth()
}

func (this *baseController) auth() {
	switch {
	case this.controllerName == "AdminController" && this.actionName == "Login":
	//do nothing
	case this.controllerName == "AdminController" && this.actionName == "Index":
		if this.validSess() {
			this.Data["logined"] = true
			this.Redirect("/admin/article/add", 302)
			return
		}
		this.Data["logined"] = false
		this.LayoutSections["login"] = "admin/login.html"
	default:
		if !this.validSess() {
			this.Data["logined"] = false
			this.LayoutSections["login"] = "admin/login.html"
			this.Redirect("/admin/login", 302)
			return
		}
		this.Data["logined"] = true
	}
}

func (this *baseController) validSess() bool {
	cookie := this.Ctx.GetCookie(controllers.C_COOKIE_NAME)
	return models.AdminModel().ValidSession(cookie)
}
