package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

// GetDB connects to MySQL (Docker or local) and returns *gorm.DB
func GetDB() *gorm.DB {
	// Load .env file (ignore error if missing)
	_ = godotenv.Load()

	// Build DSN from environment variables
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}
	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		pass = "root"
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "bookstore"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)

	// Attempt connection
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		// If host is docker service `db`, fallback to localhost (for local runs)
		if host == "db" {
			log.Printf("Initial DB connect failed; retrying with localhost: %v", err)
			host = "127.0.0.1"
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)
			db, err = gorm.Open("mysql", dsn)
			if err != nil {
				log.Fatalf("Failed to connect to MySQL after retrying with localhost: %v", err)
				return nil
			}
		} else {
			log.Fatalf("Failed to connect to MySQL: %v", err)
			return nil
		}
	}

	// Create database if it doesn't exist
	if err := db.Exec("CREATE DATABASE IF NOT EXISTS " + name).Error; err != nil {
		log.Fatalf("Could not create database: %v", err)
		return nil
	}

	// Reconnect to the newly created database
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL after creating database: %v", err)
		return nil
	}

	log.Println("Connected to MySQL successfully")
	return db
}

// GetDBInstance returns the global DB instance
func GetDBInstance() *gorm.DB {
	return db
}

// GetJWTSecret returns the JWT secret from environment or a default
func GetJWTSecret() string {
	_ = godotenv.Load()
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}
