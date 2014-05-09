package controllers

import (
	"blog/models"
	"net/http"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Index() {
	this.Layout = "layout.html"
	this.TplNames = "main/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "main/nav.html"

	as, err := models.ArticlesModel().Articles()
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusInternalServerError)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	this.Data["AS"] = as
}
