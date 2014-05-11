package admin

import "blog/models"

type AddController struct {
	baseController
}

func (this *AddController) Index() {
	this.TplNames = "admin/add.html"
}

func (this *AddController) Add() {
	title := this.GetString("title")
	content := this.GetString("content")

	err := models.ArticlesModel().Add(title, content)
	if nil != err {
		this.Abort("insert blog failed")
	}

	this.Redirect("/", 302)
}
