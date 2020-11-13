package test

import "github.com/kataras/iris"

func Test(ctx iris.Context){
	ctx.WriteString("success!!!!")
}
