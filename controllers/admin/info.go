package admin

import (
	"blog/models"
	"net/http"

	"github.com/astaxie/beego"
)

type InfoController struct {
	baseController
}

func (this *InfoController) About() {
	about, err := models.InfoModel().InfoFromName("about")
	if nil != err {
		beego.Error(err)
		this.Ctx.Output.SetStatus(http.StatusBadRequest)
		this.Ctx.Output.Body([]byte(err.Error()))
		return
	}
	if len(about) > 0 {
		this.TplNames = "admin/about.html"
		this.Data["about"] = about[0]
	}
}

func (this *InfoController) UpdateAbout() {
	about := this.GetString("about")

	err := models.InfoModel().UdateAbout(about)
	if nil != err {
		this.Abort("update blog failed")
	}

	this.Redirect("/admin/about", 302)
}
