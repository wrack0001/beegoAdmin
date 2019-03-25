package system

import (
	"apiAdmin/admin/base"
)

type CategoryController struct {
	base.BaseController
}

func (this *CategoryController) Get() {
	this.TplName = "admin/system/category.html"
}

func (this *CategoryController) Post()  {

}

// @router //admin/system/category_add/:key [get]
func (this *CategoryController) Add()  {
	this.TplName = "admin/system/category_add.html"
}

// @router //admin/system/category_show/:key [get]
func (this *CategoryController) Show()  {

}