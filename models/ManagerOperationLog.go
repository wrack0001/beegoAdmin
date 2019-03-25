package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 管理员权限
type ManagerOperationLog struct {
	Id         int64 `orm:"auto;column(id);pk"`
	MgId       int64
	Url        string `orm:"size(255)"`
	CreateTime string
}

//初始化模型
func init() {
	dbprefix := beego.AppConfig.String("dbprefix")
	orm.RegisterModelWithPrefix(dbprefix, new(ManagerOperationLog))
}
