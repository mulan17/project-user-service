package main

import (
	"net/http"
	"os"

	User "github.com/mulan17/project-user-service/internal/user"
	"github.com/mulan17/project-user-service/pkg/authentication"
	"github.com/mulan17/project-user-service/pkg/authentication_check"
	"github.com/rs/zerolog/log"
)

// func adminOnly(next http.Handler) http.Handler {
// 	return authentication_check.RoleMiddleware("admin", next)
// }

func main() {

	mux := http.NewServeMux()

	connStr := os.Getenv("POSTGRES_CONN_STR")

	userStorage, err := User.NewPostgresStorage(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	userService := User.NewService(userStorage)
	userHandler := User.NewHandler(userService)

	authHandler := &authentication.AuthHandler{
		UserStorage: userStorage,
	}

	mux.HandleFunc("/login", authHandler.Login) // Маршрут для логіну

	authenticatedRouter := http.NewServeMux()
	mux.HandleFunc("GET /users", userHandler.GetUsers) 
	mux.HandleFunc("PATCH /users/{id}", userHandler.UpdateUser)
	
	mux.HandleFunc("POST /users", userHandler.Create)
	authenticatedRouter.HandleFunc("GET /users/{id}", userHandler.GetUserById)

	authenticatedRouter.HandleFunc("/admin/block/{id}", userHandler.BlockUser)
	authenticatedRouter.HandleFunc("/admin/limit/{id}", userHandler.BlockUser)


	// mux.Handle("GET /users", authentication_check.Authenticate(authenticatedRouter))
	// mux.Handle("PUT /users/{id}", authentication_check.Authenticate(authenticatedRouter))
	mux.Handle("GET /users/{id}", authentication_check.Authenticate(authenticatedRouter))

	mux.Handle("/admin/block/{id}", authentication_check.Authenticate(authenticatedRouter))
	mux.Handle("/admin/limit/{id}", authentication_check.Authenticate(authenticatedRouter))

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}

}
