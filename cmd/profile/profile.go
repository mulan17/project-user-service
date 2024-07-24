package main

import (
	"fmt"
	"net/http"

	"github.com/mulan17/project-user-service/internal/profile"
)

func main() {
	storage := profile.NewStorage()

	users := profile.UserResource{
		S: storage,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", users.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", users.GetUserById)
	mux.HandleFunc("POST /users", users.CreateUser)
	mux.HandleFunc("PUT /users/{id}", users.UpdateUser)

	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		fmt.Println(err.Error())
	}
}
