package base

import (
	"apiAdmin/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
)

var (
	M models.Manager
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Data["logion_user_name"] = M.Name
	this.Data["tole_name"] = M.ManagerRole.Name
}

func init() {

}

type JsonFormat struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

/**
	验证是否登录
 */
func CheckLogin(ctx *context.Context) {

	var ok bool
	M.Id, ok = ctx.Input.Session("Mid").(int64)

	if !ok {
		ctx.Redirect(302, "/login")
	}
	// 通过用户ID 查询用户信息
	o := orm.NewOrm()
	err := o.Read(&M)

	if err != nil {
		beego.Error(err)
		ctx.Redirect(302, "/login")
	}

	M.ManagerRole = &models.ManagerRole{Id:M.RId}

	err = o.Read(M.ManagerRole)

	if err != nil {
		beego.Error(err)
		ctx.Redirect(302, "/login")
	}

}

// 是否POST提交
func (self *BaseController) IsPost() bool {
	return self.Ctx.Request.Method == "POST"
}
