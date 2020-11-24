package goodsInfoException

import "grpc-demo/models"

func SpecificationMarshalFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5407,
		Message: "规格序列化失败",
	}
}

func SpecificationUnMarshalFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5408,
		Message: "规格反序列化失败",
	}
}

func SpecificationParamsError() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5409,
		Message: "规格参数错误",
	}
}

func DeleteSpecificationFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "商品规格删除失败",
	}
}

func SpecificationIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5411,
		Message: "商品规格不存在",
	}
}

func SpecificationCreateError() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "商品规格创建失败",
	}
}

func SpecificationPutError() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5413,
		Message: "商品规格修改失败",
	}
}

func SpecificationExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5414,
		Message: "商品规格已存在",
	}
}