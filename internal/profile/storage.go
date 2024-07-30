package profile

import (
	"sync"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Storage struct {
	m     sync.Mutex
	users map[string]User
}

func NewStorage() *Storage {
	return &Storage{
		users: make(map[string]User),
	}
}

func (s *Storage) GetAllUsers() []User {
	users := make([]User, 0, len(s.users))

	for _, user := range s.users {
		users = append(users, user)
	}

	return users
}

func (s *Storage) CreateUser(p User) (string, bool) {
	s.m.Lock()

	defer s.m.Unlock()

	id, err := gonanoid.New()

	if err != nil {
		return "", false
	}

	p.ID = id
	s.users[p.ID] = p

	return id, true
}

func (s *Storage) GetUserById(id string) (User, bool) {
	p, ok := s.users[id]

	if !ok {
		return User{}, false
	}

	return p, true
}

func (s *Storage) UpdateUser(id string, user User) bool {
	s.m.Lock()

	defer s.m.Unlock()

	t, ok := s.users[id]

	if !ok {
		return false
	}

	user.ID = t.ID
	s.users[user.ID] = user

	return true
}
