package user

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
		Role:     "buyer",
		Email:    email,
		Password: password,
		Name:     "nil",
		Lastname: "nil",
		Status:   "active",
	}

}
