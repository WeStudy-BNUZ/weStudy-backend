package classify

import (
	"github.com/kataras/iris"
	classifyException "my_demo/exceptions/classify"
	"my_demo/models/db"
	paramsUtils "my_demo/utils/params"
)

//目前没写登陆，先不卡权限
func CreateCollage(ctx iris.Context){
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	name := params.Str("name","名称")

	collage := db.Collage{
		Name: name,
	}

	db.Driver.Create(&collage)
	ctx.JSON(iris.Map{
		"id":collage.ID,
	})

}

func PutCollage(ctx iris.Context,cid int){
	var collage db.Collage
	if err := db.Driver.GetOne("collage",cid,&collage);err != nil{
		panic(classifyException.CollageIsNotExsit())
	}

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	params.Diff(&collage)
	collage.Name = params.Str("name","名称")

	db.Driver.Save(&collage)

	ctx.JSON(iris.Map{
		"id":collage.ID,
	})
}

func DeleteCollage(ctx iris.Context,cid int){
	var collage db.Collage
	if err := db.Driver.GetOne("collage",cid,&collage);err == nil{

		//级联删除
		if err := db.Driver.Exec("delete from major where collage_id = ?",cid).Error;err == nil{
			db.Driver.Delete(collage)
		}
	}

	ctx.JSON(iris.Map{
		"id":cid,
	})
}

func ListCollage(ctx iris.Context){
	var lists []struct {
		Id         int   `json:"id"`
		Name string `json:"name"`
	}
	var count int

	table := db.Driver.Table("collage")

	limit := ctx.URLParamIntDefault("limit", 10)
	page := ctx.URLParamIntDefault("page", 1)

	table.Count(&count).Offset((page - 1) * limit).Limit(limit).Select("id, name").Find(&lists)
	ctx.JSON(iris.Map{
		"collages": lists,
		"total": count,
		"limit": limit,
		"page":  page,
	})
}

func MgetCollage(ctx iris.Context){
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	ids := params.List("ids", "id列表")

	data := make([]interface{}, 0, len(ids))
	collages := db.Driver.GetMany("collage", ids, db.Collage{})
	for _,collage  := range collages {
		func(data *[]interface{}) {
			*data = append(*data, paramsUtils.ModelToDict(collage,[]string{"ID","Name"}))
			defer func() {
				recover()
			}()
		}(&data)
	}
	ctx.JSON(data)
}
