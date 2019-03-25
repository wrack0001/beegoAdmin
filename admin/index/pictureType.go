package index

import (
	"fmt"
	"apiAdmin/admin/base"
	"apiAdmin/models"
)

type PictureTypeController struct {
	base.BaseController
}

// 列表
func (this *PictureTypeController) Get() {

	pic_type, err := models.PicTypeSelect()

	//fmt.Println(pic_type)

	if err == nil {
		TreePicType := models.GetParentList(&pic_type,0)
		fmt.Println(TreePicType)
	}

	this.TplName = "admin/picture_type/list.html"
}

// 修改
func (this *PictureTypeController) Show() {
	this.TplName = "admin/picture/show.html"
}

// 添加
func (this *PictureTypeController) Add() {
	if this.IsPost() {

	} else {
		this.TplName = "admin/picture_type/add.html"
	}

}
