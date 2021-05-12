package service

import (
	v1 "book_project/api/book/v1"
	"book_project/internal/dao"
	"context"
)

type BookService struct {
	v1.BookServerServer
}

func NewBookService() *BookService {
	return &BookService{}
}

func (s *BookService) GetBookInfo(ctx context.Context, req *v1.GetBookInfoRequest) (*v1.GetBookInfoResponse, error) {
	book, err := dao.GetBook()
	if err != nil {
		return nil, err
	}

	resp := &v1.GetBookInfoResponse{
		BookTitle:  book.Title,
		BookAuthor: book.Author,
	}

	return resp, nil
}
