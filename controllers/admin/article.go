package admin

import (
	"blog/models"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	baseController
}

func (this *ArticleController) AddView() {
	this.TplNames = "admin/add.html"
}

func (this *ArticleController) Add() {
	title := this.GetString("title")
	content := this.GetString("content")

	err := models.ArticlesModel().Add(title, content)
	if nil != err {
		this.Abort("insert blog failed")
	}

	this.Redirect("/", 302)
}

func (this *ArticleController) Del() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if nil != err {
		beego.Error(err)
		return //todo
	}

	err = models.ArticlesModel().Del(id)
	if nil != err {
		beego.Error(err) //todo
	}

	this.Redirect("/", 302)
}

func (this *ArticleController) EditView() {
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
		this.TplNames = "admin/edit.html"
		this.Data["article"] = as[0]
		this.Data["id"] = id
	}
}

func (this *ArticleController) Edit() {
	title := this.GetString("title")
	content := this.GetString("content")
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusBadRequest)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	err = models.ArticlesModel().Update(id, title, content)
	if nil != err {
		this.Abort("update blog failed")
	}

	this.Redirect("/", 302)
}
