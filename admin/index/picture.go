package index

import (
	"apiAdmin/admin/base"
	"apiAdmin/models"
)

type PictureController struct {
	base.BaseController
}
// 列表
func (this *PictureController) Get() {
	this.TplName = "admin/picture/list.html"
}
// 修改
func (this *PictureController) Show() {
	this.TplName = "admin/picture/show.html"
}
// 添加
func (this *PictureController) Add() {
	if this.IsPost(){

	}else {
		this.Data["pic_type_list"],_ = models.PicTypeSelect()
		this.TplName = "admin/picture/add.html"
	}


}