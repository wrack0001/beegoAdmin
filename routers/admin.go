package routers

import (
	"apiAdmin/admin/system"
	"github.com/astaxie/beego"
	"apiAdmin/admin/base"
	"apiAdmin/admin/index"
	"apiAdmin/admin/login"

)


func init() {
	beego.InsertFilter("/admin/*", beego.BeforeRouter, base.CheckLogin)	// 中间件
	beego.Router("/login", &login.IndexController{})

	beego.Router("/admin/outlogin",&login.IndexController{},"get:OutLogin")
    beego.Router("/admin", &index.IndexController{})
	beego.Router("/admin/picture",&index.PictureController{})
	beego.Router("/admin/picture/show/:id",&index.PictureController{},"get:Show")
	beego.Router("/admin/picture/add",&index.PictureController{},"get,post:Add")
	beego.Router("/admin/picture_type/list",&index.PictureTypeController{})
	beego.Router("/admin/picture_type/add",&index.PictureTypeController{},"get,post:Add")

	beego.Router("/admin/system/base",&system.BaseController{})

	beego.Router("/admin/system/category",&system.CategoryController{})
	beego.Router("/admin/system/data",&system.DataController{})
	beego.Router("/admin/system/shielding",&system.ShieldingController{})
	beego.Router("/admin/system/log",&system.LogController{})


}