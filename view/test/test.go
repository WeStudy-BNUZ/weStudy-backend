package test

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kataras/iris"
	"github.com/shoogoome/mutils/hash"
	paramsUtils "github.com/shoogoome/mutils/params"
	"my_demo/constants"
	"my_demo/core/cache"
)

func Test(ctx iris.Context){
	ctx.WriteString("eeeeeeeee!!!!")
}

func TestRedis(ctx iris.Context){
	token := hash.GetRandomString(6)

	if _, err := redis.Bytes(cache.Redis.Do(
		constants.DbNumberOther,
		"set",
		paramsUtils.CacheBuildKey(constants.TeamTokenModel, token),
		"1",
		60*15)); err != nil {
		panic("111")
	}

	ctx.JSON(iris.Map{
		"token" :token,
	})
}


