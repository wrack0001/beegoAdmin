package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/captcha"
)

type VerificationCode struct {
	beego.Controller
}


var cpt *captcha.Captcha

