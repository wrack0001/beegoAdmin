package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MemberLoginLog struct {
	Id         int64 `orm:"auto;"`
	MId        int64
	Ip         string `orm:"column(ip)"`
	CreateTime string
}

//初始化模型
func init() {
	dbprefix := beego.AppConfig.String("dbprefix")
	orm.RegisterModelWithPrefix(dbprefix, new(MemberLoginLog))
}

func (this *MemberLoginLog) Insert() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(this) // 插入数据
	return
}

/**
	获取上次用户登录信息
 */
func UserLogLastInfo(m_id int64) (log MemberLoginLog, err error) {
	o := orm.NewOrm()
	var user_log []MemberLoginLog
	num, err := o.Raw("SELECT * FROM beego_user_log where m_id = ?  order by id desc limit 2", m_id).QueryRows(&user_log)

	if err != nil {
		return
	}

	if num == 1 {
		log = user_log[0]
	} else {
		log = user_log[1]
	}
	return
}

/**
	统计用户登录次数(count)
 */
func UserLogCount(m_id int64) (cnt int64) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("b_member_login_log").Filter("m_id", m_id).Count() // SELECT COUNT(*) FROM USER
	fmt.Printf("Count Num: %s, %s", cnt, err)

	return
}
