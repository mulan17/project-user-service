package UserSignup

import (
	"net/http"

	User "github.com/mulan17/project-user-service/internal/user"
	"github.com/rs/zerolog/log"
)

func UserSignUp() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user", User.CreateUser)

	mux.HandleFunc("GET /users", User.GetUsers)

	error := http.ListenAndServe(":8080", mux)
	if error != nil {
		log.Fatal().Err(error).Msg("Failed to listen and serve")
	}
}
