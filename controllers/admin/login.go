package admin

import (
	"blog/controllers"
	"blog/models"

	"github.com/astaxie/beego"
)

type LoginController struct {
	baseController
}

func (this *LoginController) Index() {
	this.TplNames = "admin/login.html"
}

//Method-POST; URL-/admin/login
func (this *LoginController) Login() {
	name := this.GetString("name")
	passwd := this.GetString("passwd")

	sess, err := models.AdminModel().Login(name, passwd)
	if nil != err {
		beego.Error(err)
		this.Redirect("/admin/login", 302)
	}
	this.Ctx.SetCookie(controllers.C_COOKIE_NAME, sess)
	this.Redirect("/admin/article/add", 302)
}
