package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type storage interface {
	Create(u User)
	GetUsers() []User
	Exists(email string) bool
	UpdateUser(reqBody User, id string) bool
	GetUserById(id string) (User, bool)
	BlockUser(id string) bool
	LimitUser(id string) bool
}

type Service struct {
	s storage
}

func NewService(s storage) *Service {
	return &Service{
		s: s,
	}
}

func (s *Service) SignUp(email, password string) error {
	if s.s.Exists(email) {
		return fmt.Errorf("user already exists")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// complex logic of gathering user data
	user := New(email, string(hash))

	//user notifications: emails, sms etc
	s.s.Create(user)

	//sending internal events to notify other services that user was created
	//etc
	return nil
}

func (s *Service) GetUsers() ([]UserResponse, error) {
	users := s.s.GetUsers()
	var response []UserResponse

	for _, u := range users {
		response = append(response, UserResponse{
			ID:       u.ID,
			Email:    u.Email,
			Role:     u.Role,
			Name:     u.Name,
			Lastname: u.Lastname,
			Status:   u.Status,
		})
	}

	return response, nil
}

func (s *Service) GetUserById(id string) (User, bool) {
	user, ok := s.s.GetUserById(id)
	return user, ok
}

func (s *Service) UpdateUser(reqBody User, id string) bool {
	ok := s.s.UpdateUser(reqBody, id)
	return ok
}

func (s *Service) BlockUser(id string) bool {
	ok := s.s.BlockUser(id)
	return ok
}

func (s *Service) LimitUser(id string) bool {
	ok := s.s.LimitUser(id)
	return ok
}
