package view

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	noMatter "my_demo/view/course"
)

func RegisterCourseRouters(app *iris.Application){

	courseRouter := app.Party("ocean/course")

	courseRouter.Post("", hero.Handler(noMatter.CreateCourse))
	courseRouter.Put("/{cid:int}", hero.Handler(noMatter.PutCourse))
	courseRouter.Delete("/{cid:int}", hero.Handler(noMatter.DeleteCourse))
	courseRouter.Get("/list", hero.Handler(noMatter.ListCourse))
	courseRouter.Post("/_mget", hero.Handler(noMatter.MgetCourse))
}
