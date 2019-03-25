package system

import (
	"apiAdmin/admin/base"
)

type LogController struct {
	base.BaseController
}

func (this *LogController) Get() {
	this.TplName = "admin/system/log.html"
}

func (this *LogController) Post()  {

}

// @router //admin/system/log_show/:key [get]
func (this *LogController) show()  {

}