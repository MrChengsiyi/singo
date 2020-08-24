package service

import (
	"singo/model"
	"singo/serializer"
)

type ListBookService struct {
}

func (service *ListBookService) List() serializer.Response {
	var books []model.Book
	if err := model.DB.Find(&books).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "书籍不存在",
			Error: err.Error(),
		}
	}

	return serializer.BuildBooksResponse(books)
}
