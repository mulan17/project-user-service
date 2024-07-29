package UserSignup

import (
	"net/http"
	"os"

	User "github.com/mulan17/project-user-service/internal/user"
	"github.com/rs/zerolog/log"
)

func UserSignUp() {
	mux := http.NewServeMux()
	connStr := os.Getenv("POSTGRES_CONN_STR")

	// userStorage := User.NewInMemStorage()
	userStorage, err := User.NewPostgresStorage(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}
	userService := User.NewService(userStorage)
	userHandler := User.NewHandler(userService)

	mux.HandleFunc("POST /user", userHandler.Create)
	mux.HandleFunc("GET /users", userHandler.GetUsers)

	error := http.ListenAndServe(":8080", mux)
	if error != nil {
		log.Fatal().Err(error).Msg("Failed to listen and serve")
	}
}
