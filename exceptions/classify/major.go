package classifyException

import "my_demo/models"

func MajorIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "专业类别不存在",
	}
}
