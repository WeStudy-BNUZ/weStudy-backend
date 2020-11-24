package globalException

import "grpc-demo/models"

func IconIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "图标不存在",
	}
}
