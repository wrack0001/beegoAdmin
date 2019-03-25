package system

import (
	"apiAdmin/admin/base"
)

type DataController struct {
	base.BaseController
}

func (this *DataController) Get() {
	this.TplName = "admin/system/data.html"
}

func (this *DataController) Post()  {

}

// @router //admin/system/data_show/:key [get]
func (this *DataController) show()  {

}