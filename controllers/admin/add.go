package admin

import ()

type AddController struct {
	baseController
}

func (this *AddController) Index() {
	log.Debug("AddController:index()")

	this.Layout = "layout.html"
	this.TplNames = "admin/add.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Nav"] = "admin/nav.html"
	this.LayoutSections["CSS"] = "admin/css.html"
	this.LayoutSections["JS"] = "admin/js.html"
}
