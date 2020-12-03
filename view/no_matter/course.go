package noMatter

import (
	"github.com/kataras/iris"
	"my_demo/models/db"
	paramsUtils "my_demo/utils/params"
)

func CreateCourse(ctx iris.Context){
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	name := params.Str("name","名称")
	kind := params.Int("kind","kind")
	class := params.Int("class","class")

	course := db.Course{
		Name: name,
		Kind:kind,
		Class:class,
	}

	db.Driver.Create(&course)
	ctx.JSON(iris.Map{
		"id":course.ID,
	})

}

func PutCourse(ctx iris.Context,cid int){
	var course db.Course
	if err := db.Driver.GetOne("course",cid,&course);err != nil{
		panic("课程不存在")
	}

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	params.Diff(&course)
	course.Name = params.Str("name","名称")
	course.Kind = params.Int("kind","kind")
	course.Class = params.Int("class","class")

	db.Driver.Save(&course)

	ctx.JSON(iris.Map{
		"id":course.ID,
	})
}

func DeleteCourse(ctx iris.Context,cid int){
	var course db.Course
	if err := db.Driver.GetOne("course",cid,&course);err == nil{


			db.Driver.Delete(course)

	}

	ctx.JSON(iris.Map{
		"id":cid,
	})
}

func ListCourse(ctx iris.Context){
	var lists []struct {
		Id         int   `json:"id"`
		Name string `json:"name"`
	}
	var count int

	table := db.Driver.Table("course")

	//todo
	limit := ctx.URLParamIntDefault("limit", 10)
	page := ctx.URLParamIntDefault("page", 1)

	table.Count(&count).Offset((page - 1) * limit).Limit(limit).Select("id, name").Find(&lists)
	ctx.JSON(iris.Map{
		"courses": lists,
		"total": count,
		"limit": limit,
		"page":  page,
	})
}

func MgetCourse(ctx iris.Context){
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	ids := params.List("ids", "id列表")

	data := make([]interface{}, 0, len(ids))
	courses := db.Driver.GetMany("course", ids, db.Course{})
	for _,course  := range courses {
		func(data *[]interface{}) {
			*data = append(*data, paramsUtils.ModelToDict(course,[]string{"ID","Name","Kind","Class"}))
			defer func() {
				recover()
			}()
		}(&data)
	}
	ctx.JSON(data)
}

