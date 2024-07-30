package user

import (
	"encoding/json"
	"net/http"
)

// Структура профілю користувача
type Profile struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

// Ініціалізація користувачів
var users = map[string]Profile{
	"admin":   {Name: "Admin", Email: "admin@example.com", Status: "active"},
	"user1":   {Name: "User One", Email: "user1@example.com", Status: "active"},
	"blocked": {Name: "Blocked User", Email: "blocked@example.com", Status: "blocked"},
}

// Функція для отримання профілю користувача за ім'ям користувача
func GetProfile(username string) Profile {
	return users[username]
}

// Функція для оновлення профілю користувача за ім'ям користувача
func UpdateProfile(username string, profile Profile) {
	users[username] = profile
}

// Функція для отримання списку всіх активних покупців
func GetAllCustomers() []Profile {
	var customers []Profile
	for _, profile := range users {
		if profile.Status == "active" {
			customers = append(customers, profile)
		}
	}
	return customers
}

// Функція для блокування користувача за ім'ям користувача
func BlockCustomer(username string) {
	if user, exists := users[username]; exists {
		user.Status = "blocked"
		users[username] = user
	}
}

// Функція для перегляду профілю користувача
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	profile := GetProfile(username)
	json.NewEncoder(w).Encode(profile)
}

// Функція для редагування профілю користувача
func EditProfile(w http.ResponseWriter, r *http.Request) {
	var profile Profile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	username := r.URL.Query().Get("username")
	UpdateProfile(username, profile)
	w.Write([]byte("Profile updated successfully"))
}

