package orderException

import "grpc-demo/models"

func AddressIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "地址不存在",
	}
}

func OrderCreateFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "订单创建失败",
	}
}

func OrderDeleteFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "订单删除失败",
	}
}

func OrderMarshalFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5415,
		Message: "订单序列化失败",
	}
}

func OrderUnmarshalFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5416,
		Message: "订单反序列化失败",
	}
}

func OrderIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "订单不存在",
	}
}

func OrderPutFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "订单修改失败",
	}
}

func OrderPutStatusFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "此状态订单不能修改",
	}
}

func OrderDeleteStatusFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "此状态订单不能删除",
	}
}

func TotalNoEnough() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "库存不足",
	}
}

func TotalPutFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "库存修改失败",
	}
}

func ChildOrderIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "子订单不存在",
	}
}

func OrderNotCancel() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "此订单不能取消",
	}
}

func OrderCancelFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "订单取消失败",
	}
}

func OrderCancelNoPut() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "已取消订单不能修改",
	}
}

func OrderSelfGetFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5410,
		Message: "自动收货失败",
	}
}


