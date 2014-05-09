package admin

import (
	"blog/models"
	"net/http"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
	AuthController
}

func (this *LoginController) Index() {
	this.Layout = "layout.html"
	this.TplNames = "admin/login.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "admin/nav.html"
	this.LayoutSections["CSS"] = "admin/css.html"
	this.LayoutSections["JS"] = "admin/js.html"

	//err := this.Auth(this.Ctx)
	//if nil == err {
	//	//this.Redirect("/admin/add", http.StatusOK)
	//	http.Redirect(this.Ctx.ResponseWriter, this.Ctx.Request, "/admin/add", http.StatusOK)
	//	return
	//}
}

//Method-POST; URL-/admin/login
func (this *LoginController) Login() {
	name := this.GetString("name")
	passwd := this.GetString("passwd")

	err := models.AdminModel().Login(name, passwd)
	if nil != err {
		beego.Error(err)
		this.Abort("login failed")
	}

	sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	sess.Set(C_COOKIE_NAME, sess.SessionID())

	//this.Ctx.SetCookie(C_COOKIE_NAME, cookie)
	this.Redirect("/admin/add", http.StatusOK)
}
