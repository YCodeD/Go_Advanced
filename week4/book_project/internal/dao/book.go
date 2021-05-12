package dao

import "book_project/internal/pkg/db"

func GetBook() (book *db.Book, err error) {
	db := db.InitDB()
	err = db.Find(&book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}
