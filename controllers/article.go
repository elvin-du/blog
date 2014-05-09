package controllers

import (
	"net/http"
	"strconv"

	"blog/models"
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Index() {
	this.Layout = "layout.html"
	this.TplNames = "article/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["CSS"] = ""
	this.LayoutSections["JS"] = ""
	this.LayoutSections["Nav"] = "article/nav.html"

	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusBadRequest)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	as, err := models.ArticlesModel().ArticleFromId(id)
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusInternalServerError)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	if len(as) != 0 {
		this.Data["article"] = as[0]
		this.Data["comments"] = as
	}
}
