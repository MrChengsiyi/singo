package service

import (
	"singo/model"
	"singo/serializer"
)

type DeleteBookService struct {
}

func (service *DeleteBookService) Delete(id string) serializer.Response {
	var book model.Book
	if err := model.DB.First(&book, id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "书籍不存在",
			Error: err.Error(),
		}
	}

	err := model.DB.Delete(&book).Error
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "书籍删除失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{}
}
