package models

import (
	"errors"
	"github.com/ashwinkg/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	ISBN        string `json:"isbn"`
}

func InitDB() error {
	config.Connect()
	db = config.GetDB()
	return db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	result := db.Create(&b)
	return b, result.Error
}

func GetAllBooks() ([]Book, error) {
	var Books []Book
	result := db.Find(&Books)
	return Books, result.Error
}

func GetBookById(id int64) (*Book, error) {
	var getBook Book
	result := db.Where("ID=?", id).Find(&getBook)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &getBook, result.Error
}

func DeleteBook(id int64) error {
	var deleteBook Book
	result := db.Where("ID=?", id).Delete(&deleteBook)
	return result.Error
}

func UpdateBook(id int64, value interface{}) error {
	result := db.Model(&Book{}).Where("ID=?", id).Updates(value)
	return result.Error
}
