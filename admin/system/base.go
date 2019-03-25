package system

import (
	"apiAdmin/admin/base"
)

type BaseController struct {
	base.BaseController
}

func (this *BaseController) Get() {
	this.TplName = "admin/system/base.html"
}

func (this *BaseController) Post()  {

}

// @router //admin/system/base_show/:key [get]
func (this *BaseController) show()  {

}