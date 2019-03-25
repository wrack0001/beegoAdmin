package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 用户
type User struct {
	Id         int64  `orm:"auto"`
	Name       string `orm:"size(100)"`
	NickName   string `orm:"size(100)"`
	Mobile     string `orm:"size(20)"`
	Email      string `orm:"size(100)"`
	Password   string `orm:"size(32)"`
	Salt       string `orm:"size(6)"`
	Age        int8
	Sex        int8
	Roleid     int64
	IsDel      int32
	CreateTime string
	UpdateTime string
}

//初始化模型
func init() {
	dbprefix := beego.AppConfig.String("dbprefix")
	orm.RegisterModelWithPrefix(dbprefix, new(User))
}


//根据用户分页数据列表
func LimitList(pagesize int, pageno int) (users []User) {

	o := orm.NewOrm()
	qs := o.QueryTable("user")

	var us []User
	cnt, err := qs.Limit(pagesize, (pageno-1)*pagesize).All(&us)
	if err == nil {
		fmt.Printf("count", cnt)
	}
	return us
}

//根据用户数据总个数
func GetDataNum() int64 {

	o := orm.NewOrm()
	qs := o.QueryTable("user")

	var us []User
	num, err := qs.Filter("id__gt", 0).All(&us)
	if err == nil {
		return num
	} else {
		return 0
	}
}
