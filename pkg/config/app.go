package config

import (
	"github.com/jinzhu/gorm"
	"_ github.com/go-sql-driver/mysql"
)

var (
	db * gorm.DB
)

func GetDB() * gorm.DB {
	d, err := gorm.Open("mysql", "root:password@tcp()")
	if err != nil {
		panic("failed to connect database")
	}
	db = d
}

func Getdb() * gorm.DB {
	return db
}

//purpose of this file is to create a connection with the database
//to return the database so that the other files can use it

