package Signup

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/rs/zerolog/log"
)

type Order struct {
}

type PersonalInformation struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type User struct {
	PersonalInfo PersonalInformation
	ID           string `json:"ID"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Role         string `json: "role"`
	Orders       []Order
}

var (
	usersM sync.Mutex
	users  []User
)

func CreateUser(w http.ResponseWriter, r *http.Request, jsonData []byte) {
	var u User
	// Варіант обробки помилки якщо беремо дані з тіла запита
	// err := json.NewDecoder(r.Body).Decode(&u)

	// Варіант обробки помилки якщо дані про юзера приходять в готовому JSON парсимо його
	err := json.Unmarshal(jsonData, &u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Debug().Err(err).Msg("Failed to decode JSON")
		return
	}

	usersM.Lock()
	defer usersM.Unlock()
	users = append(users, u)
	w.WriteHeader(http.StatusCreated)
}

// Створила цю функц щоб перевірити чи відправляються юзери на сервер
func GetUsers(w http.ResponseWriter, r *http.Request) {
	usersM.Lock()
	defer usersM.Unlock()

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Debug().Err(err).Msg("Failed to encode JSON response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
