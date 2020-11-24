package goodsTagException

import "grpc-demo/models"

func PlaceTagIsNotExists() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5401,
		Message: "属地标签不存在",
	}
}

func SaleTagIsNotExists() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5402,
		Message: "销售标签不存在",
	}
}

func KindTagIsNotExists() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5403,
		Message: "种类标签不存在",
	}
}

func KindTagOverLevel() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5404,
		Message: "分类超过两层",
	}
}

func DeleteKindTagFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5405,
		Message: "种类标签删除失败",
	}
}

func KindTagMntError() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5406,
		Message: "不能挂载自身",
	}
}