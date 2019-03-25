package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 管理员权限
type ManagerAuth struct {
	Id     int64  `orm:"auto;pk"`
	Name   string `orm:"size(100)"`
	Pid    int64 `orm:"size(32);"`
	Ac     string `orm:"size(255)"`
	Lv     int8
}

//初始化模型
func init() {
	dbprefix := beego.AppConfig.String("dbprefix")
	orm.RegisterModelWithPrefix(dbprefix, new(ManagerAuth))
}
