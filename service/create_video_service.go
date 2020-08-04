package service

import (
	"singo/model"
	"singo/serializer"
)

type CreateVideoService struct {
	Title  string `form:"title" json:"title" `
	Url	   string `form:"url" json:"url"`
	Info   string `form:"info" json:"info"`
	Avatar string `form:"avatar" json:"avatar"`
}


func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title:  service.Title,
		Url:	service.Url,
		Info:   service.Info,
		Avatar: service.Avatar,
	}

	if err := model.DB.Create(&video).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}

	return serializer.BuildVideoResponse(video)
}
