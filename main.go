package main

import (
	_ "apiAdmin/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis" // 使用redis 存储session 必须引入这个方法
)

func init() {

	dbtype := beego.AppConfig.String("dbtype")
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")

	err := orm.RegisterDriver(dbtype, orm.DRMySQL)
	if err != nil {
		panic(err)
	}

	err = orm.RegisterDataBase("default", dbtype, dbuser+":"+dbpassword+"@tcp("+dbhost+":"+dbport+")/"+dbname+"?charset=utf8&loc=Asia%2FShanghai")
	if err != nil {
		panic(err)
	}
}

func main() {
	beego.Run()
}
