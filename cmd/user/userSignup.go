package UserSignup

import (
	"net/http"

	User "github.com/mulan17/project-user-service/internal/user"
	UserEmulator "github.com/mulan17/project-user-service/pkg/userEmulator"
	"github.com/rs/zerolog/log"
)

func UserSignUp() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		userJSON, err := UserEmulator.EmulateUser()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Failed to emulate user: %v", err)
			return
		}

		User.CreateUser(w, r, userJSON)
	})
	
	mux.HandleFunc("GET /users", User.GetUsers)

	error := http.ListenAndServe(":8080", mux)
	if error != nil {
		log.Fatal().Err(error).Msg("Failed to listen and serve")
	}
}
