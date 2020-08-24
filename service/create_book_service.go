package service

import (
	"singo/model"
	"singo/serializer"
)

type CreateBookService struct {
	Name  string `form:"name" json:"name" `
	Price	   float64 `form:"price" json:"price"`
	Number   int `form:"number" json:"number"`
}


func (service *CreateBookService) Create() serializer.Response {
	book := model.Book{
		Name:  service.Name,
		Price:	service.Price,
		Number:   service.Number,

	}

	if err := model.DB.Create(&book).Error; err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "书籍保存失败",
			Error: err.Error(),
		}
	}

	return serializer.BuildBookResponse(book)
}
