package goodsHomeException

import "grpc-demo/models"

func OrderExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "序号已存在",
	}
}

func SlideshowIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "该轮播图不存在",
	}
}
