package service

import (
	"singo/model"
	"singo/serializer"
)

type ShowBookService struct {

}

func (service *ShowBookService) Show(id string) serializer.Response {
	var book model.Book
	//int, _ := strconv.Atoi(id)
	//fmt.Println(int)
	if err := model.DB.First(&book,id).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "书籍不存在",
			Error: err.Error(),
		}
	}

	return serializer.BuildBookResponse(book)
}