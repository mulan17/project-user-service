package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type CreateUserRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type service interface {
	SignUp(email, password string) error
	GetUsers() ([]UserResponse, error)
	GetUserById(id string) (User, bool)
	UpdateUser(reqBody User, id string) bool
	BlockUser(id string) bool
	LimitUser(id string) bool
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

func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	response, err := h.s.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Debug().Err(err).Msg("Failed to get users")
		return
	}

	err = json.NewEncoder(w).Encode(response)
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

	ok := h.s.UpdateUser(reqBody, id)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch user"})
	}

	w.Header().Set("Content-Type", "application/json")

}

func (h Handler) BlockUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	ok := h.s.BlockUser(id)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to block user"})
	}

	w.Header().Set("Content-Type", "application/json")

}

func (h Handler) LimitUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	ok := h.s.LimitUser(id)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to limit user"})
	}

	w.Header().Set("Content-Type", "application/json")

}
