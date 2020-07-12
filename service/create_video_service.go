package service

import (
	"github.com/jinzhu/gorm"
	"singo/model"
	"singo/serializer"
)

type CreateVideoService struct {
	Title  string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Url    string `form:"url" json:"url"`
	Info   string `form:"info" json:"info" binding:"required,min=5,max=300"`
	Avatar string `form:"avatar" json:"avatar"`
}

func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Model:  gorm.Model{},
		Title:  service.Title,
		Info:   service.Info,
		Url:	service.Url,
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
