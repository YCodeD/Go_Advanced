package dao

import (
	"book_project/internal/pkg/db"
)

type Dao struct {
	GormDB db.GormDB
}

func NewDao(g db.GormDB) Dao {
	return Dao{
		GormDB: g,
	}
}

func (d *Dao) GetBook() (book *db.Book, err error) {

	err = d.GormDB.DB.Find(&book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}
