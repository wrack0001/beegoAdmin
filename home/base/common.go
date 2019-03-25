package base

import (
	"github.com/astaxie/beego/context"
)

func init() {

}

func CheckLogin(ctx *context.Context)  {
	_, ok := ctx.Input.Session("uid").(int)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}

