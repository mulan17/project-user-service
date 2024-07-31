package User

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

// var s InMemStorage

type CreateUserRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type service interface {
	SignUp(email, password string) error
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

	err = h.s.SignUp(reqBody.Email, reqBody.Password)
	if err != nil {
		if err.Error() == "user already exists" {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{"error": "User already exists"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
		}
		return
	}

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

func (h Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	user, ok := h.s.GetUserById(id)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
}

func (h Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var reqBody User

	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		fmt.Println("Failed to encode: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.s.UpdateUser(reqBody, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch user"})
		}

	w.Header().Set("Content-Type", "application/json")

}