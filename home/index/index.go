package index

import (
	"apiAdmin/admin/base"
)

type IndexController struct {
	base.BaseController
}


func (c *IndexController) Get() {
	c.Ctx.WriteString("home-index-get")
}

func (c *IndexController) Post()  {

	c.Ctx.WriteString("home-index-post")

}
