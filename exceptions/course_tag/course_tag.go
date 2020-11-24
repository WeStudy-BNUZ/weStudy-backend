package courseTagException

import "grpc-demo/models"

func CourseTagIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "课程标签不存在",
	}
}

func CourseKindIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "课程类型不存在",
	}
}