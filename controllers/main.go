package controllers

import (
	"blog/models"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
)

const (
	C_PAGE_NUM = 10
)

type MainController struct {
	baseController
}

func (this *MainController) Index() {
	this.TplNames = "main/index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "main/nav.html"
	//this.LayoutSections["CSS"] = "main/css.html"

	pageStr := this.Ctx.Input.Param(":page")
	var page int = 1
	var err error
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if nil != err {
			beego.Error(err)
			page = 1
		}
	}

	as, total, err := models.ArticlesModel().Articles(C_PAGE_NUM, int64(page))
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusInternalServerError)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}

	totalArray := make([]int, total)
	for i, _ := range totalArray {
		totalArray[i] = i + 1
	}

	this.Data["AS"] = as
	this.Data["total"] = totalArray
	this.Data["cur_page"] = page
	this.Data["pre_page"] = int(page - 1)
	this.Data["next_page"] = int(page + 1)
}
