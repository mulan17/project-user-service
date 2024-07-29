package User

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

var s InMemStorage

type CreateUserRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type service interface {
	SignUp(email, password string)
	GetUsers() []User
}

type Handler struct {
	s service
}

func NewHandler(s service) Handler {
	return Handler{
		s: s,
	}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateUserRequestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Debug().Err(err).Msg("Failed to decode JSON response")
		return
	}

	h.s.SignUp(reqBody.Email, reqBody.Password)

	w.WriteHeader(http.StatusCreated)
}

// Створила цю функц щоб перевірити чи відправляються юзери на сервер
func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.s.GetUsers()
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Debug().Err(err).Msg("Failed to encode JSON response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
