package systemException

import (
	"my_demo/models"
)

func SystemException() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status: false,
		ErrCode: 5000,
		Message: "系统错误",
	}
}




