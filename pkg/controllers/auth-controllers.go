package controllers

import (
	"encoding/json"
	"net/http"	
	"Go-bookstore/pkg/models"
)

var NewAuth models.Auth

// Login handles user login
func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var auth models.Auth
	_ = json.NewDecoder(r.Body).Decode(&auth)
	token, err := models.Login(auth.Username, auth.Password)
	if err != nil {
		http.Error(w, "login failed: "+err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
// Register handles user registration

func Register(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var auth models.Auth

	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		http.Error(w, "invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.Register(auth.Username, auth.Password); err != nil {
		http.Error(w, "registration failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "registration successful"})
}	
