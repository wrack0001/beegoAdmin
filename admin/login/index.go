package login

import (
	"apiAdmin/common"
	"apiAdmin/lib"
	"apiAdmin/models"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/captcha"
	"math"
	"strconv"
	"time"
)

type IndexController struct {
	beego.Controller
}

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4 // 设置位数
	cpt.StdWidth = 150    // 设置宽度
}

var cpt *captcha.Captcha

func (this *IndexController) Get() {

	this.TplName = "admin/login/login.html"

}

func (this *IndexController) Post() {

	// 验证 - 验证码
	if cpt.VerifyReq(this.Ctx.Request) == false { // 验证验证码
		this.Data["json"] = common.JsonFormat{Code: 405, Msg: "验证码错误"}
		this.ServeJSON()
	}

	user_name := this.GetString("user_name")
	password := this.GetString("password")
	online,_ := this.GetBool("online")
	o := orm.NewOrm()
	m := &models.Manager{Mobile: user_name}

	err := o.Read(m, "mobile")

	if err != nil {
		this.Data["json"] = common.JsonFormat{Code: 405, Msg: "用户名不存在"}
		this.ServeJSON()
	}

	password = fmt.Sprintf("%x", md5.Sum([]byte(password+m.Salt)))

	if m.Pwd != password {
		this.Data["json"] = common.JsonFormat{Code: 405, Msg: "密码错误"}
		this.ServeJSON()
	}

	err = this.StartSession().Set("Mid", m.Id)
	if err != nil {
		this.Data["json"] = common.JsonFormat{Code: 405, Msg: "登录失败"}
		this.ServeJSON()
	}

	date_time := time.Now().Format("2006-01-02 15:04:05")

	// 生成登录token
	m.Token = lib.Md5(strconv.FormatInt(m.Id, 10) + date_time)

	_, err = o.Update(m,"token")

	if online{
		this.Ctx.SetCookie("token", m.Token, math.MaxInt32)
	}


	// 将登陆信息插入 UserLog 表
	_,_ = o.Insert(&models.MemberLoginLog{
		MId: m.Id, Ip: this.Ctx.Input.IP(), CreateTime: date_time,
	})

	this.Data["json"] = common.JsonFormat{Code: 200, Msg: "成功", Data: nil}

	this.ServeJSON()

}

/**
 *	退出登录
**/
func (this *IndexController) OutLogin() {

	err := this.StartSession().Flush()
	if err == nil {
		url := beego.URLFor("login.IndexController.Get")
		this.Redirect(url, 302)
	}
}
