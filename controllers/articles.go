package controllers

import (
	"strconv"

	"blog/models"
	"github.com/astaxie/beego"
)

type ArticlesController struct {
	beego.Controller
}

func (this *ArticlesController) Index() {
	this.Layout = "layout.html"
	this.TplNames = "articles/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["CSS"] = ""
	this.LayoutSections["JS"] = ""
	this.LayoutSections["Nav"] = "articles/nav.html"

	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(C_HTTP_BAD_REQUEST)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	as, err := models.ArticlesModel().ArticleFromId(id)
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(C_HTTP_INTERNAL_ERROR)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	if len(as) != 0 {
		this.Data["article"] = as[0]
		this.Data["comments"] = as
	}
}
