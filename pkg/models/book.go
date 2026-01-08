package models

import (
	"github.com/jinzhu/gorm"
	"Go-bookstore/pkg/config"
)

var db *gorm.DB
var fallbackBooks []Book

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	// initialize and assign the DB returned by config
	db = config.GetDB()
	if db != nil {
		db.AutoMigrate(&Book{})
	}
}

func (b *Book) CreateBookHandler() (*Book, error) {
	if db == nil {
		// store in memory fallback and assign a simple incrementing ID
		b.ID = uint(len(fallbackBooks) + 1)
		fallbackBooks = append(fallbackBooks, *b)
		return b, nil
	}
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Book) GetAllBooks() ([]Book, error) {
	if db == nil {
		return fallbackBooks, nil
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
		// search fallback slice by numeric ID matching index+1 if Model ID absent
		for i := range fallbackBooks {
			// if fallbackBooks were stored without Model ID, try matching by index+1
			if int64(i+1) == Id {
				return &fallbackBooks[i], nil
			}
			if fallbackBooks[i].ID == uint(Id) {
				return &fallbackBooks[i], nil
			}
		}
		return &getBook, nil
	}
	d := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, d
}

func (b *Book) DeleteBook(Id int64) Book {
	var book Book
	if db == nil {
		// remove from fallbackBooks by index match
		for i := range fallbackBooks {
			if int64(i+1) == Id || fallbackBooks[i].ID == uint(Id) {
				book = fallbackBooks[i]
				fallbackBooks = append(fallbackBooks[:i], fallbackBooks[i+1:]...)
				break
			}
		}
		return book
	}
	db.Where("ID = ?", Id).Delete(&book)
	return book
}

//flow of this file is to create a book model
//routes -> controllers -> models -> config -> database