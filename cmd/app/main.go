package main

import (
	"net/http"
	"os"

	User "github.com/mulan17/project-user-service/internal/user"
	"github.com/rs/zerolog/log"
	// "github.com/mulan17/project-user-service/pkg/authentication"


)

func main() {

	mux := http.NewServeMux()

	connStr := os.Getenv("POSTGRES_CONN_STR")

	userStorage, err := User.NewPostgresStorage(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	userService := User.NewService(userStorage)
	userHandler := User.NewHandler(userService)

	mux.HandleFunc("POST /users", userHandler.Create)
	mux.HandleFunc("GET /users", userHandler.GetUsers)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUserById)
	mux.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)

	mux.HandleFunc("PUT /admin/block/{id}", userHandler.BlockUser)

	// mux.HandleFunc("POST /login", authentication.Login)
	
	// // Маршрут для перегляду списку покупців
	// http.HandleFunc("/admin/customers", admin.ViewCustomers)
	// // Маршрут для блокування покупців
	// http.HandleFunc("/admin/block", admin.BlockCustomer)
	// // Маршрут для перегляду профілю користувача
	// http.HandleFunc("/user/profile", user.ViewProfile)
	// // Маршрут для редагування профілю користувача
	// http.HandleFunc("/user/edit", user.EditProfile)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}

}


// // ПРИКЛАД Захищеного маршруту
// mux.HandleFunc("/protected-route", func(w http.ResponseWriter, r *http.Request) {
// 	// Використовуємо middleware Authenticate для перевірки автентифікації
// 	authentication_check.Authenticate(http.HandlerFunc(yourProtectedHandler)).ServeHTTP(w, r)
// })
