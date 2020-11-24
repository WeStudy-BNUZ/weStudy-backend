package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"my_demo/core/cache"
	"my_demo/middlewares"
	"my_demo/models/db"
	"my_demo/view"
)

func initRouter(app *iris.Application){
	view.RegisterClassifyRouters(app)
}

func main(){
	app := iris.New()

	app.UseGlobal(middlewares.AbnormalHandle,middlewares.RequestLogHandle)
	initRouter(app)

	db.InitDB()
	cache.InitRedisPool()

	app.Run(iris.Addr(":8088"),iris.WithoutServerError(iris.ErrServerClosed))
}

