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
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(C_HTTP_BAD_REQUEST)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	a, err := models.ArticlesModel().ArticleFromId(id)
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(C_HTTP_INTERNAL_ERROR)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	this.Data["title"] = a.Title
	this.Data["content"] = a.Content
	this.Layout = "layout.html"
	this.TplNames = "articles/index.html"
}
