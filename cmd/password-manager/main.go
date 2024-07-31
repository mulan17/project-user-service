package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/mulan17/project-user-service/tree/feature/admin-service/internal/admin"
	"github.com/ulan17/project-user-service/tree/feature/admin-service/hixi4/internal/user"
)

func main() {
	// Отримання порту із змінної середовища
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Порт за замовчуванням, якщо змінна середовища не встановлена
	}

	// Маршрут для перегляду списку покупців
	http.HandleFunc("/admin/customers", admin.ViewCustomers)
	// Маршрут для блокування покупців
	http.HandleFunc("/admin/block", admin.BlockCustomer)
	// Маршрут для перегляду профілю користувача
	http.HandleFunc("/user/profile", user.ViewProfile)
	// Маршрут для редагування профілю користувача
	http.HandleFunc("/user/edit", user.EditProfile)

	fmt.Printf("Starting server at :%s\n", port)
	// Запуск сервера на заданому порту
	http.ListenAndServe(":"+port, nil)
}

