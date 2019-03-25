package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 管理员
type Manager struct {
	Id          int64  `orm:"auto;column(id);pk"`
	Mobile      string `orm:"size(20)"`
	Name        string `orm:"size(100)"`
	Salt        string `orm:"size(6)"`
	Pwd         string `orm:"size(32)"`
	RId         int64
	Token       string `orm:"size(32)"`
	CreateTime  string
	ManagerRole *ManagerRole `orm:"-"`
}

//初始化模型
func init() {
	dbprefix := beego.AppConfig.String("dbprefix")
	orm.RegisterModelWithPrefix(dbprefix, new(Manager))
}
