package system

import (
	"apiAdmin/admin/base"
)

type ShieldingController struct {
	base.BaseController
}

func (this *ShieldingController) Get() {
	this.TplName = "admin/system/shielding.html"
}

func (this *ShieldingController) Post()  {

}

// @router //admin/system/shielding_show/:key [get]
func (this *ShieldingController) show()  {

}