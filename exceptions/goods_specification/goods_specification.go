package goodsSpecificationException

import "grpc-demo/models"

func TemplateMarshalFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5407,
		Message: "模版序列化失败",
	}
}

func TemplateUnMarshalFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5408,
		Message: "模版反序列化失败",
	}
}

func TemplateIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5409,
		Message: "规格模版不存在",
	}
}

func TemplateUsed() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "规格模版被使用",
	}
}
