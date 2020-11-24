package courseApplyException

import "grpc-demo/models"

func PreApplyIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "预报名不存在",
	}
}

func StatusPutFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "此状态不能修改",
	}
}
