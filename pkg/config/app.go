package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

// GetDB connects to Docker MySQL and returns *gorm.DB
func GetDB() *gorm.DB {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default DSN")
	}

	// Get DSN from environment variable
	dsn := os.Getenv("DB_DSN")

	// Connect to MySQL
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
		return nil
	}

	// Auto-create database if it doesn't exist
	if err := d.Exec("CREATE DATABASE IF NOT EXISTS bookstore").Error; err != nil {
		log.Fatalf("Could not create database: %v", err)
		return nil
	}

	// Reconnect to the newly created database
	d, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL after creating database: %v", err)
		return nil
	}

	db = d
	log.Println("Connected to MySQL successfully")
	return db
}

// Getdb returns the global DB instance
func Getdb() *gorm.DB {
	return db
}
