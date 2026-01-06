package config

import (
	"log"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

// GetDB attempts to open a DB connection using DB_DSN env var (or default);
// if it cannot connect it logs the error and returns nil instead of panicking.
func GetDB() *gorm.DB {
	// prefer explicit DSN from environment for safety
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		// reasonable default (may fail if no MySQL running)
		dsn = "root:password@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local"
	}
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("warning: failed to connect to MySQL: %v\nattempting sqlite fallback", err)
		// try sqlite in-memory fallback so the app can run without MySQL
		sd, serr := gorm.Open("sqlite3", ":memory:")
		if serr != nil {
			log.Printf("warning: sqlite fallback failed: %v\nDB is not initialized; continuing without DB", serr)
			return nil
		}
		db = sd
		return db
	}
	db = d
	return db
}

func Getdb() * gorm.DB {
	return db
}

//purpose of this file is to create a connection with the database
//to return the database so that the other files can use it

