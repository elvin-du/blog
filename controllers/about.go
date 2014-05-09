package controllers

import (
	"blog/models"
	"net/http"

	"github.com/astaxie/beego"
)

type AboutController struct {
	beego.Controller
}

func (this *AboutController) Index() {
	this.Layout = "layout.html"
	this.TplNames = "about/index.html"
	this.LayoutSections = make(map[string]string)
	//this.LayoutSections["CSS"] = "about/css.html"
	this.LayoutSections["Nav"] = "about/nav.html"

	about, err := models.InfoModel().InfoFromName("about")
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusBadRequest)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	this.Data["about"] = about[0]
}
