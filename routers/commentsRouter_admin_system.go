package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["apiAdmin/admin/system:BaseController"] = append(beego.GlobalControllerRouter["apiAdmin/admin/system:BaseController"],
		beego.ControllerComments{
			Method: "show",
			Router: `//admin/system/base_show/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiAdmin/admin/system:CategoryController"] = append(beego.GlobalControllerRouter["apiAdmin/admin/system:CategoryController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `//admin/system/category_add/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiAdmin/admin/system:CategoryController"] = append(beego.GlobalControllerRouter["apiAdmin/admin/system:CategoryController"],
		beego.ControllerComments{
			Method: "Show",
			Router: `//admin/system/category_show/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiAdmin/admin/system:DataController"] = append(beego.GlobalControllerRouter["apiAdmin/admin/system:DataController"],
		beego.ControllerComments{
			Method: "show",
			Router: `//admin/system/data_show/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiAdmin/admin/system:LogController"] = append(beego.GlobalControllerRouter["apiAdmin/admin/system:LogController"],
		beego.ControllerComments{
			Method: "show",
			Router: `//admin/system/log_show/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["apiAdmin/admin/system:ShieldingController"] = append(beego.GlobalControllerRouter["apiAdmin/admin/system:ShieldingController"],
		beego.ControllerComments{
			Method: "show",
			Router: `//admin/system/shielding_show/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
