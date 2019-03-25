package index

import (
	"apiAdmin/admin/base"
	"apiAdmin/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type IndexController struct {
	base.BaseController
}

func (c *IndexController) Get() {
	user_log, _ := models.UserLogLastInfo(base.M.Id)
	c.Data["ip"] = user_log.Ip
	c.Data["login_count"] = models.UserLogCount(base.M.Id)
	c.Data["login_time"] = user_log.CreateTime
	c.Data["now_time"] = time.Now().Format("2006-01-02 15:04:05") //当前时间
	c.Data["session_id"] = c.StartSession().SessionID()           // 获取sessionID
	c.Data["CPUnum"] = runtime.NumCPU()                           // 获取cpu数量
	c.Data["file_path"] = beego.AppPath
	c.Data["now_file_path"], _ = filepath.Abs(os.Args[0])
	c.Data["user_agent"] = c.Ctx.Input.Header("user-agent") // 获取登录浏览器信息
	//c.Data["user"] = base.U.Name		// 这里证明在这里设置的变量在，layout 中可以直接使用

	//c.TplName = "admin/index/index.html"
	c.TplName = "admin/index/index.html"
}

func (this *IndexController) Post() {
	jsoninfo := this.GetString("jsoninfo")
	this.Ctx.WriteString(jsoninfo)

	// 这个只能获取提交过来的数字类型数据。否则返回 0
	i9, _ := this.GetInt("is") // 如果get与post中同时有 is 变量，将获取 url 中的变量
	fmt.Println(i9)

	//this.Input() 这个函数是有多少变量获得多少同名的参数
	b := this.Input().Get("jsoninfo") // 获取第一个 同名变量

	data, _ := json.Marshal(b)
	this.Ctx.WriteString(string(data))
	if jsoninfo == "" {
		this.Ctx.WriteString("jsoninfo is empty")
		return
	}
	this.Ctx.WriteString(jsoninfo)

	return
}
