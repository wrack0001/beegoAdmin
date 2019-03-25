package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type PicType struct {
	Id         int64  `orm:"auto"`
	Pid        int64
	Title      string `orm:"size(50)"`
	IsDel      int64
	CreateTime string
	UpdateTime string
}

type TreePicType struct {
	Id    int64
	Title string
	Pid   int64
	Child []TreePicType
}

//初始化模型
func init() {

	dbprefix := beego.AppConfig.String("dbprefix")
	orm.RegisterModelWithPrefix(dbprefix, new(PicType))
}

// 插入数据
func (this *PicType) PicTypeInsert() (id int64, err error) {

	this.IsDel = 0
	id, err = orm.NewOrm().Insert(this)

	return
}

// 查找所有数据
func PicTypeSelect() (pic_type []PicType, err error) {
	PicType := new(PicType)
	// todo  这里的  isDel 条件有问题

	_, err = orm.NewOrm().QueryTable(PicType).Filter("IsDel", 0).All(&pic_type)
	fmt.Println(err)
	if err != nil {

	}
	return
}

// 获取无限极分类
func GetParentList(arr *[]PicType, pid int64) (l []TreePicType) {

	for _, v := range *arr {
		if v.Pid == pid {
			l = append(l, TreePicType{
				Id:    v.Id,
				Title: v.Title,
				Pid:   v.Pid,
				Child: GetParentList(arr, v.Pid),
			})
		}
	}

	return;
}
