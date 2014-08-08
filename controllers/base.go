package controllers

import (
	"blog/models"
	"blog/utils"
	"time"

	"github.com/astaxie/beego"
)

type baseController struct {
	controllerName string
	actionName     string
	beego.Controller
}

func (this *baseController) Prepare() {
	this.Layout = "layout.html"
	this.VisitorCount()
	this.auth()
	this.HotArticles()
	this.LatestArticles()
	this.LatestComments()
	this.Tags()
}

func (this *baseController) HotArticles() {
	hotAS, err := models.ArticlesModel().Hot()
	if nil != err {
		beego.Error(err)
		this.Abort(err.Error())
		return
	}
	this.Data["hot_articles"] = hotAS
}

func (this *baseController) LatestArticles() {
	latestAS, err := models.ArticlesModel().Latest()
	if nil != err {
		beego.Error(err)
		this.Abort(err.Error())
		return
	}
	this.Data["latest_articles"] = latestAS
}

func (this *baseController) LatestComments() {
	latestCS, err := models.CommentModel().Latest()
	if nil != err {
		beego.Error(err)
		this.Abort(err.Error())
		return
	}
	this.Data["latest_comments"] = latestCS
}

func (this *baseController) Tags() {
	tags, err := models.TagModel().Tags()
	if nil != err {
		beego.Error(err)
		this.Abort(err.Error())
		return
	}

	for i, _ := range tags {
		tags[i]["font_size"] = utils.Rand(8, 22)
	}
	this.Data["tags"] = tags
}

func (this *baseController) VisitorCount() {
	sess := this.GetSession("macs")
	if nil == sess {
		err := models.VisitorModel().Log(this.Ctx.Input.IP(), time.Now())
		if nil != err {
			beego.Error(err)
			this.Abort(err.Error())
		}
		this.SetSession("macs", "du")
	}
}

func (this *baseController) auth() {
	if this.validSess() {
		this.Data["admin"] = true
	} else {
		this.Data["admin"] = false
		beego.Debug("false")
	}
}

func (this *baseController) validSess() bool {
	cookie := this.Ctx.GetCookie(C_COOKIE_NAME)
	return models.AdminModel().ValidSession(cookie)
}
