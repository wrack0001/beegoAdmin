package routers

import (
	"github.com/astaxie/beego"
	"apiAdmin/home/base"
	"apiAdmin/home/index"
)


func init() {
	beego.InsertFilter("/user/*", beego.BeforeRouter, base.CheckLogin)
	//注册自动路由
	beego.AutoRouter(&index.IndexController{})

}

