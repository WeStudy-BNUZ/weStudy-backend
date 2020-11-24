package classifyException

import "my_demo/models"

func CollageIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "学院类别不存在",
	}
}
