// +build wireinject

package main

import (
	"book_project/internal/dao"
	"book_project/internal/pkg/db"
	"book_project/internal/service"
	"github.com/google/wire"
)

func initBookServer() *service.BookService {
	wire.Build(db.NewGormDB(), dao.NewDao(), service.NewBookService())
	return service.BookService{}
}
