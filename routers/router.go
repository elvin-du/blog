package routers

import (
	. "blog/controllers"
	"blog/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &MainController{}, "GET:Index")
	beego.Router("/about", &AboutController{}, "GET:Index")
	beego.Router("/contact", &ContactController{}, "GET:Index")
	beego.Router("/article/:id:int", &ArticleController{}, "GET:Index")
	beego.Router("/admin/login", &admin.LoginController{}, "GET:Index;POST:Login")
	beego.Router("/admin/add", &admin.AddController{}, "GET:Index;POST:Add")
	//beego.Router("/tags/:id:int", &MainController{})
	//beego.Router("/conmments/:id:int", &MainController{})
	//beego.Router("/date/:id:int", &MainController{})
	//beego.Router("/admin", &MainController{})

	//static file
	//beego.SetStaticPath("/ico/favicon.ico", "/static/ico/favicon.ico")
	//beego.SetStaticPath("/css", "/static/css")
	//beego.SetStaticPath("/js", "/static/js")
	//beego.SetStaticPath("/img", "/static/img")
}
