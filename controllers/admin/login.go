package admin

import "blog/models"

type LoginController struct {
	baseController
}

func (this *LoginController) Index() {
	//this.Layout = "layout.html"
	this.TplNames = "admin/login.html"
	//this.LayoutSections = make(map[string]string)
	//this.LayoutSections["Nav"] = "admin/nav.html"
	//this.LayoutSections["CSS"] = "admin/css.html"
	//this.LayoutSections["JS"] = "admin/js.html"
}

//Method-POST; URL-/admin/login
func (this *LoginController) Login() {
	log.Debug("LoginController().Login()")
	name := this.GetString("name")
	passwd := this.GetString("passwd")

	err := models.AdminModel().Login(name, passwd)
	if nil != err {
		log.Error("%s", err)
		this.Redirect("/admin/login", 302)
	}
	log.Debug("here")

	sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	//sess.Set(C_COOKIE_NAME, sess.SessionID())

	this.Ctx.SetCookie(C_COOKIE_NAME, sess.SessionID())
	this.Redirect("/admin/add", 302)
}
