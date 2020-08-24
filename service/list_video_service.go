package service

import (
	"singo/model"
	"singo/serializer"
)

type ListVideoService struct {
}

func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video  
	if err := model.DB.Find(&videos).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	return serializer.BuildVideosResponse(videos)
}
