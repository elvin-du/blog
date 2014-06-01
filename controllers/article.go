package controllers

import (
	"net/http"
	"strconv"

	"blog/models"
	"github.com/astaxie/beego"
)

type ArticleController struct {
	baseController
}

func (this *ArticleController) Index() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusBadRequest)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	//log read history
	err = models.ArticlesModel().LogReadHistory(id, this.Ctx.Input.IP())
	if nil != err {
		beego.Error(err)
	}

	this.TplNames = "article/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "article/nav.html"
	this.LayoutSections["JS"] = "article/js.html"
	this.Data["admin"] = true

	as, err := models.ArticlesModel().ArticleFromId(id)
	if nil != err {
		beego.Error(err)
		//this.Abort(err.Error())
		return
	}
	this.Data["article"] = as
	this.Data["id"] = id

	cs, err := models.CommentModel().CommentFromArticleId(id)
	if nil != err {
		beego.Error(err)
		//return
	}
	this.Data["comments"] = cs
}
