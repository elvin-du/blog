package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type CommentController struct {
	baseController
}

func (this *CommentController) Add() {
	comment := this.GetString("comment")
	nick := this.GetString("nick")
	site := this.GetString("site")
	email := this.GetString("email")
	id := this.Ctx.Input.Param(":id")
	ip := this.Ctx.Input.IP()

	err := models.CommentModel().Add(id, comment, ip, nick, site, email)
	if nil != err {
		beego.Error(err)
		this.Abort("insert comment failed")
		return
	}

	url := "/article/" + id
	this.Redirect(url, 302)
}
