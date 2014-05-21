package main

import (
	"blog/utils"

	"github.com/astaxie/beego"
)

func init() {
	beego.SetLogFuncCall(true)
	beego.SetLogger("file", `{"filename":"blog.log"}`)
	beego.AddFuncMap("excerpt", utils.Excerpt)
	beego.AddFuncMap("ymd", utils.YYYYMMDD)

}
