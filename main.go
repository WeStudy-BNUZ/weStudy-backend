package main

import (
	"github.com/kataras/iris"
	"my_demo/middlewares"
	"my_demo/view"
)

func initRouter(app *iris.Application){
	view.RegisterTestRouters(app)
}

func main(){
	app := iris.New()

	app.UseGlobal(middlewares.AbnormalHandle,middlewares.RequestLogHandle)
	initRouter(app)

	//app.Get("/get_request", func(ctx iris.Context){
	//	app.Logger().Info(ctx.Path())
	//	//ctx.WriteString(ctx.Path())
	//
	//	//user := ctx.URLParam("user")
	//	//app.Logger().Info(user)
	//	//ctx.HTML("<h1>"+user+"</h1>")
	//	ctx.JSON(iris.Map{"message":"hello","requestcode":200})
	//})
	//
	//app.Post("/post_login", func(ctx iris.Context){
	//	app.Logger().Info(ctx.Path())
	//
	//	//user := ctx.PostValue("user")
	//	//ctx.WriteString(user)
	//	var person Person
	//	if err := ctx.ReadJSON(&person);err != nil{
	//		panic(err.Error())
	//	}
	//
	//	ctx.Writef("%#+v\n",person)
	//
	//})

	app.Run(iris.Addr(":8000"),iris.WithoutServerError(iris.ErrServerClosed))
}

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
