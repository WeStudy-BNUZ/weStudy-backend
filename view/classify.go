package view

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"my_demo/view/classify"
)

func RegisterClassifyRouters(app *iris.Application){

	collageRouter := app.Party("study/collage")

	collageRouter.Post("", hero.Handler(classify.CreateCollage))
	collageRouter.Put("/{cid:int}", hero.Handler(classify.PutCollage))
	collageRouter.Delete("/{cid:int}", hero.Handler(classify.DeleteCollage))
	collageRouter.Get("/list", hero.Handler(classify.ListCollage))
	collageRouter.Post("/_mget", hero.Handler(classify.MgetCollage))

	majorRouter := app.Party("study/major")
	majorRouter.Post("", hero.Handler(classify.CreateMajor))
	majorRouter.Put("/{mid:int}", hero.Handler(classify.PutMajor))
	majorRouter.Delete("/{mid:int}", hero.Handler(classify.DeleteMajor))
	majorRouter.Get("/list", hero.Handler(classify.ListMajor))
	majorRouter.Post("/_mget", hero.Handler(classify.MgetMajor))
}
