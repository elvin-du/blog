package routers

import (
	. "blog/controllers"
	"blog/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &MainController{}, "GET:Index")
	beego.Router("/article/page/:page:int", &MainController{}, "GET:Index")
	beego.Router("/about", &AboutController{}, "GET:Index")
	beego.Router("/contact", &ContactController{}, "GET:Index")
	beego.Router("/article/:id:int", &ArticleController{}, "GET:Index")
	beego.Router("/comment/:id:int", &CommentController{}, "POST:Add")

	//admin
	beego.Router("/admin/login", &admin.LoginController{}, "GET:Index;POST:Login")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "GET:AddView;POST:Add")
	beego.Router("/admin/article/del/:id:int", &admin.ArticleController{}, "GET:Del")
	beego.Router("/admin/article/edit/:id:int", &admin.ArticleController{}, "GET:EditView;POST:Edit")
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
