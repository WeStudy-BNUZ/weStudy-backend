package view

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"my_demo/view/test"
)

func RegisterTestRouters(app *iris.Application){
	testRouter := app.Party("/test")

	testRouter.Get("", hero.Handler(test.Test))
	testRouter.Get("/redis", hero.Handler(test.TestRedis))
}
