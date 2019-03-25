package common

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
)

var Redis cache.Cache

func init() {
	var err error
	Redis, err = cache.NewCache("redis", `{"key":"user","conn":":6379","dbNum":"0"}`)

	if err != nil {
		logs.Error(err)
	}

}
