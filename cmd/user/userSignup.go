package UserSignup

import (
	"net/http"

	User "github.com/mulan17/project-user-service/internal/user"
	"github.com/rs/zerolog/log"
)

func UserSignUp() {
	mux := http.NewServeMux()
	userStorage := User.NewInMemStorage()
	userService := User.NewService(userStorage)
	userHandler := User.NewHandler(userService)

	mux.HandleFunc("POST /user", userHandler.Create)
	mux.HandleFunc("GET /users", userHandler.GetUsers)

	error := http.ListenAndServe(":8080", mux)
	if error != nil {
		log.Fatal().Err(error).Msg("Failed to listen and serve")
	}
}
