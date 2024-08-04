package authentication

import (
	"encoding/json"
	"net/http"

	"github.com/mulan17/project-user-service/internal/user"
	"github.com/mulan17/project-user-service/pkg/token"
)

// Зберігає об'єкт PostgresStorage для доступу до бази даних
type AuthHandler struct {
	UserStorage *user.PostgresStorage
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var usr user.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, "Could not parse request data", http.StatusBadRequest)
		return
	}

	usr, err = h.UserStorage.Login(usr.Email, usr.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	token, err := token.GenerateToken(usr.Email, usr.Role, usr.ID)

	if err != nil {
		http.Error(w, "Problem with generating a token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
		"token":   token,
	})
}
