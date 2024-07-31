package user

import (
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	ID       string `json:"ID"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

func New(email, password string) User {
	return User{
		ID:       time.Now().Format("20060102") + strconv.Itoa(rand.Intn(1000)),
		Role:     "buyer",
		Email:    email,
		Password: password,
		Name: "Nastya",
		Lastname: "Mulyukova",
	}

}