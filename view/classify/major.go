package classify

import (
	"github.com/kataras/iris"
	classifyException "my_demo/exceptions/classify"
	"my_demo/models/db"
	paramsUtils "my_demo/utils/params"
)

//目前没写登陆，先不卡权限
func CreateMajor(ctx iris.Context,cid int){
	var collage db.College
	if err := db.Driver.GetOne("collage",cid,&collage);err == nil{
		db.Driver.Delete(collage)
	}

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	name := params.Str("name","名称")

	major := db.Major{
		Name: name,
		CollegeID:cid,
	}

	db.Driver.Create(&major)
	ctx.JSON(iris.Map{
		"id":major.ID,
	})

}

func PutMajor(ctx iris.Context,cid int,mid int){
	var major db.Major
	if err := db.Driver.Where("id = ? and collage_id = ?",mid,cid).First(&major).Error;err != nil{
		panic(classifyException.MajorIsNotExsit())
	}

	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	params.Diff(&major)
	major.Name = params.Str("name","名称")

	db.Driver.Save(&major)

	ctx.JSON(iris.Map{
		"id":major.ID,
	})
}

func DeleteMajor(ctx iris.Context,cid int,mid int){
	var major db.Major
	if err := db.Driver.Where("id = ? and collage_id = ?",mid,cid).First(&major).Error;err == nil{
		db.Driver.Delete(major)
	}

	ctx.JSON(iris.Map{
		"id":mid,
	})
}

func ListMajor(ctx iris.Context){
	var lists []struct {
		Id         int   `json:"id"`
		Name string `json:"name"`
	}
	var count int

	table := db.Driver.Table("major")

	limit := ctx.URLParamIntDefault("limit", 10)
	page := ctx.URLParamIntDefault("page", 1)

	//学院过滤
	if collageID := ctx.URLParamIntDefault("collage_id", 0); collageID != 0 {
		table = table.Where("collage_id = ?", collageID)
	}

	table.Count(&count).Offset((page - 1) * limit).Limit(limit).Select("id, name").Find(&lists)
	ctx.JSON(iris.Map{
		"majors": lists,
		"total": count,
		"limit": limit,
		"page":  page,
	})
}

func MgetMajor(ctx iris.Context){
	params := paramsUtils.NewParamsParser(paramsUtils.RequestJsonInterface(ctx))
	ids := params.List("ids", "id列表")

	data := make([]interface{}, 0, len(ids))
	majors := db.Driver.GetMany("major", ids, db.Major{})
	for _,major  := range majors {
		func(data *[]interface{}) {
			*data = append(*data, paramsUtils.ModelToDict(major,[]string{"ID","Name","CollageID"}))
			defer func() {
				recover()
			}()
		}(&data)
	}
	ctx.JSON(data)
}

