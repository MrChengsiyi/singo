package service

import (
	"singo/model"
	"singo/serializer"
)

type UpdateBookService struct {
	Name string `form:"name" json:"name" binding:"required,min=2,max=30"`
	Price  float64 `form:"price" json:"price" binding:"required"`
	Number int `form:"number" json:"number" binding:"required"`
}

func (service *UpdateBookService) Update(id string) serializer.Response {
	var book model.Book
	if err := model.DB.First(&book, id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "书籍不存在",
			Error: err.Error(),
		}
	}
	book.Name = service.Name
	book.Price = service.Price
	book.Number = service.Number

	err := model.DB.Save(&book).Error
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}
	return serializer.BuildBookResponse(book)
}
