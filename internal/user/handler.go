package user

import (
	"encoding/json"
	"errors"
	"net/http"

	service_errors "github.com/mulan17/project-user-service/internal/errors"
	"github.com/rs/zerolog/log"
)

type CreateUserRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type service interface {
	SignUp(email, password string) error
	GetUsers() ([]UserResponse, error)
	GetUserById(id string) (User, error)
	UpdateUser(reqBody User, id string) error
	BlockUser(id string) error
	LimitUser(id string) error
}

type Handler struct {
	s service
}

func NewHandler(s service) Handler {
	return Handler{
		s: s,
	}
}

// TODO розібратись по функції щоб потіп додати помилки, бо ми тут пропусти та необмазували помилками
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
		if errors.Is(err, service_errors.ErrUserAlreadyExists) {
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

	user, err := h.s.GetUserById(id)

	if err != nil {
		log.Debug().Err(err).Msg("Failed to get user by id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		log.Debug().Err(err).Msg("Failed to encode")
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
		log.Debug().Err(err).Msg("Failed to decode")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.s.UpdateUser(reqBody, id)

	if err != nil {
		log.Debug().Err(err).Msg("Failed to update user")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch user"})
	}

	w.Header().Set("Content-Type", "application/json")

}

func (h Handler) BlockUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.s.BlockUser(id)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to block user")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to block user"})
	}

	w.Header().Set("Content-Type", "application/json")

}

func (h Handler) LimitUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.s.LimitUser(id)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to limit user")

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to limit user"})
	}

	w.Header().Set("Content-Type", "application/json")

}
