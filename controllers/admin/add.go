package admin

import "github.com/astaxie/beego"

type AddController struct {
	beego.Controller
	AuthController
}

func (this *AddController) Index() {
	beego.Debug("AddController:index()")
	err := this.Auth(this.Ctx)
	if nil != err {
		this.Redirect("/admin/login", 302)
	}

	this.Layout = "layout.html"
	this.TplNames = "admin/add.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "admin/nav.html"
	this.LayoutSections["CSS"] = "admin/css.html"
	this.LayoutSections["JS"] = "admin/js.html"
}
