package user

import (
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	ID       string `json:"ID"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Status   string `json:"status"`
}

type UserResponse struct {
	ID       string `json:"ID"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Status   string `json:"status"`
}

func New(email, password string) User {
	return User{
		ID:       time.Now().Format("20060102") + strconv.Itoa(rand.Intn(1000)),
		Role:     "buyer",
		Email:    email,
		Password: password,
		Name:     "nil",
		Lastname: "nil",
		Status:   "active",
	}

}
