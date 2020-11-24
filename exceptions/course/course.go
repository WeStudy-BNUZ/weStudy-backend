package courseException

import "grpc-demo/models"

func CourseCreateFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "课程创建失败",
	}
}

func SessionMarshalFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5415,
		Message: "课程场次序列化失败",
	}
}

func SessionUnmarshalFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5416,
		Message: "课程场次反序列化失败",
	}
}

func CourseIsNotExsit() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "课程不存在",
	}
}

func CoursePutFail() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "课程修改失败",
	}
}

func CourseNotPut() models.RestfulAPIResult {
	return models.RestfulAPIResult{
		Status:  false,
		ErrCode: 5412,
		Message: "课程未发布",
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
