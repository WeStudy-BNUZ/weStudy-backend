package goodsInfoException

import "grpc-demo/models"

func AdvanceTimeError() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "预售时间错误",
	}
}

func GetWayIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5411,
		Message: "收货方式不存在",
	}
}

func PutawayIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "上架方式不存在",
	}
}

func PutawayTimeError() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5413,
		Message: "上架时间错误",
	}
}

func GoodsIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5414,
		Message: "商品不存在",
	}
}

func OnSaleFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5414,
		Message: "上架失败",
	}
}

func PictureMarshalFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5415,
		Message: "图片列表序列化失败",
	}
}

func PictureUnmarshalFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5416,
		Message: "图片列表反序列化失败",
	}
}