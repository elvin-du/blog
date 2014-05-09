package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type AuthController struct{}

func (this *AuthController) Auth(ctx *context.Context) error {
	beego.Info(ctx.Input.Cookie(C_COOKIE_NAME))
	//beego.Info(ctx.Input.Session(C_COOKIE_NAME))

	return nil
}
