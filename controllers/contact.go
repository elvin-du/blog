package controllers

import (
	"blog/models"
	"net/http"

	"github.com/astaxie/beego"
)

type ContactController struct {
	baseController
}

func (this *ContactController) Index() {
	this.TplNames = "contact/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "contact/nav.html"

	email, err := models.InfoModel().InfoFromName("email")
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusBadRequest)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	nick, err := models.InfoModel().InfoFromName("nick")
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusBadRequest)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	this.Data["email"] = email[0]
	this.Data["nick"] = nick[0]
}
