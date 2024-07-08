package User

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

var s Storage

type CreateUserRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateUserRequestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Debug().Err(err).Msg("Failed to decode JSON response")
		return
	}

	s.Create(New(reqBody.Email, reqBody.Password))

	w.WriteHeader(http.StatusCreated)
}

// Створила цю функц щоб перевірити чи відправляються юзери на сервер
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := s.users

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Debug().Err(err).Msg("Failed to encode JSON response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
