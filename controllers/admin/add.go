package admin

import (
	"net/http"

	"github.com/astaxie/beego"
)

type AddController struct {
	beego.Controller
	AuthController
}

func (this *AddController) Index() {
	err := this.Auth(this.Ctx)
	if nil != err {
		this.Redirect("/admin/login", http.StatusOK)
		return
	}

	this.Layout = "layout.html"
	this.TplNames = "admin/add.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "admin/nav.html"
	this.LayoutSections["CSS"] = "admin/css.html"
	this.LayoutSections["JS"] = "admin/js.html"
}
