package main

import "github.com/astaxie/beego"

func init() {
	beego.SetLogFuncCall(true)
	beego.SetLogger("file", `{"filename":"blog.log"}`)
}
