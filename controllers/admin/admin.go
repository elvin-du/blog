package admin

import (
	"blog/controllers"
	"blog/models"

	"github.com/astaxie/beego"
)

type AdminController struct {
	baseController
}

func (this *AdminController) Index() {
	beego.Debug("heerer")
	this.TplNames = "admin/login.html"
}

//Method-POST; URL-/admin/login
func (this *AdminController) Login() {
	name := this.GetString("name")
	passwd := this.GetString("passwd")

	sess, err := models.AdminModel().Login(name, passwd)
	if nil != err {
		beego.Error(err)
		this.Redirect("/admin/login", 302)
		return
	}
	this.Ctx.SetCookie(controllers.C_COOKIE_NAME, sess)
	this.Redirect("/admin/article/add", 302)
}

func (this *AdminController) Logout() {
	session := this.Ctx.GetCookie(controllers.C_COOKIE_NAME)
	this.Ctx.SetCookie(controllers.C_COOKIE_NAME, "")
	err := models.AdminModel().DelSession(session)
	if nil != err {
		beego.Error(err)
	}

	this.Redirect("/admin/login", 302)
}
