package controllers

type DownloadController struct {
	baseController
}

func (this *DownloadController) Index() {
	path := this.Ctx.Input.Param(":path")
	ext := this.Ctx.Input.Param(":ext")
	file := "./static/" + path + "." + ext
	this.Ctx.Output.Download(file)
}
