package serializer

import "singo/model"

type Book struct {
	ID        uint   `json:"id"`
	Name     string `json:"name"`
	Price      float64 `json:"price"`
	Number		  int `json:"number"`
	CreatedAt int64  `json:"created_at"`
}

func BuildBook(item model.Book) Book {
	return Book{
		ID:        item.ID,
		Name:	   item.Name,
		Price:     item.Price,
		Number:      item.Number,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

func BuildBookResponse(item model.Book) Response {
	return Response{
		Data: BuildBook(item),
	}
}

func BuildBooks(items []model.Book) (books []Book) {
	for _, item := range items {
		book := BuildBook(item)
		books = append(books, book)
	}
	return books

}

func BuildBooksResponse(items []model.Book) Response {
	return Response{
		Data: BuildBooks(items),
	}
}
