package courseApplyException

import "grpc-demo/models"

func SessionIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "课程场次不存在",
	}
}

func PeopleOverLimit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "人数超过限制",
	}
}

func StatusApplyFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "此状态不能正式报名",
	}
}

func ApplyCreateFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "创建报名失败",
	}
}

func ApplyIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "报名不存在",
	}
}

func ApplyPutFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "此状态报名不能修改",
	}
}

func RelationCreateFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "关联信息创建失败",
	}
}

func PeopleError() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "人数不一致",
	}
}

func NeedNumber() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "需要身份证号码",
	}
}

