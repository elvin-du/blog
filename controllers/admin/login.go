package admin

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
	AuthController
}

func (this *LoginController) Index() {
	beego.Debug("LoginController:index()")
	this.Layout = "layout.html"
	this.TplNames = "admin/login.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "admin/nav.html"
	this.LayoutSections["CSS"] = "admin/css.html"
	this.LayoutSections["JS"] = "admin/js.html"

	err := this.Auth(this.Ctx)
	if nil == err {
		this.Redirect("/admin/add", 302)
	}
}

//Method-POST; URL-/admin/login
func (this *LoginController) Login() {
	beego.Debug("LoginController().Login()")
	name := this.GetString("name")
	passwd := this.GetString("passwd")

	err := models.AdminModel().Login(name, passwd)
	if nil != err {
		beego.Error(err)
		this.Redirect("/admin/login", 302)
	}
	beego.Debug("here")

	sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	sess.Set(C_COOKIE_NAME, sess.SessionID())

	//this.Ctx.SetCookie(C_COOKIE_NAME, cookie)
	this.Redirect("/admin/add", 302)
}
