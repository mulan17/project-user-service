package user

import (
	"fmt"
	"log"

	"github.com/mulan17/project-user-service/pkg/hashing"
	// "golang.org/x/crypto/bcrypt"
)

type storage interface {
	Create(u User) error
	GetUsers() ([]User, error)
	Exists(email string) (bool, error)
	UpdateUser(reqBody User, id string) error
	GetUserById(id string) (User, error)
	BlockUser(id string) error
	LimitUser(id string) error
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
	// if s.s.Exists(email) {
	// 	return fmt.Errorf("user already exists")
	// }
	exists, err := s.s.Exists(email)

	if exists {
		return fmt.Errorf("user already exists")
	}
	if err != nil {
		return fmt.Errorf("checking if users exist: %v", err)
	}
	log.Printf("password that we get %v", password)

	hash, err := hashing.HashPassword(password)
	log.Printf("hash that we get %v", hash)
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
	users, err := s.s.GetUsers()
	
	if err != nil {
		return nil, fmt.Errorf("getting users: %v", err)
	}
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

func (s *Service) GetUserById(id string) (User, error) {
	user, err := s.s.GetUserById(id)
	if err != nil {
		return User{}, fmt.Errorf("getting user by id: %v", err)
	}
	return user, nil
}

func (s *Service) UpdateUser(reqBody User, id string) error {
	err := s.s.UpdateUser(reqBody, id)
	if err != nil {
		return fmt.Errorf("updating user: %v", err)
	}
	return nil
}

func (s *Service) BlockUser(id string) error {
	err := s.s.BlockUser(id)
	if err != nil {
		return fmt.Errorf("blocking user: %v", err)
	}
	return nil
}

func (s *Service) LimitUser(id string) error {
	err := s.s.LimitUser(id)
	if err != nil {
		return fmt.Errorf("limiting user: %v", err)
	}
	return nil
}
