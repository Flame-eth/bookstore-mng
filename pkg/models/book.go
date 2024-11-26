package models

import (
	"github.com/Flame-eth/bookstore-mng/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.SetupDatabase()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db.Where("ID = ?", id).Find(&getBook)
	return &getBook, db
}

func (b *Book) UpdateBook(id int64) *Book {
	db.Save(&b)
	return b
}

func DeleteBook(id int64) *Book {
	var book Book
	db.Where("ID = ?", id).Delete(&book)
	return &book
}
