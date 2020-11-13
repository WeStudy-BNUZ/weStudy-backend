package middlewares

import (
	"fmt"
	"github.com/kataras/iris"
	"my_demo/models"
	logUtils "my_demo/utils/log"
	"runtime"
)

func AbnormalHandle(ctx iris.Context){
	defer func(){
		re := recover()
		if re == nil{
			return
		}
		ctx.StatusCode(iris.StatusInternalServerError)

		log := fmt.Sprintf("%v\n%s",re,stack())
		if debug,err := ctx.URLParamInt("debug");err ==nil && debug == 1{
			ctx.Text(log)
			return
		}

		switch result := re.(type){
		case models.RestfulAPIResult:
			ctx.JSON(result)
		default:
			ctx.JSON(models.RestfulAPIResult{
				Status:  false,
				ErrCode: 500,
				Message: fmt.Sprintf("系统错误：%v",result),
			})
			logUtils.Println(log)
		}
	}()

	ctx.Next()
}


//todo 打印堆栈信息
func stack() string{
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], true)])
}
