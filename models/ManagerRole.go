package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 管理员角色
type ManagerRole struct {
	Id      int64  `orm:"auto;column(id);pk"`
	Name    string `orm:"size(30)"`
	AuthIds string
	AuthAc  string
}

//初始化模型
func init() {
	dbprefix := beego.AppConfig.String("dbprefix")
	orm.RegisterModelWithPrefix(dbprefix, new(ManagerRole))
}
