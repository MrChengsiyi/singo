package service

import (
	"singo/model"
	"singo/serializer"
)

type ShowVideoService struct {
	Title  string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Info   string `form:"info" json:"info" binding:"required,min=5,max=300"`
	Avatar string `form:"avatar" json:"avatar"`
}

func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	return serializer.BuildVideoResponse(video)
}
