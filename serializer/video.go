package serializer

import "singo/model"

type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	Url		  string `json:"url"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Url:	   item.Url,
		Title:     item.Title,
		Info:      item.Info,
		Avatar:    item.AvatarURL(),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

func BuildVideoResponse(item model.Video) Response {
	return Response{
		Data: BuildVideo(item),
	}
}

func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos

}

func BuildListResponse(items []model.Video) Response {
	return Response{
		Data: BuildVideos(items),
	}
}
