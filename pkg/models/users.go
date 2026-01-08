package models

import (
	"errors"
	"sync"
	"Go-bookstore/pkg/config"
	"Go-bookstore/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

var mu sync.Mutex

type Auth struct {
   Username string `json:"username"`
   Password string `json:"password"` 
}

func Register(username, password string) error {
	mu.Lock()
	defer mu.Unlock()

	var existingUser Auth
	if err := db.Where("username = ?", username).First(&existingUser).Error; err == nil {
		return errors.New("username already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := Auth{
		Username: username,
		Password: string(hashedPassword),
	}
	if err := db.Create(&newUser).Error; err != nil {
		return err
	}
	return nil
}

func Login(username, password string) (string, error) {
	var user Auth
	if db == nil {
		// ensure DB is initialized
		db = config.GetDB()
	}
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", errors.New("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWT(username)
	if err != nil {
		return "", err
	}
	return token, nil
}

