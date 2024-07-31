//головний мейн

package main

import (
	"net/http"
	"os"

	"github.com/mulan17/project-user-service/internal/profile"
	User "github.com/mulan17/project-user-service/internal/user"
	"github.com/mulan17/project-user-service/pkg/authentication"
	"github.com/rs/zerolog/log"
)

func main() {

	storage := profile.NewStorage()

	users := profile.UserResource{
		S: storage,
	}

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

	mux.HandleFunc("POST /login", authentication.Login)

	//mux.HandleFunc("GET /users", users.GetAllUsers) get users
	mux.HandleFunc("GET /users/{id}", users.GetUserById)
	//mux.HandleFunc("POST /users", users.CreateUser)
	mux.HandleFunc("PUT /users/{id}", users.UpdateUser)

	// Маршрут для перегляду списку покупців
	http.HandleFunc("/admin/customers", admin.ViewCustomers)
	// Маршрут для блокування покупців
	http.HandleFunc("/admin/block", admin.BlockCustomer)
	// Маршрут для перегляду профілю користувача
	http.HandleFunc("/user/profile", user.ViewProfile)
	// Маршрут для редагування профілю користувача
	http.HandleFunc("/user/edit", user.EditProfile)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to listen and serve")
	}

}

// Функція для блокування користувача за ім'ям користувача
func BlockCustomer(username string) {
	if user, exists := users[username]; exists {
		user.Status = "blocked"
		users[username] = user
	}
}

// ПРИКЛАД Захищеного маршруту
// mux.HandleFunc("/protected-route", func(w http.ResponseWriter, r *http.Request) {
// 	// Використовуємо middleware Authenticate для перевірки автентифікації
// 	authentication_check.Authenticate(http.HandlerFunc(yourProtectedHandler)).ServeHTTP(w, r)
// })
