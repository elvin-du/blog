package controllers

import "github.com/astaxie/beego"

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Index() {
	this.Layout = "layout.html"
	this.TplNames = "admin/login.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "admin/nav.html"
	this.LayoutSections["CSS"] = "admin/css.html"
	this.LayoutSections["JS"] = "admin/js.html"
}

//Method-POST; URL-/admin/login
func (this *AdminController) Login() {
	//name := this.GetString("name")
	//passwd := this.GetString("passwd")

	//cookie, err := AdminModel().Login(name, passwd)
	//if nil != err {
	//	beego.Error(err)
	//	return
	//}

	//sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	//defer sess.SessionRelease()
	//sess.Set()

	//this.Ctx.SetCookie(C_COOKIE_NAME, cookie)
	//this.Redirect("/admin", C_HTTP_OK)
}
