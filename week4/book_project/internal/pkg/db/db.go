package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
}

type GormDB struct {
	DB *gorm.DB
}

func NewGormDB() GormDB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(Book{})

	return GormDB{
		DB: db,
	}
}
