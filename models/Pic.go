package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 图片
type Pic struct {
	Id         int64 `orm:"auto;pk"`
	TypeId     int64
	Thumb      string `orm:"size(255)"`
	CreateTime string
	UpdateTime string
	IsDel      int64
}

//初始化模型
func init() {
	dbprefix := beego.AppConfig.String("dbprefix")
	orm.RegisterModelWithPrefix(dbprefix, new(Pic))
}
