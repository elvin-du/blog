package admin

import (
	"blog/models"
	"net/http"

	"github.com/astaxie/beego"
)

type TagController struct {
	baseController
}

func (this *TagController) Index() {
	tags, err := models.TagModel().Tags()
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusBadRequest)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	this.TplNames = "admin/tags.html"
	this.Data["tags"] = tags
}
