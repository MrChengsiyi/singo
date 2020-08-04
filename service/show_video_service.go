package service

import (
	"singo/model"
	"singo/serializer"
)

type ShowVideoService struct {

}

func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	//int, _ := strconv.Atoi(id)
	//fmt.Println(int)
	if err := model.DB.First(&video,id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	return serializer.BuildVideoResponse(video)
}
