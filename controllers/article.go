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

func (this *ArticleController) Tag() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusBadRequest)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	as, err := models.ArticlesModel().ArticlesFromTagId(id)
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusBadRequest)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	this.TplNames = "main/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "main/nav.html"

	this.Data["AS"] = as
	this.Data["total"] = 1
	this.Data["cur_page"] = 1
	this.Data["pre_page"] = 0
	this.Data["next_page"] = 0
	this.Data["pagination"] = ""
}

func (this *ArticleController) Search() {
	key := this.GetString("key")
	beego.Debug(key)
	as, err := models.ArticlesModel().Search(key)
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusBadRequest)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	this.TplNames = "main/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "main/nav.html"

	this.Data["AS"] = as
	this.Data["total"] = 1
	this.Data["cur_page"] = 1
	this.Data["pre_page"] = 0
	this.Data["next_page"] = 0
	this.Data["pagination"] = ""
}
