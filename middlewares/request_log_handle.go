package middlewares

import (
	"github.com/kataras/iris"
	logUtils "my_demo/utils/log"
)

func RequestLogHandle(ctx iris.Context){
	logUtils.Println("Host:", ctx.RemoteAddr(), "Method:", ctx.Method(), "Path:", ctx.Path())
	ctx.Next()
}
