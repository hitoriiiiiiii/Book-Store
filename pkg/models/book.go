package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"Go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.GetDB()
	db = config.Getdb()
	if db != nil {
		db.AutoMigrate(&Book{})
	}
}

func (b *Book) CreateBook() (*Book, error) {
	if db == nil {
		return nil, errors.New("database not initialized")
	}
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Book) GetAllBooks() ([]Book, error) {
	if db == nil {
		return nil, errors.New("database not initialized")
	}
	var Books []Book
	if err := db.Find(&Books).Error; err != nil {
		return nil, err
	}
	return Books, nil
}

func (b *Book) GetBooksById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	if db == nil {
		return &getBook, nil
	}
	d := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, d
}

func (b *Book) DeleteBook(Id int64) Book {
	var book Book
	if db == nil {
		return book
	}
	db.Where("ID = ?", Id).Delete(&book)
	return book
}

//flow of this file is to create a book model
//routes -> controllers -> models -> config -> database