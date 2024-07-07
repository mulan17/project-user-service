package main

import (
	"net/http"

	Signup "github.com/mulan17/project-user-service/internal/users/feature-signup"
	"github.com/rs/zerolog/log"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		userJSON, err := Signup.EmulateUser()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Failed to emulate user: %v", err)
			return
		}

		Signup.CreateUser(w, r, userJSON)
	})
	mux.HandleFunc("GET /users", Signup.GetUsers)

	error := http.ListenAndServe(":8080", mux)
	if error != nil {
		log.Fatal().Err(error).Msg("Failed to listen and serve")
	}
}
