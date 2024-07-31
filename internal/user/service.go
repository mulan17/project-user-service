package user

import "fmt"

type storage interface {
	Create(u User)
	GetUsers() []User
	Exists(email string) bool
	UpdateUser(reqBody User, id string) bool
	GetUserById(id string) (User, bool)
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
	// complex logic of gathering user data
	user := New(email, password)

	//user notifications: emails, sms etc
	s.s.Create(user)

	//sending internal events to notify other services that user was created
	//etc
	return nil
}

func (s *Service) GetUsers() []User {
	return s.s.GetUsers()
}

func (s *Service) GetUserById(id string) (User, bool) {
	user, ok := s.s.GetUserById(id)
	return user, ok
}

func (s *Service) UpdateUser(reqBody User, id string) bool {
	ok := s.s.UpdateUser(reqBody, id)
 	return ok
}

